package main

import (
	"github.com/rs/zerolog"
	"os"
	"sid/lazy"
	"sid/sms"
	"sid/sms/twilio"
)

type Service struct {
	smsProvider *lazy.Lazy[sms.ServiceSMS]
}

func (s *Service) handleRandomRequest() error {
	err := s.smsProvider.Value().SendSMS()

	return err
}

func main() {
	logger := zerolog.New(os.Stdout)

	lazyService := lazy.Of(func() sms.ServiceSMS {
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
