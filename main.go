package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
	prevFrame := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	curFrame := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	step := 0

	for {
		for i := 0; i < len(prevFrame); i++ {
			for j := 0; j < len(prevFrame[0]); j++ {
				counter := 0
				if i == 0 && j == 0 {
					counter += prevFrame[0][1] +
						prevFrame[1][0] +
						prevFrame[1][1]
				} else if i == 0 && j == len(prevFrame[0])-1 {
					counter += prevFrame[0][len(prevFrame[0])-2] +
						prevFrame[1][len(prevFrame[0])-1] +
						prevFrame[1][len(prevFrame[0])-2]
				} else if i == len(prevFrame)-1 && j == len(prevFrame[0])-1 {
					counter += prevFrame[len(prevFrame)-2][len(prevFrame[0])-1] +
						prevFrame[len(prevFrame)-1][len(prevFrame[0])-2] +
						prevFrame[len(prevFrame)-2][len(prevFrame[0])-2]
				} else if i == len(prevFrame)-1 && j == 0 {
					counter += prevFrame[len(prevFrame)-2][0] +
						prevFrame[len(prevFrame)-1][1] +
						prevFrame[len(prevFrame)-2][1]
				} else if i == 0 {
					counter += prevFrame[i][j-1] +
						prevFrame[i+1][j-1] +
						prevFrame[i+1][j] +
						prevFrame[i+1][j+1] +
						prevFrame[i][j+1]
				} else if j == len(prevFrame[0])-1 {
					counter += prevFrame[i-1][j] +
						prevFrame[i-1][j-1] +
						prevFrame[i][j-1] +
						prevFrame[i+1][j-1] +
						prevFrame[i+1][j]
				} else if i == len(prevFrame)-1 {
					counter += prevFrame[i][j-1] +
						prevFrame[i-1][j-1] +
						prevFrame[i-1][j] +
						prevFrame[i-1][j+1] +
						prevFrame[i][j+1]
				} else if j == 0 {
					counter += prevFrame[i-1][j] +
						prevFrame[i-1][j+1] +
						prevFrame[i][j+1] +
						prevFrame[i+1][j+1] +
						prevFrame[i+1][j]
				} else {
					counter += prevFrame[i-1][j] +
						prevFrame[i-1][j+1] +
						prevFrame[i][j+1] +
						prevFrame[i+1][j+1] +
						prevFrame[i+1][j] +
						prevFrame[i+1][j-1] +
						prevFrame[i][j-1] +
						prevFrame[i-1][j-1]
				}

				if prevFrame[i][j] == 0 && counter == 3 {
					curFrame[i][j] = 1
				} else if prevFrame[i][j] == 1 && counter == 2 || counter == 3 {
					curFrame[i][j] = 1
				} else {
					curFrame[i][j] = 0
				}

				counter = 0
			}
		}

		fmt.Printf("Step: %d \n", step)

		for i := 0; i < len(prevFrame); i++ {
			for j := 0; j < len(prevFrame[0]); j++ {
				if curFrame[i][j] == 0 {
					fmt.Printf("%s ", "□")
				} else {
					fmt.Printf("%s ", "■")
				}

				prevFrame[i][j] = curFrame[i][j]
			}
			fmt.Print("\n")
		}

		step += 1
		time.Sleep(time.Second / 3)
		screen.Clear()
	}
}
