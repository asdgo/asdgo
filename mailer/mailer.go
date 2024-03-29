package mailer

import (
	"os"
	"strconv"
	"time"

	"github.com/wneessen/go-mail"
)

type Mailer struct {
	client mail.Client
	from   string
}

var Instance *Mailer

func New() {
	host := os.Getenv("MAIL_HOST")
	from := os.Getenv("MAIL_FROM_ADDRESS")
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	portEnv := os.Getenv("MAIL_HOST")

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}

	client, err := mail.NewClient(
		host,
		mail.WithTimeout(mail.DefaultTimeout),
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithPort(port),
		mail.WithUsername(username),
		mail.WithPassword(password),
	)

	if err != nil {
		panic(err)
	}

	Instance = &Mailer{
		client: *client,
		from:   from,
	}
}

func (m *Mailer) Send(recipient string, template string) error {
	msg := mail.NewMsg()

	err := msg.To(recipient)
	if err != nil {
		return err
	}

	err = msg.From(m.from)
	if err != nil {
		return err
	}

	msg.SetBodyString(mail.TypeTextPlain, template)

	for i := 1; i <= 3; i++ {
		err = m.client.DialAndSend(msg)

		if nil == err {
			return nil
		}

		if i != 3 {
			time.Sleep(2 * time.Second)
		}
	}

	return err
}
