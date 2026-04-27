package main

import (
	"fmt"
	"time"
)

// func init() {
// 	fmt.Println("This function runs even before the Main function")
// }

// func maincd() {
// 	var name string
// 	fmt.Print("Enter your name:")
// 	fmt.Scan(&name)
// 	fmt.Println("How was your day ? ", name)
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your string: ")
// 	str, _ := reader.ReadString('\n')
// 	fmt.Println(str)
// }

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter a floating point number: ")
// 	str, _ := reader.ReadString('\n')
// 	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(f)
// 	}

// }

// func main() {
// 	func() {
// 		fmt.Println("Keve a singh? darda kyu a?.. WwwaaaiinKhaalSaaaWwaainFateh!")
// 	}()

// }

// func main() {
// 	wkwf := func() {
// 		fmt.Println("Good Life saddi muteyaree!")
// 	}
// 	wkwf()
// }

// func main() {
// 	var num1 float64 = 4.5210
// 	res := float64(9) + num1
// 	fmt.Println(res)
// }

// exploring pointers

// func main() {
// 	g := 32
// 	age := 45
// 	var ptr *int = &age

// 	if ptr == nil {
// 		fmt.Println("P is nil")
// 	} else {
// 		fmt.Print(g)
// 		fmt.Println("Address : ", ptr)
// 	}
// }

// func add(b int) {
// 	b++
// 	fmt.Println(b)
// 	add(b)
// }

// import "fmt"

// func main() {

// 	name := "Swapnil_Mishra"
// 	const name2 = "Swapnil_Mishra"
// 	var g *string = &name

// 	fmt.Println(name2 == name) // true
// 	fmt.Println(g)             // Prints the memory address of 'name'

// 	i := 1

// 	for i <= 30 {
// 		i = i + 1
// 		fmt.Println(i)
// 	}

// }

// func rec(v int) int {
// 	if v == 0 {
// 		fmt.Println(v)
// 		return 1
// 	}
// 	fmt.Println(v)
// 	return rec(v - 1)
// }

// func mulRet() (int, int) {
// 	return 11, 17
// }

// func init() {
// 	fmt.Println("Let us see, if this run first or not..!")
// }

// func vSum(nums ...int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum = sum + num
// 	}
// 	return sum
// }

// func main() {
// 	for {
// 		fmt.Println("Standard infinite for loop in Go language")
// 	}
// }

// func main() {
// 	var age int
// 	fmt.Println("Enter your age: ")
// 	fmt.Scan(&age)
// 	if age >= 13 && age <= 19 {
// 		fmt.Println("Welcome to Teen Age")
// 	} else if age >= 20 {
// 		fmt.Println("You are an Adult")
// 	} else {
// 		fmt.Println("Hello kiddo!")
// 	}
// }

// func main() {
// 	var name = Shiv()()
// 	fmt.Println(name)
// }

// func Shiv() func() string {
// 	return func() string {
// 		return "Narayan"
// 	}
// }

// func parity_of_product(x int, y int) bool {
// 	result := x * y
// 	return result%2 == 0
// }

// func main() {
// 	fmt.Println(time.Now())
// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's Weekend")
// 	default:
// 		fmt.Println("It is workday")
// 	}
// }

// func main() {
// 	var tmp = 17
// 	fmt.Println(passByVal(tmp))
// 	fmt.Println(tmp)
// 	/*
// 		classic pass by value example in Go :
// 		pass by value means copy of the variable will be passed
// 		but the original variable will not change
// 		that is why the output 18 then 17
// 	*/

// }

// func passByVal(x int) int {
// 	x++
// 	return x
// }

// func numRev(x int) int {
// 	rev := 0
// 	cpx := x
// 	for cpx != 0 {
// 		rev = rev*10 + cpx%10
// 		cpx = cpx / 10
// 	}
// 	return rev
// }

// func main() {
// 	var tmp = 12
// 	fmt.Println(tmp)
// 	ptrManipulation(&tmp)
// 	fmt.Println(tmp)
// 	fmt.Println(tmp)
// }
// // simple pointer manipulation
// func ptrManipulation(ptr *int) {
// 	*ptr = 22
// }

// func main() {
// 	func() {
// 		fmt.Println("Example of an Anonymous function")
// 	}()
// }

// func main() {
// 	CaptialCities := make(map[string]string)
// 	CaptialCities["Karnataka"] = "Bengaluru"
// 	CaptialCities["Rajasthan"] = "Jaipur"
// 	CaptialCities["Kerela"] = "Thiruvananthapuram"
// 	CaptialCities["Andhra Pradesh"] = "Amaravathi"
// 	CaptialCities["Assam"] = "Dispur"
// 	fmt.Println(CaptialCities)
// }

// func main() {
// 	const (
// 		zero = iota
// 		one
// 		two
// 		three
// 		four
// 		five
// 		six
// 		seven
// 	)
// 	fmt.Println(seven + six)
// }

// func main() {
// 	const (
// 		low = iota
// 		medium
// 		average
// 	)
// 	fmt.Println((low + average) / 2)
// }

// type Laptop struct {
// 	ram   int
// 	rom   int
// 	cpu   int
// 	cache int
// 	brand string
// }

// func main() {
// 	L1 := Laptop{
// 		rom:   512,
// 		ram:   16,
// 		cache: 8,
// 		brand: "HP",
// 		cpu:   8,
// 	}
// 	fmt.Println(L1.brand)
// }

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.radius * c.radius
// }

/*
Interface polymorphs on struct types
given that the struct type implements all the methods of the interface
then the struct type is said to be implementing the interface

Basically Struct Should obey the contract of that interface
https://gobyexample.com/interfaces
*/

// func main() {
// 	circle := Circle{radius: 7}
// 	fmt.Println("Area of this Circle is", circle.Area())
// }

// func main() {
// 	res, err := otpCheck(1717)
// 	if err != nil {
// 		fmt.Println("Wrong Passcode!")
// 	} else {
// 		return res
// 	}

// }

// func otpCheck(otp int) error {
// 	if otp != 1729 {
// 		return fmt.Errorf("Wrong OTP")
// 	}
// 	return nil
// }

// Handling Errors in Go
// func main() {
// 	var otp int
// 	// Calling otpCheck and handling the result
// 	message, err := otpCheck(otp)
// 	if err != nil {
// 		fmt.Println("Error:", err) // Prints the error returned from otpCheck
// 	} else {
// 		fmt.Println(message)
// 	}
// }

// // Function to check OTP
// func otpCheck(otp int) (string, error) {
// 	// Taking user input for OTP
// 	fmt.Scanln(&otp)

// 	// Checking the OTP value
// 	if otp == 1729 {
// 		return "Code Accepted", nil
// 	}

// 	// Returning an error with fmt.Errorf for incorrect OTP
// 	return "Wrong Code", fmt.Errorf("invalid OTP, please try again")
// }

func main() {
	fmt.Println(time.Now().Clock())
}
