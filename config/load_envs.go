package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type ConfMap map[string]interface{}
var Conf ConfMap = ConfMap{}

func LoadEnvs() error {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Println("LoadEnvs - error loading .env")
        return err
    }

    fmt.Println("LoadEnvs - .env loaded")
    Conf["jwt_secret"]               = []byte(os.Getenv("jwt_secret"))
    Conf["jwt_expiration_time"], err = strconv.Atoi( os.Getenv("jwt_expiration_time"))
    db_pg_username                  := os.Getenv("db_pg_username")
    db_pg_password                  := os.Getenv("db_pg_password")
    db_pg_host                      := os.Getenv("db_pg_host")
    db_pg_port                      := os.Getenv("db_pg_port")
    db_pg_name                      := os.Getenv("db_pg_name")
    Conf["db_url"]                   =  "postgres://" +
                                                        db_pg_username + ":" +
                                                        db_pg_password + "@" +
                                                        db_pg_host + ":" +
                                                        db_pg_port + "/" +
                                                        db_pg_name
    // fmt.Printf("db_url: %v\n", Conf["db_url"])
    fmt.Println("LoadEnvs - Vars assigned")
    fmt.Println("LoadEnvs - Check DB")
    err = testdb()
    return err
}


func testdb() error {

    conn, err := PGconn()
    if err != nil {
        fmt.Printf("Config - Login - error on connecting to DB, %v", err)
    }

    defer conn.Close(context.Background())

    var query_res string
    err = conn.QueryRow( context.Background(),
                         "select 'connected'",
                         // "select user_id from tb_users limit 1",
                         ).Scan(&query_res)
    if err != nil {
        fmt.Printf("jwt - Login - error validating user: %v", err)
        return err
    }
    fmt.Printf("DB is %v", query_res)
    return nil
}



