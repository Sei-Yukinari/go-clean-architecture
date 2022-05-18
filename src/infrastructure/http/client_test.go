package http

import (
	"encoding/json"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var mockBaseURL = "https://www.mock.co.jp/"

type User struct {
	ID   int
	Name string
}

func TestRequestReturnsSuccess(t *testing.T) {
	user := &User{
		ID:   1,
		Name: "Name",
	}
	expected, _ := json.Marshal(user)
	var actual User
	endpoint := "user"
	client := New(mockBaseURL)
	//モックサーバー作成
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	t.Run("Get()", func(t *testing.T) {
		httpmock.RegisterResponder(
			http.MethodGet,
			fmt.Sprintf("%s%s", mockBaseURL, endpoint),
			httpmock.NewBytesResponder(http.StatusOK, expected))

		//HTTPクライアント作成
		request, err := client.Get(endpoint)
		assert.NoError(t, err)

		//モックサーバーにリクエスト
		response, err := client.Do(request)
		assert.NoError(t, err)

		response.To(&actual)
		assert.Equal(t, user.ID, actual.ID)
	})
	t.Run("PostWith()", func(t *testing.T) {
		httpmock.RegisterResponder(
			http.MethodPost,
			fmt.Sprintf("%s%s", mockBaseURL, endpoint),
			httpmock.NewBytesResponder(http.StatusOK, expected))
		request, err := client.PostWith(endpoint, user)
		assert.NoError(t, err)

		//モックサーバーにリクエスト
		response, err := client.Do(request)
		assert.NoError(t, err)

		response.To(&actual)
		assert.Equal(t, user.ID, actual.ID)
	})
}
