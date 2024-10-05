package queries

const (
	CreateNewUser = "INSERT INTO users (nickname, password, email, avatar) VALUES ($1, $2, $3, $4) RETURNING user_id"
	FindUser       = "SELECT user_id, nickname, password, email, avatar FROM users WHERE nickname = $1"
	UpdateUserData = "UPDATE users SET nickname = $1, password = $2, email = $3, avatar = $4 WHERE user_id = $5"
)
