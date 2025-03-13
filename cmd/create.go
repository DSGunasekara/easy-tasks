/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/DSGunasekara/easy-tasks/db"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [project name]",
	Short: "create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		createProject(projectName)
	},
}

func createProject(name string) {
	project := db.Project{Name: name}
	db.DB.Create(&project)
	fmt.Printf("✅ Project added: %s\n", name)
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
