package main

import (
	"context"
	"medblogers_base/internal"
)

func main() {
	ctx := context.Background()

	internal.New(ctx).Run(ctx)
}
