// Code generated by goa v3.1.1, DO NOT EDIT.
//
// HTTP request path constructors for the following service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"fmt"
)

// FollowFollowingPath returns the URL path to the following service follow HTTP endpoint.
func FollowFollowingPath(id int64) string {
	return fmt.Sprintf("/projects/%v/follow", id)
}

// UnfollowFollowingPath returns the URL path to the following service unfollow HTTP endpoint.
func UnfollowFollowingPath(id int64) string {
	return fmt.Sprintf("/projects/%v/unfollow", id)
}

// FollowersFollowingPath returns the URL path to the following service followers HTTP endpoint.
func FollowersFollowingPath(id int64) string {
	return fmt.Sprintf("/projects/%v/followers", id)
}
