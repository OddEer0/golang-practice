package pgxMapper

import (
	pgxUseCase "github.com/OddEer0/golang-practice/practices/pgx_practice/pgx_usecase"
	proto "github.com/OddEer0/golang-practice/protogen"
)

type UserMapper struct{}

func (u UserMapper) PureUserToGrpcResponseUser(data *pgxUseCase.PureUser) *proto.ResponseUser {
	return &proto.ResponseUser{
		Id:    string(data.Id),
		Login: data.Login,
		Email: data.Email,
	}
}

func (u UserMapper) PureUserAggregateToGrpcResponseUserAggregate(userAggregate *pgxUseCase.PureUserAggregate) *proto.ResponseUserAggregate {
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
	}
}
