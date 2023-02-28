package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"primeDemo/prime"
)

func main() {
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := prime.NewPrimeCheckerClient(conn)
	resp, err := client.IsPrime(context.Background(), &prime.PrimeRequest{Number: 17})
	if err != nil {
		log.Fatalf("Error calling IsPrime: %v", err)
	}
	fmt.Println(resp.IsPrime) // Output: true
}
