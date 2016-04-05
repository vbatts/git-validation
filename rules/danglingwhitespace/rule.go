package danglingwhitespace

import (
	"bytes"
	"fmt"

	"github.com/vbatts/git-validation/git"
	"github.com/vbatts/git-validation/validate"
)

var (
	// DanglingWhitespace is the rule for checking the presence of dangling
	// whitespaces on line endings.
	DanglingWhitespace = validate.Rule{
		Name:        "dangling-whitespace",
		Description: "checking the presence of dangling whitespaces on line endings",
		Run:         ValidateDanglingWhitespace,
	}
)

func init() {
	validate.RegisterRule(DanglingWhitespace)
}

func ValidateDanglingWhitespace(c git.CommitEntry) (vr validate.Result) {
	diff, err := git.Show(c["commit"])
	if err != nil {
		return validate.Result{Pass: false, Msg: err.Error(), CommitEntry: c}
	}

	vr.CommitEntry = c
	vr.Pass = true
	for _, line := range bytes.Split(diff, newLine) {
		if !bytes.HasPrefix(line, diffAddLine) || bytes.Equal(line, diffAddLine) {
			continue
		}
		if len(bytes.TrimSpace(line)) != len(line) {
			vr.Pass = false
			vr.Msg = fmt.Sprintf("line %q has trailiing spaces", string(line))
		}
	}
	vr.Msg = "all added diff lines do not have trailing spaces"
	return
}

var (
	newLine     = []byte("\n")
	diffAddLine = []byte("+ ")
)
