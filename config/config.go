package config

import "os"

func init() {
	SERVER_IP = os.Getenv("SERVER_IP")
	ECON_PORT = os.Getenv("ECON_PORT")
	ECON_PASSWORD = os.Getenv("ECON_PASSWORD")
	DEBUG_MODE = true
	SERVE_API = true

	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DB = os.Getenv("MYSQL_DB")
}

var (
	// SERVER_IP is the IP to the machine running the teeworlds server
	SERVER_IP = os.Getenv("SERVER_IP")
	// ECON_PORT is the port to econ (set via ec_port in autoexec.cfg)
	ECON_PORT = os.Getenv("ECON_PORT")
	// ECON_PASSWORD is the password to authenticate (set via ec_password)
	ECON_PASSWORD = os.Getenv("ECON_PASSWORD")

	// DEBUG_MODE gives additional logs
	DEBUG_MODE = true

	// SERVE_API will initialize a web server on port 8000 that serves a websocket connection
	SERVE_API = true

	// MYSQL_USER username to mysql database
	MYSQL_USER = os.Getenv("MYSQL_USER")
	// MYSQL_PASSWORD password to mysql database
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	// MYSQL_DB is the name to the database
	MYSQL_DB = os.Getenv("MYSQL_DB")
)
