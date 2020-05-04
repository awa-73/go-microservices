package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"titusdishon/coffee-shop/go-microservices/product-api/data"
)

type Products struct {
	l *log.Logger
}
type KeyProduct struct{}

	func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	//get all the available products
	lp := data.GetProducts()
	//convert to json
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to unmarshal data", http.StatusInternalServerError)
	}
}
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle put Products")
	vars :=mux.Vars(r)
	id, err:=strconv.Atoi(vars["id"])
	if err!=nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
	}
	//get the data being received from the request body and convert it into json
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
      return http.HandlerFunc(func(rw http.ResponseWriter,r*http.Request){
      	prod:= data.Product{}
		err := prod.FromJSON(r.Body)
		if err!=nil {
			http.Error(rw, "Unable to unmarshal JSON", http.StatusNotFound)
			return
		}
		err=prod.Validate()
		  if err!=nil {
			  http.Error(rw, fmt.Sprintf("error validating data: %s", err), http.StatusBadRequest)
			  return
		  }
		ctx:=context.WithValue(r.Context(), KeyProduct{}, prod)
		r= r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
