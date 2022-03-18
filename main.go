/*
Copyright © 2022 NAME HERE seungjinyu93@gmail.com
*/
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/seungjinyu/glara/cmd"
)

func main() {

	args := os.Args
	if len(args) < 3 {
		cmd.Execute()
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env")
		}
		cmd.Execute()
	}

}
