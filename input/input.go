package input

import (
	"bufio"
	"fmt"

	"os"
	"strings"
)
// cheacked two inputs on purpose just for flexibility

//inputuser was used to take inputs for all struct string data
func InputUser()string{
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil{
		panic(err)
	}
	input = strings.TrimSpace(input)
	return input
} 
//InputUser2 was used to take input for all struct int data
//It was also used for taking input for the number of students to be registed
func InputUser2() int{
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	return input
}
