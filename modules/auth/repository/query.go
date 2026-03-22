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

	qGetPasswordHashByEmail = `
		SELECT password_hash
		FROM users
		WHERE email = $1
	`

	qGetPasswordHashByPhoneNumber = `
		SELECT password_hash
		FROM users
		WHERE phonenummber = $1
	`

	qGetPasswordHashByUsername = `
		SELECT password_hash
		FROM users
		WHERE username = $1
	`
)
