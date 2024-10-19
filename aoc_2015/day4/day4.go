package main

import (
	"crypto/md5"
	"fmt"
)

func mineAdventCoin(input, prefix string) string {
	num := ""
	md5sum := md5.Sum([]byte(input + num))
	for i := 0; ; i++ {
		sum := fmt.Sprintf("%x", md5sum)
		if sum[:len(prefix)] == prefix {
			break
		}
		num = fmt.Sprintf("%d", i)
		md5sum = md5.Sum([]byte(input + num))
	}
	fmt.Printf("%x\n", md5sum)
	return num
}

func main() {
	input := "bgvyzdsv"
	// part1Prefix := "00000"
	part2Prefix := "000000"

	fmt.Println("Result: ", mineAdventCoin(input, part2Prefix))
}
