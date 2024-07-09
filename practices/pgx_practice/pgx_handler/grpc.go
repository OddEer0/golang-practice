package pgxHandler

import (
	"context"
	"errors"
	pgxInteractor "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_interactor"
	pgxMapper "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_mapper"
	pgxUseCase "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_usecase"
	proto "github.com/OddEer0/golang-practice/protogen"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	proto.UnimplementedNewsServiceServer
	interactor *pgxInteractor.Interactor
}

var (
	userMapper = pgxMapper.UserMapper{}
	postMapper = pgxMapper.PostMapper{}
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

func (g *GrpcHandler) GetUserById(ctx context.Context, request *proto.GetUserByIdRequest) (*proto.ResponseUserAggregate, error) {
	connOpt := model.UserConns{}
	for key, opt := range request.ConnOption.Conns {
		connOpt[key] = &model.ManyOpt{
			Limit:   uint(opt.Limit),
			Page:    uint(opt.Page),
			SortDir: opt.SortDir,
			SortBy:  opt.SortBy,
		}
	}
	userAggregate, err := g.interactor.UserUseCase.GetUserById(ctx, domain.Id(request.Id), connOpt)
	if err != nil {
		return nil, err
	}

	return userMapper.PureUserAggregateToGrpcResponseUserAggregate(&userAggregate), nil
}

func (g *GrpcHandler) CreateUser(ctx context.Context, request *proto.CreateUserRequest) (*proto.ResponseUser, error) {
	newUser, err := g.interactor.UserUseCase.Create(ctx, &pgxUseCase.CreateUserData{
		Login:    request.Login,
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		return nil, Catch(err)
	}

	return userMapper.PureUserToGrpcResponseUser(&newUser), nil
}

func (g *GrpcHandler) UpdateUserLogin(ctx context.Context, request *proto.UpdateUserLoginRequest) (*proto.ResponseUser, error) {
	pureUser, err := g.interactor.UserUseCase.UpdateUserLogin(ctx, domain.Id(request.Id), request.NewLogin)
	if err != nil {
		return nil, Catch(err)
	}

	return userMapper.PureUserToGrpcResponseUser(&pureUser), nil
}

func ConvertGrpcConnsToPostConns(option *proto.ConnOption) model.PostConns {
	res := model.PostConns{}
	for key, val := range option.Conns {
		res[key] = &model.ManyOpt{
			Limit:   uint(val.Limit),
			Page:    uint(val.Page),
			SortDir: val.SortDir,
			SortBy:  val.SortBy,
		}
	}

	return res
}

func (g *GrpcHandler) GetPostsByUserId(ctx context.Context, request *proto.GetPostsByUserIdRequest) (*proto.ResponseManyResponsePost, error) {
	postConn := ConvertGrpcConnsToPostConns(request.ConnOption)
	postAggregates, err := g.interactor.PostUseCase.GetPostsByUserId(ctx, domain.Id(request.UserId), postConn)
	if err != nil {
		return nil, Catch(err)
	}

	return postMapper.PostAggregatesToGrpcResponseManyResponsePost(postAggregates), nil
}

func (g *GrpcHandler) GetPostById(ctx context.Context, request *proto.GetPostByIdRequest) (*proto.ResponsePostAggregate, error) {
	postConn := ConvertGrpcConnsToPostConns(request.ConnOption)
	postAggregate, err := g.interactor.PostUseCase.GetPostById(ctx, domain.Id(request.Id), postConn)
	if err != nil {
		return nil, Catch(err)
	}

	return postMapper.PostAggregateToResponsePostAggregate(postAggregate), nil
}

func NewGrpcHandler(interactor *pgxInteractor.Interactor) proto.NewsServiceServer {
	return &GrpcHandler{
		interactor: interactor,
	}
}
