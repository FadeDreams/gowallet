package app

import (
	"encoding/json"
	//"encoding/xml"
	"fmt"
	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/service"
	"net/http"
	//"github.com/golang-jwt/jwt"
	//"strings"
	//"github.com/jmoiron/sqlx"
	//"golang.org/x/crypto/bcrypt"
	//"time"
)

var secretkey string = "secret"

type UserHandlers struct {
	service service.IUserService
}

type User struct {
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (ch *UserHandlers) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = ch.service.SignUp(domain.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Role:     newUser.Role,
	})

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created successfully")
}

func (ch *UserHandlers) loginUser(w http.ResponseWriter, r *http.Request) {
	var loginDetails User
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&loginDetails)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	email := loginDetails.Email
	password := loginDetails.Password

	// Authenticate user
	token, err := ch.service.SignIn(email, password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Return the JWT token in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func (ch *UserHandlers) IsAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authenticate using the service method
		_ = ch.service.IsAuthorized(handler)
	}
}

//func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
//return func(w http.ResponseWriter, r *http.Request) {
//tokenHeader := r.Header.Get("Authorization")

//if tokenHeader == "" {
//fmt.Println("No Token Found")
//http.Error(w, "No Token Found", http.StatusUnauthorized)
//return
//}

//fmt.Println("Token Found")

//// Extract the token from the "Bearer <token>" format
//tokenString := strings.Split(tokenHeader, " ")[1]

//var mySigningKey = []byte(secretkey)

//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//return nil, fmt.Errorf("There was an error in parsing token.")
//}
//return mySigningKey, nil
//})

//if err != nil {
//fmt.Println("Error parsing token:", err)
//http.Error(w, "Invalid Token", http.StatusUnauthorized)
//return
//}

//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//fmt.Println("Valid Token")

//if role, exists := claims["role"].(string); exists {
//r.Header.Set("Role", role)
//}

//handler.ServeHTTP(w, r)
//} else {
//fmt.Println("Invalid Token")
//http.Error(w, "Invalid Token", http.StatusUnauthorized)
//}
//}
//}
//func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
//return func(w http.ResponseWriter, r *http.Request) {
//tokenHeader := r.Header.Get("Authorization")

//if tokenHeader == "" {
//fmt.Println("No Token Found")
//http.Error(w, "No Token Found", http.StatusUnauthorized)
//return
//}

//fmt.Println("Token Found")

//// Extract the token from the "Bearer <token>" format
//tokenString := strings.Split(tokenHeader, " ")[1]

//var mySigningKey = []byte(secretkey)

//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//return nil, fmt.Errorf("There was an error in parsing token.")
//}
//return mySigningKey, nil
//})

//if err != nil {
//fmt.Println("Error parsing token:", err)
//http.Error(w, "Invalid Token", http.StatusUnauthorized)
//return
//}

//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//fmt.Println("Valid Token")

//if role, exists := claims["role"].(string); exists {
//r.Header.Set("Role", role)
//}

//handler.ServeHTTP(w, r)
//} else {
//fmt.Println("Invalid Token")
//http.Error(w, "Invalid Token", http.StatusUnauthorized)
//}
//}
//}
