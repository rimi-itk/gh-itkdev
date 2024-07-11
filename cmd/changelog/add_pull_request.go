/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package changelog

import (
	"log"

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
			log.Printf("\n\nadd-pull-request: viper: %q\nentryTemplate: %q\n\n", viper.AllSettings(), entryTemplate)
			log.Println(viper.GetString("changelog.add-pull-request.entry-template"))
			changelog.AddPullRequest(changelogName, entryTemplate)
		},
	}
)

func init() {
	log.Println("add-pull-request")

	viperPrefix := "changelog.add-pull-request."
	flagName := "entry-template"
	addPullRequestCmd.Flags().StringVarP(&entryTemplate, flagName, "", entryTemplate, "pull request entry template")
	viper.BindPFlag(viperPrefix+flagName, addPullRequestCmd.Flags().Lookup(flagName))

	ChangelogCmd.AddCommand(addPullRequestCmd)
}
