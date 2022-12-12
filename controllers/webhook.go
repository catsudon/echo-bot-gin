package controllers

import (
	// "context"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"echo-bot-gin/configs"
	"echo-bot-gin/models"
	"echo-bot-gin/responses"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
var lineAccessToken string = configs.EnvLineAccessToken()

func HandleWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var msg models.LineMessage

		//validate the request body
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		replyToken := msg.Events[0].ReplyToken
		text := models.Text{
			Type : "text",
			Text : "ข้อความเข้ามา : " + msg.Events[0].Message.Text ,
		}

		replyMsg := models.ReplyMessage{
			ReplyToken: replyToken,
			Messages: []models.Text{
				text,
			},
		}
		go replyMessageLine(replyMsg)

		//use the validator library to validate required fields
		// (I don't think we have it in this case)
		if validationErr := validate.Struct(&msg); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": replyToken}})
	}
}

func replyMessageLine(Message models.ReplyMessage) error {
	value, _ := json.Marshal(Message)

	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+lineAccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return err
}
