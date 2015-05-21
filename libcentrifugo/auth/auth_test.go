package auth

import (
	"testing"
)

func TestGenerateClientToken(t *testing.T) {
	var (
		secretKey  = "secret"
		projectKey = "project"
		user       = "user"
		timestamp  = "1430669930"
	)
	tokenWithInfo := GenerateClientToken(secretKey, projectKey, user, timestamp, "{}")
	if len(tokenWithInfo) != 64 {
		t.Error("sha256 token length must be 64")
	}
}

func TestCheckClientToken(t *testing.T) {
	var (
		secretKey     = "secret"
		projectKey    = "project"
		user          = "user"
		timestamp     = "1430669930"
		info          = "{}"
		providedToken = "token"
	)
	result := CheckClientToken(secretKey, projectKey, user, timestamp, info, providedToken)
	if result {
		t.Error("provided token is wrong, but check passed")
	}
	correctToken := GenerateClientToken(secretKey, projectKey, user, timestamp, info)
	result = CheckClientToken(secretKey, projectKey, user, timestamp, info, correctToken)
	if !result {
		t.Error("correct client token must pass check")
	}
}

func TestGenerateApiSign(t *testing.T) {
	var (
		secretKey   = "secret"
		projectKey  = "project"
		encodedData = "{}"
	)
	sign := GenerateApiSign(secretKey, projectKey, encodedData)
	if len(sign) != 64 {
		t.Error("sha256 sign length must be 64")
	}
}

func TestCheckApiSign(t *testing.T) {
	var (
		secretKey    = "secret"
		projectKey   = "project"
		encodedData  = "{}"
		providedSign = "sign"
	)
	result := CheckApiSign(secretKey, projectKey, encodedData, providedSign)
	if result {
		t.Error("provided sign is wrong, but check passed")
	}
	correctSign := GenerateApiSign(secretKey, projectKey, encodedData)
	result = CheckApiSign(secretKey, projectKey, encodedData, correctSign)
	if !result {
		t.Error("correct sign must pass check")
	}
}

func TestGenerateChannelSign(t *testing.T) {
	var (
		secretKey   = "secret"
		client      = "client"
		channel     = "channel"
		channelData = "{}"
	)
	sign := GenerateChannelSign(secretKey, client, channel, channelData)
	if len(sign) != 64 {
		t.Error("sha256 sign length must be 64")
	}
}

func TestCheckChannelSign(t *testing.T) {
	var (
		secretKey    = "secret"
		client       = "client"
		channel      = "channel"
		channelData  = "{}"
		providedSign = "sign"
	)
	result := CheckChannelSign(secretKey, client, channel, channelData, providedSign)
	if result {
		t.Error("provided sign is wrong, but check passed")
	}
	correctSign := GenerateChannelSign(secretKey, client, channel, channelData)
	result = CheckChannelSign(secretKey, client, channel, channelData, correctSign)
	if !result {
		t.Error("correct sign must pass check")
	}
}
