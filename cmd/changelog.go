package cmd

import (
	"fmt"

	"github.com/rimi-itk/gh-itkdev/changelog"
	"github.com/spf13/cobra"
)

// changelogCmd represents the changelog command
var (
	create bool

	fuckingChangelog        bool
	pullRequestItemTemplate string = `* [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`

	release      string
	baseBranches      = []string{"develop", "main", "master"}
	commit       bool = false

	changelogName string = "CHANGELOG.md"

	changelogCmd = &cobra.Command{
		Use:   "changelog",
		Short: "Update changelog",
		Run: func(cmd *cobra.Command, args []string) {
			if create {
				changelog.Create(changelogName)
			} else if fuckingChangelog {
				changelog.FuckingChangelog(changelogName, pullRequestItemTemplate)
			} else if release != "" {
				changelog.Release(release, baseBranches, changelogName, commit)
			} else {
				cmd.Usage()
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(changelogCmd)

	changelogCmd.Flags().BoolVarP(&create, "create", "", false, fmt.Sprintf("create a changelog (%q) if it does not exist", changelogName))

	changelogCmd.Flags().BoolVarP(&fuckingChangelog, "fucking-changelog", "", false, "add missing pull request entry to changelog")
	changelogCmd.Flags().StringVarP(&pullRequestItemTemplate, "item-template", "", pullRequestItemTemplate, "pull request item template")

	changelogCmd.Flags().StringVarP(&release, "release", "", "", "create a release branch with updated changelog")
	changelogCmd.Flags().StringSliceVarP(&baseBranches, "base", "", baseBranches, "base branch for release")
	changelogCmd.Flags().BoolVarP(&commit, "commit", "", commit, "commit changes")

	changelogCmd.Flags().StringVarP(&changelogName, "changelog", "", changelogName, "changelog name")
}