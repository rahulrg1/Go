package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"go_valid/models"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

func createConnection() *sql.DB {

	err := godotenv.Load(".env")
	var (
		host     = os.Getenv("DB_HOST")
		port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	if err != nil {
		log.Fatal("Error loading .env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, er := sql.Open("postgres", psqlInfo)
	if er != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("connection not created %v", err)
	}

	return db
}

func CreateForm(form *models.Form) error {
	db := createConnection()
	defer db.Close()
	if form.Phonenumber != 0 {
		phonenumber := fmt.Sprint(form.Phonenumber)
		match, _ := regexp.MatchString("^[0-9]{10}$", phonenumber)
		if !match {
			return errors.New("not a valid Phone Number")
		}
	}
	if form.Email != "" {
		match, _ := regexp.MatchString(`^(.*)@(.*)([\.])com$`, form.Email)
		if !match {
			return errors.New("not a valid Email")
		}
	}

	sqlStatement := `Select email from form WHERE email=$1`
	var email string
	row := db.QueryRow(sqlStatement, form.Email)
	row.Scan(&email)

	// fmt.Printf("%v", email)
	if email != "" {
		return errors.New("this email is already entered , Please add unique only! ")
	}
	sqlStatement = `INSERT INTO form(firstname,lastname,email,phonenumber,city) VALUES($1,$2,$3,$4,$5)`
	_, err := db.Exec(sqlStatement, form.Firstname, form.Lastname, form.Email, form.Phonenumber, form.City)

	if err != nil {
		log.Fatalf("Insertion failed ! : %v", err)
	}
	fmt.Println("data inserted!")
	return nil
}

func GetAllData() ([]models.Form, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `SELECT * FROM form`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute %v", err)
	}
	defer rows.Close()
	var forms []models.Form

	for rows.Next() {
		var form models.Form
		err = rows.Scan(&form.Firstname, &form.Lastname, &form.Email, &form.Phonenumber, &form.City)
		if err != nil {
			log.Fatalf("retrive data failed : %v", err)
		}
		forms = append(forms, form)
	}
	return forms, err
}

func GetDataByEmail(email string) ([]models.Form, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `select firstname, lastname, email, phonenumber ,city from form where email=$1`
	var forms []models.Form

	rows, err := db.Query(sqlStatement, email)
	defer rows.Close()

	for rows.Next() {
		var form models.Form
		err = rows.Scan(&form.Firstname, &form.Lastname, &form.Email, &form.Phonenumber, &form.City)
		if err != nil {
			log.Fatalf("retrival failed :%v", err)
		}
		forms = append(forms, form)
	}
	return forms, err
}
func GetDataByCity(city string) ([]models.Form, error) {
	db := createConnection()
	defer db.Close()
	sqlStatement := `select firstname, lastname, email, phonenumber, city from form where city=$1`
	var forms []models.Form

	rows, err := db.Query(sqlStatement, city)
	defer rows.Close()

	for rows.Next() {
		var form models.Form
		err = rows.Scan(&form.Firstname, &form.Lastname, &form.Email, &form.Phonenumber, &form.City)
		if err != nil {
			log.Fatalf("retrival failed :%v", err)
		}
		forms = append(forms, form)
	}
	return forms, err
}

func UpdateForm(form *models.Form) error {
	db := createConnection()
	defer db.Close()
	sqlStatement := `Select email from form WHERE email=$1`
	row := db.QueryRow(sqlStatement, form.Email)
	var email string
	row.Scan(&email)

	if email == "" {
		return errors.New("this email does not exist, Please provide email which exists! ")
	}
	if form.Phonenumber != 0 {
		phonenumber := fmt.Sprint(form.Phonenumber)
		match, _ := regexp.MatchString("^[0-9]{10}$", phonenumber)
		if !match {
			return errors.New("not a valid Phone Number")
		}
	}
	var err error
	if form.City != "" {
		sqlStatement := `update form set firstname=$1,lastname=$2,phonenumber=$3,city=$4 where email=$5`
		_, err = db.Exec(sqlStatement, form.Firstname, form.Lastname, form.Phonenumber, form.City, form.Email)
	} else {

		sqlStatement := `update form set firstname=$1,lastname=$2,phonenumber=$3 where email=$4`
		_, err = db.Exec(sqlStatement, form.Firstname, form.Lastname, form.Phonenumber, form.Email)
	}

	if err != nil {
		log.Fatalf("Upadte failed ! : %v", err)
		return err
	}
	fmt.Println("Data Updated Successfully!")
	return nil
}

func DeleteForm(email string) error {
	db := createConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM form WHERE email=$1`
	_, err := db.Exec(sqlStatement, email)
	if err != nil {
		log.Fatalf("Deletion failed! : %v", err)
		return errors.New("Provided email does not exist")
	}
	fmt.Println("data deleted!")
	return nil
}
