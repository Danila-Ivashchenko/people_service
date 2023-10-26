package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	postgresUser    string
	postgresPass    string
	postgresHost    string
	postgresPort    string
	postgresDB      string
	postgresSSLMode string
	env             string
	httpPort        string
	httpHost        string
	ageUrl          string
	genderUrl       string
	nationalityUrl  string
}

func (c *config) GetPsqlURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.postgresUser, c.postgresPass, c.postgresHost, c.postgresPort, c.postgresDB, c.postgresSSLMode)
}

func (c *config) GetHTTPPort() string {
	return c.httpPort
}

func (c *config) GetEnv() string {
	return c.env
}

func (c *config) GetAgeApiURL() string {
	return c.ageUrl
}
func (c *config) GetGenderApiURL() string {
	return c.genderUrl
}
func (c *config) GetNationalityApiURL() string {
	return c.nationalityUrl
}

func LoadEnv(filenames ...string) error {
	const op = "pkg.config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func GetConfig() *config {
	cfg := &config{
		postgresUser:    "",
		postgresPass:    "",
		postgresHost:    "localhost",
		postgresPort:    "27017",
		postgresDB:      "",
		env:             "local",
		postgresSSLMode: "disable",
		httpHost:        "localhost",
		ageUrl:          "https://api.agify.io/",
		genderUrl:       "https://api.genderize.io/",
		nationalityUrl:  "https://api.nationalize.io/",
	}

	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	ssl := os.Getenv("POSTGRES_SSL_MODE")
	env := os.Getenv("ENV")
	httpPort := os.Getenv("HTTP_PORT")
	httpHost := os.Getenv("HTTP_HOST")
	ageUrl := os.Getenv("AGE_API_URL")
	genderUrl := os.Getenv("GENDER_API_URL")
	nationalityUrl := os.Getenv("NATIONALITY_API_URL")

	if env != "" {
		cfg.env = env
	}
	if httpPort != "" {
		cfg.httpPort = httpPort
	}
	if httpHost != "" {
		cfg.httpHost = httpHost
	}
	if user != "" {
		cfg.postgresUser = user
	}
	if pass != "" {
		cfg.postgresPass = pass
	}
	if host != "" {
		cfg.postgresHost = host
	}
	if port != "" {
		cfg.postgresPort = port
	}
	if db != "" {
		cfg.postgresDB = db
	}
	if ssl != "" {
		cfg.postgresSSLMode = ssl
	}
	if ageUrl != "" {
		cfg.ageUrl = ageUrl
	}
	if genderUrl != "" {
		cfg.genderUrl = genderUrl
	}
	if nationalityUrl != "" {
		cfg.nationalityUrl = nationalityUrl
	}

	return cfg
}
