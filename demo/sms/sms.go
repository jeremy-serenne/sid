package sms

type ServiceSMS interface {
	SendSMS() error
}
