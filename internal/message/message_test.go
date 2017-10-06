package message_test

import (
	"testing"

	"github.com/gabelula/messagebird_api/internal/message"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestValidate(t *testing.T) {
	originator_long := "This is Gabriela is too long"
	originator_valid := "Gabriela"

	body_empty := ""
	body_valid := "This is a valid body."

	recipient_short := 123
	recipient_valid := 123456778

	t.Log("Given the need to validate a message")
	{
		t.Log("Given no recipient")
		{
			no_recipient_message := message.Message{Body: body_valid, Originator: originator_valid}
			err := no_recipient_message.Validate()
			if err == nil {
				t.Errorf("Expected an error, %v", err)
			}
		}

		t.Log("Given a originator with more than 11 characters")
		{
			long_originator_message := message.Message{Body: body_valid, Originator: originator_long, Recipient: recipient_valid}
			err := long_originator_message.Validate()
			if err == nil {
				t.Errorf("Expected an error, %v", err)
			}
		}

		t.Log("Given a recipient with more less than 9 numbers")
		{
			short_recipient_message := message.Message{Body: body_valid, Originator: originator_valid, Recipient: recipient_short}
			err := short_recipient_message.Validate()
			if err == nil {
				t.Errorf("Expected an error, %v", err)
			}
		}

		t.Log("Given no originator")
		{
			no_originator_message := message.Message{Body: body_valid, Recipient: recipient_valid}
			err := no_originator_message.Validate()
			if err == nil {
				t.Errorf("Expected an error, %v", err)
			}
		}

		t.Log("Given an empty message")
		{
			empty_message := message.Message{Body: body_empty, Originator: originator_valid, Recipient: recipient_valid}
			err := empty_message.Validate()
			if err == nil {
				t.Errorf("Expected an error, %v", err)
			}
		}

		t.Log("Given a valid message")
		{
			valid_message := message.Message{Body: body_valid, Originator: originator_valid, Recipient: recipient_valid}
			err := valid_message.Validate()
			if err != nil {
				t.Errorf("Expected an empty error, %v", err)
			}
		}
	}
}

func TestProcess(t *testing.T) {
	body_more_than_160 := "This are more than one message.This are more than one message.This are more than one message.This are more than one message.This are more than one message.This are more than one message.This are more than one message."
	body_less_than_160 := "This is just a message."
	body_with_160 := "This message has 160 characters.This message has 160 characters.This message has 160 characters.This message has 160 characters.This message has 160 characters."

	t.Log("Given the need to process a message")
	{
		t.Log("Given a message that has more than 160 characters")
		{
			more_than_160_message := message.Message{Body: body_more_than_160}
			ms, err := more_than_160_message.Process()
			if err != nil {
				t.Errorf("Expected an empty error. Got %v", err)
			}
			if len(ms) < 2 {
				t.Error("Expected the message to be divided")
			}
		}

		t.Log("Given a message that has less than 160 characters")
		{
			less_than_160_message := message.Message{Body: body_less_than_160}
			ms, err := less_than_160_message.Process()
			if err != nil {
				t.Errorf("Expected an empty error. Got %v", err)
			}
			if len(ms) > 1 {
				t.Errorf("Expected the message to be only 1, %v ", ms)
			}
		}

		t.Log("Given a message that has 160 characters")
		{
			len_160_message := message.Message{Body: body_with_160}
			ms, err := len_160_message.Process()
			if err != nil {
				t.Errorf("Expected an empty error. Got %v", err)
			}
			if len(ms) > 1 {
				t.Errorf("Expected the message to be only 1, %v ", ms)
			}
		}
	}
}

func TestSend(t *testing.T) {
	t.Skip("TODO.")
}
