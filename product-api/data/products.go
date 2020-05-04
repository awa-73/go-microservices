package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID  int `json:"id"`
	Name string  `json:"name" validate:"required"`
	Description string   `json:"description" `
	Price float32 `json:"price" validate:"gt=0"`
	SKU string `json:"sku" validate:"required,sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

func (p *Product) Validate() error  {
	validate:=validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return  validate.Struct(p)
}
func validateSKU(fl validator.FieldLevel) bool {
      re:=regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
      matches:=re.FindAllString(fl.Field().String(), -1)
	if len(matches)!=1 {
		return false
	}
	if fl.Field().String() == "invalid" {
		return false
	}
	return true
}
//convert the data to return into json
type Products[]*Product

func (p *Products) ToJSON(w io.Writer) error {
	return  json.NewEncoder(w).Encode(p)
}
//from json
func (p *Product)FromJSON(r io.Reader) error  {
	e:= json.NewDecoder(r)
	return e.Decode(p)
}
func GetProducts() Products {
   return productList
}
func AddProduct(p *Product)  {
   p.ID= getNext()
   productList =append(productList,p)
}
func UpdateProduct(id int, p *Product)error  {
   _, pos,err:= findProduct(id)
	if err!=nil {
		return err
	}
   p.ID=id
   productList[pos]=p
   return nil
}
var ErrProductNotFound = fmt.Errorf("product not found")
func findProduct(id int) (*Product,int, error) {
	for l,p:=range productList {
		if p.ID==id{
			return p,l, nil
	}
	}
	return nil,-1, ErrProductNotFound
}

func getNext() int  {
	lp:= productList[len(productList)-1]
	return lp.ID+1
}
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Really nice latte",
		Price:       200,
		SKU:         "abc232",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Mango Juice",
		Description: "Tamu mango juice",
		Price:       100,
		SKU:         "abc232",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
}