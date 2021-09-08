package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"calculator/calculate/calculatepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const precision = 3

type server struct {
	calculatepb.UnimplementedCalculationServiceServer
}

func main() {
	fmt.Println("Server is running ...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculationServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//Implement gRPC Calculation method
func (serv *server) Calculation(ctx context.Context, req *calculatepb.CalculationRequest) (*calculatepb.CalculationResponse, error) {
	fmt.Printf("Calculation invoked with %v\n", req)

	a := req.GetCalc().GetA()
	b := req.GetCalc().GetB()
	method := req.GetCalc().GetMethod()
	precision := precision

	if prec := req.GetCalc().GetPrecision(); prec != nil {
		precision = int(prec.Value)
	}

	out, err := serv.calc(a, b, method)
	return &calculatepb.CalculationResponse{Result: toFixed(out, precision)}, err
}

// Perform Calculation for two numbers, and an operator
func (serv *server) calc(a float64, b float64, method calculatepb.Calculation_Method) (float64, error) {

	// This produces IEEE rounding errors like 0.1 * 0.1 = 0.010000000000000002
	// Could probably be addressed with math/big
	// For our purposes, rounding to nth digit is fine

	switch method {
	case calculatepb.Calculation_ADD:
		return a + b, nil
	case calculatepb.Calculation_SUB:
		return a - b, nil
	case calculatepb.Calculation_MULT:
		return a * b, nil
	case calculatepb.Calculation_DIV:
		if b == 0 {
			return 0, status.Error(codes.InvalidArgument, "Forbidden: Division by 0")
		}
		return a / b, nil
	case calculatepb.Calculation_SQD:
		return math.Pow(a, b), nil
	case calculatepb.Calculation_ROOT:
		if b == 0 {
			return 0, status.Error(codes.InvalidArgument, "Forbidden: 0th root undefined")
		}
		return math.Pow(a, 1/b), nil
	default:
		return 0, status.Error(codes.InvalidArgument, "Forbidden: Unknown Operator")
	}
}

//Round a number by given precision and return it
func toFixed(num float64, precision int) float64 {
	// From https://stackoverflow.com/questions/18390266/how-can-we-truncate-float64-type-to-a-particular-precision
	output := math.Pow(10, float64(precision))
	return float64(int(num*output+math.Copysign(0.5, num*output))) / output
}
