package main

import (
	"fmt"
	"math/big"
	"sort"
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*
* Make a slice containing the unique values from the slice that receives
* @intSlice slice from we are going to get the unique values
*/
func unique(intSlice []int) []int{
	keys := make(map[int]bool)
	list := []int{}
	for _,entry := range intSlice{
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list,entry)
		}
	}
	return list
}

/**
* @n Is an int that states the limit number where the primes number can occurr.
* @l Is the length of the list of values in the ciphertext.
* @primeProducts list of the prime products of the ciphertext
*/
func cryptopangram(n int, l int, primeProducts []int) string {
	// define alphabet
	alphabet := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// create a list for the ciphered primes and empty message
	primes :=make([]int, l+1)
	message := ""

	// create an auxiliar variable to store the GCD. it initialize in 0
	factor := new(big.Int)
	// calculate the gcd of the two first non equal ciphered primes
	var index int
	index = -1
	for i := 0; i < len(primeProducts)-1; i++ {
		if primeProducts[i] != primeProducts[i+1]{
			factor.GCD(big.NewInt(1), big.NewInt(1), big.NewInt(int64(primeProducts[i])), big.NewInt(int64(primeProducts[i+1])))
			index = i+1
			primes[index] = int(factor.Int64())
			break
		}
	}
	// calculate backpropagation
	for i := index; i > 0; i-- {
		prev_factor := primeProducts[i-1] / primes[i]
		primes[i-1] = prev_factor
	}
	// calculate frontpropagation
	for i:=index; i < len(primeProducts); i++ {

		next_factor := primeProducts[i] / primes[i]
		primes[i+1] = next_factor
	}

	// remove repetitions
	unique_primes := unique(primes[:])

	// copy into new list and sort it
	sort.Ints(unique_primes)
	unique_primes_sorted := unique_primes

	// map the sorted list to each letter in a dictionary
	 prime_map := make(map[int]string)
	 for i := range unique_primes_sorted {
		 prime_map[unique_primes_sorted[i]] = alphabet[i]
	 }

	 // map the calculated factors to their respective letters
	 for _, prime := range primes {
		 message += prime_map[prime]
	 }

	 return message
}

/*****************************************************************************/

func test_cryptopangrams() {

	var n_cases, n, l int
	//var primeProducts string
	// reads first line, the number of cases
	fmt.Scanf("%d\n", &n_cases)

	// iterate the number of cases to test
	for i := 0; i < n_cases; i++ {
		// read line with N and L
		fmt.Scanf("%d %d\n", &n, &l)

		// read line with list of integers
		reader := bufio.NewReader(os.Stdin)
		primeProductsString, _ := reader.ReadString('\n')
		primeProductsString = strings.Trim(primeProductsString,"\n")
		primeProductsString = strings.Trim(primeProductsString,"\r")
		primeProductsSlice := strings.SplitN(primeProductsString, " ", -1)

		primeProducts := make([]int, 0)
		for _, val := range primeProductsSlice {
			aux,_ := strconv.Atoi(val)
			primeProducts = append(primeProducts, aux)
		}
		//fmt.Println(primeProductsSlice)
		//fmt.Printf("%T\n", primeProductsSlice)
		//fmt.Println(len(primeProductsSlice))
		//fmt.Println(primeProducts)
		//fmt.Printf("%T\n", primeProducts)
		//fmt.Println(len(primeProducts))
		//fmt.Scanf("%s\n", &primeProducts)
		//fmt.Println(n, l, primeProducts)
		 message := cryptopangram(n, l, primeProducts)
		 fmt.Printf("Case #%d: %s\n",i+1, message)
	}

}

func main() {
	test_cryptopangrams()
}
