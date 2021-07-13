package lcg

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownIncrement(t *testing.T) {
	assert := assert.New(t)

	const a, m = 81853448938945944, 9223372036854775783
	states := []*big.Int{big.NewInt(4501678582054734753), big.NewInt(4371244338968431602)}
	bigM := big.NewInt(m)
	bigA := big.NewInt(a)
	multiplier, increment, modulus, err := CrackUnknownIncrement(states, bigM, bigA)

	assert.NoError(err)
	assert.EqualValues(81853448938945944, multiplier.Int64())
	assert.EqualValues(7247473133432955167, increment.Int64())
	assert.EqualValues(9223372036854775783, modulus.Int64())
}

func TestUnknownIncrementAndMultiplier(t *testing.T) {
	assert := assert.New(t)

	const m = 9223372036854775783
	states := []*big.Int{big.NewInt(6473702802409947663), big.NewInt(6562621845583276653), big.NewInt(4483807506768649573)}
	bigM := big.NewInt(m)
	multiplier, increment, modulus, err := CrackUnknownMultiplier(states, bigM)

	assert.NoError(err)
	assert.EqualValues(6068601099849884345, multiplier.Int64())
	assert.EqualValues(8366172131088513789, increment.Int64())
	assert.EqualValues(9223372036854775783, modulus.Int64())
}

func TestUnknownIncrementAndMultiplierAndModulus(t *testing.T) {
	assert := assert.New(t)

	states := []*big.Int{big.NewInt(2818206783446335158), big.NewInt(3026581076925130250), big.NewInt(136214319011561377),
		big.NewInt(359019108775045580), big.NewInt(2386075359657550866), big.NewInt(1705259547463444505), big.NewInt(2102452637059633432)}
	multiplier, increment, modulus, err := CrackUnknownModulus(states)

	assert.NoError(err)
	assert.EqualValues(302080878814014441, multiplier.Int64())
	assert.EqualValues(3613230612905734352, increment.Int64())
	assert.EqualValues(4611686018427387847, modulus.Int64())
}

func TestNextLCG(t *testing.T) {
	assert := assert.New(t)

	const a, c, m, seed = 672257317069504227, 7382843889490547368, 9223372036854775783, 2300417199649672133
	bigA := big.NewInt(a)
	bigC := big.NewInt(c)
	bigM := big.NewInt(m)
	state := big.NewInt(seed)

	nextState := NextLCG(bigA, bigC, bigM, state)
	assert.EqualValues(2071270403368304644, nextState.Int64())

	nextState = NextLCG(bigA, bigC, bigM, nextState)
	assert.EqualValues(5907618127072939765, nextState.Int64())

	nextState = NextLCG(bigA, bigC, bigM, nextState)
	assert.EqualValues(5457707446309988294, nextState.Int64())
}

func BenchmarkUnknownIncrement(b *testing.B) {
	const a, m = 81853448938945944, 9223372036854775783
	states := []*big.Int{big.NewInt(4501678582054734753), big.NewInt(4371244338968431602)}
	bigM := big.NewInt(m)
	bigA := big.NewInt(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, _ = CrackUnknownIncrement(states, bigM, bigA)
	}
}

func BenchmarkUnknownIncrementAndMultiplier(b *testing.B) {
	const m = 9223372036854775783
	states := []*big.Int{big.NewInt(6473702802409947663), big.NewInt(6562621845583276653), big.NewInt(4483807506768649573)}
	bigM := big.NewInt(m)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, _ = CrackUnknownMultiplier(states, bigM)
	}
}

func BenchmarkUnknownIncrementAndMultiplierAndModulus(b *testing.B) {
	states := []*big.Int{big.NewInt(2818206783446335158), big.NewInt(3026581076925130250), big.NewInt(136214319011561377),
		big.NewInt(359019108775045580), big.NewInt(2386075359657550866), big.NewInt(1705259547463444505), big.NewInt(2102452637059633432)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, _ = CrackUnknownModulus(states)
	}
}

func BenchmarkNextLCG(b *testing.B) {
	const a, c, m, seed = 672257317069504227, 7382843889490547368, 9223372036854775783, 2300417199649672133
	bigA := big.NewInt(a)
	bigC := big.NewInt(c)
	bigM := big.NewInt(m)
	state := big.NewInt(seed)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NextLCG(bigA, bigC, bigM, state)
	}
}
