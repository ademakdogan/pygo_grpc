package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "github.com/ademakdogan/pygo_grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	xValue := flag.String("x", "1", "x_value")
	yValue := flag.String("y", "1", "y_value")
	flag.Parse()
	addr := "localhost:9001"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewMathClient(conn)
	start := time.Now()
	intX, _ := strconv.ParseInt(*xValue, 10, 32)
	intY, _ := strconv.ParseInt(*yValue, 10, 32)
	req := &pb.MathRequest{

		X: int32(intX),
		Y: int32(intY),
	}

	resp, err := client.OrderMath(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

	duration := time.Since(start)
	fmt.Println(duration)
}
