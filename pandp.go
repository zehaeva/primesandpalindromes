package main

import (
	"fmt"
	"github.com/caleblloyd/primesieve"
	"math/big"
	"time"
)

func pwheel() {

}

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func psieve(bound int) []bool {
	sieve := make([]bool, bound)

	sieve[0] = false
	sieve[1] = false

	for i:=2; i < bound; i++ {
		sieve[i] = true
	}

	for i:=2; i < bound; i++ {
		if sieve[i] {
			for j:=i+i; j < bound; j+=i {
				sieve[j] = false
			}
		}
	}

	return sieve
}

func generate(process func(x int)) {
	for i := 0; i < 10; i++ {
		process(i)
	}
}

func main() {
	c := make(chan uint64)

	primesieve.Channel(c)
	start := time.Now()

	var x big.Int
	var one = big.NewInt(1)
	var twentyfour = big.NewInt(24)

	for {
		fmt.Scanln()
		p := <-c
		x.SetUint64(p)
		x.Mul(&x, &x)
		x.Sub(&x, one)
		x.Div(&x, twentyfour)
		t := x.String()
		if isPalindrome(t) {
			fmt.Printf("palindrome: %v prime: %v time:%v\n", x.String(), p, time.Now().Sub(start))
		}
	}

	fmt.Println()
	// 2 3 5 7 11 13 17 19 23 29
}

//func main() {
//	start := time.Now()
//
//	process := func(x int) {
//		fmt.Println("Processing", x)
//		y := ((x * x) - 1) / 24
//		if isPalindrome(strconv.Itoa(y)) {
//			fmt.Printf("palindrome: %v prime: %v time:%v\n", y, x, time.Now().Sub(start))
//		}
//	}
//	generate(process)
//
//	fmt.Print("Finished")
//}


