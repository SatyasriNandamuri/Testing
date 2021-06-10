// You are given two arbitrarily large numbers,
// stored one digit at a time in a slice.
// The first must be added to the second,
// and the second must be reversed before addition.
//
// The goal is to calculate the sum of those two sets of values.
//
// IMPORTANT NOTE:
// - The input can be any lengths (i.e: it can be 20+ digits long).
// - num1 and num2 can be different lengths.
//
// Sample Inputs:
// num1 = 123456
// num2 = 123456
//
// Sample Output:
// Result: 777777
//
// We would also like to see a demonstration of appropriate unit tests
// for this functionality.

package main

import (
	"fmt"
	"strings"
)

// Caluclate sum and return string
func CalSum(num1 []int, num2 []int, n1 int, n2 int, flag bool) string {
	sum := make([]int, n1) // final output integer variable
	var i, j int
	if flag { // to check and set variable as per which integer array is bigger, as because we want to start from 0 for second whereas last from first
		i, j = 0, n2-1
	} else {
		i, j = n1-1, 0
	}
	k, carry, s := n1-1, 0, 0 // handle carry
	for n2 > 0 {              // loop till shortest array where both of them have values
		s = num1[i] + num2[j] + carry
		sum[k] = (s % 10)
		carry = int(s / 10)
		n1, n2, k = n1-1, n2-1, k-1
		if flag {
			i, j = i+1, j-1
		} else {
			i, j = i-1, j+1
		}
	}
	n1 = k
	for n1 >= 0 { // leftover elements from the largest array after looping through elements while matching and summing up with shortest array
		s = num1[i] + carry
		sum[k] = (s % 10)
		carry = int(s / 10)
		if flag {
			i += 1
		} else {
			i -= 1
		}
		n1, k = n1-1, k-1
	} // final carry case, if its non zero then append it to first position
	if carry == 1 {
		sum = append(sum, 0)
		copy(sum[1:], sum[:])
		sum[k+1] = 1
	}

	return strings.Trim(strings.Replace(fmt.Sprint(sum), " ", "", -1), "[]") // convert to string as per requested output
}

// function provided to us to perform addition with 2 num integer array inputs
func Add(num1 []int, num2 []int) string {
	// implement this
	n1 := len(num1)
	n2 := len(num2)
	var op string
	if n1 >= n2 {
		op = CalSum(num1, num2, n1, n2, false)
	} else {
		op = CalSum(num2, num1, n2, n1, true)
	}
	return op
}

func main() {
	num1 := []int{}
	num2 := []int{}

	num1length := 6
	for i := 1; i <= num1length; i++ {
		num1 = append(num1, i)
	}

	num2length := 6
	for i := 1; i <= num2length; i++ {
		num2 = append(num2, i)
	}

	result := Add(num1, num2)

	fmt.Println("Result:", result)
}
