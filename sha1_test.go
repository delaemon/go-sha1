package sha1

import "testing"

type Test struct {
	in string
	out string
}

var tests = []Test{
	{"The quick brown fox jumps over the lazy dog",
	"2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"},
	{"The quick brown fox jumps over the lazy cog",
	"de9f2c7fd25e1b3afad3e85a0bd17d9b100db4b3"},
	{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	{"testing\n", "9801739daae44ec5293d4e1f53d3f4d2d426d91c"},
	{"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	"025ecbd5d70f8fb3c5457cd96bab13fda305dc59"},
}

func TestSha1(t *testing.T) {
	for i, test := range tests {
		out := Sha1(test.in)
		if out != test.out {
			t.Errorf("#%d: Sha1(%s) = %s; want %s", i, test.in, out, test.out)
		}
	}
}
