package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	/*
		dddlengthMins := 60
		cawLengthMins := 120

		if dddlengthMins > cawLengthMins {
			fmt.Println("DDD is longer thand CAW")
		} else if dddlengthMins == cawLengthMins {
			fmt.Print("Same Length")
		} else {
			fmt.Println("CAW longer than DDD")
		}

		switch "kubernetes" {
		case "kubernetes":
			fmt.Println("Case 1. kubernetes with lower case")

		case "Kubernetes":
			fmt.Println("Case 2. kubernetes with Capital case")

		case "K8s":
			fmt.Println("Case 3. kubernetes with short name")

		default:
			fmt.Println("Default")

		}*/
	fmt.Println(time.Now().Weekday())
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Even", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd", tmpNum)

	}

	_, err := os.Open("./test1.txt")
	if err != nil {
		fmt.Println("This is the error code:", err)
	}
	fmt.Println("This is the error code:", err)
}
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)

}
