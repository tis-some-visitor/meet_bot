package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"meet_bot/conversations"
	"meet_bot/dbfellow"

	"meet_bot/envar"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	var err error
	log.Print("jdbfg")

	err = envar.LoadEnvironmentals()
	if err != nil {
		log.Panic(err)
	}

	conversations.Bot, err = tgbotapi.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Panic(err)
	}

	conversations.Bot.Debug = true

	log.Printf("Authorized on account %s", conversations.Bot.Self.UserName)

	var hook string
	var port string

	if os.Args == nil || len(os.Args) == 1 {

		hook = os.Getenv("hook")
		port = os.Getenv("port")
		dbfellow.InitDB("db")

	} else if os.Args[1] == "local" {

		hook = os.Getenv("hooklocal")
		port = "5000"
		dbfellow.InitDB("local")
	}

	wh, err := tgbotapi.NewWebhook(hook + os.Getenv("token"))
	if err != nil {
		log.Fatalf("Can't set the webhook: %s", err)
	}

	_, err = conversations.Bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := conversations.Bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	if port == "" {
		log.Fatal("port not set")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(CORSMiddleware(), gin.Logger())
	router.POST("/"+conversations.Bot.Token, webhookHandler)
	router.POST("/app", webhookHandlerWeb)

	err = router.Run(":" + port)
	if err != nil {
		log.Printf("Failed to attach a router: %s", err)
	}

}

func webhookHandlerWeb(c *gin.Context) {

	if c.Request.Body != nil {
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		g := make(map[string]string)
		json.Unmarshal(bodyBytes, &g)
		log.Print(g)
	}

}

func webhookHandler(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Paniced: %s", err)
			debug.PrintStack()
		}
	}()

	defer c.Request.Body.Close()

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		log.Println(err)
		return
	}

	conversations.NewCycleUpdates(update)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-credentials, access-control-allow-methods, access-control-allow-headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
