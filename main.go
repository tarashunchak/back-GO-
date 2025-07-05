package main

import (
	DB_CONN "NeoCom/connection"
	"NeoCom/handlers"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	DB_CONN.Conn.ConnectToDB()

	http.HandleFunc("/chat_messages", handlers.ChatMessagesHandler)
	http.HandleFunc("/corporate_chats", handlers.ChatsHandler)
	http.HandleFunc("/employees_page", handlers.EmployeesPageHandler)
	http.HandleFunc("/customers_page", handlers.CustomersPageHandler)
	http.HandleFunc("/tariffs_page", handlers.TariffsPageHandler)
	http.HandleFunc("/login_page", handlers.LoginPageHandler)
	http.HandleFunc("/dashboard_page", handlers.DashboardHandler)
	http.HandleFunc("/dashboard_page_req_history", handlers.DashboardReqTableHandler)
	http.HandleFunc("/dashboard_page_req_series", handlers.DashboardReqSeriesHandler)

	http.ListenAndServe(":8080", nil)
}
