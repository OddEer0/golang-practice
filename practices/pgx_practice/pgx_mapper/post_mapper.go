package pgxMapper

import (
	proto "github.com/OddEer0/golang-practice/protogen"
	"github.com/OddEer0/golang-practice/resources/aggregate"
	"github.com/OddEer0/golang-practice/resources/model"
)

type PostMapper struct{}

func (p PostMapper) PostAggregateToResponsePostAggregate(d *aggregate.Post) *proto.ResponsePostAggregate {
	post := &proto.ResponsePostAggregate{
		Value: &proto.ResponsePost{
			Id:      string(d.Value.Id),
			Title:   d.Value.Title,
			Content: d.Value.Content,
		},
	}
	if d.Owner != nil {
		post.Owner = &proto.ResponseUser{
			Id:    string(d.Owner.Id),
			Email: d.Owner.Email,
			Login: d.Owner.Login,
		}
	}

	if d.Comments != nil {
		comments := make([]*proto.ResponseComment, len(d.Comments))
		for _, c := range d.Comments {
			comments = append(comments, &proto.ResponseComment{
				Id:      string(c.Id),
				Message: c.Message,
			})
		}
	}

	return post
}

func (p PostMapper) PostAggregatesToGrpcResponseManyResponsePost(data []*aggregate.Post) *proto.ResponseManyResponsePost {
	posts := make([]*proto.ResponsePostAggregate, 0, len(data))
	for _, d := range data {

		posts = append(posts, p.PostAggregateToResponsePostAggregate(d))

	}
	return &proto.ResponseManyResponsePost{
		Posts: posts,
	}
}

func (p PostMapper) PostToGrpcResponsePost(post *model.Post) *proto.ResponsePost {
	return &proto.ResponsePost{
		Id:      string(post.Id),
		Title:   post.Title,
		Content: post.Content,
	}
}
