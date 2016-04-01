package git

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestCommitEntry(t *testing.T) {
	c, err := HeadCommit()
	if err != nil {
		t.Fatal(err)
	}
	cr, err := Commits(c)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cr {
		for _, cV := range FieldNames {
			found := false
			for k := range c {
				if k == cV {
					found = true
				}
			}
			if !found {
				t.Errorf("failed to find field names: %q", c)
			}
		}
	}
}

func TestMarshal(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/commits.json")
	if err != nil {
		t.Fatal(err)
	}
	cr := []CommitEntry{}
	if err := json.Unmarshal(buf, &cr); err != nil {
		t.Error(err)
	}
	for _, c := range cr {
		for _, cV := range FieldNames {
			found := false
			for k := range c {
				if k == cV {
					found = true
				}
			}
			if !found {
				t.Errorf("failed to find field names: %q", c)
			}
		}
	}
}
