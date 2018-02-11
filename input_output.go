package main

import(
    "fmt"
    "bufio"
    "os"
)

func main() {
    // edad := 24
    // nombre := "Jonathan"
    // hijos := false
    // sueldo := 10000.0
    //
    // fmt.Printf("Mi edad ed %d y mi nombre es %v, y dicen que tengo hijos pero es %t y mi sueldo es de %.2f\n", edad, nombre, hijos, sueldo)

    var edad int
    var nombre string

    fmt.Print("Cual es tu edad: ")
    fmt.Scanf("%d", &edad)

    fmt.Print("Cual es tu nombre: ")
    fmt.Scanf("%v", &nombre)


    fmt.Printf("Mi edad ed %d y mi nombre es %v \n", edad, nombre)

    // with bufio
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Ingresa tu nombre de usuario: ")
    username,err := reader.ReadString('\n')

    if(err != nil) {
        fmt.Println(err)
    } else {
        fmt.Println("Hola " + username)
    }

}
