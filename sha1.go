package sha1
import "fmt"

func Sha1(message string) string {

	//Initialize variables:
	var h0 uint32 = 0x67452301
	var h1 uint32 = 0xEFCDAB89
	var h2 uint32 = 0x98BADCFE
	var h3 uint32 = 0x10325476
	var h4 uint32 = 0xC3D2E1F0

	// Pre-processing:
	m := []byte(message)
	m = append(m, byte(0x80))
	m = append(m, byte(0x00 * ((56 - (len(message) * 8 + 1) % 64) % 64)))

	words := make([]uint32, 80, 80)
	for i := 0; i < 16; i++ {
		chunk := message[0:4]
		words[i] = uint32(chunk[3] | (chunk[2] << 8) | (chunk[1] << 16) | (chunk[0] << 24))
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

	var f, k uint32
	for i := 0; i < 80; i++ {
		if 0 <= i && i <= 19 {
			f = d ^ (b & (c ^ d))
			k = 0x5A827999
		} else if 20 <= i && i <= 39 {
			f = b ^ c ^ d
			k = 0x6ED9EBA1
		} else if 40 <= i && i <= 59 {
			f = (b & c) ^ (b & d) ^ (c & d)
			k = 0x8F1BBCDC
		} else if 60 <= i && i <= 79 {
			f = b ^ c ^ d
			k = 0xCA62C1D6
		}

		a, b, c, d, e = LeftRotate(a, uint32(5)) + f + e + k + words[i] & 0xffffffff, a, LeftRotate(b, uint32(30)), c, d
	}

	h0 = (h0 + a) & 0xffffffff
	h1 = (h1 + b) & 0xffffffff
	h2 = (h2 + c) & 0xffffffff
	h3 = (h3 + d) & 0xffffffff
	h4 = (h4 + e) & 0xffffffff

	return fmt.Sprintf("%08x%08x%08x%08x%08x", h0, h1, h2, h3, h4)
}

func LeftRotate(n, b uint32) uint32 {
	return ((n << b) | (n >> (32 - b))) & 0xffffffff
}
