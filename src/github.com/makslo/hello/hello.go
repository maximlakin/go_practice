package main

// import (
//           "fmt"

//           "github.com/makslo/newmath"
// )

// func main() {
//   fmt.Printf("Hello, world. Sqrt(2)=%v\n", newmath.Sqrt(2))
// }

import (
    "fmt"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, "Hello!")
}

func main() {
    var h Hello
    http.ListenAndServe("localhost:4000", h)
}