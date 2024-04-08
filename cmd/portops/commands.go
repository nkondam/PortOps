package cmd

import (
	"fmt"
	"os"
	"portops/pkg/action"
	"strconv"

	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:     "kill [pid]",
	Aliases: []string{"k [pid]"},
	Short:   "Kill a process by its PID",

	Run: func(cmd *cobra.Command, args []string) {
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid PID: ", err)
			os.Exit(1)
		}

		err = action.KillProcess(pid)
		if err != nil {
			fmt.Println("Error killing process: ", err)
			os.Exit(1)
		}
	},
}

var releaseCmd = &cobra.Command{
	Use:     "release [port]",
	Aliases: []string{"r [pid]"},
	Short:   "Releases a process by its Port",

	Run: func(cmd *cobra.Command, args []string) {
		port, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid PID: ", err)
			os.Exit(1)
		}

		err = action.ReleasePort(port)
		if err != nil {
			fmt.Println("Error Reasleasing Port: ", err)
			os.Exit(1)
		}
	},
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all the running process",

	Run: func(cmd *cobra.Command, args []string) {
		err := action.ListProcesses()
		if err != nil {
			fmt.Println("Error Reasleasing Port: ", err)
			os.Exit(1)
		}
	},
}
