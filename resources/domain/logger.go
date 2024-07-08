package domain

import "context"

type Logger interface {
	Info(ctx context.Context)
}
