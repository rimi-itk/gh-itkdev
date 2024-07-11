/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package changelog

import (
	"github.com/rimi-itk/gh-itkdev/changelog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var (
	base   = "develop"
	commit = false

	addReleaseCmd = &cobra.Command{
		Use:   "add-release",
		Short: "Create a release branch and update the changelog",
		Long: `Create a release branch named release/«release» and update the changelog as per https://keepachangelog.com/en/1.1.0/

Examples:

gh itkdev changelog add-release 0.1.0
gh itkdev changelog add-release v0.1.0
`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			release := args[0]

			changelog.AddRelease(release, base, changelogName, commit)
		},
	}
)

func init() {
	viperPrefix := "changelog.add-release."
	flagName := "base"
	addReleaseCmd.Flags().StringVarP(&base, flagName, "", base, "base branch for release")
	viper.BindPFlag(viperPrefix+flagName, addReleaseCmd.Flags().Lookup(flagName))

	flagName = "commit"
	addReleaseCmd.Flags().BoolVarP(&commit, flagName, "", commit, "commit changes")
	viper.BindPFlag(viperPrefix+flagName, addReleaseCmd.Flags().Lookup(flagName))

	ChangelogCmd.AddCommand(addReleaseCmd)
}
