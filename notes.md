## Go Programming by Derek Banas ##
8/26

var
const
bool
int (gazillion variations)
uint (gazillion variations)
float32
float64
func
package

for loops:

for i <= 10 {
	fmt.Println(i)
	i++
}

for j := 0; j < 5; j++ {
	fmt.Println(j)
}

if/else statements

if yourAge >= 16 {
	fmt.Println("You can drive")
} else if yourAge >= 18 {
	fmt.Println("You can vote")
} else {
	fmt.Println("You can have fun")
}

[only does first true statement, if there's a second true statement it won't notice]




