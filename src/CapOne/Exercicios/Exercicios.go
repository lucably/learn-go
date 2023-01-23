//Exercicios referentes aos videos 16 à 21
package main

import "fmt"

//Exercicio: Nivel #1 - 1
/*
func main()  {
	x := 42;
	y := "James bound";
	z := true;

	fmt.Println(x, y, z);
	fmt.Println(x);
	fmt.Println(y);
	fmt.Println(z);
}
*/

//Exercicio: Nivel #1 - 2
/*
var x int;
var y string;
var z bool;

func main()  {

	fmt.Printf("%v\n %v\n %v\n",x, y, z);

}
*/

//Exercicio: Nivel #1 - 3
/*
var x int  = 42;
var y string  = "James Bond";
var z bool = true;

func main()  {

	s := fmt.Sprint(x, "\t", y, "\t", z, "\n");
	fmt.Printf(s);

}
*/
//Exercicio: Nivel #1 - 4
/*
type idade int 
var x idade;

func main()  {
	fmt.Printf("%v\n %T\n ", x, x);

	x  = 42;
	fmt.Println(x);
}

*/

//Exercicio: Nivel #1 - 5

type idade int 
var x idade;
var y int;
func main()  {
	x  = 42;
	y = int(x);
	fmt.Println(y);
}
