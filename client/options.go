package client

import "github.com/tmconsulting/amadeus-golang-sdk/logger"

// Option describes a functional option for configuring AmadeusClient.
type Option func(client *AmadeusClient)

//SetURL sets client connection url
func SetURL(url string) Option {
	return func(c *AmadeusClient) {
		c.service.Client.Url = url
	}
}

// SetUser sets client user (originator)
func SetUser(originator string) Option {
	return func(c *AmadeusClient) {
		c.service.Client.User = originator
	}
}

//SetPassword sets client password
func SetPassword(pass string) Option {
	return func(c *AmadeusClient) {
		c.service.Client.Pass = pass
	}
}

//SetAgent sets client office id
func SetAgent(officeID string) Option {
	return func(c *AmadeusClient) {
		c.service.Client.Agent = officeID
	}
}

func SetLogger(l logger.LogWriter) Option {
	return func(c *AmadeusClient) {
		c.service.Client.Logger = logger.NewLogger(l)
	}
}
