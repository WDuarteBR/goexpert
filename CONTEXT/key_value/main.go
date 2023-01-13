package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "token", "123pws")
	lerContext(ctx)

}

func lerContext(ctx context.Context) { // obs.: sempre o context dever ser o primeiro parâmetro
	token := ctx.Value("token")
	fmt.Println(token)
}
