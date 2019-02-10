package main

import (
	"flag"
	"fmt"
	"github.com/shogo82148/go-shuffle"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numerics = "0123456789"
	symbols  = "?./+,;=%@#$!()'"
)

var (
	tmpPassword []string
	numPassword int  = 12
	numLetters  int  = 8
	numNumerics int  = 3
	numSymbols  int  = 1
	displayHelp bool = false
)

// RandStrings provides a random string of letters
func RandStrings(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandNums provides a random string of numerics
func RandNums(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numerics[rand.Intn(len(numerics))]
	}
	return string(b)
}

// RandSymbols provides a random string of symbols
func RandSymbols(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

func getParams() {
	lenPtr := flag.Int("l", numPassword, "Password length")
	numPtr := flag.Int("n", numNumerics, "Numerics count")
	symPtr := flag.Int("s", numSymbols, "Symbols count")
	helpPtr := flag.Bool("h", false, "Display help")
	flag.Parse()
	displayHelp = bool(*helpPtr)
	numPassword = int(*lenPtr)
	numNumerics = int(*numPtr)
	numSymbols = int(*symPtr)
	numLetters = int(*lenPtr) - numNumerics - numSymbols
}

func main() {
	getParams()
	if displayHelp {
		fmt.Println(".:HELP:.")
		fmt.Println("--------")
		fmt.Println(" -l		password's length (int)")
		fmt.Println(" -n		numerics' count (int)")
		fmt.Println(" -s		symbols' count (int)")
		fmt.Println("Example: gen_password -l 24 -n 6 -s 4")
		os.Exit(0)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	s := RandStrings(numLetters)
	for i := 0; i < len(s); i++ {
		tmpPassword = append(tmpPassword, string(s[i]))
	}
	shuffle.Strings(tmpPassword)
	n := RandNums(numNumerics)
	for i := 0; i < len(n); i++ {
		tmpPassword = append(tmpPassword, string(n[i]))
	}
	shuffle.Strings(tmpPassword)
	sym := RandSymbols(numSymbols)
	for i := 0; i < len(sym); i++ {
		tmpPassword = append(tmpPassword, string(sym[i]))
	}
	shuffle.Strings(tmpPassword)
	fmt.Println(strings.Join(tmpPassword, ""))
}
