/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cs/cmd/jerry"
	"cs/cmd/spike"
	"cs/cmd/tom"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(tom.TomCmd)
	rootCmd.AddCommand(jerry.JerryCmd)
	rootCmd.AddCommand(spike.SpikeCmd)
}
