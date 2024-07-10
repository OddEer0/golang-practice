package pgxHandler

import (
	"errors"
	"github.com/OddEer0/golang-practice/resources/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Catch(err error) error {
	var domainErr *domain.Error
	if errors.As(err, &domainErr) {
		code := codes.Internal
		switch domainErr.Code {
		case domain.ErrInternalCode:
			code = codes.Internal
		case domain.ErrNotFoundCode:
			code = codes.NotFound
		}
		return status.Error(code, domainErr.Message)
	}
	return status.Error(codes.Internal, err.Error())
}
