package handler

import (
	"context"
	pgxInteractor "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_interactor"
	proto "github.com/OddEer0/golang-practice/protogen"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
)

type GrpcHandler struct {
	proto.UnimplementedNewsServiceServer
	interactor *pgxInteractor.Interactor
}

func (g GrpcHandler) GetUserById(ctx context.Context, request *proto.GetUserByIdRequest) (*proto.ResponseUserAggregate, error) {
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

	posts := make([]*proto.ResponsePost, 0, len(userAggregate.Posts))
	for _, post := range userAggregate.Posts {
		posts = append(posts, &proto.ResponsePost{
			Id:      string(post.Id),
			Title:   post.Title,
			Content: post.Content,
		})
	}

	return &proto.ResponseUserAggregate{
		Value: &proto.ResponseUser{
			Id:    string(userAggregate.Value.Id),
			Login: userAggregate.Value.Login,
			Email: userAggregate.Value.Email,
		},
		Posts: posts,
	}, nil
}

func (g GrpcHandler) CreateUser(ctx context.Context, request *proto.CreateUserRequest) (*proto.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcHandler) UpdateUserLogin(ctx context.Context, request *proto.UpdateUserLoginRequest) (*proto.ResponseUser, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcHandler) GetPostsByUserId(ctx context.Context, request *proto.GetPostsByUserIdRequest) (*proto.ResponseManyResponsePost, error) {
	//TODO implement me
	panic("implement me")
}

func (g GrpcHandler) GetPostById(ctx context.Context, request *proto.GetPostByIdRequest) (*proto.ResponsePostAggregate, error) {
	//TODO implement me
	panic("implement me")
}

func NewGrpcHandler(interactor *pgxInteractor.Interactor) proto.NewsServiceServer {
	return &GrpcHandler{
		interactor: interactor,
	}
}
