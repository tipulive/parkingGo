package helper

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Email struct {
	To      []string
	Subject string
	Body    string
	HTML    bool
}

// SendEmail automatically handles Implicit TLS (465) and STARTTLS (587)
func SendEmail(cfg SMTPConfig, email Email) error {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	contentType := "text/plain; charset=UTF-8"
	if email.HTML {
		contentType = "text/html; charset=UTF-8"
	}

	msg := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: %s\r\n\r\n"+
			"%s\r\n",
		cfg.From,
		strings.Join(email.To, ", "),
		email.Subject,
		contentType,
		email.Body,
	))

	var c *smtp.Client
	var err error

	tlsConfig := &tls.Config{
		ServerName: cfg.Host,
		MinVersion: tls.VersionTLS12,
	}

	// --- Try Implicit TLS first (port 465 or others that support it)
	c, err = tryImplicitTLS(addr, cfg.Host, tlsConfig)
	if err != nil {
		// --- Fallback: Try plain connection + STARTTLS upgrade
		c, err = tryStartTLS(addr, cfg.Host, tlsConfig)
		if err != nil {
			return fmt.Errorf("SMTP connection failed (TLS & STARTTLS): %w", err)
		}
	}

	defer c.Quit()

	// Auth
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	if err = c.Auth(auth); err != nil {
		return fmt.Errorf("SMTP authentication failed: %w", err)
	}

	// Mail headers
	if err = c.Mail(cfg.From); err != nil {
		return err
	}
	for _, to := range email.To {
		if err = c.Rcpt(to); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	if _, err = w.Write(msg); err != nil {
		return err
	}
	return w.Close()
}

// tryImplicitTLS attempts a direct TLS connection (for port 465)
func tryImplicitTLS(addr, host string, tlsConfig *tls.Config) (*smtp.Client, error) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}, "tcp", addr, tlsConfig)
	if err != nil {
		if strings.Contains(err.Error(), "handshake") || strings.Contains(err.Error(), "first record") {
			return nil, errors.New("not implicit TLS")
		}
		return nil, err
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// tryStartTLS attempts a plain SMTP connection and upgrades to TLS
func tryStartTLS(addr, host string, tlsConfig *tls.Config) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}
	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(tlsConfig); err != nil {
			client.Close()
			return nil, fmt.Errorf("STARTTLS failed: %w", err)
		}
	} else {
		client.Close()
		return nil, errors.New("server does not support STARTTLS")
	}
	return client, nil
}
