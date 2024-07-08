package pgxRepository

const (
	CreatePostQuery = `
		INSERT INTO posts (id, ownerId, title, content, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5);
	`
	GetPostByIdQuery = `
		SELECT id, ownerId, title, content, updatedAt, createdAt FROM posts WHERE id = $1;
	`
	GetPostByQueryAsc = `
		SELECT id, ownerId, title, content, updatedAt, createdAt FROM posts WHERE id = $1
		ORDER BY $2 ASC;
		LIMIT $3 OFFSET $4;
	`
	GetPostByQueryDesc = `
		SELECT id, ownerId, title, content, updatedAt, createdAt FROM posts WHERE id = $1
		ORDER BY $2 ASC;
		LIMIT $3 OFFSET $4;
	`
)
