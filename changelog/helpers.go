package changelog

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/cli/go-gh"
)

func isChanged(name string) bool {
	cmd := exec.Command("git", "diff", "--exit-code", name)

	return cmd.Run() != nil
}

func gitDiff(files []string) {
	cmd := exec.Command("git", append([]string{"diff"}, files...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func gitCommit(files []string, message string) {
	cmd := exec.Command("git", append([]string{"add"}, files...)...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "commit", "--message", message)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "log", "--patch-with-stat", "--max-count", "1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
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
