package changelog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuckingChangelog(t *testing.T) {
	defaultItemTemplate := `* [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`
	testCases := []struct {
		changelog    string
		pullRequest  PullRequest
		itemTemplate string
		expected     string
	}{
		{
			`## [Unreleased]
`,
			PullRequest{
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
			PullRequest{
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
			PullRequest{
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
			PullRequest{
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
		file, _ := os.CreateTemp("", "CHANGELOG.md")
		defer os.Remove(file.Name())

		name := file.Name()
		file.Write([]byte(testCase.changelog))

		assert.FileExists(t, name)

		addPullRequest(name, testCase.pullRequest, testCase.itemTemplate)

		bytes, _ := os.ReadFile(name)
		actual := string(bytes)
		assert.Equal(t, testCase.expected, actual)
	}
}
