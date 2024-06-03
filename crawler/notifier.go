package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const TOKEN string = "1814067387:AAH5IZRir2kGNbTD6ZhmAOAkpffUppujvT4"

func SendTelegram(text, photoUrl, caption string) {
	fmt.Println(photoUrl)
	data, _ := json.Marshal(map[string]string{
		"chat_id": "-1001367184617",
		"text":    text,
		"photo":   photoUrl,
		"caption": caption,
	})
	resp, err := http.Post("https://api.telegram.org/bot"+TOKEN+"/sendPhoto", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Failed to send request...")
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	}

}
