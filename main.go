/*
Copyright © 2022 NAME HERE seungjinyu93@gmail.com

*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/seungjinyu/glara/cmd"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	cmd.Execute()
}
