package service

import (
	"context"
	"fmt"
	"horsy_api/config"
	"net/smtp"
	"time"

	"go.uber.org/zap"
)

const (
	sendMailTimeout = 10 * time.Second
)

type EmailService struct {
	cfg config.EmailConfig
}

func (s *EmailService) SendCreatedAccountMail(ctx context.Context, email string, password string) error {
	subject := "Your account has been created"
	body := fmt.Sprintf(
		`Hello,

Your account has been successfully created.

Login: %s
Password: %s

Please change your password after first login.

Regards,
Horsy Team
`, email, password)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/plain; charset=\"utf-8\"\r\n"+
			"\r\n%s",
		s.cfg.From,
		email,
		subject,
		body,
	))

	zap.L().Info(string(msg))
	return nil

	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)

	auth := smtp.PlainAuth(
		"",
		s.cfg.Username,
		s.cfg.Password,
		s.cfg.Host,
	)

	done := make(chan error, 1)

	go func() {
		done <- smtp.SendMail(
			addr,
			auth,
			s.cfg.From,
			[]string{email},
			msg,
		)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		if err != nil {
			zap.L().Error(err.Error())
		}
		return err
	case <-time.After(sendMailTimeout):
		return ErrSendMailTimeout
	}
}

func NewEmailService(cfg config.EmailConfig) *EmailService {
	return &EmailService{cfg: cfg}
}
