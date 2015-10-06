package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/vbatts/git-validation/rules/dco"
	_ "github.com/vbatts/git-validation/rules/shortsubject"
	"github.com/vbatts/git-validation/validate"
)

var (
	flCommitRange = flag.String("range", "", "use this commit range instead")
	flListRules   = flag.Bool("list-rules", false, "list the rules registered")
	flRun         = flag.String("run", "", "comma delimited list of rules to run. Defaults to all.")
	flVerbose     = flag.Bool("v", false, "verbose")
	flDebug       = flag.Bool("D", false, "debug output")
	flDir         = flag.String("d", ".", "git directory to validate from")
)

func main() {
	flag.Parse()

	if *flDebug {
		os.Setenv("DEBUG", "1")
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

	runner, err := validate.NewRunner(*flDir, rules, *flCommitRange, *flVerbose)
	if err != nil {
		log.Fatal(err)
	}

	if err := runner.Run(); err != nil {
		log.Fatal(err)
	}
	_, fail := runner.Results.PassFail()
	if fail > 0 {
		fmt.Printf("%d issues to fix\n", fail)
		os.Exit(1)
	}

}
