package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/nobbs/uptime-kuma-api/pkg/action"
	"github.com/nobbs/uptime-kuma-api/pkg/client"
)

type Config struct {
	Host   string
	Port   int
	Secure bool

	Username string
	Password string
	Token    string

	JWT string
}

func NewConfig() (*Config, error) {
	host := os.Getenv("UPTIME_KUMA_HOST")
	if host == "" {
		return nil, fmt.Errorf("UPTIME_KUMA_HOST env var must be set")
	}

	portStr := os.Getenv("UPTIME_KUMA_PORT")
	if portStr == "" {
		return nil, fmt.Errorf("UPTIME_KUMA_PORT env var must be set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	secureStr := os.Getenv("UPTIME_KUMA_SECURE")
	if secureStr == "" {
		return nil, fmt.Errorf("UPTIME_KUMA_SECURE env var must be set")
	}

	secure, err := strconv.ParseBool(secureStr)
	if err != nil {
		return nil, err
	}

	username := os.Getenv("UPTIME_KUMA_USERNAME")
	password := os.Getenv("UPTIME_KUMA_PASSWORD")
	token := os.Getenv("UPTIME_KUMA_TOKEN")

	jwt := os.Getenv("UPTIME_KUMA_JWT")

	if username == "" && password == "" && token == "" && jwt == "" {
		return nil, fmt.Errorf("UPTIME_KUMA_USERNAME, UPTIME_KUMA_PASSWORD, UPTIME_KUMA_TOKEN or UPTIME_KUMA_JWT env vars must be set")
	}

	return &Config{
		Host:     host,
		Port:     port,
		Secure:   secure,
		Username: username,
		Password: password,
		Token:    token,
		JWT:      jwt,
	}, nil
}

func main() {
	// create new slog logger instance
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// create new config instance
	config, err := NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// create new client instance
	c, err := client.NewClient(config.Host, config.Port, config.Secure)
	if err != nil {
		log.Fatalln(err)
	}

	switch {
	case config.Username != "" && config.Password != "":
		// login with username and password
		if _, err = action.Login(c, config.Username, config.Password, config.Token); err != nil {
			log.Fatalln(err)
		}
	case config.JWT != "":
		if err = action.LoginByToken(c, config.JWT); err != nil {
			log.Fatalln(err)
		}
	default:
		log.Fatalln("no login credentials provided")
	}

	// get info
	info, _ := c.State().Info()
	slog.Info("info", slog.Any("info", info))

	// wait indefinitely
	select {}
}
