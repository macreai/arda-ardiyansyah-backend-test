package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findSelfNumbers(1, 5000))
}

func d(n int) int {

	sum := n

	for _, digit := range strconv.Itoa(n) {
		digitValue, _ := strconv.Atoi(string(digit))

		sum += digitValue
	}

	return sum
}

func findSelfNumbers(fromNumber, toNumber int) int {

	generated := make([]bool, toNumber+1)
	selfNumbersSum := 0

	for i := fromNumber; i <= toNumber; i++ {
		generatedValue := d(i)
		if generatedValue <= toNumber {
			generated[generatedValue] = true
		}
	}

	for i := fromNumber; i <= toNumber; i++ {
		if !generated[i] {
			selfNumbersSum += i
		}
	}

	return selfNumbersSum

}

// Sebagian bilangan dapat memiliki lebih dari satu generator. Contohnya:
// Bilangan 101 memiliki dua generator yaitu bilangan 91 dan 100. Buktinya:
// d(91) = 91 + 9 + 1 = 101
// d(100) = 100 + 1 + 0 + 0 = 101
// Bilangan 818 juga memiliki dua generator yaitu bilangan 796 dan 805. Buktinya:
// d(796) = 796 + 7 + 9 + 6 = 818
// d(805) = 805 + 8 + 0 + 5 = 818

// Kelompok bilangan yang memiliki lebih dari satu generator (seperti 101 dan 818) disebut dengan junction-numbers.

// Sedangkan bagi kelompok bilangan yang tidak memiliki generator sama sekali, maka inilah yang disebut dengan self-numbers.

// Sebagai contoh, ada 13 bilangan dari rentang 1-100 yang termasuk kedalam kelompok self-numbers yaitu 1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, dan 97.

// Sekarang tugas Anda adalah:

// Buatlah sebuah kode program yang akan memberikan output berupa hasil penjumlahan kelompok self-numbers dari rentang bilangan 1 sampai dengan 5000
