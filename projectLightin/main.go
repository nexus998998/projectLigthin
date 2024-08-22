package main 


import (
  "fmt"
  "net/http"
  "database/sql"
  _"github.com/mattn/go-sqlite3" 
)
var err error
var database *sql.DB


func HandleErr(w http.ResponseWriter , message string ) {
  if err != nil {
    fmt.Fprint(w , message , err)
  }
}

func loginRoute(w http.ResponseWriter , r *http.Request ) {
  if r.Method == "POST" {
    err = r.ParseForm()
    HandleErr("error , failed to parse form : " , err)
    

    usernameInputted := r.FormValue("username")
    passwordInputted := r.FormValue("password")
    
    var username string 
    var password string 
    var UserID int  
    row := database.QueryRow("
    SELECT username , password , userID FROM Users WHERE username = ? , password = ? ; 
    " , usernameInputted , passwordInputted)
    

    err = row.Scan(&username , &password , &UserID )
    
    if err != nil {
      if err == sql.ErrNoRows {
        fmt.Fprint(w , "incorrect information ")
        return 
      }
      checkErr(w , "error querying row : ")
      return
    }
    cookie := http.Cookie{
      Name : 
    }




    return
  }


  return
}

func mainRoute(w http.ResponseWriter , r *http.Request) {
  fmt.Fprint(w , "hello world") 
}


func main() {
  databse , err = sql.Open("sqlite3" , "main.db")
  defer database.Close() 
  checkErr("cannot open database : " , err )
  


  var MuxHandler http.ServeMux 
  MuxHandler.HandleFunc("/" , mainRoute)
  

  var server = http.Server {
    Handler : &MuxHandler , 
    Addr : ":8080",
  }

  err = server.ListenAndServe() 
  if err != nil {
    fmt.Println("error cannot init server : " , err )
  }

}
