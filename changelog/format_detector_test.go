package changelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullRequestTemplateDetector(t *testing.T) {
	testCases := []struct {
		changelog string
		expected  string
	}{
		{
			`- [PR-87](https://example.com/pr/87)
  Test
`,
			`- [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`,
		},

		{
			`- [#87](https://example.com/pr/87)
  Test
`,
			`- [#{{ .Number }}]({{ .Url }})
  {{ .Title }}`,
		},

		{
			`* [PR-87](https://example.com/pr/87)
  Test
`,
			`* [PR-{{ .Number }}]({{ .Url }})
  {{ .Title }}`,
		},
	}

	for _, testCase := range testCases {
		actual, _ := DetectPullRequestEntryFormat(testCase.changelog)

		assert.Equal(t, testCase.expected, actual)
	}
}
