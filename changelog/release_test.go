package changelog

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRelease(t *testing.T) {
	testCases := []struct {
		changelog string
		release   string
		expected  string
	}{
		{
			`## [Unreleased]

[Unreleased]: https://example.com/
`,
			"v0.0.0",
			`## [Unreleased]

## [v0.0.0] - %TODAY%

[Unreleased]: https://example.com/compare/v0.0.0...HEAD
[v0.0.0]: https://example.com/releases/tag/v0.0.0
`,
		},

		{
			`## [Unreleased]

* [PR-42](https://example.com/pr/42)
  Added the meaning

## [v0.0.0] - 2001-01-01

[Unreleased]: https://example.com/compare/v0.0.0...HEAD
[v0.0.0]: https://example.com/releases/tag/v0.0.0
`,
			"v0.1.0",
			`## [Unreleased]

## [v0.1.0] - %TODAY%

* [PR-42](https://example.com/pr/42)
  Added the meaning

## [v0.0.0] - 2001-01-01

[Unreleased]: https://example.com/compare/v0.1.0...HEAD
[v0.1.0]: https://example.com/compare/v0.0.0...v0.1.0
[v0.0.0]: https://example.com/releases/tag/v0.0.0
`,
		},

		{
			`## [Unreleased]

* [PR-42](https://example.com/pr/42)
  Added the meaning

## [v0.1.0] - 2002-01-01

## [v0.0.0] - 2001-01-01

[Unreleased]: https://example.com/compare/v0.1.0...HEAD
[v0.1.0]: https://example.com/compare/v0.0.0...v0.1.0
[v0.0.0]: https://example.com/releases/tag/v0.0.0
`,
			"v0.1.1",
			`## [Unreleased]

## [v0.1.1] - %TODAY%

* [PR-42](https://example.com/pr/42)
  Added the meaning

## [v0.1.0] - 2002-01-01

## [v0.0.0] - 2001-01-01

[Unreleased]: https://example.com/compare/v0.1.1...HEAD
[v0.1.1]: https://example.com/compare/v0.1.0...v0.1.1
[v0.1.0]: https://example.com/compare/v0.0.0...v0.1.0
[v0.0.0]: https://example.com/releases/tag/v0.0.0
`,
		},
	}

	for _, testCase := range testCases {
		file, _ := os.CreateTemp("", "CHANGELOG.md")
		defer os.Remove(file.Name())

		name := file.Name()
		file.Write([]byte(testCase.changelog))

		assert.FileExists(t, name)

		updateChangelog(name, testCase.release)

		bytes, _ := os.ReadFile(name)
		actual := string(bytes)
		assert.Equal(t, strings.ReplaceAll(testCase.expected, "%TODAY%", time.Now().Format("2006-01-02")), actual)
	}
}
