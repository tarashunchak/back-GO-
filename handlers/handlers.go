package handlers

import (
	CHATS "NeoCom/chats"
	CUST "NeoCom/customers"
	DASHBOARD "NeoCom/dashboard"
	EMPL "NeoCom/employees"
	LOGIN "NeoCom/login_and_auth"
	TARIFFS "NeoCom/tariffs"
	"encoding/json"
	"net/http"
)

func ChatMessagesHandler(w http.ResponseWriter, r *http.Request) {
	chat_id := r.URL.Query().Get("chat_id")
	messages := CHATS.FindAllMessagesFromChat(chat_id)
	if messages == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func ChatsHandler(w http.ResponseWriter, r *http.Request) {
	empl_id := r.URL.Query().Get("empl_id")
	chats := CHATS.FindAllChatsByEmplID(empl_id)
	if chats == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(chats)
}

func EmployeesPageHandler(w http.ResponseWriter, r *http.Request) {
	tmp_bool := r.URL.Query().Get("is_search_result")

	var employees []EMPL.Employee
	if tmp_bool == "true" {
		str := r.URL.Query().Get("str")
		employees = EMPL.SelectEmployeesBy(str)
	} else {
		employees = EMPL.SelectAllEmployees()
	}

	if employees == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func CustomersPageHandler(w http.ResponseWriter, r *http.Request) {
	customers := CUST.SelectAllCustomers()
	if customers == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func TariffsPageHandler(w http.ResponseWriter, r *http.Request) {
	tariffs := TARIFFS.SelectAllTariffs()
	if tariffs == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(tariffs)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	current_user := LOGIN.Authenticate(username, password)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(current_user)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	interval := r.URL.Query().Get("interval")
	cust_table_content := DASHBOARD.DashboardCustomersTable(interval)
	if cust_table_content == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(cust_table_content)
}

func DashboardReqTableHandler(w http.ResponseWriter, r *http.Request) {
	interval := r.URL.Query().Get("interval")
	req_table_content := DASHBOARD.DashboardRequestsHistory(interval)
	if req_table_content == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(req_table_content)
}

func DashboardReqSeriesHandler(w http.ResponseWriter, r *http.Request) {
	interval := r.URL.Query().Get("interval")
	req_series_content := DASHBOARD.DashboardRequestsSeries(interval)
	if req_series_content == nil {
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(req_series_content)
}
