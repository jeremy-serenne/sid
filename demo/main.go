package main

import (
	"os"
	"sid"
	"sid/demo/sms"
	"sid/demo/sms/twilio"

	"github.com/rs/zerolog"
)

type Service struct {
	smsProvider *sid.Lazy[sms.ServiceSMS]
}

func (s *Service) handleRandomRequest() error {
	err := s.smsProvider.Value().SendSMS()

	return err
}

func main() {
	logger := zerolog.New(os.Stdout)

	lazyService := sid.Of(func() sms.ServiceSMS {
		return twilio.New(&logger)
	})

	s := Service{
		smsProvider: lazyService,
	}

	err := s.handleRandomRequest()
	if err != nil {
		panic(err)
	}
}
