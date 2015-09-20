package sha1

import "fmt"

func TestGenerate() {
	r := Sha1()
	fmt.Println(r)
}
