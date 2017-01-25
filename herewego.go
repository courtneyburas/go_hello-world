package main

import "fmt"

func main() {
	fmt.Println("Hello World") //Println references format package
}

/* go run name.go 
	how to run go file */

/* godoc fmt Println  (godoc package function)
	print out a ton of info */

/* variables - type can't change once defined
	data-type: int (integer), string, bool (boolean)
		
		const = constant
*/

var age int = 40

var (
		varA = 2
		varB = 3
	)

var favNum float64 = 1.61800

randNum := 1

var myName string = "Courtney Buras"

fmt.Println(len(myName)) // length

fmt.Println( myName + " is a robot") // must have space to concat

var isOver40 bool = true

/* loops
	if/else - first condition met, nothing below even considered, even if true!!
*/

for i <= 10 {

	fmt.Println(i)
	i++
}

for j := 0; j < 5; j++ {
	fmt.Println(j)
}

yourAge := 18

if yourAge >= 16 {
	fmt.Println("You can drive")
} else if yourAge >= { 18
	fmt.Println("You can vote")
} else {
	fmt.Println("You can have fun")
}

switch youAge { // limited number of possibilities
	case 16 : fmt.Println("Go Drive")
	case 18 : fmt.Println("Go Vote")
	default : fmt.Println("Go Have Fun")
}

// arrays

var favNum2[5] float64
				// or
favNum3 := [5]float64 {1,2,3,4,5}

for i, value := range favNum3{
	fmt.Println(value, i)
}

/* if don't want index to print, replace with underscore:

for _, value := range favNum3{
	fmt.Println(value)
} */

// slice (like array but when defined forget size

numSlice := []int {5,4,3,2,1}

numSlice2 := numeSlice[3:5] // get indexes 3 and 4, ignore 5
// if don't have first number, assume start from begin
// if dont have last number, assume start from end

numSlice3 := make([]int, 5, 10) //10 set of values, values undefined
copy(numSlice3, numSlice) //taking values from other Slice

numSlice3 = append(numSlice, 0, -1) //add onto end of slice

// map - collection of key value pairs

mapMap := make(map[string] int) // map[key] value


~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// Get started with Go by Google Developers on YouTube

// go run
// go test
// go build
// go fmt
// go get code - fetch packages from remote repos

/* must have a workspace
	sub-workspace:
	bin - executable binaries
	pkg - compiled object files
	src - source code
*/

// all Go code belongs to a package - all start with package main?
// import - declares file dependencies (like 'bundle?')

// Go can do what Nokogiri does in Ruby..



	


