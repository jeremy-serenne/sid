package twilio

import "github.com/rs/zerolog"

type Twilio struct {
	logger *zerolog.Logger
}

func New(log *zerolog.Logger) *Twilio {
	return &Twilio{
		logger: log,
	}
}

func (t *Twilio) SendSMS() error {
	t.logger.Info().Msg("Send SMS with Twilio")
	return nil
}
