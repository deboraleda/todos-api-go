package models

import "todos-api/dbconfig"

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	User_id     int64  `json:"user_id"`
}

func Update(id int64, todo Todo) (int64, error) {
	res, err := dbconfig.DB.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Insert(todo Todo, userId int64) (id int64, err error) {
	sql := `INSERT INTO todos (title, description, done, user_id) VALUES ($1, $2, $3, $4) RETURNING id`

	err = dbconfig.DB.QueryRow(sql, todo.Title, todo.Description, todo.Done, userId).Scan(&id)

	return
}

func Get(id int64) (todo Todo, err error) {
	row := dbconfig.DB.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.User_id)

	return
}

func GetAllFromUser(userId int64) (todos []Todo, err error) {
	rows, err := dbconfig.DB.Query(`SELECT * FROM todos WHERE user_id=$1`, userId)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.User_id)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}

func Delete(id int64) (int64, error) {
	res, err := dbconfig.DB.Exec(`DELETE FROM todos WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
