package main
import ("fmt")

// type constants
const A int = 1

// untyped Constants
const B = "Sumons"

// Multiple constants declaration
const (
	AA = 1
	BB = 3.1415
	CC = "Hi!"
)

func main(){
	fmt.Println(A)
	fmt.Println(B)

	fmt.Println("------------")
	
	fmt.Println(AA)
	fmt.Println(BB)
	fmt.Println(CC)
}