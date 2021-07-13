package lcg

import (
	"errors"
	"math/big"
)

// https://web.archive.org/web/20201112025614/https://tailcall.net/blog/cracking-randomness-lcgs/
// https://github.com/nccgroup/featherduster/blob/master/cryptanalib/modern.py

var (
	errMinTwoStates   = errors.New("At least two states are required")
	errMinThreeStates = errors.New("At least three states are required")
	errMinFiveStates  = errors.New("At least five states are required")
)

// CrackUnknownIncrement derives the increment from at least 2 states, the multiplier and the modulus.
func CrackUnknownIncrement(states []*big.Int, bigM, bigA *big.Int) (multiplier *big.Int, increment *big.Int, modulus *big.Int, err error) {
	if len(states) < 2 {
		err = errMinTwoStates
		return
	}

	// increment = (states[1] - states[0]*multiplier) % modulus
	bigC := new(big.Int)
	bigC.Mul(states[0], bigA)
	bigC.Sub(states[1], bigC)
	bigC.Mod(bigC, bigM)
	multiplier, increment, modulus = bigA, bigC, bigM
	return
}

// CrackUnknownMultiplier derives the multiplier and increment from at least 3 states and the modulus.
func CrackUnknownMultiplier(states []*big.Int, bigM *big.Int) (multiplier *big.Int, increment *big.Int, modulus *big.Int, err error) {
	if len(states) < 3 {
		err = errMinThreeStates
		return
	}

	bigA := new(big.Int)
	var bigStateModInv big.Int

	// multiplier = (states[2] - states[1]) * modinv(states[1] - states[0], modulus) % modulus
	bigStateModInv.Sub(states[1], states[0])
	bigStateModInv.ModInverse(&bigStateModInv, bigM)
	bigA.Sub(states[2], states[1])
	bigA.Mul(bigA, &bigStateModInv)
	bigA.Mod(bigA, bigM)

	multiplier, increment, modulus, err = CrackUnknownIncrement(states, bigM, bigA)
	return
}

// CrackUnknownModulus derives the multiplier, increment and the modulus from at least 5 states.
func CrackUnknownModulus(states []*big.Int) (multiplier *big.Int, increment *big.Int, modulus *big.Int, err error) {
	if len(states) < 5 {
		err = errMinFiveStates
		return
	}

	// diffs = [s1 - s0 for s0, s1 in zip(states, states[1:])]
	diffs := make([]*big.Int, len(states)-1, len(states)-1)
	for i, size := 0, len(states)-1; i < size; i++ {
		diffs[i] = new(big.Int).Sub(states[i+1], states[i])
	}

	// zeroes = [t2*t0 - t1*t1 for t0, t1, t2 in zip(diffs, diffs[1:], diffs[2:])]
	zeroes := make([]*big.Int, len(diffs)-2, len(diffs)-2)
	for i, size := 0, len(diffs)-2; i < size; i++ {
		arg1 := new(big.Int).Mul(diffs[i+2], diffs[i])
		arg2 := new(big.Int).Mul(diffs[i+1], diffs[i+1])
		zeroes[i] = arg1.Sub(arg1, arg2)
	}

	// modulus = abs(reduce(gcd, zeroes))
	bigM, big1 := zeroes[0], big.NewInt(1)
	for i, size := 1, len(zeroes); i < size; i++ {
		bigM.GCD(big1, big1, bigM, zeroes[i])
	}

	multiplier, increment, modulus, err = CrackUnknownMultiplier(states, bigM)
	return
}

// NextLCG returns the next state of the LCG.
func NextLCG(bigA, bigC, bigM *big.Int, state *big.Int) *big.Int {
	bigState := new(big.Int).SetUint64(state.Uint64())
	bigState.Mul(bigState, bigA)
	bigState.Add(bigState, bigC)
	bigState.Mod(bigState, bigM)
	return bigState
}
