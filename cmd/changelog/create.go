/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package changelog

import (
	"github.com/rimi-itk/gh-itkdev/changelog"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a changelog if it does not exist",
	Long:  `Create a changelog if it does not exist. The changelog will used the format defined by https://keepachangelog.com/en/1.1.0/.`,
	Run: func(cmd *cobra.Command, args []string) {
		changelog.Create(changelogName)
	},
}

func init() {
	ChangelogCmd.AddCommand(createCmd)
}
