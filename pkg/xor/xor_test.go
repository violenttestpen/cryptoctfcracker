package xor

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAngstromCTFExclusiveCipher(t *testing.T) {
	assert := assert.New(t)

	const ciphertext = "ae27eb3a148c3cf031079921ea3315cd27eb7d02882bf724169921eb3a469920e07d0b883bf63c018869a5090e8868e331078a68ec2e468c2bf13b1d9a20ea0208882de12e398c2df60211852deb021f823dda35079b2dda25099f35ab7d218227e17d0a982bee7d098368f13503cd27f135039f68e62f1f9d3cea7c"
	key := []byte{237, 72, 133, 93, 102}
	cipherbytes, _ := hex.DecodeString(ciphertext)
	output := FastXOR(cipherbytes, key)

	assert.Contains(string(output), "actf{who_needs_aes_when_you_have_xor}")
}

func BenchmarkXOR(b *testing.B) {
	const ciphertext = "ae27eb3a148c3cf031079921ea3315cd27eb7d02882bf724169921eb3a469920e07d0b883bf63c018869a5090e8868e331078a68ec2e468c2bf13b1d9a20ea0208882de12e398c2df60211852deb021f823dda35079b2dda25099f35ab7d218227e17d0a982bee7d098368f13503cd27f135039f68e62f1f9d3cea7c"
	key := []byte{237, 72, 133, 93, 102}
	cipherbytes, _ := hex.DecodeString(ciphertext)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = XOR(cipherbytes, key)
	}
}

func BenchmarkFastXOR(b *testing.B) {
	const ciphertext = "ae27eb3a148c3cf031079921ea3315cd27eb7d02882bf724169921eb3a469920e07d0b883bf63c018869a5090e8868e331078a68ec2e468c2bf13b1d9a20ea0208882de12e398c2df60211852deb021f823dda35079b2dda25099f35ab7d218227e17d0a982bee7d098368f13503cd27f135039f68e62f1f9d3cea7c"
	key := []byte{237, 72, 133, 93, 102}
	cipherbytes, _ := hex.DecodeString(ciphertext)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FastXOR(cipherbytes, key)
	}
}
