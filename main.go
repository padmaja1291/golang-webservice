package main
import (
  "fmt"
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
)

func getToken()(string, string){
	currentTime := time.Now()
	formattedTime := currentTime.Format("1504")
	previousMin := currentTime.Add(-1*time.Minute).Format("1504")
	fmt.Println("\nformattedTime",formattedTime)
	fmt.Println("\npreviousMin",previousMin)
	return formattedTime, previousMin

}

func PostMethod(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token := c.PostForm("token")

	user := "c137@onecause.com"
	pass := "#th@nH@rm#y#r!$100%D0p#"
	formattedTime, previousMin := getToken();
	
	if (username == user && password == pass && (token == formattedTime || token==previousMin)){
		c.JSON(http.StatusOK , gin.H{"status":  "verified"})
	}else{
		c.JSON(401, gin.H{"error": "invalid user"})
	}
  }

func main() {
	router := gin.Default()
	router.POST("/", PostMethod)
	listenPort := "8081"
	router.Run(":"+listenPort)
  }