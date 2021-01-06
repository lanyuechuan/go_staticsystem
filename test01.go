package main
import "net/http"

func main() {
    fs := http.FileServer(http.Dir("../"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8000", nil)
}
