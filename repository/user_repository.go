package repository

import (
	"database/sql"
	//"go-restapi-v2/database"
	"go-restapi-v2/model"
)

type UserRepository struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

func (r UserRepository) GetAllUsers() (*[]model.User, error) {
	var users []model.User

	//MySQLDatabase := database.NewMySQLDatabase()
	//rows, err := MySQLDatabase.Instance.Query("SELECT * FROM users")
	rows, err := r.database.Query("SELECT * FROM users")

	if err != nil {
		panic(err.Error())
	}

	//defer rows.Close()

	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Nombre, &u.Edad, &u.Email); err != nil {
			panic(err.Error())
		}
		users = append(users, u)
	}

	return &users, nil
}

//func (r UserRepository) GetUser() (*model.User, error) {
//
//}
//
//func (r UserRepository) CreateUser(u *model.User) (*model.User, error) {
//
//}
//
//func (r UserRepository) UpdateUser(id int) (*model.User, error) {
//
//}
//
//func (r UserRepository) DeleteUser(id int) error {
//
//}
