package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world from go")

	fmt.Println("\nEnter your age")
	var ageData int
	fmt.Scanln(&ageData)

	// processing input for output
	ifElseExample(ageData)
	switchExample(ageData)
	forLoopExample(ageData)

	// for range
	experience := []string{"beginner", "intermediate", "pro"}
	forRangeExample(experience)

	// goto
	goToExample()

	cmdArgs()
}

func ifElseExample(age int) {
	if age > 22 {
		fmt.Println("You must be in a job/business")
	} else if age < 18 {
		fmt.Println("You must be in school")
	} else {
		fmt.Println("You must be in college")
	}
}

func switchExample(age int) {
	switch age {
	case 10:
		fmt.Println("age <= 10")
	case 20:
		fmt.Println("age <= 20")
	case 30:
		fmt.Println("age <= 30")
	}
}

func goToExample() {
	var name string

StartOver:

	fmt.Println("You are not Harsh")
	fmt.Println("\nEnter your name")
	fmt.Scanln(&name)

	if name != "Harsh" {
		goto StartOver
	} else {
		fmt.Println("Hi Harsh!!")
	}
}

func forLoopExample(age int) {
	for i := 0; i < age; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("")
}

func forRangeExample(experience []string) {
	for i, e := range experience {
		fmt.Println("index =", i, "experience =", e)
	}
}

func cmdArgs() {
	fmt.Println("\n", os.Args[0])
	fmt.Println("argument from console is", os.Args[1])
}
