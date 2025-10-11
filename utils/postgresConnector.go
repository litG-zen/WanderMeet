package utils

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	REGISTERED_USER_QUERY = "select * from users where phone = $1"
)

var (
	DATABASE_URL = "postgres://postgres:lalit1234@localhost:5432/wandermeet?sslmode=disable" //os.Getenv("DATABASE_URL")
	DbClient     *DBClient
)

type DBClient struct {
	conn  *pgx.Conn
	ctx   context.Context
	mutex sync.RWMutex
}

func (db *DBClient) initialize() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	fmt.Printf("DATABASE_URL :%v", DATABASE_URL)
	if DATABASE_URL == "" {
		return fmt.Errorf("DATABASE	 URL is not defined")
	}

	conn, err := pgx.Connect(db.ctx, DATABASE_URL)
	if err != nil {
		return fmt.Errorf("Error connecting to DB %v", err)
	}

	db.conn = conn
	return nil
}

func NewDBConnection() (*DBClient, error) {
	var initErr error
	redisOnce.Do(func() {
		DbClient = &DBClient{ctx: context.Background()}
		initErr = DbClient.initialize()
	})

	if initErr != nil {
		return nil, initErr
	}
	fmt.Printf("\n****Database connection Initiated****\n")
	return DbClient, nil
}

func GetDBInstance() (*DBClient, error) {
	if DbClient == nil {
		return NewDBConnection()
	}
	return DbClient, nil
}

func CloseDBConnection() {
	DbClient.conn.Close(DbClient.ctx)
	fmt.Printf("\n****Database connection Closed****\n")

}

// Database specific query calls.

// 1 - Check existing phone number when the user is trying to register.
func IsExistingUser(phone string) error {
	db_inst, err := GetDBInstance()
	fmt.Printf("Connection ok!")
	if err != nil {
		fmt.Errorf("Error while accessing Database", err)
		return err
	} else {
		row := db_inst.conn.QueryRow(db_inst.ctx, REGISTERED_USER_QUERY).Scan(phone)
		fmt.Printf("\n Row data%v", row)
	}
	return nil
}
