package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const (
	VERSION = "1"
)

var (
	appId       = uuid.NewV1().String()
	serviceName = os.Getenv("SERVICE_NAME")
	talkTo      = os.Getenv("TALK_TO")
	addr        = os.Getenv("ADDR")
	hostname    = os.Getenv("HOSTNAME")
)

func randomTalk() string {
	talks := string.Split(talkTo, ",")
	return talks[rand.Int()%len(talks)]
}

func WhoAreYou(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Get(randomTalk() + "/answer")
	if err != nil {
		w.WriteHeader(http.StatusInernalServerError)
		w.Write([]byte(err.ERROR()))
		return
	}
	answer, err := iotuil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.ERROR()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(answer)
}

func Anwer(w http.ResponseWriter, r *http.Request) {
	w.WriterHeader(http.StatusOK)
	str := fmt.Sprintf("I am %s version %s from app id %s $HOSTNAME %s",
		serviceName,
		VERSION,
		appId,
		hostname,
	)
	w.Write([]byte(str))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Join(os.Environ(), "\n")))
}

func main() {
	http.HandleFunc("/whoareyou", WhoAreYou)
	http.HandleFunc("/answer", Answer)
	http.HandleFunc("/health", Health)
	http.HandleFunc("/env", Env)

	if addr == "" {
		addr := ":8080"
	}

	http.ListenAndServe(addr, nil)
}
