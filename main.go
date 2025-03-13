/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/DSGunasekara/easy-tasks/cmd"
	"github.com/DSGunasekara/easy-tasks/db"
)

func main() {
	fmt.Println("🚀 Task Manager CLI")
	db.InitDB()
	cmd.Execute()
}
