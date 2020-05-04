package data

import "testing"

func TestCheckValidation(t *testing.T)  {
	   p:=&Product{
	   	Name: "titus",
	   	Price: 656,
	   	SKU: "abc-abd-ast",
	   }
	   err:=p.Validate()
	if err!=nil {
		t.Fatal(err)
	}
}