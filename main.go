Here's an improved version of the provided Go code:

```go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/image/:text", GetImageByText)
	log.Println("Server is running on port 8080")
	r.Run(":8080")
}

func GetImageByText(c *gin.Context) {
	c.JSON(200, map[string]string{"image": "some_image_text"})
}
```

In this corrected version:

- The `log.Println` statement has been replaced with a single line `log.Println("Server is running on port 8080")`, as it's better practice to log messages at the beginning of your main function or in a separate logger.

This code serves two purposes: 
1. It sets up a Gin router and defines an endpoint `/image/:text` that will handle GET requests.
2. It logs a message indicating whether the server is running on port 8080, but does not actually start the server, which would be done by calling `r.Run(":8080")`. The `log.Println` statement prints the message to the console. 

Note: To run this code, you need to have Go and Gin installed in your environment. You can install Gin with the following command:
```bash
go get -u github.com/gin-gonic/gin
```

Also note that the server will exit when it receives a request. If you want to keep it running indefinitely (for example, to catch any errors), you'll need to call `r.Run(":8080")` multiple times in your main function or add additional code to handle any exceptions.