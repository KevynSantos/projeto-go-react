package main

import "fmt"

func main() {
	saudacao := ola("Kevyn")
	fmt.Println(saudacao)
}

func ola(nome string) string {
	return fmt.Sprintf("Olá, %s! Seja bem-vindo ao Go!", nome)
}
