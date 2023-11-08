package main

import "fmt"

func main() {

	fmt.Println("for[v1]")
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("for[v2] (while)")
	/*
		no := 1
		for no < 100 {
			no += no
		}
		fmt.Println("no :", no)
	*/

	for no := 1; no < 100; {
		no += no
		fmt.Println("no :", no)
	}

	fmt.Println("for[v3] (infinite)")
	/*
		no := 1
		for {
			no += no
			fmt.Println("no :", no)
			if no >= 128 {
				break
			}
		}
	*/

	for no := 1; ; {
		no += no
		fmt.Println("no :", no)
		/*
			if no >= 128 {
				break
			}
		*/
		if no < 128 {
			continue
		}
		break
	}

	// Labels: To control the outer loop from an inner loop

I_LOOP:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("================")
				continue I_LOOP
			}
		}
	}
}
