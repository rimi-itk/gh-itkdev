package changelog

import (
	"fmt"
	"os"
	"regexp"
)

func DetectPullRequestEntryFormat(changelog string) (string, error) {
	if _, err := os.Stat(changelog); err == nil {
		b, err := os.ReadFile(changelog)
		if err != nil {
			return "", fmt.Errorf("cannot read file %s", changelog)
		}
		changelog = string(b)
	}

	templates := getPullRequestEntryTemplates()

	for _, template := range templates {
		if template.pattern.MatchString(changelog) {
			return template.template, nil
		}
	}

	if len(templates) < 1 {
		return "", fmt.Errorf("cannot detect pull request entry format")
	}

	return templates[len(templates)-1].template, nil
}

func getPullRequestEntryTemplates() []struct {
	pattern  *regexp.Regexp
	template string
} {
	return []struct {
		pattern  *regexp.Regexp
		template string
	}{
		{
			regexp.MustCompile("(?m)^- \\[PR-[0-9]+\\]\\([^)]+\\)\n  .+$"),
			`- [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`,
		},

		{
			regexp.MustCompile("(?m)^- \\[#[0-9]+\\]\\([^)]+\\)\n  .+$"),
			`- [#{{ .Number }}]({{ .Url }})
  {{ .Title }}`,
		},

		// The default – must come last.
		{
			regexp.MustCompile("."),
			`* [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`,
		},
	}
}
