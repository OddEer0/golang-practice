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
)