package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Product struct {
	ID          int
	Name        string
	Category    string
	Summary     string
	Description string
	Price       float64
}

func main() {

	apiRoot := "/api"

	http.HandleFunc(apiRoot+"/products", func(w http.ResponseWriter, r *http.Request) {
		products := []Product{
			Product{ID: 1, Name: "Telefon", Category: "Teknoloji", Summary: "asd", Description: "asd", Price: 200},
			Product{ID: 2, Name: "Kazak", Category: "Giyim", Summary: "asd", Description: "asd", Price: 15},
			Product{ID: 3, Name: "Laptop", Category: "Teknoloji", Summary: "asd", Description: "asd", Price: 200},
		}

		message := products
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))

	})

	http.ListenAndServe(":8080", nil)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
}
