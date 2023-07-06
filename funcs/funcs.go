package funcs

//func (cart Objects.Transaction)Add_items(name string, qty uint32, price uint64)
import (
	"database/sql"
	"fmt"
)

/*type bday struct {
	id    int
	name  string
	year  int
	month int
	day   int
}*/

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "windowslogin123"
	dbname   = "postgres"
)

//var connStr = "postgresql://postgres:windowslogin123@localhost:5432/postgres/todos?sslmode=disable"

/* func connect() (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}*/

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Readquery() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}
