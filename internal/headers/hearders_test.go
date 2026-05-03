package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHearderparse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	host, hostOK := headers.Get("host")
	assert.True(t, hostOK)
	assert.Equal(t, "localhost:42069", host)
	dasdsad, notOk := headers.Get("dasdsad")
	assert.False(t, notOk)
	assert.Equal(t, "", dasdsad)
	assert.Equal(t, 25, n)
	assert.True(t, done)

	// // Test: muitlple value
	headers = NewHeaders()
	data = []byte("Set-Person: lane-loves-go\r\nSet-Person: prime-loves-zig\r\nSet-Person: tj-loves-ocaml\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	das, dasOk := headers.Get("dasdsad")
	assert.False(t, dasOk)
	assert.Equal(t, "", das)
	setPerson, setPersonOK := headers.Get("Set-Person")
	assert.True(t, setPersonOK)
	assert.Equal(t, "lane-loves-go,prime-loves-zig,tj-loves-ocaml", setPerson)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	headers = NewHeaders()
	data = []byte("H©st: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)
}
