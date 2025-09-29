package main

import "os"

const (
	// PORT             = rand.IntN(65535) `This will not work as Golang returns a value error namely
	//                                         [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.'
	PORT            = 8000 // Adding server spinup on random for erytime, avoid attacks
	INVALID_API_MSG = "Invalid API KEY"
	LOG_EXP_DELTA   = 86400 //24*60*60
	OTP_SIZE        = 4
	API_KEY         = "f902b6d8-cfc9-4e62-84aa-1d5c44fe27ec"
)

var (
	DATABSE_URL = os.Getenv("DATABASE_URL")
	REDIS_URL   = os.Getenv("REDIS_URL")
)
