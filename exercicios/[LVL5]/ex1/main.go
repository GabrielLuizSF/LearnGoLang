package main

import "fmt";

type people struct{
  nome      string
  sobrenome string
  sabores_de_sorvete []string
}


func main(){

  customer1 := people{
    nome: "João" ,
    sobrenome:"" ,
    sabores_de_sorvete: []string{"Napolitano","Flocos"},
  }
  customer2 := people{
    nome: "Maria",
    sobrenome: "",
    sabores_de_sorvete: []string{"Chocolate","Morango"},
  }
   pedidos(customer1.nome,customer1.sabores_de_sorvete);
   pedidos(customer2.nome,customer2.sabores_de_sorvete);

  }

func pedidos(name string, slice []string){
  fmt.Println("[Cliente]:",name);
  for indice , value := range slice{
    if indice == 0{
    fmt.Println("Sorvete:",value);
  }else{
    fmt.Println("Sorvete:",value,"\n ");
  }
  }
}
