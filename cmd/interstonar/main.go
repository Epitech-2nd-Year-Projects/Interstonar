package main

import (
	"flag"
	"fmt"
	"interstonar/internal/config"
	"interstonar/internal/global"
	"interstonar/internal/local"
	"interstonar/internal/utils"
	"os"
	"strconv"
)

func main() {
	globalMode := flag.Bool("global", false, "")
	localMode := flag.Bool("local", false, "")
	flag.Parse()

	args := flag.Args()
	if len(args) != 7 {
		utils.DisplayHelp()
		os.Exit(84)
	}

	configFile := args[0]

	position := make([]float64, 3)
	velocity := make([]float64, 3)
	for i := 0; i < 3; i++ {
		var err error
		position[i], err = strconv.ParseFloat(args[i+1], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "strconv.ParseFloat(args[%d+1], 64): %v\n", i, err)
			os.Exit(84)
		}
		velocity[i], err = strconv.ParseFloat(args[i+4], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "strconv.ParseFloat(args[%d+4], 64): %v\n", i, err)
			os.Exit(84)
		}
	}

	if *globalMode {
		err := runGlobal(configFile, position, velocity)
		if err != nil {
			fmt.Fprintf(os.Stderr, "runGlobal(%s, %v, %v): %v\n", configFile, position, velocity, err)
			os.Exit(84)
		}
	} else if *localMode {
		err := runLocal(configFile, position, velocity)
		if err != nil {
			fmt.Fprintf(os.Stderr, "runLocal(%s, %v, %v): %v\n", configFile, position, velocity, err)
			os.Exit(84)
		}
	} else {
		utils.DisplayHelp()
		os.Exit(84)
	}
}

func runGlobal(configFile string, position, velocity []float64) error {
	conf, err := config.LoadGlobalConfig(configFile)
	if err != nil {
		return fmt.Errorf("config.LoadGlobalConfig: %v", err)
	}
	global.Simulate(conf, position, velocity)
	return nil
}

func runLocal(configFile string, position, velocity []float64) error {
	conf, err := config.LoadLocalConfig(configFile)
	if err != nil {
		return fmt.Errorf("config.LoadLocalConfig: %v", err)
	}
	local.Simulate(conf, position, velocity)
	return nil
}
