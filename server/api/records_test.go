package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"

	"github.com/fieldkit/cloud/server/tests"
)

func TestGetMissingDataRecord(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/records/data/%d", 0), nil)
	req.Header.Add("Authorization", e.NewAuthorizationHeader())
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusNotFound, rr.Code)
}

func TestGetMetaRecord(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	fd, err := e.AddStations(1)
	assert.NoError(err)

	ar, err := e.AddMetaAndData(fd.Stations[0], fd.Owner)
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/records/meta/%d", ar.MetaRecords[0].ID), nil)
	req.Header.Add("Authorization", e.NewAuthorizationHeader())
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusOK, rr.Code)

	ja := jsonassert.New(t)
	ja.Assertf(rr.Body.String(), `
	{
		"meta": {
			"id": "<<PRESENCE>>",
			"number": "<<PRESENCE>>",
			"provision_id": "<<PRESENCE>>",
			"raw": "<<PRESENCE>>",
			"time": "<<PRESENCE>>"
		}
	}`)
}

func TestGetDataRecord(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	fd, err := e.AddStations(1)
	assert.NoError(err)

	ar, err := e.AddMetaAndData(fd.Stations[0], fd.Owner)
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/records/data/%d", ar.DataRecords[0].ID), nil)
	req.Header.Add("Authorization", e.NewAuthorizationHeader())
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusOK, rr.Code)

	ja := jsonassert.New(t)
	ja.Assertf(rr.Body.String(), `
	{
		"meta": {
			"id": "<<PRESENCE>>",
			"number": "<<PRESENCE>>",
			"provision_id": "<<PRESENCE>>",
			"raw": "<<PRESENCE>>",
			"time": "<<PRESENCE>>"
		},
		"data": {
			"time": "<<PRESENCE>>",
			"id": "<<PRESENCE>>",
			"location": "<<PRESENCE>>",
			"meta": "<<PRESENCE>>",
			"number": "<<PRESENCE>>",
			"provision_id": "<<PRESENCE>>",
			"raw": "<<PRESENCE>>"
		}
	}`)
}

func TestGetDataRecordResolved(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	fd, err := e.AddStations(1)
	assert.NoError(err)

	ar, err := e.AddMetaAndData(fd.Stations[0], fd.Owner)
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/records/data/%d/resolved", ar.DataRecords[0].ID), nil)
	req.Header.Add("Authorization", e.NewAuthorizationHeader())
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusOK, rr.Code)

	ja := jsonassert.New(t)
	ja.Assertf(rr.Body.String(), `
	{
		"filtered": {
			"filters": "<<PRESENCE>>",
			"record": "<<PRESENCE>>"
		},
		"meta": {
			"id": "<<PRESENCE>>",
			"number": "<<PRESENCE>>",
			"provision_id": "<<PRESENCE>>",
			"raw": "<<PRESENCE>>",
			"time": "<<PRESENCE>>"
		},
		"data": {
			"time": "<<PRESENCE>>",
			"id": "<<PRESENCE>>",
			"location": "<<PRESENCE>>",
			"meta": "<<PRESENCE>>",
			"number": "<<PRESENCE>>",
			"provision_id": "<<PRESENCE>>",
			"raw": "<<PRESENCE>>"
		}
	}`)
}

func TestGetDataRecordFiltered(t *testing.T) {
	assert := assert.New(t)
	e, err := tests.NewTestEnv()
	assert.NoError(err)

	fd, err := e.AddStations(1)
	assert.NoError(err)

	ar, err := e.AddMetaAndData(fd.Stations[0], fd.Owner)
	assert.NoError(err)

	api, err := NewTestableApi(e)
	assert.NoError(err)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/records/data/%d/filtered", ar.DataRecords[0].ID), nil)
	req.Header.Add("Authorization", e.NewAuthorizationHeader())
	rr := tests.ExecuteRequest(req, api)

	assert.Equal(http.StatusOK, rr.Code)

	// NOTE This always just returns {}
	ja := jsonassert.New(t)
	ja.Assertf(rr.Body.String(), `
	{
	}`)
}
