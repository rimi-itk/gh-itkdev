package changelog

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	changelogName = "CHANGELOG.md"
	ChangelogCmd  = &cobra.Command{
		Use:   "changelog",
		Short: "Update changelog",
	}
)

func init() {
	ChangelogCmd.PersistentFlags().StringVarP(&changelogName, "changelog", "", changelogName, "changelog name")
	viper.BindPFlag("changelog.changelog", ChangelogCmd.PersistentFlags().Lookup("changelog"))
}
