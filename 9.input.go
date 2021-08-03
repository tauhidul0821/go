package main
import ("fmt")

func main(){
	var i int
	var j string


	fmt.Print("What is your name : ")
	fmt.Scan(&j)

	fmt.Println()


	fmt.Print("What is your age: ")

	fmt.Scan(&i)

	fmt.Println()

	fmt.Println("your name ", j,"and you are", i ,"Years old")

}