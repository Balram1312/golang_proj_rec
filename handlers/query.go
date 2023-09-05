package handlers
const(
	insertNewUser = `INSERT INTO users (Username,Password) VALUES ($1,$2) RETURNING id`
	
)