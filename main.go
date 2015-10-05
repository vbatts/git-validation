package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vbatts/git-validation/git"
	_ "github.com/vbatts/git-validation/rules/dco"
	_ "github.com/vbatts/git-validation/rules/shortsubject"
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

	if *flVerbose {
		git.Verbose = true
	}

	if *flListRules {
		for _, r := range validate.RegisteredRules {
			fmt.Printf("%q -- %s\n", r.Name, r.Description)
		}
		return
	}

	// reduce the set being run
	rules := validate.RegisteredRules
	if *flRun != "" {
		rules = validate.FilterRules(rules, validate.SanitizeFilters(*flRun))
	}

	// Guess the commits we're working with
	var commitrange string
	if *flCommitRange != "" {
		commitrange = *flCommitRange
	} else {
		var err error
		commitrange, err = git.FetchHeadCommit()
		if err != nil {
			commitrange, err = git.HeadCommit()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// collect the entries
	c, err := git.Commits(commitrange)
	if err != nil {
		log.Fatal(err)
	}

	// run them and show results
	results := validate.Results{}
	for _, commit := range c {
		fmt.Printf(" * %s %s ... ", commit["abbreviated_commit"], commit["subject"])
		vr := validate.Commit(commit, rules)
		results = append(results, vr...)
		if _, fail := vr.PassFail(); fail == 0 {
			fmt.Println("PASS")
		} else {
			fmt.Println("FAIL")
		}
		for _, r := range vr {
			if *flVerbose {
				if r.Pass {
					fmt.Printf("ok %s\n", r.Msg)
				} else {
					fmt.Printf("not ok %s\n", r.Msg)
				}
			} else if !r.Pass {
				fmt.Printf("not ok %s\n", r.Msg)
			}
		}
	}
	_, fail := results.PassFail()
	if fail > 0 {
		fmt.Printf("%d issues to fix\n", fail)
		os.Exit(1)
	}
}
