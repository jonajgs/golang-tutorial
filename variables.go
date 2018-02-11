package main

import "fmt"
import "strconv"

func main(){

    edad_string := "24"

    edad,_ := strconv.Atoi(edad_string)

    edad_str := strconv.Itoa(edad)


    fmt.Println("Mi edad es " + edad_str)
}
