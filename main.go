package main
import (
  "fmt"
  "time"
  "net/http"
  "github.com/gin-gonic/gin"
)

//getToken provides  time in the 2 digit hour and 2 digit minute format. eg. 1504
//function returns current time and a minute before that
func getToken()(string, string){
	currentTime := time.Now()
	formattedTime := currentTime.Format("1504")
	previousMin := currentTime.Add(-1*time.Minute).Format("1504")
	fmt.Println("\nformattedTime",formattedTime)
	return formattedTime, previousMin

}

//validates user using username, password and token
//functin returns Http status code 200  OK if user is valid
//functin returns Http status code 401 UNAUTHORIZED if user is invalid
func loginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	token := c.PostForm("token")

	user := "c137@onecause.com"
	pass := "#th@nH@rm#y#r!$100%D0p#"
	formattedTime, previousMin := getToken();
	
	if (username == user && password == pass && (token == formattedTime || token==previousMin)){
		c.JSON(http.StatusOK , gin.H{"status":  "valid"})
	}else{
		c.JSON(401, gin.H{"message": "Invalid Credentials"})
	}
  }

  func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/login", loginUser)
	listenPort := "8081"
	router.Run(":"+listenPort)
  }