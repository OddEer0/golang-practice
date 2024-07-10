package pgxRepository

const (
	GetAllCommentsIdContentByUserId = `
		SELECT id, message FROM comments WHERE ownerId = $1;
	`
	GetCommentByIdQuery = `
		SELECT id, ownerId, postId, message, updatedAt, createdAt FROM comments WHERE id = $1;
	`
	GetCommentsByPostIdQueryPart1 = `
		SELECT id, ownerId, postId, message, updatedAt, createdAt FROM comments WHERE postId = $1
		ORDER BY `
	GetCommentsByPostIdQueryPart2 = `
		LIMIT $2 OFFSET $3;
	`
	GetCommentsByOwnerIdQueryPart1 = `
		SELECT id, ownerId, postId, message, updatedAt, createdAt FROM comments WHERE ownerId = $1
		ORDER BY `
	GetCommentsByOwnerIdQueryPart2 = `
		LIMIT $2 OFFSET $3;
	`
	CreateCommentQuery = `
		INSERT INTO comments(id, ownerId, postId, message, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	DeleteCommentByIdQuery = `
		DELETE FROM comments WHERE id = $1;
	`
	UpdateCommentByIdQuery = `
		UPDATE comments SET message = $1, updatedAt = $2 WHERE id = $3;
	`
)
