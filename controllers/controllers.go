package controllers

import (
	"log"
	"net/http"
	"p2src/database"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       string
	Username string
	Link     string
}

func Home(c *gin.Context) {
	var user User
	var userSAR []User
	var userVscript []User
	username := c.Request.URL.Query().Get("user")
	sql := `SELECT * FROM "sar" WHERE username = $1`
	rows, err := database.DB.Query(sql, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Link)
		if err != nil {
			log.Fatal(err)
		}
		userSAR = append(userSAR, user)
	}
	sql = `SELECT * FROM "vscript" WHERE username = $1`
	rows, err = database.DB.Query(sql, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Link)
		if err != nil {
			log.Fatal(err)
		}
		userVscript = append(userVscript, user)
	}
	c.HTML(http.StatusOK, "infractions.html", gin.H{
		"username":               username,
		"stateSAR":               len(userSAR) != 0,
		"stateVscript":           len(userVscript) != 0,
		"infractionCountSAR":     len(userSAR),
		"infractionCountVscript": len(userVscript),
		"userSAR":                userSAR,
		"userVscript":            userVscript,
	})
}

func Add(c *gin.Context) {
	username := c.Request.URL.Query().Get("user")
	link := c.Request.URL.Query().Get("link")
	itype := c.Request.URL.Query().Get("type")
	status := false
	if username != "" && link != "" && itype != "" {
		if itype == "sar" {
			sql := `INSERT INTO "sar" (username, link) VALUES ($1, $2);`
			database.DB.Exec(sql, username, link)
			status = true
		} else if itype == "vscript" {
			sql := `INSERT INTO "vscript" (username, link) VALUES ($1, $2);`
			database.DB.Exec(sql, username, link)
			status = true
		}
	}
	c.HTML(http.StatusOK, "infractions-submit.html", gin.H{
		"status":   status,
		"username": username,
		"link":     link,
		"type":     itype,
	})
}
