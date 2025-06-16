package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var esp32conn *websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Reseiver string `json:"phone"`
	Text     string `json:"text"`
}

func send_message_to_esp32(conn *websocket.Conn, message Message) {
	conn.WriteJSON(message)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Cannot connect to websocket")
		return
	}
	esp32conn = conn
}

func check_new_messages(db *sql.DB) {
	var curr_max_message_id uint
	db.QueryRow("SELECT MAX(id) AS max_id FROM messages;").Scan(&curr_max_message_id)

	for {
		var new_max_id uint
		var chat_id uint
		var sender_id uint
		var message_text string
		db.QueryRow("SELECT MAX(id) AS max_id, chat_id, sender_participant_id AS sender_id"+
			", text FROM messages;").Scan(&new_max_id, &chat_id, &sender_id, &message_text)

		if curr_max_message_id < new_max_id {
			curr_max_message_id = new_max_id
			var partner_id uint
			var reference_id uint
			var role string
			db.QueryRow("SELECT cp.participants_id, p.role, p.reference_id FROM chat_participants "+
				"WHERE cp.chat_id = $1 AND cp.participants_id != $2", chat_id, sender_id).Scan(&partner_id, &role, &reference_id)
			if role == "customer" {
				var phone string
				db.QueryRow("SELECT phone FROM customers WHERE id = $1", reference_id).Scan(&phone)

				message := Message{
					Reseiver: phone,
					Text:     message_text,
				}

				send_message_to_esp32(&websocket.Conn{}, message)
			}
		}

	}
}

func main() {
	connStr := "postgresql://neondb_owner:npg_qILNuP6Diz1Z@ep-divine-sun-a83zg48v-pooler.eastus2.azure.neon.tech/neondb?sslmode=require"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("Error to connect")
	}
	http.HandleFunc("/ws", handleConnections)
	http.ListenAndServe(":8080", nil)
	check_new_messages(db)
}
