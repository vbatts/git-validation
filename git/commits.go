package git

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Commits returns a set of commits.
// If commitrange is a git still range 12345...54321, then it will be isolated set of commits.
// If commitrange is a single commit, all ancestor commits up through the hash provided.
func Commits(commitrange string) ([]CommitEntry, error) {
	cmdArgs := []string{"git", "log", prettyFormat + formatCommit, commitrange}
	if debug() {
		logrus.Infof("[git] cmd: %q", strings.Join(cmdArgs, " "))
	}
	output, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).Output()
	if err != nil {
		return nil, err
	}
	commitHashes := strings.Split(strings.TrimSpace(string(output)), "\n")
	commits := make([]CommitEntry, len(commitHashes))
	for i, commitHash := range commitHashes {
		c, err := LogCommit(commitHash)
		if err != nil {
			return commits, err
		}
		commits[i] = *c
	}
	return commits, nil
}

// CommitEntry represents a single commit's information from `git`
type CommitEntry map[string]string

var (
	prettyFormat         = `--pretty=format:`
	formatSubject        = `%s`
	formatBody           = `%b`
	formatCommit         = `%H`
	formatAuthorName     = `%aN`
	formatAuthorEmail    = `%aE`
	formatCommitterName  = `%cN`
	formatCommitterEmail = `%cE`
	formatSigner         = `%GS`
	formatCommitNotes    = `%N`
	formatMap            = `{"commit": "%H", "abbreviated_commit": "%h", "tree": "%T", "abbreviated_tree": "%t", "parent": "%P", "abbreviated_parent": "%p", "refs": "%D", "encoding": "%e", "sanitized_subject_line": "%f", "verification_flag": "%G?", "signer_key": "%GK", "author_date": "%aD" , "committer_date": "%cD" }`
)

// LogCommit assembles the full information on a commit from its commit hash
func LogCommit(commit string) (*CommitEntry, error) {
	buf := bytes.NewBuffer([]byte{})
	cmdArgs := []string{"git", "log", "-1", prettyFormat + formatMap, commit}
	if debug() {
		logrus.Infof("[git] cmd: %q", strings.Join(cmdArgs, " "))
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(strings.Join(cmd.Args, " "))
		return nil, err
	}
	c := CommitEntry{}
	output := buf.Bytes()
	if err := json.Unmarshal(output, &c); err != nil {
		fmt.Println(string(output))
		return nil, err
	}

	// any user provided fields can't be sanitized for the mock-json marshal above
	for k, v := range map[string]string{
		"subject":         formatSubject,
		"body":            formatBody,
		"author_name":     formatAuthorName,
		"author_email":    formatAuthorEmail,
		"committer_name":  formatCommitterName,
		"committer_email": formatCommitterEmail,
		"commit_notes":    formatCommitNotes,
		"signer":          formatSigner,
	} {
		output, err := exec.Command("git", "log", "-1", prettyFormat+v, commit).Output()
		if err != nil {
			return nil, err
		}
		c[k] = strings.TrimSpace(string(output))
	}

	return &c, nil
}

func debug() bool {
	return len(os.Getenv("DEBUG")) > 0
}

// FetchHeadCommit returns the hash of FETCH_HEAD
func FetchHeadCommit() (string, error) {
	cmdArgs := []string{"git", "rev-parse", "--verify", "FETCH_HEAD"}
	if debug() {
		logrus.Infof("[git] cmd: %q", strings.Join(cmdArgs, " "))
	}
	output, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// HeadCommit returns the hash of HEAD
func HeadCommit() (string, error) {
	cmdArgs := []string{"git", "rev-parse", "--verify", "HEAD"}
	if debug() {
		logrus.Infof("[git] cmd: %q", strings.Join(cmdArgs, " "))
	}
	output, err := exec.Command(cmdArgs[0], cmdArgs[1:]...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
