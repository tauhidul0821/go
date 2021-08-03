package main
import ("fmt")

func main(){
	var i,j string ="a","b"
	n:=5

	fmt.Print(i)
	fmt.Print(j,"\n")

	fmt.Println("-------------")

	fmt.Println(i,j)


	fmt.Println("----type and value print---")

	fmt.Printf(" -------->n has value: %v and type: %T",n,n);
	fmt.Println();
	fmt.Printf(" -------->j has value: %v and type: %T",j,j);

	fmt.Println();


}