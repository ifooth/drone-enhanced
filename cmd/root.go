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
		Use:   "botprime",
		Short: "A spider utils for funny",
		Long:  `BotPrime is a spider utils for funny`,
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
		Short: "Print the version number of BotPrime",
		Long:  `All software has versions. This is BotPrime's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Print("botprime"))
		},
	}
	return cmd
}
