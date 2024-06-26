package app

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	"github.com/fadedreams/gowallet/domain"
	"github.com/fadedreams/gowallet/service"
	"github.com/jmoiron/sqlx"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

func Start() {

	router := mux.NewRouter()
	dbClient := getDbClient()
	//dbGormClient := getDbGormClient()

	//ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryStub())}
	ch := ClientHandlers{service.NewClientService(domain.NewClientRepositoryDb(dbClient))}
	wh := WalletHandlers{service.NewWalletService(domain.NewWalletRepositoryDb(dbClient))}
	uh := UserHandlers{service.NewUserService(domain.NewUserRepositoryDb(dbClient))}

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/clients", ch.getAllClients).Methods(http.MethodGet)
	router.HandleFunc("/clients", ch.createClient).Methods(http.MethodPost)

	router.HandleFunc("/wallets", wh.newWallet).Methods(http.MethodPost)
	router.
		HandleFunc("/clients/{client_id:[0-9]+}/wallet/{wallet_id:[0-9]+}", wh.MakeTransaction).
		Methods(http.MethodPost)

	router.HandleFunc("/users", uh.createUser).Methods(http.MethodPost)
	router.HandleFunc("/login", uh.loginUser).Methods(http.MethodPost)

	//router.HandleFunc("/user", IsAuthorized(UserIndex)).Methods("GET")
	router.HandleFunc("/user", uh.IsAuth(UserIndex)).Methods("GET")

	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))

}

func createClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["client_id"])
}

var (
	dbClient     *sqlx.DB
	dbClientOnce sync.Once
)

func getDbClient() *sqlx.DB {
	dbClientOnce.Do(func() {
		// Update the driver name and connection string for PostgreSQL
		// Log the connection string
		log.Println("Connecting to database with connection string:", "user=postgres password=postgres dbname=dbwallet sslmode=disable")
		client, err := sqlx.Open("postgres", "user=postgres password=postgres dbname=dbwallet sslmode=disable")
		if err != nil {
			panic(err)
		}
		// See "Important settings" section.
		client.SetConnMaxLifetime(time.Minute * 3)
		client.SetMaxOpenConns(10)
		client.SetMaxIdleConns(10)

		dbClient = client
	})
	return dbClient
}

//func getDbGormClient() *gorm.DB {
//db, err := gorm.Open(postgres.New(postgres.Config{DSN: "host=localhost user=postgres password=postgres dbname=dbwallet port=5432 sslmode=disable TimeZone=Asia/Shanghai"}), // data source name, refer https://github.com/jackc/pgx PreferSimpleProtocol: true,                                                                                                          // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
//&gorm.Config{})

//if err != nil {
//log.Fatalln("Invalid database url")
//}

//fmt.Println("Database connection successful.", db)
//return db

//}
