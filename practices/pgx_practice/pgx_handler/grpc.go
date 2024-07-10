package pgxHandler

import (
	"context"
	pgxInteractor "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_interactor"
	pgxMapper "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_mapper"
	pgxUseCase "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_usecase"
	proto "github.com/OddEer0/golang-practice/protogen"
	"github.com/OddEer0/golang-practice/resources/domain"
	"github.com/OddEer0/golang-practice/resources/model"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

var (
	userMapper    = pgxMapper.UserMapper{}
	postMapper    = pgxMapper.PostMapper{}
	commentMapper = pgxMapper.CommentMapper{}
)

type GrpcHandler struct {
	proto.UnimplementedNewsServiceServer
	interactor *pgxInteractor.Interactor
}

func (g *GrpcHandler) GetUserByQuery(ctx context.Context, request *proto.GetUserByQueryRequest) (*proto.ResponseManyUserAggregate, error) {
	panic("implement me")
}

func (g *GrpcHandler) DeleteUserById(ctx context.Context, id *proto.Id) (*proto.Empty, error) {
	slog.Info("request data", slog.Any("request", id))
	err := g.interactor.UserRepository.DeleteById(ctx, domain.Id(id.Value))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (g *GrpcHandler) DeletePostById(ctx context.Context, id *proto.Id) (*proto.Empty, error) {
	slog.Info("request data", slog.Any("request", id))
	err := g.interactor.PostRepository.DeleteById(ctx, domain.Id(id.Value))
	if err != nil {
		return nil, Catch(err)
	}
	return &proto.Empty{}, nil
}

func (g *GrpcHandler) UpdatePostById(ctx context.Context, request *proto.UpdatePostByIdRequest) (*proto.ResponsePost, error) {
	slog.Info("request data", slog.Any("request", request))
	res, err := g.interactor.PostRepository.UpdateById(ctx, &model.Post{
		Id:      domain.Id(request.Id),
		Title:   request.Title,
		Content: request.Content,
	})
	if err != nil {
		return nil, Catch(err)
	}
	return postMapper.PostToGrpcResponsePost(res), nil
}

func (g *GrpcHandler) GetCommentsByPostId(ctx context.Context, request *proto.GetCommentByIdRequest) (*proto.ResponseManyCommentAggregate, error) {
	slog.Info("request data", slog.Any("request", request))
	connOpt := model.CommentConns{}
	for key, opt := range request.ConnOption.Conns {
		connOpt[key] = &model.ManyOpt{
			Limit:   uint(opt.Limit),
			Page:    uint(opt.Page),
			SortDir: opt.SortDir,
			SortBy:  opt.SortBy,
		}
	}
	res, err := g.interactor.CommentUseCase.GetCommentsByPostId(ctx, domain.Id(request.Id), &model.ManyOpt{
		Limit:   uint(request.Option.Limit),
		Page:    uint(request.Option.Page),
		SortDir: request.Option.SortDir,
		SortBy:  request.Option.SortBy,
	}, connOpt)
	if err != nil {
		return nil, Catch(err)
	}

	return &proto.ResponseManyCommentAggregate{
		Comments: commentMapper.CommentAggregatesToGrpcResponseCommentAggregates(res),
	}, nil
}

func (g *GrpcHandler) GetCommentsByOwnerId(ctx context.Context, request *proto.GetCommentByIdRequest) (*proto.ResponseManyCommentAggregate, error) {
	slog.Info("request data", slog.Any("request", request))
	commentConn := model.CommentConns(ConvertGrpcConnsToPostConns(request.ConnOption))
	res, err := g.interactor.CommentUseCase.GetCommentsByOwnerId(ctx, domain.Id(request.Id), &model.ManyOpt{
		Limit:   uint(request.Option.Limit),
		Page:    uint(request.Option.Page),
		SortDir: request.Option.SortDir,
		SortBy:  request.Option.SortBy,
	}, commentConn)
	if err != nil {
		return nil, Catch(err)
	}

	return &proto.ResponseManyCommentAggregate{
		Comments: commentMapper.CommentAggregatesToGrpcResponseCommentAggregates(res),
	}, nil
}

func (g *GrpcHandler) UpdateCommentById(ctx context.Context, request *proto.UpdateCommentByIdRequest) (*proto.ResponseComment, error) {
	slog.Info("request data", slog.Any("request", request))
	res, err := g.interactor.CommentRepository.UpdateById(ctx, &model.Comment{})
	if err != nil {
		return nil, Catch(err)
	}
	return commentMapper.CommentToGrpcResponseComment(res), nil
}

func (g *GrpcHandler) DeleteCommentById(ctx context.Context, id *proto.Id) (*proto.Empty, error) {
	slog.Info("request data", slog.Any("request", id))
	err := g.interactor.CommentRepository.DeleteById(ctx, domain.Id(id.Value))
	if err != nil {
		return nil, Catch(err)
	}
	return &proto.Empty{}, nil
}

func (g *GrpcHandler) GetUserById(ctx context.Context, request *proto.GetUserByIdRequest) (*proto.ResponseUserAggregate, error) {
	slog.Info("request data", slog.Any("request", request))
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
	slog.Info("request data", slog.Any("request", request))
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
	slog.Info("request data", slog.Any("request", request))
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
	slog.Info("request data", slog.Any("request", request))
	postConns := ConvertGrpcConnsToPostConns(request.ConnOption)
	postAggregates, err := g.interactor.PostUseCase.GetPostsByUserId(ctx, domain.Id(request.UserId), &model.ManyOpt{
		Limit:   uint(request.Option.Limit),
		Page:    uint(request.Option.Page),
		SortDir: request.Option.SortDir,
		SortBy:  request.Option.SortBy,
	}, postConns)
	if err != nil {
		return nil, Catch(err)
	}

	return postMapper.PostAggregatesToGrpcResponseManyResponsePost(postAggregates), nil
}

func (g *GrpcHandler) GetPostById(ctx context.Context, request *proto.GetPostByIdRequest) (*proto.ResponsePostAggregate, error) {
	slog.Info("request data", slog.Any("request", request))
	postConn := ConvertGrpcConnsToPostConns(request.ConnOption)
	postAggregate, err := g.interactor.PostUseCase.GetPostById(ctx, domain.Id(request.Id), postConn)
	if err != nil {
		return nil, Catch(err)
	}

	return postMapper.PostAggregateToResponsePostAggregate(&postAggregate), nil
}

func (g *GrpcHandler) CreatePost(ctx context.Context, request *proto.CreatePostRequest) (*proto.ResponsePost, error) {
	slog.Info("request data", slog.Any("request", request))
	res, err := g.interactor.PostRepository.Create(ctx, &model.Post{
		Id:        domain.Id(uuid.New().String()),
		Content:   request.Content,
		OwnerId:   domain.Id(request.OwnerId),
		Title:     request.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, Catch(domain.ErrInternal)
	}

	return postMapper.PostToGrpcResponsePost(res), nil
}

func (g *GrpcHandler) CreateComment(ctx context.Context, request *proto.CreateCommentRequest) (*proto.ResponseComment, error) {
	slog.Info("request data", slog.Any("request", request))
	comment, err := g.interactor.CommentRepository.Create(ctx, &model.Comment{
		Id:        domain.Id(uuid.New().String()),
		OwnerId:   domain.Id(request.OwnerId),
		PostId:    domain.Id(request.PostId),
		Message:   request.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, Catch(domain.ErrInternal)
	}
	return &proto.ResponseComment{
		Id:      string(comment.Id),
		Message: comment.Message,
	}, nil
}

func NewGrpcHandler(interactor *pgxInteractor.Interactor) proto.NewsServiceServer {
	return &GrpcHandler{
		interactor: interactor,
	}
}
