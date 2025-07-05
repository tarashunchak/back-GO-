package LOGIN

import (
	DB_CONN "NeoCom/connection"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
)

func HashSHA256(s string) string {
	hash := sha256.Sum256([]byte(s)) // повертає [32]byte
	return hex.EncodeToString(hash[:])
}

type Guest struct {
	Email    string
	Password string
	Id       int
}

type CurrentUser struct {
	UID     int `json:"user_id"`
	Empl_Id int `json:"empl_id"`
	Part_Id int `json:"part_id"`
}

func IsExist(user *Guest) bool {
	err := DB_CONN.Conn.DB.QueryRow("SELECT id FROM users WHERE username = $1;", user.Email).Scan(&user.Id)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func IsPasswordCorrect(user *Guest) bool {
	var bin int
	err := DB_CONN.Conn.DB.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2", user.Email, user.Password).Scan(&bin)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func Authenticate(username string, password string) CurrentUser {
	hashed_pass := HashSHA256(password)
	curr_guest := Guest{
		username,
		hashed_pass,
		0,
	}
	var curr_user CurrentUser
	if !IsExist(&curr_guest) {
		curr_user = CurrentUser{
			0, 0, 0,
		}
	} else if !IsPasswordCorrect(&curr_guest) {
		curr_user = CurrentUser{
			0, 0, 0,
		}
	} else {
		var is_online bool
		err := DB_CONN.Conn.DB.QueryRow("SELECT u.id AS user_id, "+
			"u.empl_id AS empl_id, p.id AS part_id, "+
			"u.is_online AS is_online "+
			"FROM participants p "+
			"JOIN users u ON u.id = p.reference_id "+
			"WHERE p.role = 'employee' "+
			"AND u.username = $1 "+
			"AND u.password = $2;", curr_guest.Email, curr_guest.Password).Scan(&curr_user.UID, &curr_user.Empl_Id, &curr_user.Part_Id, &is_online)
		if err == sql.ErrNoRows || is_online {
			curr_user = CurrentUser{
				0, 0, 0,
			}
		}
	}
	return curr_user
}
