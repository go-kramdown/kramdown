// Copyright (c) the go-kramdown/kramdown authors
//
// SPDX-License-Identifier: BSD-3-Clause

package kramdown

import "testing"

// TestAbbreviationDegenerateKeys is the regression for two fuzz findings: a
// whitespace-only abbreviation key ("*[ ]:") recursed forever, and a key that is
// not valid UTF-8 made regexp.MustCompile panic.
func TestAbbreviationDegenerateKeys(t *testing.T) {
	for _, src := range []string{"*[ ]:\r&", "*[ ]: x\n\ntext here\n", "*[\xfc]: x\n\na\xfc b\n"} {
		_ = ToHTML(src, nil)
	}
}

// FuzzToHTMLBlocks checks that the block parser (lists, tables, task lists)
// never panics or hangs on arbitrary input, in both the kramdown and GFM
// dialects.
func FuzzToHTMLBlocks(f *testing.F) {
	for _, s := range []string{
		"- a\n- b\n\n| x | y |\n| - | - |\n| 1 | 2 |\n",
		"- [x] done\n- [ ] todo\n",
		"1. a\n\n^\n| x |\n| - |\n",
		"- a\n  - b\n\n\n| x | y |\n",
	} {
		f.Add(s)
	}
	gfm := gfmOpts()
	f.Fuzz(func(t *testing.T, src string) {
		_ = ToHTML(src, nil)
		_ = ToHTML(src, gfm)
	})
}
