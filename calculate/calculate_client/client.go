package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"calculator/calculate/calculatepb"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	a      *float64
	b      *float64
	method *string
	prec   *uint
)

func init() {
	a = flag.Float64("a", 1, "First Number")
	b = flag.Float64("b", 1, "Second Number")
	method = flag.String("method", "add", "Operator to use (add, sub, mult, div, sqd, root)")
	prec = flag.Uint("prec", 3, "Precision for rounding (round to nth digit)")
}

func main() {
	flag.Parse()

	cc, err := grpc.Dial("server:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	cmethod := mapMethod(method)
	c := calculatepb.NewCalculationServiceClient(cc)
	rpcCalc(c, float64(*a), float64(*b), cmethod, *prec)
}

//Maps a method as string to a Calculation_method choice
func mapMethod(method *string) calculatepb.Calculation_Method {

	//TODO: Maybe use Enum for client side too, e.g. go.chromium.org/luci/common/flag/flagenum
	switch *method {
	case "add":
		return calculatepb.Calculation_ADD
	case "sub":
		return calculatepb.Calculation_SUB
	case "mult":
		return calculatepb.Calculation_MULT
	case "div":
		return calculatepb.Calculation_DIV
	case "sqd":
		return calculatepb.Calculation_SQD
	case "root":
		return calculatepb.Calculation_ROOT
	default:
		fmt.Printf("Invalid method: %v. Default to 'add'\n", *method)
		return calculatepb.Calculation_ADD
	}
}

// Run gRPC request by providing a client, two numbers, a method and preicsion. Receive a response and print it to Stdout
func rpcCalc(c calculatepb.CalculationServiceClient, a float64, b float64, cmethod calculatepb.Calculation_Method, prec uint) {
	req := &calculatepb.CalculationRequest{
		Calc: &calculatepb.Calculation{
			A:         a,
			B:         b,
			Method:    cmethod,
			Precision: &wrappers.UInt32Value{Value: uint32(prec)}, //Precision is optional but we always pass it
		},
	}

	res, err := c.Calculation(context.Background(), req)

	if err != nil {
		fmt.Println("Problem with request!")
		st, _ := status.FromError(err)
		log.Fatalf("Error code: %v, Error Message: %v", st.Code(), st.Message())
	}

	fmt.Printf("Response: %v\n", res.Result)
}
