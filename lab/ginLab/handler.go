package ginLab

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {

	// c.JSON(http.StatusOK, gin.H{"message": "This is Gin Lab"})
	c.String(http.StatusOK, "This is Gin Lab")
}

func showParameter(c *gin.Context) {

	user := c.Param("user")
	c.String(http.StatusOK, "User: %s", user)

}

func showQuery(c *gin.Context) {

	// user := c.Query("user")
	user := c.DefaultQuery("user", "Peter")
	c.String(http.StatusOK, "User: %s", user)

}

// type rBody struct {
// 	A string
// }

// func (r rBody) String() string {
// 	return fmt.Sprintf("A: %s", r.A)
// }

func showBody(c *gin.Context) {

	var body map[string]interface{}
	err := c.Bind(&body)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Body:", body)
	}

	// rBody := rBody{}
	// err := c.BindJSON(&rBody)
	// if err != nil {
	// 	fmt.Println("Error:", err.Error())
	// } else {
	// 	fmt.Println("Body:", rBody)
	// }

	c.String(200, "success")

}
