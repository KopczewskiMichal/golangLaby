package main

// go run -p 1 z1.go   - uruchomienie go run ale kompilacja na 1 wątku
import (
	"fmt"
)

func main() {
	// loop_helper(uint64(18_446_744_073_709_551_615))
	maxFor := 1
	maxCount := 0
	limit := 10_000_000
	sum := uint64(0)
	for i := 1; i <= limit; i++ {
		count := loop_helper(i)

		sum += uint64(count)
		if count > maxCount {
			maxCount = count
			maxFor = i
			}
		}
		
	fmt.Printf("Dotychczas największe maksimum dla %d : %d\n", maxFor, maxCount)
	fmt.Printf("Średnia wynosi: %f", float64(sum)/float64(limit))
}

func rec_helper(n uint64) {
	fmt.Println(n)
	if n == 4 {
		fmt.Println("Obecnie n = 4")
		fmt.Println("Zakończono działanie programu")
	} else {
		switch n % 2 {
		case 0:
			rec_helper(n / 2)
		case 1:
			rec_helper(n*3 + 1)
		}
	}
}

func loop_helper(n int) int {
	counter := 0
	for n != 4 {
		counter++
		switch n % 2 {
		case 0:
			n = n / 2
		case 1:
			n = n*3 + 1
		}
		// if counter >= 1000 {
		// 	fmt.Println("Przekroczyliśmy 1000 iteracji")
		// 	break
		// }
	}
	return counter
}

// colatz
// Jeżeli liczba jest parzysta dzielimy przez 2
// Jeśli liczba nie jest parzysta mnożymy przez 3 i dodajemy 1
