package sms

import (
	"encoding/json"

	"github.com/trstoyan/Alertify/kafka"
)

func ProduceSMSMessage(producer *kafka.Writer, payload SMSPayload) error {
	msg, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return kafka.ProduceMessage(producer, nil, msg)
}
