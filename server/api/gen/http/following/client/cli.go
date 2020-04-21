// Code generated by goa v3.1.1, DO NOT EDIT.
//
// following HTTP client CLI support package
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"fmt"
	"strconv"

	following "github.com/fieldkit/cloud/server/api/gen/following"
)

// BuildFollowPayload builds the payload for the following follow endpoint from
// CLI flags.
func BuildFollowPayload(followingFollowID string, followingFollowAuth string) (*following.FollowPayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(followingFollowID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var auth *string
	{
		if followingFollowAuth != "" {
			auth = &followingFollowAuth
		}
	}
	v := &following.FollowPayload{}
	v.ID = &id
	v.Auth = auth

	return v, nil
}

// BuildUnfollowPayload builds the payload for the following unfollow endpoint
// from CLI flags.
func BuildUnfollowPayload(followingUnfollowID string, followingUnfollowAuth string) (*following.UnfollowPayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(followingUnfollowID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var auth *string
	{
		if followingUnfollowAuth != "" {
			auth = &followingUnfollowAuth
		}
	}
	v := &following.UnfollowPayload{}
	v.ID = &id
	v.Auth = auth

	return v, nil
}

// BuildFollowersPayload builds the payload for the following followers
// endpoint from CLI flags.
func BuildFollowersPayload(followingFollowersID string, followingFollowersPage string) (*following.FollowersPayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(followingFollowersID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	var page *int64
	{
		if followingFollowersPage != "" {
			val, err := strconv.ParseInt(followingFollowersPage, 10, 64)
			page = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for page, must be INT64")
			}
		}
	}
	v := &following.FollowersPayload{}
	v.ID = &id
	v.Page = page

	return v, nil
}
