package users

import (
	"tchipify/users/internal/helpers"
	"tchipify/users/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	users := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.ID, &data.InscriptionDate, &data.Name, &data.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return users, err
}

func GetUserById(id int) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id)
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.ID, &user.InscriptionDate, &user.Name, &user.Username )
	if err != nil {
		return nil, err
	}
	return &user, err
}


func CreateUser(newUser models.User) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	result, err := db.Exec("INSERT INTO users (username, name, inscription_date) VALUES (?, ?, ?)",
		newUser.Username, newUser.Name, newUser.InscriptionDate)

	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Update the User object with the generated ID
	newUser.ID = int(id)
	return &newUser, nil
}
 
func UpdateUserById(updatedUser models.User) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	// Update the user in the database
	_, err = db.Exec("UPDATE users SET artist=?, file_name=?, published_date=?, title=? WHERE id=?",
		updatedUser.Username, updatedUser.Name, updatedUser.InscriptionDate, updatedUser.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserById(id int) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

