package main

import (
    "net/http"

    "github.com/gorilla/mux"
	"time"
	"log"
	"encoding/json"
)

func healthHandler(w http.ResponseWriter, r*http.Request){
	log.Println("Checking application health")
	response := map[string]string{
		"status": "HiJacked",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}
func main() {
    r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}



// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	title := vars["title"]
// 	page := vars["page"]

// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// 	// type in: http://localhost/books/TheGreatGatsby/page/42
// })

// package main

// import (
//     "fmt"
// 	"log"
//     "net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
//     http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
//     http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("Web server has started")
//     http.ListenAndServe(":80", nil)
// }











// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	geo "github.com/Brad-99/go-microservices/geometry"

// 	"rsc.io/quote"
// )

// func rectProps(length, width float64) (area1, perimeter float64) {
// 	area1 = length * width
// 	perimeter = 2 * (length + width)
// 	return
// }

// func main() {
// 	x := 10
// 	var y, z = 2, 3
// 	var name string
// 	isWorking := false

// 	name = "Puppy"
// 	fmt.Println("Test")
// 	fmt.Println(quote.Go())
// 	fmt.Println(x, y, z, name, isWorking)
// 	fmt.Println("Type of name %T and size is %d\n", name, unsafe.Sizeof(name))

// 	a, p := rectProps(1, 2)
// 	fmt.Println("Area is %f and perimeter is %f", a, p)

// 	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
// 	fmt.Println(daysOfTheMonth)

// 	area := geo.Area(1, 2)
// 	diag := geo.Diagonal(1, 2)
// 	fmt.Println(area, diag)

// }
