package pgxRepository

const (
	GetAllCommentsIdContentByPostId = `
		SELECT id, content FROM comments WHERE postId = $1;
	`
	GetCommentByIdQuery = `
		SELECT id, ownerId, postId, message, updatedAt, createdAt FROM comments WHERE id = $1;
	`
	GetCommentsByPostIdQueryAsc = `
		SELECT id, ownerId, postId, message, updatedAt, createdAt FROM comments 
		WHERE postId = $1
		ORDER BY $2 ASC
		LIMIT $3 OFFSET $4;
	`
	GetCommentsByPostIdQueryDesc = `
		SELECT id, ownerId, postId, message, updatedAt, createdAt FROM comments 
		WHERE postId = $1
		ORDER BY $2 DESC
		LIMIT $3 OFFSET $4;
	`
	CreateCommentQuery = `
		INSERT INTO comments(id, ownerId, postId, message, updatedAt, createdAt);
	`
	DeleteCommentByIdQuery = `
		DELETE FROM comments WHERE id = $1;
	`
	UpdateCommentByIdQuery = `
		UPDATE comments SET message = $1, updatedAt = $2 WHERE id = $3;
	`
)
