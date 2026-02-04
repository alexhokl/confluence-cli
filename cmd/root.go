package cmd

import (
	"fmt"

	"github.com/alexhokl/helper/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var noColor bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "confluence-cli",
	Short:             "A CLI application to talk to Confluence APIs",
	SilenceUsage:      true,
	PersistentPreRunE: validateConfiguration,
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.confluence-cli.yaml)")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "Disable color output")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	cli.ConfigureViper(cfgFile, "confluence-cli", false, "")
}

func validateConfiguration(_ *cobra.Command, _ []string) error {
	if viper.GetString("email") == "" {
		return fmt.Errorf("email is not configured")
	}
	if viper.GetString("api_key") == "" {
		return fmt.Errorf("api_key is not configured")
	}
	if viper.GetString("organization") == "" {
		return fmt.Errorf("organization is not configured")
	}

	return nil
}
