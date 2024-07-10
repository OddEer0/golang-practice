package pgxRepository

const (
	CreateUserQuery = `
		INSERT INTO users (id, login, email, password, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	GetUserById = `
		SELECT id, login, email, password, updatedAt, createdAt FROM users WHERE id = $1;
	`
	DeleteUserById = `
		DELETE FROM users WHERE id = $1;
	`
	UpdateUserLoginById = `
		UPDATE users SET login = $1 WHERE id = $2
		RETURNING id, login, email, password, updatedAt, createdAt;
	`
	UpdateUserById = `
		UPDATE users SET login = $1, email = $2, password = $3, updatedAt = $4 WHERE id = $5
	`
	GetUserByQueryPart1 = `
		SELECT id, login, email, password, updatedAt, createdAt FROM posts
		ORDER BY `
	GetUserByQueryPart12 = `
		LIMIT $1 OFFSET $2;
	`
)
