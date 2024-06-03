/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"gopkg.in/yaml.v2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show active config",
	Run: func(cmd *cobra.Command, args []string) {
		yaml, err := yaml.Marshal(viper.AllSettings())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(yaml))
	},
}
