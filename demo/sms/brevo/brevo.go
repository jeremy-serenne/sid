package brevo

import "github.com/rs/zerolog"

type Brevo struct {
	logger *zerolog.Logger
}

func New(log *zerolog.Logger) *Brevo {
	return &Brevo{
		logger: log,
	}
}

func (t *Brevo) SendSMS() error {
	t.logger.Info().Msg("Send SMS with Brevo")
	return nil
}
