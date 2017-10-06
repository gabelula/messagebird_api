package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gabelula/messagebird_api/internal/message"
)

// Index to show the home page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing Exercise to write an API")
}

// MessageSend formats the message, send it to MessageBird and retun a response
func MessageSend(w http.ResponseWriter, r *http.Request) {
	log.Print("Started : MessageSend")

	var themessage message.Message

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: CHECK TO SEE HOW THIS IS NOT WORKING
	err = json.Unmarshal(body, &themessage)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}
	}

	err = themessage.Send()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatal(err)
		}

		log.Printf("Error: %v", err)
	}

	if err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Message Sent")
	}

	log.Print("Completed : MessageSend")
}
