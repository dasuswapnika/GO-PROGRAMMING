package main

import (
	"fmt"
	"student/student"
)

func main() {
	s := student.Student{
		Name: "Swapnika",
		Age:  20,
	}

	fmt.Println("Student Details:")
	s.Display()
}
