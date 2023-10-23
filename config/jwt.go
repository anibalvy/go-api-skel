package config

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// type User struct {
//     Username string `json:"username"`
//     Password string `json:"password"`
// }

func Get_token(c *fiber.Ctx) error {

    payload := struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }{}
    if err := c.BodyParser(&payload); err != nil {
        return c.SendStatus( fiber.StatusUnauthorized)
    }

    conn, err := PGconn()
    if err != nil {
        fmt.Printf("jwt - Login - error on connecting to DB, %v", err)
    }
    defer conn.Close(context.Background())


    var query_res string
    err = conn.QueryRow( context.Background(),
                         "select public.fn_user_validate($1, $2)",
                         payload.Username,
                         payload.Password).
                         Scan(&query_res)
    if err != nil {
        fmt.Printf("jwt - Login - error validating user: %v", err)
        return c.SendStatus( fiber.StatusUnauthorized)  // I'm not giving info to user because sec.
    }

    var user_validation map[string]interface{}
    json.Unmarshal( []byte(query_res), &user_validation)
    if user_validation["valid"] != true {
        fmt.Printf("jwt - Login - error validating user: %v", err)
        return c.SendStatus( fiber.StatusUnauthorized)
    }

    // Token creation
    // make claim
    // user_expiration_time := user_validation["options"]["token_expiration_time"]
    user_expiration_time := Conf["jwt_expiration_time"].(int)

    expiration := time.Now().Add(time.Hour * time.Duration( user_expiration_time))
    claims := jwt.MapClaims{
        "name": user_validation["username"],
        "rol":  user_validation["rol"],
        "exp":  expiration.Unix(),
    }
    // create token
    new_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // encoded token
    token, err := new_token.SignedString(Conf["jwt_secret"])
    if err != nil {
        fmt.Printf("jwt - Login - error creating token: %v", err)
        return c.SendStatus(fiber.StatusInternalServerError)

    }

    return c.JSON(fiber.Map{"token": token, "exp": Conf["jwt_expiration_time"].(int)})
}

func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
