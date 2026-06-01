package repository

const (
	qUpdateUserRole = `
		UPDATE users
		SET role_id = $1
		WHERE id = $2
	`

	qUpdateTokenByUserID = `
		UPDATE users
		SET token_version = token_version + 1
		WHERE user_id = $1
	`
)
