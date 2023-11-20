package models

import "todos-api/dbconfig"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func InsertUser(user User) (id int64, err error) {
	sql := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`

	err = dbconfig.DB.QueryRow(sql, user.Username, user.Password).Scan(&id)

	return
}

func GetUser(id int64) (user User, err error) {
	row := dbconfig.DB.QueryRow(`SELECT * FROM users WHERE id=$1`, id)

	err = row.Scan(&user.ID, &user.Username, &user.Password)

	return
}

func GetByUserName(userName string, pass string) (user User, err error) {
	row := dbconfig.DB.QueryRow(`SELECT * FROM users WHERE username=$1 AND password=$2`, userName, pass)

	err = row.Scan(&user.ID, &user.Username, &user.Password)

	return
}
