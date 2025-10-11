package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/litG-zen/WanderMeet/auth"
	"github.com/litG-zen/WanderMeet/logs"
	"github.com/litG-zen/WanderMeet/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// created a simple GET API using gin
// reference site : https://gin-gonic.com/en/docs/quickstart/

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"owner":   "Lit",
	})

	log_string := fmt.Sprintf("%v %v %v", time.Now(), c.FullPath(), http.StatusOK)
	logs.Logger(log_string, false)
}

func AsyncAPIHandler(c *gin.Context) {
	fmt.Println("\n\n Async API hander called")

	token := c.GetHeader("API-KEY")

	if token != API_KEY {
		c.JSON(http.StatusUnauthorized, gin.H{"message": INVALID_API_MSG})
		log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusUnauthorized, INVALID_API_MSG)
		logs.Logger(log_string, true)
		return
	}

	go func() {
		// I have added this goroutine to mimick a blocking function like email-sending, heavy calculation
		// which is usually triggered in an async way, this is that attempt.
		time.Sleep(10000 * time.Millisecond)
		print("\n\n\n Sleep ended \n\n\n")
	}()

	c.JSON(http.StatusOK, gin.H{
		"message": "this is a successful response of AsyncAPI call",
	})

	log_string := fmt.Sprintf("%v %v %v %v", time.Now(), c.FullPath(), http.StatusUnauthorized, "SUCCESS")
	logs.Logger(log_string, false)
}

func NearbyUsersFetch(c *gin.Context) {
	/*
		Flow :
		  Constraints check
		   - parse JWT token to get the userid
		   - check for user_details presence in DB
		   - fetch the nearby users in configurable radius, based on user's pro-version from GeoRedis/Postgis(depends)
		   - return the details in response
	*/
	utils.GetRedisInstance()
	c.JSON(http.StatusOK, gin.H{
		"message": "nearby_users available",
	})
}

func RegisterUser(c *gin.Context) {
	var validator RegistrationBody

	// Bind Json to struct
	if err := c.ShouldBindBodyWith(&validator, binding.JSON); err != nil {
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

		// phone number already exists check
		fmt.Printf("Calling check")
		reg_err := utils.IsExistingUser(validator.Phone)
		if reg_err != nil {
			fmt.Errorf("Registration error check", reg_err)
		}

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
}
