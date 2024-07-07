package aggregate

import "github.com/OddEer0/golang-practice/resources/model"

type (
	User struct {
		Value *model.User
		Posts []*model.Post
		Comments []*model.Comment
	}

	Post struct {
		Value *model.Post
		Owner *model.User
		Comments []*model.Comment
	}

	Comment struct {
		Value *model.Comment
		Owner *model.User
		Comments []*model.Comment
	}
)
