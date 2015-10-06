# git-validation

A way to do per git commit validation

## install

```bash
vbatts@valse ~ (master) $ go get -u github.com/vbatts/git-validation
```

## usage

```bash
vbatts@valse ~/src/vb/git-validation (master) $ git-validation 
 * d614ccf "*: run tests in a runner" ... PASS
 * b9413c6 "shortsubject: add a subject length check" ... PASS
 * 5e74abd "*: comments and golint" ... PASS
 * 07a982f "git: add verbose output of the commands run" ... PASS
 * 03bda4b "main: add filtering of rules to run" ... PASS
 * c10ba9c "Initial commit" ... PASS
vbatts@valse ~/src/vb/git-validation (master) $ git-validation -v
 * d614ccf "*: run tests in a runner" ... PASS
  - PASS - has a valid DCO
  - PASS - commit subject is 72 characters or less! *yay*
 * b9413c6 "shortsubject: add a subject length check" ... PASS
  - PASS - has a valid DCO
  - PASS - commit subject is 72 characters or less! *yay*
 * 5e74abd "*: comments and golint" ... PASS
  - PASS - has a valid DCO
  - PASS - commit subject is 72 characters or less! *yay*
 * 07a982f "git: add verbose output of the commands run" ... PASS
  - PASS - has a valid DCO
  - PASS - commit subject is 72 characters or less! *yay*
 * 03bda4b "main: add filtering of rules to run" ... PASS
  - PASS - has a valid DCO
  - PASS - commit subject is 72 characters or less! *yay*
 * c10ba9c "Initial commit" ... PASS
  - PASS - has a valid DCO
  - PASS - commit subject is 72 characters or less! *yay*
```

Here's a failure:
```bash
vbatts@valse ~/src/vb/git-validation (master) $ git-validation 
 * 49f51a8 "README: adding install and usage" ... FAIL
  - FAIL - does not have a valid DCO
 * d614ccf "*: run tests in a runner" ... PASS
 * b9413c6 "shortsubject: add a subject length check" ... PASS
 * 5e74abd "*: comments and golint" ... PASS
 * 07a982f "git: add verbose output of the commands run" ... PASS
 * 03bda4b "main: add filtering of rules to run" ... PASS
 * c10ba9c "Initial commit" ... PASS
1 issues to fix
vbatts@valse ~/src/vb/git-validation (master) $ echo $?
1
```
