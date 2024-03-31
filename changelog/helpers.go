package changelog

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/cli/go-gh"
)

func isChanged(name string) bool {
	cmd := exec.Command("git", "diff", "--exit-code", name)

	return cmd.Run() != nil
}

func showDiff(name string) {
	cmd := exec.Command("git", "diff", name)
	output, _ := cmd.CombinedOutput()
	fmt.Println(string(output))
}

func getRepositoryUrl() (string, error) {
	repository, err := gh.CurrentRepository()
	if err != nil {
		return "", fmt.Errorf("error getting repository: %s", err)
	}

	url := fmt.Sprintf("https://%s/%s/%s", repository.Host(), repository.Owner(), repository.Name())

	return url, nil
}

type pullRequest struct {
	Number int
	Title  string
	Url    string
}

func getPullRequest() (pullRequest, error) {
	var pr pullRequest

	content, _, err := gh.Exec("pr", "view", "--json", "number,title,url")
	if err != nil {
		return pr, err
	}

	err = json.Unmarshal(content.Bytes(), &pr)

	return pr, err
}
