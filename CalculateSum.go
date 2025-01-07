package main
import "fmt"
//this explain the local and Global variables
var total int //this is Global Variable
func main(){
    var sum int
	{

	
	count:=5
	sum = calculateSum(&count,2)
	fmt.Println(sum)
	}
	total =sum
	{
		count:=10
		sum = calculateSum(&count,3)
		fmt.Println(sum)
	}
	total := total + sum
  fmt.Println("Total:", total)
}
func calculateSum(count *int,num int)int{
    
	sum:=*count*num
	return sum
}