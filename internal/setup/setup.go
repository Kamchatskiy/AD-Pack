package setup

import (
	"context"
)

type ctxStr string

func Context() (context.Context, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxStr("LogoEnabled"), true)
	return ctx, nil
}
