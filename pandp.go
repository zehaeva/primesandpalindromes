package main

import (
	"fmt"
	"math"
	"strconv"
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

func main() {
	start := time.Now()
	bound := int(math.Pow(10, 5))
	sieve := psieve(bound)
	primes := make([]int, 0, bound)

	for i:=0;i < bound; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	for _, x := range primes {
		y := ((x * x) - 1) / 24
		if isPalindrome(strconv.Itoa(y)) {
			fmt.Printf("palindrome: %v prime: %v time:%v\n", y, x, time.Now().Sub(start))
		}
	}

	fmt.Print("Finished")
}


