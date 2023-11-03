package main

import (
	// "fmt"
	// "os"
	"testLogin/database"
	"testLogin/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func login(c *gin.Context) {
	var req models.LoginJson
	var logstat models.LoginStatus
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	Username := req.Username
	Password := req.Password

	db := database.DBConn
	rows, err := db.Query("SELECT * FROM UserInfo WHERE username = $1", Username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var dbusername string
	var dbpassword string

	for rows.Next() {
		err := rows.Scan(&dbusername, &dbpassword)
		if err != nil {
			panic(err)
		}
		if dbusername == Username {
			if dbpassword == Password {
				//success
				logstat.Username = req.Username
				logstat.Message = "Success"
				logstat.Success = true
			} else {
				//password incorrect
				logstat.Username = req.Username
				logstat.Message = "Password incorrect"
				logstat.Success = false
			}
		}
	}
	// logstat.User = req
	// logstat.Message = "Username not found"
	// logstat.Success = false
	// println("Username not found")

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	c.JSON(200, logstat)
	return
}

func register(c *gin.Context) {
	var req models.RegistrationJson
	var response models.RegistrationStatus
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	var username string
	username = req.Username
	var password string
	password = req.Password

	sqlStatement := `
		INSERT INTO UserInfo (username, password)
		VALUES ($1, $2)`

	db := database.DBConn
	_, err := db.Exec(sqlStatement, username, password)
	if err != nil {
		response.Username = req.Username
		response.Success = false
		response.Message = "Registration failed db error"
	}

	response.Username = req.Username
	response.Success = true
	response.Message = "Registration Successful"

	c.JSON(200, response)
	return
}

func main() {
	err := database.InitDB()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	router.Use(cors.New(config))

	router.POST("/register", register)
	router.POST("/login", login)

	router.Run(":5000")
}
