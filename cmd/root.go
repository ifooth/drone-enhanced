package main

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "drone_enhanced",
		Short: "drone ci enhanced server",
		Long:  `A drone ci full enhanced server`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.drone_enhanced.yml)")

	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(ServerCmd())
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of drone_enhanced",
		Long:  `All software has versions. This is drone_enhanced's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Print("drone_enhanced"))
		},
	}
	return cmd
}
