package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"shop.com/v1/data"
)

//swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	//in:body
	Body []data.Product
}
type Products struct {
	l *log.Logger
}
type KeyProduct struct{}
	func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
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
