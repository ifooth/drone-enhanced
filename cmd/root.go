package main

import (
	"fmt"

	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "drone-enhanced",
		Short: "drone ci enhanced server",
		Long:  `A drone ci pipeline as code enhanced server`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(ServerCmd())
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show application version",
		Long:  `All software has versions. This is drone-enhanced's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Print("drone-enhanced"))
		},
	}
	return cmd
}
