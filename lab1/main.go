package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	start := time.Now()
	// changeAfterRemove := flag.Bool("strategy", false, "Should player change box after one was taken?")
	rounds := flag.Uint("c", 1, "Number of rounds to play")
	numberOfBoxes := flag.Int("n", 3, "Number of boxes, musst be 3 or grater")
	removeBoxesCount := flag.Int("remove", 3, "Number of boxes to remove after player picks")
	flag.Parse()

	// fmt.Println(serialTest(*rounds, *changeAfterRemove))
	fmt.Println(serialTest(*rounds, false, *numberOfBoxes, *removeBoxesCount)*100, "%")
	fmt.Println(serialTest(*rounds, true, *numberOfBoxes, *removeBoxesCount)*100, "%")
	elapsed := time.Since(start)
	fmt.Printf("Gra toczy się dla %d pudełek\n", *numberOfBoxes)
	fmt.Printf("Execution took %s", elapsed)
}

func playSimpleGame(changeAfterRemove bool) bool {
	pickBox := func() int8 {
		return int8(rand.IntN(2))
	}

	var boxes [3]bool
	boxes[rand.IntN(3)] = true
	switch pickedBoxIndex := pickBox(); changeAfterRemove {
	case true:
		return !boxes[pickedBoxIndex]

	case false:
		return boxes[pickedBoxIndex]

	}
	fmt.Println("Return wymagany przez składnie")
	return false
}

func playNBoxexGame(changeAfterRemove bool, numberOfBoxes int, removeBoxesCount int) bool {
	boxes := make([]bool, numberOfBoxes)
	boxes[rand.IntN(numberOfBoxes)] = true
	switch pickedBoxIndex := rand.IntN(numberOfBoxes); changeAfterRemove {
	case true:
		// Tworzymy maskę
		mask := make([]bool, numberOfBoxes)
		for removedCount := 0; removedCount < removeBoxesCount; {
			randomIndex := rand.IntN(numberOfBoxes)
			if !boxes[randomIndex] && !mask[randomIndex] {
				mask[randomIndex] = true
				removedCount++
			}
		}
		//Losujemy pudełko które ostatecznie wybierzemy
		newPickedIndex := rand.IntN(numberOfBoxes)
		for {
			if !mask[newPickedIndex] && newPickedIndex != pickedBoxIndex {
				break
			} else {

				newPickedIndex = rand.IntN(numberOfBoxes)
			}
		}
		return boxes[newPickedIndex]

	// Jeśli nie zmieniamy wyboru po odsłonięciu pudełek nie musimy wiedzieć któr otwarto
	case false:
		return boxes[pickedBoxIndex]

	}
	fmt.Println("Return wymagany przez składnie")
	return false
}

func serialTest(gamesToPlay uint, changeAfterRemove bool, numberOfBoxes int, removeBoxesCount int) float64 {
	positiveCount := uint(0)
	switch numberOfBoxes {
	case 3:
		for range gamesToPlay {
			if playSimpleGame(changeAfterRemove) {
				positiveCount++
			}
		}
	default:
		for range gamesToPlay {
			if playNBoxexGame(changeAfterRemove, numberOfBoxes, removeBoxesCount) {
				positiveCount++
			}
		}
	}
	fmt.Printf("Pozytywne wyniki: %d Negatywne: %d dla strategi %t\n", positiveCount, gamesToPlay-positiveCount, changeAfterRemove)
	return float64(positiveCount) / float64(gamesToPlay)
}
