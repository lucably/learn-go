package main

import "fmt"

//var x = 10 => Variavel global.

// :=  => serve para declaração de tipo.

var variavelGlobal = 55;

func DeclarationExemple() {

	//Decaração da variavel
	numero := 10;
	texto := "Hello";

	//Mostrando as variaveis.
	fmt.Printf("Numero: %v, Tipo: %T \n", numero, numero);
	fmt.Printf("Texto: %v, Tipo: %T \n", texto, texto);

	/*
		Reatribuindo de forma errada => numero := 20 (Erro pois para usar ":=" necessita
      	de pelo menos uma nova declaração para ser usada.).
	*/
	numero = 30;
	fmt.Printf("Numero: %v, Tipo: %T \n", numero, numero);

	//Reatribuindo de forma correta. Repara que declarei uma nova variavel (numero2).

	numero, numero2 := 60, 30.5;
	fmt.Printf("Numero: %v, Tipo: %T \n", numero, numero);
	fmt.Printf("Numero2: %v, Tipo: %T \n", numero2, numero2);
	fmt.Printf("variavelGlobal: %v, Tipo: %T \n", variavelGlobal, variavelGlobal);
}

func fmtExemplePrint() {
	//Sprint => Valor nao é jogado na tela e sim em uma variavel.
	x := "Hello";
	y := "World";

	result := fmt.Sprint(x, " ", y, "\n");

	fmt.Print(result);

}

type idade int 

func criandoMeuTipo() {

	var minhaIdade idade = 20 
	fmt.Print(minhaIdade, "\n")

	comparaIdade := 20 //Valor do tipo INT.

	/*
	porem se tentar comparar o tipo int com o tipo idade ele da divergencia.

		fmt.Print(minhaIdade == comparaIdade)

	para corrigir basta transformar a variavel do tipo IDADE em INT.
	*/
	fmt.Print(int(minhaIdade) == comparaIdade)

}

func main() {
	DeclarationExemple();
	fmtExemplePrint();
	criandoMeuTipo();
}