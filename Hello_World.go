package main
import "fmt" //Comment
// this is how function is created 
func PrintMessage(st string)string {
	return st
}
func main(){
	fmt.Println("Hello World!")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	fmt.Println(PrintMessage("Hello"))
	//this is how delecarled the variable
	var number float32 =5.2
	fmt.Println(number)
	fmt.Println(int(number))
	//for const dont need the string
	const B = "hello"
    //here ,however we can write explit string
	const C string = "World"
	fmt.Println(B)
	fmt.Println(C)
	const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int= 1, 2, 3, 4, 5, 6
	var decision bool       // Declaring a boolean variable
    fmt.Println(decision)
	//this is short form
	count:="Hello"
	fmt.Println(count)
	
}
