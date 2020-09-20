// Code generated by goa v3.1.2, DO NOT EDIT.
//
// collection HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	collection "github.com/fieldkit/cloud/server/api/gen/collection"
)

// BuildAddPayload builds the payload for the collection add endpoint from CLI
// flags.
func BuildAddPayload(collectionAddBody string, collectionAddAuth string) (*collection.AddPayload, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(collectionAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"description\": \"Nisi dolor.\",\n      \"name\": \"Dolor quo.\",\n      \"private\": true,\n      \"tags\": \"Voluptates maxime quae fuga nobis repellat debitis.\"\n   }'")
		}
	}
	var auth string
	{
		auth = collectionAddAuth
	}
	v := &collection.AddCollectionFields{
		Name:        body.Name,
		Description: body.Description,
		Tags:        body.Tags,
		Private:     body.Private,
	}
	res := &collection.AddPayload{
		Collection: v,
	}
	res.Auth = auth

	return res, nil
}

// BuildUpdatePayload builds the payload for the collection update endpoint
// from CLI flags.
func BuildUpdatePayload(collectionUpdateBody string, collectionUpdateCollectionID string, collectionUpdateAuth string) (*collection.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(collectionUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"description\": \"Quae ut voluptatem veritatis autem alias suscipit.\",\n      \"name\": \"Dignissimos vero dolorem.\",\n      \"private\": false,\n      \"tags\": \"Repellat et aperiam minima cum.\"\n   }'")
		}
	}
	var collectionID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionUpdateCollectionID, 10, 32)
		collectionID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for collectionID, must be INT32")
		}
	}
	var auth string
	{
		auth = collectionUpdateAuth
	}
	v := &collection.AddCollectionFields{
		Name:        body.Name,
		Description: body.Description,
		Tags:        body.Tags,
		Private:     body.Private,
	}
	res := &collection.UpdatePayload{
		Collection: v,
	}
	res.CollectionID = collectionID
	res.Auth = auth

	return res, nil
}

// BuildGetPayload builds the payload for the collection get endpoint from CLI
// flags.
func BuildGetPayload(collectionGetCollectionID string, collectionGetAuth string) (*collection.GetPayload, error) {
	var err error
	var collectionID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionGetCollectionID, 10, 32)
		collectionID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for collectionID, must be INT32")
		}
	}
	var auth *string
	{
		if collectionGetAuth != "" {
			auth = &collectionGetAuth
		}
	}
	v := &collection.GetPayload{}
	v.CollectionID = collectionID
	v.Auth = auth

	return v, nil
}

// BuildListMinePayload builds the payload for the collection list mine
// endpoint from CLI flags.
func BuildListMinePayload(collectionListMineAuth string) (*collection.ListMinePayload, error) {
	var auth string
	{
		auth = collectionListMineAuth
	}
	v := &collection.ListMinePayload{}
	v.Auth = auth

	return v, nil
}

// BuildAddStationPayload builds the payload for the collection add station
// endpoint from CLI flags.
func BuildAddStationPayload(collectionAddStationCollectionID string, collectionAddStationStationID string, collectionAddStationAuth string) (*collection.AddStationPayload, error) {
	var err error
	var collectionID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionAddStationCollectionID, 10, 32)
		collectionID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for collectionID, must be INT32")
		}
	}
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionAddStationStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = collectionAddStationAuth
	}
	v := &collection.AddStationPayload{}
	v.CollectionID = collectionID
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}

// BuildRemoveStationPayload builds the payload for the collection remove
// station endpoint from CLI flags.
func BuildRemoveStationPayload(collectionRemoveStationCollectionID string, collectionRemoveStationStationID string, collectionRemoveStationAuth string) (*collection.RemoveStationPayload, error) {
	var err error
	var collectionID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionRemoveStationCollectionID, 10, 32)
		collectionID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for collectionID, must be INT32")
		}
	}
	var stationID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionRemoveStationStationID, 10, 32)
		stationID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for stationID, must be INT32")
		}
	}
	var auth string
	{
		auth = collectionRemoveStationAuth
	}
	v := &collection.RemoveStationPayload{}
	v.CollectionID = collectionID
	v.StationID = stationID
	v.Auth = auth

	return v, nil
}

// BuildDeletePayload builds the payload for the collection delete endpoint
// from CLI flags.
func BuildDeletePayload(collectionDeleteCollectionID string, collectionDeleteAuth string) (*collection.DeletePayload, error) {
	var err error
	var collectionID int32
	{
		var v int64
		v, err = strconv.ParseInt(collectionDeleteCollectionID, 10, 32)
		collectionID = int32(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for collectionID, must be INT32")
		}
	}
	var auth string
	{
		auth = collectionDeleteAuth
	}
	v := &collection.DeletePayload{}
	v.CollectionID = collectionID
	v.Auth = auth

	return v, nil
}