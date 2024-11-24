package main

import (
	"database/sql"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	. "main/.gen/pankan_db/portal/table"
	"os"

	"main/.gen/pankan_db/portal/model"
)

func isProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "pankan_db"
)

func query(db *sql.DB) []struct{ model.InteractiveRooms } {
	stmt := SELECT(
		InteractiveRooms.ID,
		InteractiveRooms.DisplayName,
		InteractiveRooms.CreatedAt,
		InteractiveRooms.Password,
	).FROM(
		InteractiveRooms,
	)
	var dest []struct {
		model.InteractiveRooms
	}
	err := stmt.Query(db, &dest)
	if err != nil {
		panic(err)
	}
	return dest
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db := ConnectToDB()
	defer func(db *sql.DB) {
		panicOnError(db.Close())
	}(db)

	testPassword := "no"
	// panicOnError(insertRoom("test", testPassword, db))

	v := query(db)
	for _, v := range v {
		fmt.Println(
			fmt.Sprintf("id=%s, display_name=\"%s\", created_at=\"%s\", valid_password=%t", v.ID, v.DisplayName, v.CreatedAt, verifyPassword(testPassword, v.Password)),
		)

	}
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 14)
}

func verifyPassword(givenPassword string, hashedPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(givenPassword))
	return err == nil
}

func insertRoom(displayName string, password string, db *sql.DB) error {
	var insertResult []struct{}
	hashedPassword, err := hashPassword(password)
	panicOnError(err)
	return InteractiveRooms.INSERT(
		InteractiveRooms.DisplayName,
		InteractiveRooms.Password,
	).VALUES(
		displayName,
		hashedPassword,
	).Query(db, &insertResult)
}

func ConnectToDB() *sql.DB {
	var connectString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", connectString)
	panicOnError(err)
	return db
}
