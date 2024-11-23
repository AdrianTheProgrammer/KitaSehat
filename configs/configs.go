package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ENV struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
	jwtkey   string
	supabase string
}

func DBConnect() (*gorm.DB, ENV) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return nil, ENV{}
	}

	env := ENV{
		user:     os.Getenv("user"),
		password: os.Getenv("password"),
		host:     os.Getenv("host"),
		port:     os.Getenv("port"),
		dbname:   os.Getenv("dbname"),
		jwtkey:   os.Getenv("jwtkey"),
		supabase: os.Getenv("supabase"),
	}

	connectionString := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v", env.user, env.password, env.host, env.port, env.dbname)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return db, env
}
