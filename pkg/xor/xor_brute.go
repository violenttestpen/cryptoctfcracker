package xor

// XOR performs a wrapping multibyte XOR encryption on the buffer.
func XOR(buf, key []byte) []byte {
	cipherbytes := make([]byte, len(buf))
	for i := range cipherbytes {
		cipherbytes[i] = buf[i] ^ key[i%len(key)]
	}
	return cipherbytes
}

// FastXOR performs a wrapping multibyte XOR encryption on the buffer.
func FastXOR(buf, key []byte) []byte {
	cipherbytes := make([]byte, len(buf))
	var j int
	for i := range cipherbytes {
		cipherbytes[i] = buf[i] ^ key[j]
		j++
		if j == len(key) {
			j = 0
		}
	}
	return cipherbytes
}

// IsPrintable is false if any of the provided bytes is greater than 0x7f.
func IsPrintable(buf []byte) bool {
	for _, b := range buf {
		if b > 0x7f {
			return false
		}
	}
	return true
}
