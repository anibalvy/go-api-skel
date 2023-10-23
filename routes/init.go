package routes

import (
	"api_skel/config"
    "api_skel/routes/v1"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    api_v1 := app.Group("v1")
    api_v1.Use(jwtware.New(jwtware.Config{
        // TokenLookup: "header:Authorization",  //is the default, not need to declare, use auth scheme Bearer, authScheme param is only for header
        // TokenLookup: "cookie:jwt,header:Authorization",  // to allow seamless authorization from browsers
        // TokenLookup: "cookie:jwt",  // to allow seamless authorization from browsers
        // TokenLookup: "query:jwt",  //  other method
        // TokenLookup: "param:jwt",  // other method
        TokenLookup: "header:Authorization,cookie:jwt",  // to allow multiple authorization methods
        AuthScheme:  "Bearer",
		SigningKey:  jwtware.SigningKey{
        JWTAlg:      jwtware.HS256,
        Key:         config.Conf["jwt_secret"]}, // []byte( config.Conf["jwt_secret"].(string)

    }))

    api_v1.Get("users", v1.GetUserList)


}

