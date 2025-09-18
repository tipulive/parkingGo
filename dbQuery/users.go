package dbQuery

import (
	"log"
	"parking/db"
	"parking/model"
)

func CreateUser(user model.User) error {
	_, err := db.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}
func AdminLogin(admin model.Admin) (bool, []model.Admin) {

	//_, err := db.DB.Query("select *from users where email=? and name=?", user.Email, user.Name)
	// Clean input

	rows, err := db.DB.Query("SELECT id,uid, name, email,PhoneNumber,created_at FROM admins WHERE name=?", admin.Name)

	if err != nil {

		return false, nil
	}
	defer rows.Close()

	var admins []model.Admin
	for rows.Next() {
		var u model.Admin
		//if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		if err := rows.Scan(&u.ID, &u.UID, &u.Name, &u.Email, &u.PhoneNumber, &u.CreatedAt); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		admins = append(admins, u)
	}
	if len(admins) < 1 {
		return false, admins
	}
	return true, admins
}
func UserLogin(user model.User) (bool, []model.User) {

	//_, err := db.DB.Query("select *from users where email=? and name=?", user.Email, user.Name)
	// Clean input

	rows, err := db.DB.Query("SELECT id,uid, name, email,PhoneNumber,created_at FROM users WHERE name=?", user.Name)

	if err != nil {

		return false, nil
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		//if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		if err := rows.Scan(&u.ID, &u.UID, &u.Name, &u.Email, &u.PhoneNumber, &u.CreatedAt); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		users = append(users, u)
	}
	if len(users) < 1 {
		return false, users
	}
	return true, users
}
func GetUsers() ([]model.User, error) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

func UpdateUser(user model.User) error {
	_, err := db.DB.Exec("UPDATE users SET name=?, email=? WHERE id=?", user.Name, user.Email, user.ID)
	return err
}

func DeleteUser(id int) error {
	_, err := db.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
