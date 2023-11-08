package main

import "fmt"

func main() {
	/* rank by score
	0-3 => poor
	4-7 => good
	8-9 => very good
	10 => excellect
	otherwise => invalid score
	*/
	/*
		score := 6
		switch score {
		case 0:
			fmt.Println("Poor")
		case 1:
			fmt.Println("Poor")
		case 2:
			fmt.Println("Poor")
		case 3:
			fmt.Println("Poor")
		case 4:
			fmt.Println("Good")
		case 5:
			fmt.Println("Good")
		case 6:
			fmt.Println("Good")
		case 7:
			fmt.Println("Good")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")
		}
	*/

	switch score := 6; score {
	case 0, 1, 2, 3:
		fmt.Println("Poor")
	case 4, 5, 6, 7:
		fmt.Println("Good")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	// imitate nested if

	switch no := 25; {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%3 == 0:
		fmt.Printf("%d is divisible by 3\n", no)
	case no%2 == 1:
		fmt.Printf("%d is an odd number\n", no)
	}

	// fallthrough
	switch no := 4; no {
	case 0:
		fmt.Println(" == 0")
		fallthrough
	case 1:
		fmt.Println(" <= 1")
		fallthrough
	case 2:
		fmt.Println(" <= 2")
		fallthrough
	case 3:
		fmt.Println(" <= 3")
		fallthrough
	case 4:
		fmt.Println(" <= 4")
		fallthrough
	case 5:
		fmt.Println(" <= 5")
		fallthrough
	case 6:
		fmt.Println(" <= 6")
		// fallthrough
	case 7:
		fmt.Println(" <= 7")
		fallthrough
	case 8:
		fmt.Println(" <= 8")
	}

	// fallthrough applied
	switch spotifySubscription := "super"; spotifySubscription {
	case "super":
		fmt.Println("Family of 3")
		fallthrough
	case "pro":
		fmt.Println("Private play list")
		fallthrough
	case "premium":
		fmt.Println("No ads")
		fallthrough
	case "free":
		fmt.Println("Listen to music")
	}
}
