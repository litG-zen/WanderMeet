package main

type RegistrationBody struct {
	Name          string   `json:"name"`
	City          string   `json:"city"`
	Placesvisited []string `json:""places_visited"`
	Age           int      `json:"age"`
	Hobbies       []string `json:"hobbies"`
	Email         string   `json"email"`
	Phone         string   `json:"phone"`
	Latitude      float32  `json:"latitude"`
	Longitude     float32  `json:"longitude"`
}

type LoginBody struct {
	Phone     string  `json:"phone"`
	Otp       int32   `json"otp"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
