package main

import "fmt"

type User struct {
    edad int
    nombre,apellido string
}

func(usuario User) nombre_completo() string {
    return usuario.nombre + " " + usuario.apellido
}

func(this *User) setName(nombre string) {
    this.nombre = nombre
}

// campo anonimos

type Human struct {
    name string
}

type tutor struct {
    Human
}

func main() {

    uriel := new(User)

    uriel.nombre = "Uriel"
    uriel.apellido = "Gutierrez"

    uriel.setName("Marcos")
    fmt.Println(uriel.nombre_completo())

}
