package CHATS

import DB_CONN "NeoCom/connection"

type Message struct {
	Message_text string `json:"message_text"`
	TimeStamp    string `json:"timestamp"`
	Sender_id    uint   `json:"sender_id"`
}

func FindAllMessagesFromChat(chat_id string) []Message {
	rows, err := DB_CONN.Conn.DB.Query("SELECT m.text AS message_text, "+
		"TO_CHAR(m.timestamp, 'HH24:MI') AS timestamp, "+
		"p.reference_id AS sender_id "+
		"FROM messages m "+
		"JOIN participants p ON p.id = m.sender_participant_id "+
		"WHERE m.chat_id = $1;", chat_id)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.Message_text, &msg.TimeStamp, &msg.Sender_id)
		if err != nil {
			return nil
		}

		messages = append(messages, msg)
	}

	return messages
}
