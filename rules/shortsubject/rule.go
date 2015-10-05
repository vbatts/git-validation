package shortsubject

import (
	"github.com/vbatts/git-validation/git"
	"github.com/vbatts/git-validation/validate"
)

var (
	// DcoRule is the rule being registered
	ShortSubjectRule = validate.Rule{
		Name:        "short-subject",
		Description: "commit subjects are strictly less than 90 (github ellipsis length)",
		Run:         ValidateShortSubject,
	}
)

func init() {
	validate.RegisterRule(ShortSubjectRule)
}

func ValidateShortSubject(c git.CommitEntry) (vr validate.Result) {
	if len(c["subject"]) >= 90 {
		vr.Pass = false
		vr.Msg = "commit subject exceeds 90 characters"
		return
	}
	vr.Pass = true
	vr.Msg = "commit subject is not more than 90 characters"
	return
}
