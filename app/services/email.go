package services

import (
	"bytes"
	"github.com/batudal/derisk/app/config"
	"github.com/batudal/derisk/app/schema"
	"github.com/resendlabs/resend-go"
	"text/template"
)

func RequestEmail(cfg *config.Config, customer *schema.Customer, email_type string) error {
	request := schema.EmailRequest{
		Customer: schema.Customer{
			Email:        customer.Email,
			CustomerType: customer.CustomerType,
		},
		EmailType: email_type,
	}
	cfg.EmailRequests <- request
	return nil
}

func ListenEmailRequests(cfg *config.Config) {
	for {
		select {
		case request := <-cfg.EmailRequests:
			switch request.EmailType {
			case "join-beta-list":
				err := sendEmail(
					cfg,
					"./views/emails/betalist.html",
					request.Customer,
					"De-risk Beta List",
					[]resend.Tag{
						{
							Name:  "beta-list",
							Value: "subscribe",
						},
					},
				)
				if err != nil {
					// todo: proper logging
					println(err)
				}
			case "feedback":
				err := sendEmail(
					cfg,
					"./views/emails/feedback.html",
					request.Customer,
					"Thank you for your feedback!",
					[]resend.Tag{
						{
							Name:  "feedback",
							Value: "landing-page",
						},
					},
				)
				if err != nil {
					// todo: proper logging
					println(err)
				}
			default:
				println("Wrong email type")
			}
		}
	}
}

func sendEmail(cfg *config.Config, temp string, customer schema.Customer, subject string, tags []resend.Tag) error {
	html, err := templateToHTML(temp, customer)
	if err != nil {
		return err
	}
	err = send(cfg, customer, html, subject, tags)
	if err != nil {
		return err
	}
	return nil
}

func templateToHTML(temp string, customer schema.Customer) (string, error) {
	t, err := template.ParseFiles(temp)
	if err != nil {
		return "", err
	}
	body := new(bytes.Buffer)
	err = t.Execute(body, customer)
	if err != nil {
		return "", err
	}
	return body.String(), nil
}

func send(cfg *config.Config, customer schema.Customer, html string, subject string, tags []resend.Tag) error {
	_, err := cfg.Rs.Emails.Send(&resend.SendEmailRequest{
		From:    "Founder <batu@de-risk.app>",
		To:      []string{customer.Email},
		Subject: subject,
		Html:    html,
		Tags:    tags,
		Headers: map[string]string{
			"List-Unsubscribe": "https://de-risk.app/unsubscribe-beta-list",
		}})
	if err != nil {
		return err
	}
	return nil
}
