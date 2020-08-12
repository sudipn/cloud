// Code generated by goa v3.1.2, DO NOT EDIT.
//
// csv views
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// UserExports is the viewed result type that is projected based on a view.
type UserExports struct {
	// Type to project
	Projected *UserExportsView
	// View to render
	View string
}

// ExportStatus is the viewed result type that is projected based on a view.
type ExportStatus struct {
	// Type to project
	Projected *ExportStatusView
	// View to render
	View string
}

// UserExportsView is a type that runs validations on a projected type.
type UserExportsView struct {
	Exports []*ExportStatusView
}

// ExportStatusView is a type that runs validations on a projected type.
type ExportStatusView struct {
	ID          *int64
	CreatedAt   *int64
	CompletedAt *int64
	Progress    *float32
	URL         *string
	Kind        *string
	Args        interface{}
}

var (
	// UserExportsMap is a map of attribute names in result type UserExports
	// indexed by view name.
	UserExportsMap = map[string][]string{
		"default": []string{
			"exports",
		},
	}
	// ExportStatusMap is a map of attribute names in result type ExportStatus
	// indexed by view name.
	ExportStatusMap = map[string][]string{
		"default": []string{
			"id",
			"createdAt",
			"completedAt",
			"kind",
			"progress",
			"url",
			"args",
		},
	}
)

// ValidateUserExports runs the validations defined on the viewed result type
// UserExports.
func ValidateUserExports(result *UserExports) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUserExportsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateExportStatus runs the validations defined on the viewed result type
// ExportStatus.
func ValidateExportStatus(result *ExportStatus) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateExportStatusView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateUserExportsView runs the validations defined on UserExportsView
// using the "default" view.
func ValidateUserExportsView(result *UserExportsView) (err error) {
	if result.Exports == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("exports", "result"))
	}
	for _, e := range result.Exports {
		if e != nil {
			if err2 := ValidateExportStatusView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateExportStatusView runs the validations defined on ExportStatusView
// using the "default" view.
func ValidateExportStatusView(result *ExportStatusView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "result"))
	}
	if result.Progress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("progress", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Args == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("args", "result"))
	}
	return
}
