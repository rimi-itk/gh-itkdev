package cmd

import (
	"fmt"

	"github.com/rimi-itk/gh-itkdev/changelog"
	"github.com/spf13/cobra"
)

// changelogCmd represents the changelog command
var (
	create        bool
	changelogName string = "CHANGELOG.md"

	changelogCmd = &cobra.Command{
		Use:   "changelog",
		Short: "Update changelog",
		Run: func(cmd *cobra.Command, args []string) {
			if create {
				changelog.Create(changelogName)
			} else {
				cmd.Usage()
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(changelogCmd)

	changelogCmd.Flags().BoolVarP(&create, "create", "", false, fmt.Sprintf("create a changelog (%q) if it does not exist", changelogName))
	changelogCmd.Flags().StringVarP(&changelogName, "changelog", "", changelogName, "changelog name")
}
