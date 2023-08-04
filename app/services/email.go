package services

import (
	"bytes"
	"text/template"

	"github.com/batudal/derisk/app/config"
	"github.com/resendlabs/resend-go"
)

func JoinBetaListEmail(cfg *config.Config, email string, customer_type string) error {
	t, err := template.ParseFiles("./views/emails/betalist.html")
	if err != nil {
		return err
	}
	body := new(bytes.Buffer)
	err = t.Execute(body, struct {
		Email        string
		CustomerType string
	}{
		Email:        email,
		CustomerType: customer_type,
	})
	if err != nil {
		return err
	}
	_, err = cfg.Rs.Emails.Send(&resend.SendEmailRequest{
		From:    "De-risk <hi@de-risk.app>",
		To:      []string{email},
		Subject: "De-risk Beta List",
		Html:    body.String(),
	})
	if err != nil {
		return err
	}
	// delete later
	return nil
}
