package CHATS

import (
	DB_CONN "NeoCom/connection"
	"fmt"
)

type CorporateChat struct {
	UserID     uint   `json:"user_id"`
	ChatID     uint   `json:"chat_id"`
	PartFName  string `json:"partner_fname"`
	PartLName  string `json:"partner_lname"`
	EmplID     uint   `json:"empl_id"`
	ProfilePic string `json:"profile_pic"`
	PartID     uint   `json:"part_id"`
}

func FindAllChatsByEmplID(empl_id string) []CorporateChat {
	rows, err := DB_CONN.Conn.DB.Query("SELECT DISTINCT u.id AS user_id, c.id AS chat_id, "+
		"e.first_name AS partner_fname, e.last_name AS partner_lname, "+
		"e.id AS empl_id, e.photo AS profile_pic, p2.id AS part_id "+
		"FROM chats c "+
		"JOIN chat_participants cp1 ON cp1.chat_id = c.id "+
		"JOIN chat_participants cp2 ON cp2.chat_id = c.id "+
		"JOIN participants p1 ON p1.id = cp1.participants_id "+
		"JOIN participants p2 ON p2.id = cp2.participants_id "+
		"JOIN users u ON u.id = p2.reference_id "+
		"JOIN employees e ON e.id = u.empl_id "+
		"WHERE (p1.role = 'employee' AND p2.role = 'employee') "+
		"AND (p1.reference_id = $1 AND p2.reference_id != $1);", empl_id)

	if err != nil {
		fmt.Println("FildAllChatsByEmplId query error")
		return nil
	}
	defer rows.Close()

	var chats []CorporateChat
	for rows.Next() {
		var chat CorporateChat

		err := rows.Scan(&chat.UserID, &chat.ChatID, &chat.PartFName, &chat.PartLName, &chat.EmplID, &chat.ProfilePic, &chat.PartID)
		if err != nil {
			return nil
		}

		chats = append(chats, chat)
	}

	return chats
}
