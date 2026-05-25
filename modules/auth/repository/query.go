package repository

const (
	qRegister = `
		INSERT INTO users (username, email, phonenumber, first_name, last_name, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	qLogin = `
		SELECT COUNT(1)
		FROM users
		WHERE 
	`

	qGetHashedPasswordByKey = `
		SELECT password_hash
		FROM users
		WHERE %s = $1
	`

	qBaseUpdateTokenVersion = `
		UPDATE users
		SET token_version = token_version + 1
		WHERE %s = $1 
		RETURNING token_version
	`
)
