package main

import "fmt"

const tempFusaoK = 273
const tempEbulicaoK = 373

func main() {
	tempEbulicaoC := tempEbulicaoK - 273
	tempFusaoC := tempFusaoK - 273

	fmt.Println("A temperatura de ebulição em graus Celsius é: ", tempEbulicaoC)
	fmt.Println("A temperatura de fusão em graus Celsius é: ", tempFusaoC)
}
