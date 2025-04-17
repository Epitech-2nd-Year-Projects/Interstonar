package utils

import "fmt"

func DisplayHelp() {
	fmt.Println("USAGE")
	fmt.Println("\t./interstonar [--global | --local] CONFIG FILE Px Py Vx Vy Vz")
	fmt.Println("\nDESCRIPTION")
	fmt.Println("\t--global  Launch program in global scene mode. The CONFIG_FILE will describe a scene containing massive spherical moving bodies of which at least one is a goal.")
	fmt.Println("\n\t--local   Launch program in local scene mode. The CONFIG_FILE will describe a scene containing massless motionless shapes.")
	fmt.Println("\n\tPi\tPosition coordinates of the rock")
	fmt.Println("\n\tVi\tVelocity vector of the rock")
	fmt.Println("\n\tCONFIG_FILE TOML configuration file describing a scene")
}
