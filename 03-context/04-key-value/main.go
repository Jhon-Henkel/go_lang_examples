package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "1234")
	bookHotel(ctx, "Hilton")
}

// variaveis de contexto por convenção são em primeiro
func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
