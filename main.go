package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	model "./models"
	"github.com/fatih/color"
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
