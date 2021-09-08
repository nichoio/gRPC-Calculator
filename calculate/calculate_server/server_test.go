package main

import (
	"calculator/calculate/calculatepb"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Test basic math operations
func TestCalc(t *testing.T) {

	tables := []struct {
		a      float64
		b      float64
		method calculatepb.Calculation_Method
		result float64
	}{
		{1, 2, calculatepb.Calculation_ADD, 3},
		{1.5, 1.5, calculatepb.Calculation_ADD, 3},
		{-1, 1000, calculatepb.Calculation_ADD, 999},
		{1, 2, calculatepb.Calculation_SUB, -1},
		{1.5, 1.5, calculatepb.Calculation_SUB, 0},
		{2132422, 446565, calculatepb.Calculation_SUB, 1685857},
		{0, 2, calculatepb.Calculation_MULT, 0},
		{1.5, 1.5, calculatepb.Calculation_MULT, 2.25},
		{-433343, 2.4567, calculatepb.Calculation_MULT, -1064593.7481},
		{4, 2, calculatepb.Calculation_DIV, 2},
		{1000000, -1.25, calculatepb.Calculation_DIV, -800000},
		{3.8, 4.5, calculatepb.Calculation_DIV, 0.8444444444444444},
		{2, 3, calculatepb.Calculation_SQD, 8},
		{15, 8, calculatepb.Calculation_SQD, 2562890625},
		{2, 2, calculatepb.Calculation_ROOT, 1.4142135623730951},
		{123456, 6, calculatepb.Calculation_ROOT, 7.056435349859053},
	}

	s := server{}

	for _, table := range tables {
		result, err := s.calc(table.a, table.b, table.method)

		if err != nil {
			t.Errorf("Error: %v", err)
		} else if result != table.result {
			t.Errorf("Incorrect. Got %v instead of %v", result, table.result)
		}
	}
}

//Test if division or root by zero raises expected error
func TestCalcError(t *testing.T) {

	tables := []struct {
		a      float64
		b      float64
		method calculatepb.Calculation_Method
		code   codes.Code
	}{
		{2, 0, calculatepb.Calculation_DIV, codes.InvalidArgument},
		{2, 0, calculatepb.Calculation_ROOT, codes.InvalidArgument},
	}

	s := server{}

	for _, table := range tables {
		result, err := s.calc(table.a, table.b, table.method)
		st, _ := status.FromError(err)

		if st.Code() != table.code {
			t.Errorf("Unexpected error: Got %v instead of %v", err, table.code)
		}
		if result != 0 {
			t.Errorf("Incorrect. Got %v instead of 0", result)
		}
	}
}

//Test basic rounding
func TestToFixed(t *testing.T) {

	tables := []struct {
		input     float64
		output    float64
		precision int
	}{
		{2.424242, 2.4, 1}, {2.424242, 2, 0},
		{-1234.5678, -1234.57, 2}, {-1234.5678, -1234.568, 3},
		{-1234.99, -1235, 1}, {-1234.99, -1235, 0},
	}

	for _, table := range tables {
		output := toFixed(table.input, table.precision)
		if output != table.output {
			t.Errorf("Incorrect. Got %v instead of %v", output, table.output)
		}
	}
}
