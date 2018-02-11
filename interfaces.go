package main

import "fmt"

type User interface {
    Permissions() int
    Name() string
}

type Admin struct {
    nombre string
}

func (this Admin) Permissions() int {
    return 5
}

func (this Admin) Name() string {
    return this.nombre
}

func auth(user User) string {
    if user.Permissions() > 4 {
        return user.Name() + " tiene permisos de administrador"
    }

    return ""
}

func main( ) {
    admin := Admin{"Uriel"}

    fmt.Println(auth(admin))
}
