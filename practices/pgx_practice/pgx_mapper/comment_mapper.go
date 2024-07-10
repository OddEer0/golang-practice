package pgxMapper

import (
	proto "github.com/OddEer0/golang-practice/protogen"
	"github.com/OddEer0/golang-practice/resources/aggregate"
	"github.com/OddEer0/golang-practice/resources/model"
)

type CommentMapper struct{}

func (c CommentMapper) CommentToGrpcResponseComment(comment *model.Comment) *proto.ResponseComment {
	return &proto.ResponseComment{
		Id:      string(comment.Id),
		Message: comment.Message,
	}
}

func (c CommentMapper) CommentAggregatesToGrpcResponseCommentAggregates(data []*aggregate.Comment) []*proto.ResponseCommentAggregate {
	res := make([]*proto.ResponseCommentAggregate, 0, len(data))

	for _, comment := range data {
		aggr := &proto.ResponseCommentAggregate{
			Value: &proto.ResponseComment{
				Id:      string(comment.Value.Id),
				Message: comment.Value.Message,
			},
		}
		if comment.Owner != nil {
			aggr.Owner = &proto.ResponseUser{
				Id:    string(comment.Owner.Id),
				Email: comment.Owner.Email,
				Login: comment.Owner.Login,
			}
		}
		if comment.Post != nil {
			aggr.Post = &proto.ResponsePost{
				Id:      string(comment.Post.Id),
				Title:   comment.Post.Title,
				Content: comment.Post.Content,
			}
		}
		res = append(res, aggr)
	}

	return res
}
