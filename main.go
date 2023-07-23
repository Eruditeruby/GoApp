package main

import (
	 "net/http"
	"log"
	// "os"
        "fmt"
	 "database/sql"
    _ "github.com/lib/pq"
_ "github.com/jackc/pgx/v4/stdlib"
)

const (
    host     = "10.152.183.183"
    port     = 5432
    user     = "postgress"
    password = "postgress"
    dbname   = "user_db"
)


type sandbox struct {
//id int
name string
age int
}
// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello Ruby hi!</h1>"))
// }


var db *sql.DB
func main() {
 //        fmt.Println("Hello Ruby2")  
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "3000"
	// }

	// mux := http.NewServeMux()

	// mux.HandleFunc("/", indexHandler)
	// http.ListenAndServe(":"+port, mux)

  // code for connection

	 psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        //Open database
         db, err := sql.Open("postgres", psqlconn)
          CheckError(err)
     
        //Close database
        defer db.Close()
 
        // check db
        err = db.Ping()
        CheckError(err)
  
        fmt.Println("Connected!")

	  // insert
    // hardcoded
       insertStmt := `insert into "users"("name", "age") values('Jacob', 20)`
       _, e := db.Exec(insertStmt)
       CheckError(e)
   }

func retrieveRecord(w http.ResponseWriter, r *http.Request) {

// checks if the request is a "GET" request
if r.Method != "GET" {
http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
return
}

// We assign the result to 'rows'
rows, err := db.Query("SELECT * FROM users")

if err != nil {
http.Error(w, http.StatusText(500), http.StatusInternalServerError)
return
}
defer rows.Close()


// creates placeholder of the sandbox
snbs := make([]sandbox, 0)


// we loop through the values of rows
for rows.Next() {
snb := sandbox{}
err := rows.Scan(&snb.name, &snb.age)
if err != nil {
log.Println(err)
http.Error(w, http.StatusText(500), 500)
return
}
snbs = append(snbs, snb)
}

if err = rows.Err(); err != nil {
http.Error(w, http.StatusText(500), 500)
return
}

// loop and display the result in the browser
for _, snb := range snbs {
fmt.Fprintf(w, "%d %s %s %d\n", snb.name, snb.age)
}

}

      func CheckError(err error) {
        if err != nil {
        panic(err)
       }
  }
