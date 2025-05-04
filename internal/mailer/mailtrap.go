package mailer

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
	"time"

	"gopkg.in/gomail.v2"
)

type mailtrapClient struct {
	smtpUser string
	apiKey   string
}

func NewMailTrapClient(apiKey, smtpUser string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}

	return mailtrapClient{
		smtpUser: smtpUser,
		apiKey:   apiKey,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {

	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", "from@example.com")
	message.SetHeader("To", "to@example.com")
	message.SetHeader("Subject", subject.String())
	message.AddAlternative("text/html", body.String())

	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, m.smtpUser, m.apiKey)

	// Exponential backoff retry logic
	var retryErr error
	for i := range 3 {
		retryErr = dialer.DialAndSend(message)
		if retryErr != nil {
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}

		return 200, nil
	}

	return -1, fmt.Errorf("failed to send email after %d attempts, error %v", maxRetries, retryErr)
}
