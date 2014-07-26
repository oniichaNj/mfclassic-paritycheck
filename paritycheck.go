package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "file.mfd")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	defer f.Close()
	check(err)

	buffer := make([]byte, 5)
	_, err = f.Read(buffer)
	check(err)

	fmt.Printf("4 first bytes (UID): % X\n", buffer[0:4])
	result := buffer[0] ^ buffer[1] ^ buffer[2] ^ buffer[3]
	if result == buffer[4] {
		fmt.Printf("Parity byte (5) is %X, just as expected.\n", result)
	} else {
		fmt.Printf("Parity byte (5) should be %X, but currently is %X\n", result, buffer[4])
		fmt.Printf("Fix? (y/N) ")
		var a byte
		fmt.Scanf("%c", &a)
		if a == 121 || a == 89 {
			buffer[4] = result
			f, err := os.OpenFile(os.Args[1], os.O_RDWR, 0660)
			defer f.Close()
			check(err)
			_, err = f.Write(buffer)
		} else {
			os.Exit(0)
		}
	}
}
