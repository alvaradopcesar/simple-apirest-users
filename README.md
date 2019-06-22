## API REST SIMPLE USERS
Apirest simple of users with go (golang) and mysql

<center>

![image](https://i.ya-webdesign.com/images/golang-vector-1.png)

</center>

#### INSTALLATION & CONFIGURATION

```go
type Database struct {
    User string `json:"user"`
    Dbname string `json:"dbname"`
    Password string `json:"password"`
}

func DatabaseConfig() *sql.DB {
    var database Database
    database.User = "root"
    database.Dbname = "apirest_users"
    database.Password = "lavidaesunasola2018"

    db, err := sql.Open("mysql", fmt.Sprinf("%s:%s@/%s?charset=utf8", database.User, database.Dbname, database.Password))

    if err != nil {
        log.Fatal("Error al conectar a la database: ", err.Error())
    }

    return db
}

```
