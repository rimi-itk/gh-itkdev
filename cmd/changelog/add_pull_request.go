/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package changelog

import (
	"github.com/rimi-itk/gh-itkdev/changelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	entryTemplate = `* [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`
	addPullRequestCmd = &cobra.Command{
		Use:   "add-pull-request",
		Short: "Add missing pull request entry to changelog",
		Run: func(cmd *cobra.Command, args []string) {
			changelog.AddPullRequest(changelogName, entryTemplate)
		},
	}
)

func init() {
	viperPrefix := "changelog.add-pull-request."
	flagName := "entry-template"
	addPullRequestCmd.Flags().StringVarP(&entryTemplate, flagName, "", entryTemplate, "pull request entry template")
	viper.BindPFlag(viperPrefix+flagName, createCmd.Flags().Lookup(flagName))

	ChangelogCmd.AddCommand(addPullRequestCmd)
}
