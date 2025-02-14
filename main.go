package main

import "fmt"

/* 
A conjectura de Collatz é uma conjectura matemática que recebeu este nome em referência 
ao matemático alemão Lothar Collatz, que foi o primeiro a propô-la, em 1937.

Esta conjectura aplica-se a qualquer número natural inteiro, e diz-nos para, se este número 
for par, o dividir por 2 (/2), e se for impar, para multiplicar por 3 e adicionar 1 (*3+1). 
Desta forma, por exemplo, se a sequência iniciar com o número 5 ter-se-á: 5; 16; 8; 4; 2; 1. 
A conjectura apresenta uma regra dizendo que, qualquer número natural inteiro[1], quando 
aplicado a esta regra, eventualmente sempre chegará a 4, que se converte em 2 e termina em 1
*/



func collatz(number, steps int)(int){

	fmt.Println("numero atual: ", number)
	if number == 1{

		return steps
	}
	if number % 2 == 0{
		
		return collatz(number/2, steps+1)
	}

		return collatz((number*3)+1, steps+1)
}




func main(){


	resultado:= collatz(9876, 0)
    fmt.Println("passos:", resultado)
	
	
}