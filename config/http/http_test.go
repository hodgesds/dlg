package http

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPRequestInvalidBodyFile(t *testing.T) {
	bodyFile := "not-a-file"
	p := &Payload{
		BodyFile: &bodyFile,
	}
	_, err := p.Request(context.Background())
	require.Error(t, err)
}

func TestHTTPRequestValidBodyFile(t *testing.T) {
	f, err := ioutil.TempFile("", "test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	body := []byte("test")
	_, err = f.Write(body)
	require.NoError(t, err)
	require.NoError(t, f.Sync())
	require.NoError(t, f.Close())
	bodyFile := f.Name()

	p := &Payload{
		BodyFile: &bodyFile,
	}

	r, err := p.Request(context.Background())
	require.NoError(t, err)
	reqBody := make([]byte, len(body))
	_, err = r.Body.Read(reqBody)
	require.NoError(t, err)
	require.Equal(t, body, reqBody)
}
