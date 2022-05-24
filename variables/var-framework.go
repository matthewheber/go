package main

import (
	"fmt"
)

func main() {
	name := "Matthew"
	course := "Getting started with Kubernetes"

	fmt.Println("\n Hi", name, "your current course is ", course)
	updateCourse(course)

	fmt.Println("You're currently watching", course)

}

func updateCourse(course string) string {
	course = "geeting Started with Docker"
	fmt.Println("Updated couse to ", course)
	return course

}
