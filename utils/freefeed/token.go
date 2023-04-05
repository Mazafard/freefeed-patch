package freefeed

import (
	"encoding/json"
	"feed/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"strings"
)

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	user, err := RequestByToken(tokenString)
	if err != nil {
		return err
	}
	id, err := uuid.FromString(user.Users.ID)
	if err != nil {
		return err
	}

	c.Set("UserID", user.Users.ID)
	userContext := models.User{
		Username: user.Users.Username,
		ID:       id,
	}
	models.DB.FirstOrCreate(&userContext)

	return nil
}

// Extract token from Authorization header
func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return bearerToken
	}
	return ""
}

func RequestByToken(t string) (u models.UserData, err error) {
	url := "https://freefeed.net/v2/users/whoami"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return u, err
	}

	req.Header.Set("Authorization", t)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return u, err
	}

	defer resp.Body.Close()
	fmt.Println("Response Status Code:", resp.StatusCode)
	var user models.UserData
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error:", err)
		return u, err
	}

	return user, nil

}
