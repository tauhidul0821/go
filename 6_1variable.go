package main
import ("fmt")

var aa int;
var bb int =2;
var cc = 3
dd :=5

func main(){
	var student1 string = "john" // type is string
	var student2 = "Jane" // type is inferred 

	x:= 2 // type is inferred

	fmt.Println("this is student 1 = ",student1)
	fmt.Println("this is student 2 = ",student2)
	fmt.Println("this is X = ",x)

	fmt.Println("------------------");

	var a string;
	var b int;
	var c bool;

	fmt.Println("a is ->",a)
	fmt.Println("b is ->",b)
	fmt.Println("c is ->",c)

	fmt.Println("------------------");

	var name string;
	name = "sumons";
	fmt.Println(name);

	fmt.Println("------------------");

	fmt.Println(aa);
	fmt.Println(bb);
	fmt.Println(cc);



}