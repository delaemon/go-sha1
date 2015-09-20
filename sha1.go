package sha1

func Sha1(message string) {
	h0 := 0x67452301
	h1 := 0xEFCDAB89
	h2 := 0x98BADCFE
	h3 := 0x10325476
	h4 := 0xC3D2E1F0

	block := []byte(message)
	words := make([]int, 0, 80)
	for i, chunk := range block {
		words[i] = (chunk[3] | (chunk[2] << 8) | (chunk[1] << 16) | (chunk[0] << 24))
	}

	for i := 16; i < 80; i++ {
		n := words[i - 3] ^ words[i - 8] ^ words[i - 14] ^ words[i - 16];
		words[i] = LeftRotate(n, 1)
	}

	a := h0
	b := h1
	c := h2
	d := h3
	e := h4
}

func LeftRotate(n, b int) int {
	return ((n << b) | (n >> (32 - b))) & 0xffffffff
}
