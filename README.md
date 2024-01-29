TagDesc
=======

An emurator of `git describe --tags` on the DVCS [Jujutsu](https://martinvonz.github.io/jj/) ( [github.com/martinvonz/jj](https://github.com/martinvonz/jj) )

It reads and parses the output of `jj log --no-graph -r "latest(tags())::"`

PowerShell version
------------------

Copy tagdesc.ps1 to your work directory and write On Makefile(GNU):

```
VERSION:=$(shell powershell -noprofile -ex unrestricted -file tagdesc.ps1)
```

Go version
----------

Copy tagdesc.go to your work directory and write the following line on Makefile(GNU):

```
VERSION:=$(shell go run tagdesc.go)
```

`tagdesc.go` is ignored by `go build` because `tagdesc.go` has the header `//go:build run` .
