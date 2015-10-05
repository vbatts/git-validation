package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vbatts/git-validation/git"
	_ "github.com/vbatts/git-validation/rules/dco"
	"github.com/vbatts/git-validation/validate"
)

var (
	flCommitRange = flag.String("range", "", "use this commit range instead")
	flListRules   = flag.Bool("list-rules", false, "list the rules registered")
	flRun         = flag.String("run", "", "comma delimited list of rules to run. Defaults to all.")
	flVerbose     = flag.Bool("v", false, "verbose")
)

func main() {
	flag.Parse()

	if *flListRules {
		for _, r := range validate.RegisteredRules {
			fmt.Printf("%q -- %s\n", r.Name, r.Description)
		}
		return
	}

	rules := validate.RegisteredRules
	if *flRun != "" {
		rules = validate.FilterRules(rules, validate.SanitizeFilters(*flRun))
	}

	var commitrange string
	if *flCommitRange != "" {
		commitrange = *flCommitRange
	} else {
		var err error
		commitrange, err = git.FetchHeadCommit()
		if err != nil {
			log.Fatal(err)
		}
	}

	c, err := git.Commits(commitrange)
	if err != nil {
		log.Fatal(err)
	}

	results := validate.Results{}
	for _, commit := range c {
		fmt.Printf(" * %s %s ... ", commit["abbreviated_commit"], commit["subject"])
		vr := validate.Commit(commit, rules)
		results = append(results, vr...)
		if _, fail := vr.PassFail(); fail == 0 {
			fmt.Println("PASS")
			if *flVerbose {
				for _, r := range vr {
					if r.Pass {
						fmt.Printf("  - %s\n", r.Msg)
					}
				}
			}
		} else {
			fmt.Println("FAIL")
			// default, only print out failed validations
			for _, r := range vr {
				if !r.Pass {
					fmt.Printf("  - %s\n", r.Msg)
				}
			}
		}
	}
	_, fail := results.PassFail()
	if fail > 0 {
		fmt.Printf("%d issues to fix\n", fail)
		os.Exit(1)
	}
}
