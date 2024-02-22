package main

import "fmt"

func main() {
	var row int
	var col int

	fmt.Scanln(&row, &col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if i < row/2 {

				if j < col/2 {

					if j == 0 {
						fmt.Print("*")
					} else {
						fmt.Print("  ")
					}

				} else if j == col/2 {
					fmt.Print(" *")
				} else {

					if i == 0 {
						fmt.Print(" *")
					}
				}
			} else if i == row/2 {
				fmt.Print("* ")
			} else {

				if j == col/2 || j == col-1 {
					fmt.Print("* ")
				} else if i == row-1 {

					if j <= col/2 || j == col-1 {
						fmt.Print("* ")
					} else {
						fmt.Print("  ")
					}
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
}
