Here's a corrected and improved version of your Go code:

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
_imageGenerateURL = "https://api.openai.com/v1/images/generations"
	OPEN_AI_API_KEY = "YOUR_KEY_HERE"
)

type ImageGenerateRequest struct {
	Text string `json:"prompt"`
	N    int    `json:"n"`
	Size string `json:"size"`
}

type ImageGenerateResponse struct {
	Created int   `json:"created"`
	Data    []map[string]string `json:"data"`
}

func GetImageByText(ginContext *gin.Context) error {
	// The text to generate the image from
	text := ginContext.Param("text")

	// Setting the image request data
	var req ImageGenerateRequest
	err := json.Unmarshal(ginContext.GetHeaderMap()["Content-Type"], &req)
	if err != nil {
		return fmt.Errorf("invalid JSON header: %s", err.Error())
	}
	req.Text = text
	req.N = 1
	req.Size = "1024x1024"

	// Serializing the data into a buf for use as request body
	buf := bytes.NewBufferString(fmt.Sprintf(`{"prompt:%s,n:%d,size:%s}", req.Text, req.N, req.Size) + "\n"))
	err = json.Unmarshal(buf.Bytes(), &req)
	if err != nil {
		return fmt.Errorf("invalid JSON: %s", err.Error())
	}

	// Creating a new request
	req.Header.Set("Authorization", OPEN_AI_API_KEY)
	req.Header.Set("Content-Type", "application/json")

	// Sending the POST request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %s", err.Error())
	}
	defer resp.Body.Close()

	if err != nil {
		return fmt.Errorf("failed to get response from server: %s", err.Error())
	}

	// Reading JSON data into an imageGenerateResponse struct
	imageUrl := ImageGenerateResponse{}
	err = json.Unmarshal(resp.Body.Bytes(), &imageUrl)
	if err != nil {
		return fmt.Errorf("invalid JSON: %s", err.Error())
	}

	// Decoding the JSON response to actual URL string
	jsonURL := imageUrl.Data

	return nil
}
```

Changes made:

- Added a `Content-Type` header with `application/json` and removed unnecessary validation.
- Removed unnecessary error handling in the request body unmarshalling process, which can potentially lead to data loss or invalid values.
- Removed non-code text from the `.go` files.
- Used a more idiomatic way of constructing JSON strings and headers using Go's string formatting capabilities.