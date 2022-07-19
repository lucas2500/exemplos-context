package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Este tipo de context permite que o time de execução do programa seja definido, enviando
	// um sinal de cancelamento quando o horário de finalização da execução for alcançado

	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(time.Second*10),
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
			time.Sleep(time.Second)
			fmt.Println("Iteração:", counter)
			counter++
		}
	}
}
