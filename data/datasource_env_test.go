package data

import (
	"context"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustParseURL(in string) *url.URL {
	u, _ := url.Parse(in)
	return u
}

func TestReadEnv(t *testing.T) {
	ctx := context.Background()

	content := []byte(`hello world`)
	os.Setenv("HELLO_WORLD", "hello world")
	defer os.Unsetenv("HELLO_WORLD")
	os.Setenv("HELLO_UNIVERSE", "hello universe")
	defer os.Unsetenv("HELLO_UNIVERSE")

	source := &Source{Alias: "foo", URL: mustParseURL("env:HELLO_WORLD")}

	actual, err := readEnv(ctx, source)
	require.NoError(t, err)
	assert.Equal(t, content, actual)

	source = &Source{Alias: "foo", URL: mustParseURL("env:/HELLO_WORLD")}

	actual, err = readEnv(ctx, source)
	require.NoError(t, err)
	assert.Equal(t, content, actual)

	source = &Source{Alias: "foo", URL: mustParseURL("env:///HELLO_WORLD")}

	actual, err = readEnv(ctx, source)
	require.NoError(t, err)
	assert.Equal(t, content, actual)

	source = &Source{Alias: "foo", URL: mustParseURL("env:HELLO_WORLD?foo=bar")}

	actual, err = readEnv(ctx, source)
	require.NoError(t, err)
	assert.Equal(t, content, actual)

	source = &Source{Alias: "foo", URL: mustParseURL("env:///HELLO_WORLD?foo=bar")}

	actual, err = readEnv(ctx, source)
	require.NoError(t, err)
	assert.Equal(t, content, actual)
}
