package main

import (
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/violenttestpen/cryptoctfcracker/pkg/xor"
)

var (
	key, hexString, file, grep string
	printable                  bool
)

func main() {
	flag.StringVar(&key, "key", "", "Known plaintext string")
	flag.StringVar(&hexString, "hex", "", "Hex-encoded cipherstring")
	flag.StringVar(&file, "file", "", "Path to cipherbytes")
	flag.StringVar(&grep, "grep", "", "Known plaintext string to reverse filter output list")
	flag.BoolVar(&printable, "printable", false, "Keep only outputs with printable characters")
	flag.Parse()

	if key == "" || (hexString == "" && file == "") {
		flag.Usage()
		panic(errors.New("Missing arguments"))
	}

	var err error
	var cipherbytes []byte
	if hexString != "" {
		cipherbytes, err = hex.DecodeString(hexString)
	} else if file != "" {
		if fileInfo, err := os.Stat(file); err == nil && !fileInfo.IsDir() {
			cipherbytes, err = os.ReadFile(file)
		}
	}

	if err != nil {
		panic(err)
	}

	magic := []byte(key)
	n := len(magic)
	for i, length := 0, len(cipherbytes); i <= length-n; i++ {
		key := xor.XOR(cipherbytes[i:i+n], magic)
		output := xor.XOR(cipherbytes, key)
		if !printable || xor.IsPrintable(output) {
			if grep == "" || strings.Contains(string(output), grep) {
				fmt.Println(key, '/', string(output))
			}
		}
	}
}
