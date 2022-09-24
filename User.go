package main

import (
	"net/http"
    "gorm.io/gorm"
)
var DB *gorm.DB
var err error
const DNS = "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
type User Struct{
	//to convert struct to model
	gorm.Model
	FirstName string   'json:"firstname"'
	LastName  string   'json:"lastname"'
	Email     string   'json:"email"'
}

func InitialMigration(){
  DB , err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
  if(err != nil){
	fmt.println(err.Error())
	panic("Cannot connect to DB")
  }
  DB.AutoMigrate(&User{})

}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewDecoder(w).Encode(&user)
	
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
}
