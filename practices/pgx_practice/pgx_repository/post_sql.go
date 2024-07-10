package pgxRepository

const (
	CreatePostQuery = `
		INSERT INTO posts (id, ownerId, title, content, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	GetPostByIdQuery = `
		SELECT id, ownerId, title, content, updatedAt, createdAt FROM posts WHERE id = $1;
	`
	GetPostByOwnerIdQueryPart1 = `
		SELECT id, ownerId, title, content, updatedAt, createdAt FROM posts WHERE ownerId = $1
		ORDER BY `
	GetPostByOwnerIdQueryPart12 = `
		LIMIT $2 OFFSET $3;
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
