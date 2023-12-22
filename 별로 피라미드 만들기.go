package main

import "fmt"

func main() {
	printPyramid(5)
}

// 위쪽으로 쌓인 삼각형 출력
func printPyramid(height int) {
	for i := 1; i <= height; i++ {
		// 별표 출력
		for star := 1; star <= i; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 아래쪽으로 쌓인 삼각형 출력
	for i := height - 1; i >= 1; i-- {
		// 별표 출력
		for star := 1; star <= i; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	printDiamond(4)
}
func printDiamond(height int) {
	// 위쪽으로 쌓인 삼각형 출력
	for i := 0; i < height; i++ {
		// 공백 출력
		for space := 0; space < height-i-1; space++ {
			fmt.Print(" ")
		}
		// 별표 출력
		for star := 0; star < 2*i+1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 아래쪽으로 쌓인 역삼각형 출력
	for i := height - 2; i >= 0; i-- {
		// 공백 출력
		for space := 0; space < height-i-1; space++ {
			fmt.Print(" ")
		}
		// 별표 출력
		for star := 0; star < 2*i+1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 위쪽으로 쌓인 삼각형 출력
	for i := 0; i < height; i++ {
		// 공백 출력
		for space := 0; space < height-i-1; space++ {
			fmt.Print(" ")
		}
		// 별표 출력
		for star := 0; star < 2*i+1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	// 아래쪽으로 쌓인 역삼각형 출력
	for i := height - 2; i >= 0; i-- {
		// 공백 출력
		for space := 0; space < height-i-1; space++ {
			fmt.Print(" ")
		}
		// 별표 출력
		for star := 0; star < 2*i+1; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
