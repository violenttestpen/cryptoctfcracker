package main

import (
	"errors"
	"flag"
	"fmt"
	"math/big"
	"strings"

	"github.com/violenttestpen/cryptoctfcracker/pkg/lcg"
)

var (
	a, c, m string
	next    uint64
	states  string
)

func main() {
	flag.StringVar(&a, "a", "", "Multiplier")
	flag.StringVar(&c, "c", "", "Increment")
	flag.StringVar(&m, "m", "", "Modulus")
	flag.StringVar(&states, "states", "", "Comma-separated values of states in ascending order")
	flag.Uint64Var(&next, "next", 0, "Generate the next `n` sequence of states")
	flag.Parse()

	if states == "" {
		flag.Usage()
		return
	}

	stateSlice := make([]*big.Int, strings.Count(states, ",")+1)
	for i, seed := range strings.Split(strings.ReplaceAll(states, " ", ""), ",") {
		if s, ok := new(big.Int).SetString(seed, 10); ok {
			stateSlice[i] = s
		} else {
			panic(errors.New("Invalid seed value"))
		}
	}

	var err error
	var multiplier, increment, modulus *big.Int
	if m == "" {
		fmt.Println("Cracking unknown modulus...")
		multiplier, increment, modulus, err = lcg.CrackUnknownModulus(stateSlice)
	} else if a == "" && m != "" {
		fmt.Println("Cracking unknown multiplier...")
		bigM, ok := new(big.Int).SetString(m, 10)
		if ok {
			multiplier, increment, modulus, err = lcg.CrackUnknownMultiplier(stateSlice, bigM)
		}
	} else if a != "" && c == "" && m != "" {
		fmt.Println("Cracking unknown increment...")
		bigM, okM := new(big.Int).SetString(m, 10)
		bigA, okA := new(big.Int).SetString(a, 10)
		if okM && okA {
			multiplier, increment, modulus, err = lcg.CrackUnknownIncrement(stateSlice, bigM, bigA)
		}
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Multiplier (a):", multiplier)
	fmt.Println("Increment  (c):", increment)
	fmt.Println("Modulus    (m):", modulus)

	size := len(stateSlice)
	if next > 0 && size > 0 {
		lastState := stateSlice[size-1]
		fmt.Println("Generating the next", next, "states after", lastState)
		for i := uint64(0); i < next; i++ {
			lastState = lcg.NextLCG(multiplier, increment, modulus, lastState)
			fmt.Println(i+1, lastState)
		}
	}
}
