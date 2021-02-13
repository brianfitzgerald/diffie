package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		println("not enough args!")
		return
	}
	aString := args[0]
	bString := args[1]

	dmp := diffmatchpatch.New()

	diff := dmp.DiffMain(aString, bString, false)

	out := ""
	out += findSubstr(aString, diff)
	out = findSubstr(bString, diff)

	println(out)

}

func findSubstr(str string, diffs []diffmatchpatch.Diff) string {
	diffIdx := []int{}
	outStr := ""
	for i, d := range diffs {
		if d.Type == diffmatchpatch.DiffInsert {
			if i < 1 {
				continue
			}
			idx := strings.Index(str, d.Text)
			diffIdx = append(diffIdx, idx)
			if d.Type == diffmatchpatch.DiffDelete && diffs[i+1].Type == diffmatchpatch.DiffInsert {
				diffText := fmt.Sprintf("%s|%s at char %d\n", d.Text, diffs[i+1].Text, idx)
				outStr += diffText
			}
		}
	}
	return outStr
}
