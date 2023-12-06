package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Android", "iOS", "React Native", "Go"}
	fmt.Println(courseName)
	courseName = append(courseName, "Flutter", "Kotlin", "Swift", "Objective-C")
	fmt.Println("After append", courseName)

	courseWeb := courseName[4:8]
	fmt.Println("courseWeb", courseWeb)
	courseWeb2 := courseName[:3]
	fmt.Println("courseWeb2", courseWeb2)

}
