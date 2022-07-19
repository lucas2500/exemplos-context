package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Este tipo de context permite que um timeout para a execução do programa seja definido

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)

	defer cancel()

	PrintUntilCancel(ctx)
}

func PrintUntilCancel(ctx context.Context) {

	counter := 0

	for {

		select {
		case <-ctx.Done():
			fmt.Println("Solicitação de cancelamento recebida, saindo...")
			return

		default:
			time.Sleep(time.Second * 2)
			fmt.Println("Iteração:", counter)
			counter++
		}
	}
}
