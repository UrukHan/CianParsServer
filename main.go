package main

import (
	"CianParser/code"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type rec struct {
	Link string `json:"link"`
	Max int `json:"max"`
}

func parseResult(context *gin.Context) {
	var newReq rec
	if err := context.BindJSON(&newReq); err != nil {
		return
	}
	//fmt.Printf("%v", newReq.Links)
	//fmt.Println("\n", newReq.Link, "\n")
	resp := code.Response(newReq.Link, newReq.Max)
	context.IndentedJSON(http.StatusOK, resp)
	data, _ := json.MarshalIndent(resp, "", " ")
	_ = ioutil.WriteFile("test.json", data, 0644)
}

func helthCheck(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, map[string]interface{}{"code": 200, "status": "OK"})
}

func main() {
	fmt.Println("Listening on PORT 8000")
	router := gin.Default()
	router.GET("/health", helthCheck)
	router.POST("/req", parseResult)
	router.Run("0.0.0.0:8000")

	//code.Test()
}
