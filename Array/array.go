package main

import "fmt"

func main() {
	var words1 []string

	var hello string = "Hello "
	var im string = "I'm "
	var getting string = "getting "

	words1 = append(words1, hello)
	words1 = append(words1, im)
	words1 = append(words1, getting)



	var words2 = [3]string{ "World! ", "just ", "started." }

	var phrase string = words1[0] + words2[0] + words1[1] + words2[1] + words1[2] + words2[2]

	fmt.Println(phrase)

}