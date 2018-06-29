package main

import (
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

var patcher = []byte(`package runtime

func GoID() int64 {
    return getg().goid
}
`)

func main() {
	pkg, _ := build.Default.Import("runtime", "", build.FindOnly)
	ioutil.WriteFile(path.Join(pkg.Dir, "proc_id.go"), patcher, os.ModePerm)
	exec.Command("go", "install", "runtime").CombinedOutput()
}
