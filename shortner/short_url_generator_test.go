package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.google.com"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.youtube.com/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://www.facebook.com"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "jTa4L57P")
	assert.Equal(t, shortLink_2, "d66yfx7N")
	assert.Equal(t, shortLink_3, "dhZTayYQ")
