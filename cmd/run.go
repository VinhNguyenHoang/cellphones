/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"cs/internal/jerry"
	"cs/internal/tom"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  `Run command will bootstrap a HTTP service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		svName, err := cmd.Flags().GetString("service")
		if err != nil {
			return err
		}
		fmt.Printf("run called with %s...\n", svName)
		switch svName {
		case "tom":
			sv, err := tom.NewServer()
			if err != nil {
				return err
			}

			return sv.Start()

		case "jerry":
			sv := jerry.Server{Address: "0.0.0.0:30001"}
			err = sv.Start()
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().StringP("service", "s", "", "Name of the service to be run")
}
