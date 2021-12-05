package main

import "fmt"

func Break() {
	//break = menghentikan seluruh perulangan
	//continue = menghentikan perulangan yang berjalan,
	//dan lanjut ke perulangan selanjutnya

	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke", i)
		if i%2 == 0 {
			break
		}
	}

}
