package routing
import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Createhandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Created!")
}

func Route() {
	http.HandleFunc("/Test", handler)
	http.HandleFunc("/Create", Createhandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}