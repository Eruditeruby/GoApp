package main

import (
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
       // insertStmt := `insert into "users"("name", "age") values('Jacob', 20)`
       // _, e := db.Exec(insertStmt)
       // CheckError(e)

	rows, err := db.Query(`SELECT "name", "age" FROM "users"`)
        CheckError(err)
 
       defer rows.Close()
       for rows.Next() {
       var name string
       var age int
 
       err = rows.Scan(&name, &age)
       CheckError(err)
 
    fmt.Println(name, age)
       }
 
     CheckError(err)
   }




      func CheckError(err error) {
        if err != nil {
        panic(err)
       }
  }
