package resourcemodel

type (
	User struct {
		Id       string `json:"id"`
		Login    string `json:"login"`
		Password string `json:"-"`
	}

	Post struct {
		Id      string `json:"id"`
		OwnerId string `json:"owner_id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	PostAggregate struct {
		Post     *Post
		Owner    *User
		Comments []*Comment
	}

	Comment struct {
		Id      string `json:"id"`
		OwnerId string `json:"owner_id"`
		PostId  string `json:"post_id"`
	}

	AggregateOption struct {
		Conn []string
	}

	QueryOption struct {
		Page  int
		Limit int
	}
)
