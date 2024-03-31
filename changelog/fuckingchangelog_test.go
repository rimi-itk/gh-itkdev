package changelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuckingChangelog(t *testing.T) {
	defaultItemTemplate := `* [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`
	testCases := []struct {
		changelog    string
		pullRequest  pullRequest
		itemTemplate string
		expected     string
	}{
		{
			`## [Unreleased]
`,
			pullRequest{
				Number: 87,
				Title:  "Test",
				Url:    "https://example.com/pr/87",
			},
			defaultItemTemplate,
			`## [Unreleased]

* [PR-87](https://example.com/pr/87)
  Test
`,
		},

		{
			`## [Unreleased]

[Unreleased]: https://example.com/
`,
			pullRequest{
				Number: 87,
				Title:  "Test",
				Url:    "https://example.com/pr/87",
			},
			defaultItemTemplate,
			`## [Unreleased]

* [PR-87](https://example.com/pr/87)
  Test

[Unreleased]: https://example.com/
`,
		},

		{
			`## [Unreleased]

* [PR-42](https://example.com/pr/42)
  Added the meaning
`,
			pullRequest{
				Number: 87,
				Title:  "Test",
				Url:    "https://example.com/pr/87",
			},
			defaultItemTemplate,
			`## [Unreleased]

* [PR-87](https://example.com/pr/87)
  Test
* [PR-42](https://example.com/pr/42)
  Added the meaning
`,
		},

		{
			`## [Unreleased]

- [#42](https://example.com/pr/42): Added the meaning
`,
			pullRequest{
				Number: 87,
				Title:  "Test",
				Url:    "https://example.com/pr/87",
			},
			`- [#{{ .Number }}]({{ .Url }}): {{ .Title }}`,
			`## [Unreleased]

- [#87](https://example.com/pr/87): Test
- [#42](https://example.com/pr/42): Added the meaning
`,
		},
	}

	for _, testCase := range testCases {
		actual, _ := addPullRequest(testCase.changelog, testCase.pullRequest, testCase.itemTemplate)

		assert.Equal(t, testCase.expected, actual)
	}
}
