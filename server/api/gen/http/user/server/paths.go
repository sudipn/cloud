// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the user service.
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"fmt"
)

// RolesUserPath returns the URL path to the user service roles HTTP endpoint.
func RolesUserPath() string {
	return "/roles"
}

// DeleteUserPath returns the URL path to the user service delete HTTP endpoint.
func DeleteUserPath(userID int32) string {
	return fmt.Sprintf("/admin/users/%v", userID)
}
