package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/litG-zen/WanderMeet/auth"
	"github.com/litG-zen/WanderMeet/logs"

	"github.com/gin-gonic/gin"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func main() {
	// Setting DEBUG mode ON
	gin.SetMode(gin.DebugMode)

	app := gin.Default()

	// Defining folder to load templates from
	app.LoadHTMLGlob("templates/*")
	app.Static("/media", "./media")

	app.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "rendering HTML",
		})

		log_string := fmt.Sprintf("%v %v %v", time.Now(), c.FullPath(), http.StatusOK)
		logs.Logger(log_string, false)
	})

	app.GET("/ping", PingHandler)

	app.GET("/health", func(c *gin.Context) { //anonymous function approach

		greetings := []string{
			"Hello mate!",
			"Namaste",
			"Hii",
			"I am healthy and I know it.",
		}

		c.JSON(http.StatusOK, gin.H{
			"message": greetings[rand.Intn(len(greetings))],
		})
	})

	app.GET("/async/api", AsyncAPIHandler)

	app.POST("/register", func(c *gin.Context) {
		var validator RegistrationBody

		// Bind Json to struct
		if err := c.ShouldBindBodyWithJSON(&validator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusAccepted, "invalid payload")
			logs.Logger(log_string, true)
			return
		} else {
			// ToDo : Add user registration flow.
			/*

				Data Constains are satisfied at JSON Binding stage itself.

				Few check before proceeding with registration.
				  - Phone number already exists
				  - Empty name
				  - Email validation using regex
				  - Gender to be Male/Female (sorry I don't support non-binaries.)
				  - Places visited array should not have empty strings
				  - lat/long is valid

				Registration Process
				  - User table entry creation
				  - Set latlong to redis for later geo-redis queries
				  - Send a welcome email and otp email to the user
				  - Send an OTP sms
				  - Log activity
			*/
			c.JSON(
				http.StatusOK,
				gin.H{
					"message": "data received",
					"data":    validator,
					"auth_tokens": gin.H{
						"access_token":  string(auth.GenerateAuthToken(1234)),
						"refresh_token": string(auth.GenerateRefreshToken(1234)),
					},
				})
			log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusAccepted, "Valid user")
			logs.Logger(log_string, false)

		}
	})

	app.POST("/login", func(c *gin.Context) {
		var validator LoginBody
		if err := c.ShouldBindBodyWithJSON(&validator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusAccepted, "invalid payload")
			logs.Logger(log_string, true)
			return
		} else {
			// ToDo : Add user login flow.
			/*

				Data Constains are satisfied at JSON Binding stage itself.

				Few check before proceeding with registration.
				  - Phone number is valid and already exists
				  - User is a registered user
				  - lat/long is valid
				  - Valid OTP of 4 digit.

				Login Process
				  - Login table entry creation
				  - Set latlong to redis for later geo-redis queries
				    - perform geoquery to get all nearby registered users.
				  - Log activity
			*/
			c.JSON(
				http.StatusOK,
				gin.H{
					"message": "data received",
					"data":    validator,
					"auth_tokens": gin.H{
						"access_token":  string(auth.GenerateAuthToken(1234)),
						"refresh_token": string(auth.GenerateRefreshToken(1234)),
					},
				})
			log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusAccepted, "Valid user")
			logs.Logger(log_string, false)
		}

	})

	app.POST("/get_nearby_users", NearbyUsersFetch)

	app.Run(":" + fmt.Sprint(PORT)) // listen and serve on 0.0.0.0:<random_port> (for windows "localhost:<rand_port>")
}
