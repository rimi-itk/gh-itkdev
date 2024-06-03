package cmd

import (
	"fmt"

	"github.com/rimi-itk/gh-itkdev/cmd/changelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-itkdev",
	Short: "GitHub CLI helper for ITK Development",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(changelog.ChangelogCmd)
	rootCmd.AddCommand(configCmd)
}

func initConfig() {
	viper.SetConfigName(".gh-itkdev")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Unable to read config: ", err)
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
