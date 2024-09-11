package cmd

import (
	"fmt"
	"log"

	"github.com/rimi-itk/gh-itkdev/cmd/changelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-itkdev",
	Short: "GitHub CLI helper for ITK Development",
	Annotations: map[string]string{
		// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#creating-a-plugin
		cobra.CommandDisplayNameAnnotation: "gh itkdev",
	},
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
			log.Fatalf("Unable to read config file %s: %s\n", viper.ConfigFileUsed(), err)
		}
	} else {
		fmt.Printf("Using config file %s\n", viper.ConfigFileUsed())
	}
}
