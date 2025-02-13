package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	f := flag.Bool("f", false, "Generate a float number")
	flag.Parse()
	if *f {
		upperBound := "1.0"
		if flag.NArg() > 0 {
			upperBound = flag.Arg(0)
		}
		generateFloat(upperBound)
	} else {
		upperBound := "100"
		if flag.NArg() > 0 {
			upperBound = flag.Arg(0)
		}
		generateInt(upperBound)
	}
}

func generateFloat(upperBound string) {
	upperBoundFloat, err := strconv.ParseFloat(upperBound, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid upper bound: %s\n", err)
		os.Exit(1)
	}

	randomFloat := rand.Float64() * upperBoundFloat
	fmt.Println(randomFloat)
}

func generateInt(upperBound string) {
	upperBoundInt, err := strconv.Atoi(upperBound)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid upper bound: %s\n", err)
		os.Exit(1)
	}

	randomNumber := rand.Intn(upperBoundInt)
	fmt.Println(randomNumber)
}
