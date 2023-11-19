package domain

import (
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"

	"fmt"
	"time"
)

var secretkey string = "secret"

type UserRepositoryDb struct {
	db *sqlx.DB
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func (d UserRepositoryDb) SignUp(user User) error {
	user.Password, _ = GeneratehashPassword(user.Password)

	// Set default role if user.Role is empty
	if user.Role == "" {
		user.Role = "user"
	}

	_, err := d.db.Exec(`
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
	`, user.Name, user.Email, user.Password, user.Role)

	if err != nil {
		return err
	}

	return nil
}

func (d UserRepositoryDb) SignIn(email, password string) (string, error) {
	var storedPassword string
	var storedRole string

	// Retrieve the stored password and role for the given email
	err := d.db.QueryRow(`
		SELECT password, role FROM users
		WHERE email = $1
	`, email).Scan(&storedPassword, &storedRole)

	if err != nil {
		fmt.Println("User not found or another database error")
		return "", err // User not found or another database error
	}

	// Verify the provided password against the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		fmt.Println("Password does not match")
		return "", err // Passwords do not match
	}

	// Authentication successful

	// Generate JWT token
	token, err := GenerateJWT(email, storedRole)
	if err != nil {
		fmt.Println("Failed to generate JWT token")
		return "", err
	}

	// You may want to use the retrieved role for further authorization logic.
	// For now, let's print the role as an example.
	fmt.Println("User Role:", storedRole)
	fmt.Println(token)

	return token, nil
}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func (d UserRepositoryDb) IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			fmt.Println("No Token Found")
			http.Error(w, "No Token Found", http.StatusUnauthorized)
			return
		}

		fmt.Println("Token Found")

		// Extract the token from the "Bearer <token>" format
		tokenString := strings.Split(tokenHeader, " ")[1]

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			fmt.Println("Error parsing token:", err)
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("Valid Token")

			if role, exists := claims["role"].(string); exists {
				r.Header.Set("Role", role)
			}

			handler.ServeHTTP(w, r)
		} else {
			fmt.Println("Invalid Token")
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
		}
	}
}

func NewUserRepositoryDb(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{db: dbClient}
}
