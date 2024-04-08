package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portops",
	Short: "po",
	Long:  "A CLI tool to manage processes and ports",

	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(killCmd)
	rootCmd.AddCommand(releaseCmd)
	rootCmd.AddCommand(listCmd)
}
