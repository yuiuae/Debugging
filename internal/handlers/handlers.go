// Copyright 2023 Serhii Khrystenko. All rights reserved.

/*
Package handler implements user password verification.

This package is designed as an example of the Godoc
documentation and does not have any functionality:)
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/yuiuae/Debugging/internal/db"
	"github.com/yuiuae/Debugging/pkg/hasher"
)

// calls per hour allowed by the user
var callperhour int = 100

// token validity time (minutes)
var tokentime = 60000

// secret key for token
var tokenSecretKey = "SecretYouShouldHide"

// Table with users
// var usersTable = map[string]*UserInfo1{}
type ConnInfo struct {
	WS       *websocket.Conn
	Token    string
	ExpireAt int64
}

var connTable = map[string]*ConnInfo{}

// Create a struct that models the structure for a user creating
// Request
type CrRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Response
type CrResponse struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

// Create new use and add to user table
func UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorlog(w, "Only POST method allowed ", http.StatusBadRequest)
		return
	}
	req := &CrRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		errorlog(w, "Bad request, empty username or password1", http.StatusBadRequest)
		return
	}
	if len(req.UserName) <= 4 {
		errorlog(w, "Username should be at least 4 characters", http.StatusBadRequest)
		return
	}

	if len(req.Password) <= 8 {
		errorlog(w, "Password should be at least 8 characters", http.StatusBadRequest)
		return
	}
	// _, b := usersTable[req.UserName]
	// if b {
	// 	errorlog(w, "A user with this name already exists", http.StatusConflict)
	// 	return
	// }

	// Hash the password using the bcrypt algorithm
	hashedPassword, err := hasher.HashPassword_bcrypt(req.Password)
	if err != nil {
		errorlog(w, "Internal Server Error (hash error)", http.StatusInternalServerError)
		return
	}

	// Generate UUID
	uid, err := uuid.NewRandom()
	if err != nil {
		errorlog(w, "Internal Server Error (UUID error)", http.StatusInternalServerError)
		return
	}

	// add new user to to user table
	status, err := db.AddNewUser(req.UserName, hashedPassword, uid.String())
	if err != nil {
		errorlog(w, err.Error(), status)
		return
	}

	// Create response
	resp := &CrResponse{uid.String(), req.UserName}
	err = json.NewEncoder(w).Encode(&resp) //&resp
	if err != nil {
		errorlog(w, "Internal Server Error (json Encoder error)", http.StatusInternalServerError)
		return
	}
	// usersTable[req.UserName] = &UserInfo1{hashedPassword, uid.String(), ""}
}

// Create a struct that models the structure for a user login
// Request
type LogRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Response
type LogResponse struct {
	URL string `json:"URL"`
}

// Check username and passoword in user table
func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorlog(w, "Only POST method allowed", http.StatusBadRequest)
		return
	}

	req := &LogRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		errorlog(w, "Bad request, empty username or password", http.StatusBadRequest)
		return
	}

	_, ok, err := db.PassVerify(req.UserName, req.Password)
	if err != nil {
		errorlog(w, "error PassVerify", http.StatusInternalServerError)
		return
	}
	if !ok {
		errorlog(w, "Invalid username/password", http.StatusBadRequest)
		return
	}
	ctime := time.Now().UTC().Add(time.Minute * time.Duration(tokentime)).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// claims["Id"] = connTable[req.UserName].Token
	claims["username"] = req.UserName
	claims["exp"] = ctime
	fmt.Println(claims)
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(tokenSecretKey))
	connTable[req.UserName] = &ConnInfo{nil, tokenString, ctime}

	if err != nil {
		http.Error(w, "Internal Server Error (jwt Encoder)", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("ws://localhost:8080/chat?token=%s", connTable[req.UserName].Token)
	resp := &LogResponse{url}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-Rate-Limit", strconv.Itoa(callperhour))
	w.Header().Add("X-Expires-After", strconv.Itoa(int(connTable[req.UserName].ExpireAt)))
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		errorlog(w, "Internal Server Error (json Encoder)", http.StatusInternalServerError)
		return
	}
}
