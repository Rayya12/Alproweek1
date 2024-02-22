package main

import (
	"fmt"
)

func main() {
	var n, r, sp, ast int
	/* input ganjil */
	fmt.Scanln(&n)

	/* for loop bagian atas*/
	for r = 1; r <= n; r++ {
		/*spasi konstan*/
		for sp = 1; sp <= 2*(n+2); sp++ {
			fmt.Print(" ")
		}
		/*spasi segitiga*/
		for sp = n; sp > r; sp-- {
			fmt.Print(" ")
		}
		/*spasi asterisk*/
		for ast = 1; ast <= r; ast++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	/*bagian tengah*/
	/*bagian tengah atas*/
	for r = 1; r <= (n+3)/2; r++ { /*for loop upmid*/
		for sp = 1; sp <= 3*(r-1)+1; sp++ { /*for loop space triangle*/
			fmt.Print(" ")
		}
		for ast = 3 * (n + 1); ast >= 3*(r-1)+1; ast-- { /*for loop ast*/
			fmt.Print("* ")
		}
		fmt.Println()
	}

	//lower mid
	for r = 1; r < n/2-1; r++ { // loop untuk low mid
		for sp = 3 * (n + 1) / 2; sp >= r; sp-- { // loop untuk space trapezoid
			fmt.Print(" ")
		}
		for ast = 1; ast <= (3*n+2)/2+r; ast++ { // loop untuk asterisk
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// bottom part
	for r = 1; r <= (n+1)/2; r++ { // for loop untuk bottom
		for sp = (2*n + 7) / 2; sp >= r; sp-- { // for loop untuk spasi paling kiri
			fmt.Print(" ")
		}
		for ast = n; ast > 2*(r-1); ast-- { // for loop untuk asterisk kiri
			fmt.Print("* ")
		}
		for sp = 1; sp <= 10*(r-1)+2; sp++ { // for loop untuk spasi tengah
			fmt.Print(" ")
		}
		for ast = n; ast > 2*(r-1); ast-- { // for loop untuk asterisk kanan
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
