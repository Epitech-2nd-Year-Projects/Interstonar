package config

import (
	"bufio"
	"fmt"
	. "interstonar/internal/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadGlobalConfig(filename string) (*GlobalConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%s): %v", filename, err)
	}
	defer file.Close()

	config := &GlobalConfig{
		Bodies: []GlobalBody{},
	}

	s := bufio.NewScanner(file)
	var currentBody *GlobalBody
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if line == "[[bodies]]" {
			if currentBody != nil {
				config.Bodies = append(config.Bodies, *currentBody)
			}
			currentBody = &GlobalBody{}
			continue
		}

		if currentBody != nil && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "name":
				re := regexp.MustCompile(`"([^"]*)"`)
				matches := re.FindStringSubmatch(value)
				if len(matches) > 1 {
					currentBody.Name = matches[1]
				}
			case "mass":
				currentBody.Mass, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("strconv.ParseFloat(%s): %v", value, err)
				}
			case "radius":
				value = strings.ReplaceAll(value, "_", "")
				currentBody.Radius, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("strconv.ParseFloat(%s): %v", value, err)
				}
			case "goal":
				currentBody.Goal = (value == "true")
			case "position":
				currentBody.Position, err = parseVector(currentBody.Position, value)
				if err != nil {
					return nil, fmt.Errorf("parseVector(%s): %v", value, err)
				}
			case "direction":
				currentBody.Direction, err = parseVector(currentBody.Direction, value)
				if err != nil {
					return nil, fmt.Errorf("parseVector(%s): %v", value, err)
				}
			}
		}
	}

	if currentBody != nil {
		config.Bodies = append(config.Bodies, *currentBody)
	}
	return config, nil
}

func LoadLocalConfig(filename string) (*LocalConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%s): %v", filename, err)
	}
	defer file.Close()

	config := &LocalConfig{
		Bodies: []LocalShape{},
	}

	s := bufio.NewScanner(file)
	var currentShape *LocalShape
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if line == "[[bodies]]" {
			if currentShape != nil {
				config.Bodies = append(config.Bodies, *currentShape)
			}
			currentShape = &LocalShape{}
			continue
		}

		if currentShape != nil && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "type":
				re := regexp.MustCompile(`"([^"]*)"`)
				matches := re.FindStringSubmatch(value)
				if len(matches) > 1 {
					currentShape.Type = matches[1]
				}
			case "radius":
				currentShape.Radius, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("strconv.ParseFloat(%s): %v", value, err)
				}
			case "height":
				currentShape.Height, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("strconv.ParseFloat(%s): %v", value, err)
				}
			case "inner_radius":
				currentShape.InnerRadius, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("strconv.ParseFloat(%s): %v", value, err)
				}
			case "outer_radius":
				currentShape.OuterRadius, err = strconv.ParseFloat(value, 64)
				if err != nil {
					return nil, fmt.Errorf("strconv.ParseFloat(%s): %v", value, err)
				}
			case "position":
				currentShape.Position, err = parseVector(currentShape.Position, value)
				if err != nil {
					return nil, fmt.Errorf("parseVector(%s): %v", value, err)
				}
			case "sides":
				currentShape.Sides, err = parseVector(currentShape.Sides, value)
				if err != nil {
					return nil, fmt.Errorf("parseVector(%s): %v", value, err)
				}
			}
		}
	}

	if currentShape != nil {
		config.Bodies = append(config.Bodies, *currentShape)
	}
	return config, nil
}

func parseVector(vec Vector3, value string) (Vector3, error) {
	reX := regexp.MustCompile(`x\s*=\s*([-+]?\d*\.?\d+(?:[eE][-+]?\d+)?)`)
	reY := regexp.MustCompile(`y\s*=\s*([-+]?\d*\.?\d+(?:[eE][-+]?\d+)?)`)
	reZ := regexp.MustCompile(`z\s*=\s*([-+]?\d*\.?\d+(?:[eE][-+]?\d+)?)`)

	matchesX := reX.FindStringSubmatch(value)
	matchesY := reY.FindStringSubmatch(value)
	matchesZ := reZ.FindStringSubmatch(value)

	if len(matchesX) > 1 {
		var err error
		vec.X, err = strconv.ParseFloat(matchesX[1], 64)
		if err != nil {
			return Vector3{}, fmt.Errorf("strconv.ParseFloat(%s, 64): %v", matchesX[1], err)
		}
	}
	if len(matchesY) > 1 {
		var err error
		vec.Y, err = strconv.ParseFloat(matchesY[1], 64)
		if err != nil {
			return Vector3{}, fmt.Errorf("strconv.ParseFloat(%s, 64): %v", matchesY[1], err)
		}
	}
	if len(matchesZ) > 1 {
		var err error
		vec.Z, err = strconv.ParseFloat(matchesZ[1], 64)
		if err != nil {
			return Vector3{}, fmt.Errorf("strconv.ParseFloat(%s, 64): %v", matchesZ[1], err)
		}
	}
	return vec, nil
}
