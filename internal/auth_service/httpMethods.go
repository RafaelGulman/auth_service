package auth_service

import (
	"encoding/json"
	"net/http"
)

// Логика записи создания аккаунта
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAccount Account

	err := json.NewDecoder(r.Body).Decode(&newAccount) //запрос декодировали в объект
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]bool{"successfull": false})
		return
	}
	if newAccount.Login == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]bool{"successfull": false})
		return
	}
	/*
		Логика работы с Бд
		(перевод пароля в хэш и запихивание в бд)
	*/
	accounts = append(accounts, newAccount) //Пока заглушка
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]bool{"successfull": true})
}

func EnterAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newAccount Account

	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]bool{"successfull": false})
		return
	}
	if !CheckLogin(newAccount) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]bool{"successfull": false})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]bool{"successfull": true})
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
