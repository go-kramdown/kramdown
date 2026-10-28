// Copyright (c) the go-kramdown/kramdown authors
//
// SPDX-License-Identifier: BSD-3-Clause

package kramdown

import "testing"

// TestTableAfterBlock asserts byte-exact parity with the reference kramdown gem
// (kramdown 2.5.2, default options): a pipe table that follows a list (whose
// trailing blank line is re-emitted as an ElBlank) sits at a block boundary and
// is parsed as a table, per kramdown's after_block_boundary?. The expected
// outputs were produced by running the gem itself.
func TestTableAfterBlock(t *testing.T) {
	cases := []struct{ name, src, want string }{
		{"bq_then_table", "> q\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<blockquote>\n  <p>q</p>\n</blockquote>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"code_then_table", "    code\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<pre><code>code\n</code></pre>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"dl_then_table", "term\n: def\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<dl>\n  <dt>term</dt>\n  <dd>def</dd>\n</dl>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"list_eob_table", "- a\n- b\n\n^\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>a</li>\n  <li>\n    <p>b</p>\n  </li>\n</ul>\n<p>| x | y |\n| - | - |\n| 1 | 2 |</p>\n"},
		{"list_table_no_blank", "- a\n- b\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>a</li>\n  <li>b\n| x | y |\n| - | - |\n| 1 | 2 |</li>\n</ul>\n"},
		{"list_then_para_then_table", "- a\n\npara\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>a</li>\n</ul>\n\n<p>para</p>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"loose_list_table", "- a\n\n- b\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>\n    <p>a</p>\n  </li>\n  <li>\n    <p>b</p>\n  </li>\n</ul>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"nested_then_table", "- a\n  - b\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>a\n    <ul>\n      <li>b</li>\n    </ul>\n  </li>\n</ul>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"ol_table", "1. a\n2. b\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ol>\n  <li>a</li>\n  <li>b</li>\n</ol>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"two_blanks", "- a\n\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>a</li>\n</ul>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
		{"ul_table", "- a\n- b\n\n| x | y |\n| - | - |\n| 1 | 2 |\n", "<ul>\n  <li>a</li>\n  <li>b</li>\n</ul>\n\n<table>\n  <thead>\n    <tr>\n      <th>x</th>\n      <th>y</th>\n    </tr>\n  </thead>\n  <tbody>\n    <tr>\n      <td>1</td>\n      <td>2</td>\n    </tr>\n  </tbody>\n</table>\n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := ToHTML(c.src, nil); got != c.want {
				t.Errorf("mismatch\n got: %q\nwant: %q", got, c.want)
			}
		})
	}
}
