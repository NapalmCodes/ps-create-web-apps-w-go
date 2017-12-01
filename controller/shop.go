package controller

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"psgoweb/model"
	"psgoweb/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
	productTemplate  *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
	http.HandleFunc("/products/", s.handleProduct)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		w.Header().Add("Content-Type", "text/html")
		s.shopTemplate.Execute(w, vm)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	vm := viewmodel.NewShopDetail(products)
	w.Header().Add("Content-Type", "text/html")
	s.categoryTemplate.Execute(w, vm)
}

func (s shop) handleProduct(w http.ResponseWriter, r *http.Request) {
	productPattern, _ := regexp.Compile(`/products/(\d+)`)
	matches := productPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		productID, _ := strconv.Atoi(matches[1])
		product, err := model.GetProduct(productID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		vm := viewmodel.NewProduct(product)
		w.Header().Add("Content-Type", "text/html")
		s.productTemplate.Execute(w, vm)

	} else {
		http.NotFound(w, r)
	}
}
