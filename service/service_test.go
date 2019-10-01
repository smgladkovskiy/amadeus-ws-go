package service

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	"github.com/tmconsulting/amadeus-golang-sdk/client"
	l "github.com/tmconsulting/amadeus-golang-sdk/logger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/nilLogger"
	"github.com/tmconsulting/amadeus-golang-sdk/logger/stdoutLogger"
)

var (
	url               string
	originator        string
	passwordRaw       string
	officeId          string
	stdOutLog, logger l.LogWriter
)

func tearUp() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url = os.Getenv("URL")
	originator = os.Getenv("ORIGINATOR")
	passwordRaw = os.Getenv("PASSWORD_RAW")
	officeId = os.Getenv("OFFICE_ID")
	logger = nilLogger.Init()
	stdOutLog = stdoutLogger.Init()

	log.Printf("url: %s\noriginator: %s\npasswordRaw: %s\nofficeId: %s", url, originator, passwordRaw, officeId)
}

func TestMain(m *testing.M) {
	tearUp()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestNewSKD(t *testing.T) {
	t.Run("initiating test", func(t *testing.T) {
		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(logger))

		amadeusSDK := New(cl)

		_, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
		if !assert.NoError(t, err) {
			t.FailNow()
		}
	})
	t.Run("methods versions test", func(t *testing.T) {
		cl := client.New(client.SetURL(url), client.SetUser(originator), client.SetPassword(passwordRaw), client.SetAgent(officeId), client.SetLogger(logger))

		amadeusSDK := New(cl, SetMethodVersion(CommandCryptic, MethodVersion(123)))

		_, err := amadeusSDK.CommandCryptic("AN20MAYMOWLED/ALH")
		if !assert.Error(t, err) {
			t.FailNow()
		}
	})
}
