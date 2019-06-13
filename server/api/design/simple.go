package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("simple", func() {
	Security(JWT, func() {
		Scope("api:access")
	})

	Action("my features", func() {
		Routing(GET("my/simple/features"))
		Response(NotFound)
		Params(func() {
			Param("page", Integer)
		})
		Response(OK, func() {
			Media(PagedGeoJSON)
		})
	})

	Action("my csv data", func() {
		Routing(GET("my/simple/data/csv"))
		Response(NotFound)
		Response(OK, func() {
			Status(200)
		})
	})
})