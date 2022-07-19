package main

import (
	"context"
	"fmt"
)

func main() {

	// Este tipo de context permite que uma chave/valor sejam definidos na criação do mesmo.
	// A chave/valor informados são imutáveis dentro do context atual, sendo necessário a criação
	// de um novo para que os mesmos possam ser alterados

	ctx := context.WithValue(
		context.Background(),
		"context1",
		"value1",
	)

	PrintContextValue(ctx)
}

func PrintContextValue(ctx context.Context) {

	fmt.Println("Context 1:", ctx.Value("context1"))

	NewCtx := context.WithValue(
		ctx,
		"context2",
		"value2",
	)

	fmt.Println("Context 2:", NewCtx.Value("context2"))
}
