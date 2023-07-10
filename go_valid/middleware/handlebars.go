package middleware

import (
	"fmt"
	"go_valid/models"
	"log"
	"net/http"

	"encoding/json"

	"go_valid/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type responseValidation struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllData(w http.ResponseWriter, r *http.Request) {

	data, err := controllers.GetAllData()
	if err != nil {
		log.Fatalf("data retrival failed! %v", err)
	}
	json.NewEncoder(w).Encode(data)
}

func CreateForm(w http.ResponseWriter, r *http.Request) {
	var form models.Form
	json.NewDecoder(r.Body).Decode(&form)
	if form.Firstname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Firstname is balnk"})
		return
	} else if form.Lastname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Lastname is blank"})
		return
	} else if form.Email == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Email is blank"})
		return
	} else if form.Phonenumber == 0 {
		json.NewEncoder(w).Encode(responseValidation{Error: "Phone Number is blank"})
		return
	}

	err := controllers.CreateForm(&form)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Validation failed : %v ", err))
	}
	json.NewEncoder(w).Encode(form)
}
func GetDataByEmail(w http.ResponseWriter, r *http.Request) {
	var email string
	params := mux.Vars(r)
	email = params["email"]
	forms, err := controllers.GetDataByEmail(email)
	if err != nil {
		log.Fatalf("data retrival failed :%v", err)
	}
	json.NewEncoder(w).Encode(forms)
}
func GetDataByCity(w http.ResponseWriter, r *http.Request) {

	city := r.URL.Query().Get("location")
	forms, err := controllers.GetDataByCity(city)
	if err != nil {
		log.Fatalf("data retrival failed :%v", err)
	}
	if len(forms) != 0 {
		json.NewEncoder(w).Encode(forms)
	} else {
		json.NewEncoder(w).Encode("data with this city name does not exists! .")
	}
}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	var form models.Form
	json.NewDecoder(r.Body).Decode(&form)

	fmt.Println(form.Email)
	form.Email = email
	fmt.Println(form.Email)

	if form.Firstname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Firstname is balnk"})
		return
	} else if form.Lastname == "" {
		json.NewEncoder(w).Encode(responseValidation{Error: "Lastname is blank"})
		return
	} else if form.Phonenumber == 0 {
		json.NewEncoder(w).Encode(responseValidation{Error: "Phone Number is blank"})
		return
	}

	err := controllers.UpdateForm(&form)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Update failed : %v ", err))
	} else {
		json.NewEncoder(w).Encode("Update Successful!")
	}
}

func DeleteForm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	err := controllers.DeleteForm(email)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Deletion failed : %v ", err))
	} else {
		json.NewEncoder(w).Encode("Deletion successful!")
	}
}
