package domain

// Import the PostgreSQL driver
import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	//"time"
)

type ClientRepositoryDb struct {
	client *sqlx.DB
}

func (d ClientRepositoryDb) FindAll() ([]Client, error) {
	findAllSql := "select client_id, name, city, zipcode,  status from clients"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying clients table " + err.Error())
		return nil, err
	}

	clients := make([]Client, 0)
	for rows.Next() {
		var c Client
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while scanning clients " + err.Error())
			return nil, err
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func (d ClientRepositoryDb) IsClientTableExists() (bool, error) {
	// Check if the clients table exists in the database
	var tableExists bool
	err := d.client.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM   information_schema.tables
			WHERE  table_name = 'clients'
		)
	`).Scan(&tableExists)

	if err != nil {
		log.Println("Error checking table existence:", err)
		return false, err
	}

	return tableExists, nil
}

func (d ClientRepositoryDb) CreateClient(newClient Client) error {
	// Check if the table exists
	tableExists, err := d.IsClientTableExists()
	if err != nil {
		log.Println("Error checking table existence:", err)
		return err
	}

	if !tableExists {
		// Create the clients table if it does not exist
		_, err := d.client.Exec(`
			CREATE TABLE clients (
				id SERIAL PRIMARY KEY,
				name VARCHAR(100),
				city VARCHAR(100),
				zipcode VARCHAR(20),
				status VARCHAR(20)
			)
		`)
		if err != nil {
			log.Println("Error creating clients table:", err)
			return err
		}

		log.Println("Clients table created successfully")
	}

	// Insert the new client into the clients table
	_, err = d.client.Exec(`
		INSERT INTO clients (name, city, zipcode, status)
		VALUES ($1, $2, $3, $4)
	`, newClient.Name, newClient.City, newClient.Zipcode, newClient.Status)

	if err != nil {
		log.Println("Error creating client:", err)
		return err
	}

	log.Println("Client created successfully")
	return nil
}

func NewClientRepositoryDb(dbClient *sqlx.DB) ClientRepositoryDb {
	// Update the driver name and connection string for PostgreSQL
	// Log the connection string
	//log.Println("Connecting to database with connection string:", "user=postgres password=postgres dbname=dbwallet sslmode=disable")
	//client, err := sql.Open("postgres", "user=postgres password=postgres dbname=dbwallet sslmode=disable")
	//if err != nil {
	//panic(err)
	//}
	//// See "Important settings" section.
	//client.SetConnMaxLifetime(time.Minute * 3)
	//client.SetMaxOpenConns(10)
	//client.SetMaxIdleConns(10)

	return ClientRepositoryDb{dbClient}
}
