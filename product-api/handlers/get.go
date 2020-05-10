package handlers

import (
	"net/http"
	"shop.com/v1/data"
)

// swagger:route GET /shop/v1/products listProducts
// Returns a list of products
// responses:
//    200:productsResponse
// GetProducts returns the list of all products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	//get all the available products
	lp := data.GetProducts()
	//convert to json
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to unmarshal data", http.StatusInternalServerError)
	}
}
