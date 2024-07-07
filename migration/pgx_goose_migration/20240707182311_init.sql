-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	login VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
	updatedAt DATE NOT NULL,
    createdAt DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS posts (
	id UUID PRIMARY KEY,
	title VARCHAR(255),
	content VARCHAR(25000),
	updatedAt DATE NOT NULL,
    createdAt DATE NOT NULL,
	ownerId UUID REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments (
	id UUID PRIMARY KEY,
	message VARCHAR(1020),
	updatedAt DATE NOT NULL,
    createdAt DATE NOT NULL,
	ownerId UUID REFERENCES users(id),
	postId UUID REFERENCES posts(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
