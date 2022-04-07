package client_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		require.Equal(t, `{"query":"last_project(n:$n){data}","variables":{"n":1}}`, string(b))

		err = json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"info": "bob, 1",
			},
		})
		if err != nil {
			panic(err)
		}
	})

	c := client.New(h)

	var resp struct {
		Info string
	}

	c.MustPost("last_project(n:$n){data}", &resp, client.Var("n", 1))

	require.Equal(t, "bob, 1", resp.Info)
}
