package infraestructure

import (
	"database/sql"
	"demo/src/users/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}

func (m *MySQL) Save(user domain.User) (domain.User, error) {
	result, err := m.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return domain.User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}
	user.Id = int32(id)
	user = domain.User{Id: int32(id), Name: user.Name, Email: user.Email, Password: user.Password}
	return user, nil
}

func (m *MySQL) GetAll() ([]domain.User, error) {
	rows, err := m.db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (m *MySQL) GetByID(id int32) (domain.User, error) {
	row := m.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var user domain.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *MySQL) Update(id int32, user domain.User) (domain.User, error) {
	result, err := m.db.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return domain.User{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (m *MySQL) Delete(id int32) (int64, error) {
	result, err := m.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (m *MySQL) Login(email string, password string) (domain.LoginResponse, error) {
	row := m.db.QueryRow("SELECT id, name, email FROM users WHERE email = ? AND password = ?", email, password)
	var login domain.LoginResponse
	err := row.Scan(&login.Id, &login.Name, &login.Email)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	return login, nil
}

