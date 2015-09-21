package sha1

import (
	"testing"
)

func TestSha1(t *testing.T) {
	var r string
	r = Sha1("still unfinished")
	t.Logf("%s",r)
}
