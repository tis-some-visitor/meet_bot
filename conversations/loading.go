package conversations

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//converts image code string to a FileBytes struct
func ConvertStringToPhoto(imageCode string, chatID int64) tgbotapi.FileBytes {

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageCode))

	img2, err := png.Decode(reader)
	if err != nil {
		log.Printf("Can't decode image code for chat %v: %v", chatID, err)
	}

	var buf bytes.Buffer
	options := jpeg.Options{}
	options.Quality = 100

	err = jpeg.Encode(&buf, img2, &options)
	if err != nil {
		log.Printf("Can't buffer image for chat %v: %v", chatID, err)
	}
	data := buf.Bytes()

	file := tgbotapi.FileBytes{
		Name:  "image.jpg",
		Bytes: data,
	}

	return file
}

//downloads an image of max size send by user, converts it into string of code ready for inserting into database.
func Download(u tgbotapi.Update) string {

	var mystrings string

	if u.Message.Photo != nil {

		a := u.Message.Photo
		//выбираем в фотосайз наибольший возможный размер фотки
		big := a[len(a)-1]
		//medium := a[len(a)-3]

		mystrings = big.FileID

	} else if u.Message.Document != nil {

		if u.Message.Document.MimeType == "image/jpg" || u.Message.Document.MimeType == "image/jpeg" || u.Message.Document.MimeType == "image/png" {

			mystrings = u.Message.Document.FileID
		}

	}

	conf := tgbotapi.FileConfig{FileID: mystrings}
	file, err := Bot.GetFile(conf)
	link := file.Link(os.Getenv("token"))

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(link)
	if err != nil {
		log.Println("Error in Get, func Download: ", err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in ReadAll, func Download: ", err)
	}

	contentType := http.DetectContentType(data)

	switch contentType {
	case "image/png":
		fmt.Println("Image type is already PNG.")
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			log.Println("unable to decode jpeg: %w", err)
		}

		var buf bytes.Buffer
		if err := png.Encode(&buf, img); err != nil {
			log.Println("unable to encode png: %w", err)
		}

		data = buf.Bytes()
	default:
		log.Printf("unsupported content type: %s", contentType)
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(data)
	return imgBase64Str
}
