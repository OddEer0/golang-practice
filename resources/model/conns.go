package model

type (
	ManyOpt struct {
		Limit uint
		Page uint
	}

	PostConns map[string]*ManyOpt
	UserConns map[string]*ManyOpt
	CommentConns map[string]*ManyOpt
)
