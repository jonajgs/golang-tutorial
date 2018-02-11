package main

import (
    "net/http"
    "encoding/json"
)

type Curso struct {
    title string `json:"titulo"`
    videosCount int `json:"numero_videos"`
}
type Cursos []Curso

func main() {
    http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
        cursos := Cursos{
            Curso{"Curso de go", 30},
            Curso{"Curso de ruby", 30},
            Curso{"Curso de swift", 30},
        }

        json.NewEncoder(response).Encode(cursos)
    })

    http.ListenAndServe(":8000", nil)
}
