package main

import (
	"fmt"
	"html/template"
	"os"
)

type dataPenis struct {
	Name string
	Age  int
}

func main() {

	dataNome := dataPenis{Name: "Vitor", Age: 12}

	fmt.Println("Experimental rodando!")
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, dataNome)
}
