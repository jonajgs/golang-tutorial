package main

import "fmt"

func main() {
    var arreglo [10]int

    fmt.Println(arreglo)

    arreglo2 := [3]int{1,2}
    fmt.Println(arreglo2)

    fmt.Printf("La cantidad es %d\n", len(arreglo))

    // arreglos mutidimensionales

    var matriz [2][3]int

    fmt.Println(matriz)
}
