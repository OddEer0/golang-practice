package domain

type Logger interface {
	Info(msg string, attrs ...any)
}
