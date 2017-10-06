package message

// ^ TODO: find a better name for this package

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/messagebird/go-rest-api"
)

const (
	limitCharacters = 160 // 153 + User Data Header
	originatorName  = "Gabriela"
)

// Message is the structure for the message we are getting
type Message struct {
	Header     [7]byte // It is being calculated if it is a concatenated message.
	Body       string  `json:"body"`
	Originator string  `json:"originator"` // Be careful that limit of the name is 11.
	Recipient  int     `json:"recipient"`  // This is a phone number
}

// Validate make sure no empty or incorrect parameter values are send to MessageBird.
func (message *Message) Validate() error {

	var err error

	// Validate Recipient does not have less than 191 characters
	if len(strconv.Itoa(message.Recipient)) < 9 {
		err = fmt.Errorf("When validating: Recipient has less than 9 numbers")
	}

	// Originator and Recipients are required
	if message.Originator == "" || message.Recipient == 0 {
		err = fmt.Errorf("When validating: There is no originator or recipient")
	}

	// Validate Originator does not have more than 11 characters
	if len(message.Originator) > 11 {
		err = fmt.Errorf("When validating: Originator has more than 11 characters")
	}

	// Validate Body not Null
	if message.Body == "" {
		err = fmt.Errorf("When validating: There is an empty message")
	}

	return err
}

// Process concatenated SMS: When an incoming message content/body is longer than 160 characters,
// split it into multiple parts.
func (message *Message) Process() ([]Message, error) {
	var ms []Message
	var m string

	// TODO: check on how len works with UTF-8 string
	if len(message.Body) > limitCharacters {

		limitWithoutUDH := limitCharacters - 7
		referenceNumber := rand.Int()
		// Split into limitCharacters - 7 characters.
		for i := 0; i < len(message.Body); i = i + limitWithoutUDH {
			if i+limitCharacters < len(message.Body) {
				m = string(message.Body[i : i+limitWithoutUDH])
			} else {
				m = string(message.Body[i:])
			}

			// UDH header
			var udh [7]byte
			// doesn't include itself, its header length
			udh[0] = 0x05
			// SAR identifier
			udh[1] = 0x00
			// SAR length
			udh[2] = 0x03
			// create reference number (same for all messages)
			udh[3] = byte(referenceNumber)
			// total number of segments
			udh[4] = byte(len(message.Body) / limitWithoutUDH)
			// segment number
			udh[5] = byte(i + limitWithoutUDH)

			ms = append(ms, Message{Body: m, Recipient: message.Recipient, Originator: message.Originator, Header: udh})
		}
	} else {
		ms = append(ms, *message)
	}

	return ms, nil
}

// Send is sending data to messagebird API
func (message *Message) Send() error {
	log.Printf("Send: Started")

	// Accepts SMS messages

	// Validations:
	// Input Validation: Make sure no empty or incorrect parameter values are send to MessageBird.
	err := message.Validate()
	if err != nil {
		return err
	}

	// Process:
	// If the message has more than the limit of characters, split it
	ms, err := message.Process()
	if err != nil {
		return err
	}

	// Send the received message to the MessageBird REST API
	// Outgoing messages won't exceed one API request per second (neither when concatenated mesages
	// OR when receiving requests)

	// It is using the messagebird go rest api library
	// Documentation: https://developers.messagebird.com/docs/messaging

	mbClient := messagebird.New(os.Getenv("MESSAGEBIRD_API_KEY"))

	// message header for concatenated messages

	// Send only if there is balance
	balance, err := mbClient.Balance()
	if err != nil {
		return err
	}

	if balance.Amount > 0 {
		log.Printf("The balance in the account is: %v.", balance.Amount)
		for _, m := range ms {
			// Rate of sending messages
			time.Sleep(1000 * time.Millisecond)

			fmt.Printf("DEBUG %v", m.Header)
			// typedetails := make(map[string]interface{})

			// if len(m.Header) > 1 {
			// 	typedetails["uhf"] = m.Header
			// }

			// Adds the UDH
			params := &messagebird.MessageParams{
				Reference: "GO Developer Test"}
			//				TypeDetails: typedetails}

			// Sending the message
			message, err := mbClient.NewMessage(m.Originator, []string{strconv.Itoa(m.Recipient)}, m.Body, params)
			if err != nil {
				log.Printf("Error %v \n", err)
				log.Printf("Returned object %v \n", message)

				return err
			}
		}
	}

	log.Printf("Send: Completed.")
	return nil
}
