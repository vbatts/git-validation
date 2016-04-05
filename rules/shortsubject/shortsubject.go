package shortsubject

import (
	"github.com/vbatts/git-validation/git"
	"github.com/vbatts/git-validation/validate"
)

var (
	// ShortSubjectRule is the rule being registered
	ShortSubjectRule = validate.Rule{
		Name:        "short-subject",
		Description: "commit subjects are strictly less than 90 (github ellipsis length)",
		Run:         ValidateShortSubject,
	}
)

func init() {
	validate.RegisterRule(ShortSubjectRule)
}

// ValidateShortSubject checks that the commit's subject is strictly less than
// 90 characters (preferably not more than 72 chars).
func ValidateShortSubject(c git.CommitEntry) (vr validate.Result) {
	if len(c["subject"]) >= 90 {
		vr.Pass = false
		vr.Msg = "commit subject exceeds 90 characters"
		return
	}
	vr.Pass = true
	if len(c["subject"]) > 72 {
		vr.Msg = "commit subject is under 90 characters, but is still more than 72 chars"
	} else {
		vr.Msg = "commit subject is 72 characters or less! *yay*"
	}
	return
}
