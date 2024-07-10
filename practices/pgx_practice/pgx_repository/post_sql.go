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
	GetPostsBodyByUserIdQuery = `
		SELECT id, title, content FROM posts WHERE id = $1;
	`
	DeletePostByIdQuery = `
		DELETE FROM posts WHERE id = $1;
	`
	UpdatePostById = `
		UPDATE posts SET title = $1, content = $2, updatedAt = $3 WHERE id = $4;
	`
)
