package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	model "./models"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	page := model.Page{ID: 1, Name: "Ürünler", Description: "Ürün Listesi", URI: "/products"}
	products := loadProducts()

	var newProducts []model.Product

	for _, product := range products {
		/* for _, interestMapping := range interestMapping {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		} */
		newProducts = append(newProducts, product)
	}

	viewModel := model.ProductViewModel{Page: page, Products: newProducts}
	t, _ := template.ParseFiles("template/page.html")
	t.Execute(w, viewModel)

}

func loadProducts() []model.Product {
	bytes, _ := ioutil.ReadFile("josn/products.json")
	var products []model.Product
	json.Unmarshal(bytes, &products)
	return products
}
