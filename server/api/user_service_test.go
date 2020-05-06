package api

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fieldkit/cloud/server/tests"
)

func TestLoginGood(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	user, err := e.AddUser("goodgoodgood")
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	rbody := strings.NewReader(fmt.Sprintf(`{ "email": "%s", "password": "goodgoodgood" }`, user.Email))
	req, _ := http.NewRequest("POST", "/login", rbody)
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusNoContent, rr.Code)
}

func TestLoginBad(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	user, err := e.AddUser("goodgoodgood")
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	rbody := strings.NewReader(fmt.Sprintf(`{ "email": "%s", "password": "wrongpassword" }`, user.Email))
	req, _ := http.NewRequest("POST", "/login", rbody)
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusUnauthorized, rr.Code)
}

func TestGetAvailableRoles(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	req, _ := http.NewRequest("GET", "/roles", nil)
	req.Header.Add("Authorization", e.NewAuthorizationHeader())
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusOK, rr.Code)
}