package main

import (
	"fmt"
	"sync"
	"runtime"
)

/*
Goroutines: Linhas encadeamentos de funções
ex: varios sistemas separados rodando independentemente;
WaitGroups:
wg  sync.WaitGroup == Grupo de Espera;
func Add(){} : Define quantas goroutines
func Done(){}: Finilizar o processo
func Wait(){}: Aguarda a finalização das goroutines
ex: esperar os sistemas serem executados
*/
var(
	waitGroup sync.WaitGroup
)

func main(){
	//printNumCPUs();
	//printGoroutines();
	waitGroup.Add(2);
	go rocket();
	go whoami();
	//printGoroutines();
	waitGroup.Wait()
}

func rocket(){
	fmt.Printf("🚀");
	waitGroup.Done();
}
func printNumCPUs(){fmt.Println("[CPUs]:",runtime.NumCPU());}
func printGoroutines(){fmt.Println("[Goroutines]:",runtime.NumGoroutine());}
func whoami(){
	name := "";
	fmt.Printf("Digite seu nome:");
	fmt.Scanf("%s",&name);
	fmt.Println("Hello",name);
	waitGroup.Done();
} 