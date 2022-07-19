package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Este tipo de context interrompe a execução do programa ao receber o sinal de cancelamento

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	go PrintUntilCancel(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second)

}

func PrintUntilCancel(ctx context.Context) {

	counter := 0

	for {

		select {
		case <-ctx.Done():
			fmt.Println("Solicitação de cancelamento recebida, saindo...")
			return

		default:
			time.Sleep(time.Second)
			fmt.Println("Iteração:", counter)
			counter++
		}
	}
}
