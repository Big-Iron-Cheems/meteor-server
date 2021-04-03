package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"meteor-server/pkg/auth"
	"meteor-server/pkg/core"
	"meteor-server/pkg/db"
)

func LoginHandler(c *gin.Context) {
	token, err := auth.Login(c.Query("name"), c.Query("password"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong name or password."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func LogoutHandler(c *gin.Context) {
	auth.Logout(core.GetAccountID(c))
	c.JSON(http.StatusOK, gin.H{})
}

func AccountInfoHandler(c *gin.Context) {
	account, err := db.GetAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get account."})
		return
	}

	c.JSON(http.StatusOK, account)
}

func McAccountHandler(c *gin.Context) {
	// Get Minecraft UUID
	/*id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID."})
		return
	}

	// Get account
	account, err := db.GetAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get account."})
		return
	}

	if c.Request.Method == "POST" { // Add Minecraft account
		err = account.AddMcAccount(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else { // Remove Minecraft account
		account.RemoveMcAccount(id)
	}

	c.JSON(http.StatusOK, gin.H{})*/
}
