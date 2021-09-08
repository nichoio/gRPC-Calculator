package main

import (
	"calculator/calculate/calculatepb"
	"testing"
)

//Test mapping from strings to Calculation enum
func TestMapMethod(t *testing.T) {

	tables := []struct {
		str    string
		method calculatepb.Calculation_Method
	}{
		{"add", calculatepb.Calculation_ADD},
		{"sub", calculatepb.Calculation_SUB},
		{"mult", calculatepb.Calculation_MULT},
		{"div", calculatepb.Calculation_DIV},
		{"sqd", calculatepb.Calculation_SQD},
		{"root", calculatepb.Calculation_ROOT},
	}

	for _, table := range tables {
		method := mapMethod(&table.str)
		if method != table.method {
			t.Errorf("Incorrect. Got %v instead of %v", method, table.method)
		}
	}
}
