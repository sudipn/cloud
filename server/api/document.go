package api

import (
	"strconv"

	"github.com/goadesign/goa"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/data"
)

func DocumentType(document *data.Document) *app.Document {
	locationType := &app.Point{
		Type:        "Point",
		Coordinates: document.Location.Coordinates(),
	}

	documentType := &app.Document{
		ID:        strconv.FormatInt(document.ID, 10),
		InputID:   int(document.InputID),
		Timestamp: int(document.Timestamp.Unix()),
		Location:  locationType,
		Data:      document.Data,
	}

	if document.TeamID != nil {
		teamID := int(*document.TeamID)
		documentType.TeamID = &teamID
	}

	if document.UserID != nil {
		userID := int(*document.UserID)
		documentType.UserID = &userID
	}

	return documentType
}

func DocumentsType(documents []*data.Document) *app.Documents {
	documentCollection := make([]*app.Document, len(documents))
	for i, document := range documents {
		documentCollection[i] = DocumentType(document)
	}

	return &app.Documents{
		Documents: documentCollection,
	}
}

type DocumentControllerOptions struct {
	Backend *backend.Backend
}

// DocumentController implements the document resource.
type DocumentController struct {
	*goa.Controller
	options DocumentControllerOptions
}

func NewDocumentController(service *goa.Service, options DocumentControllerOptions) *DocumentController {
	return &DocumentController{
		Controller: service.NewController("DocumentController"),
		options:    options,
	}
}

func (c *DocumentController) List(ctx *app.ListDocumentContext) error {
	documents, err := c.options.Backend.ListDocuments(ctx, ctx.Project, ctx.Expedition, "")
	if err != nil {
		return err
	}

	return ctx.OK(DocumentsType(documents.Documents))
}

func (c *DocumentController) ListID(ctx *app.ListIDDocumentContext) error {
	documents, err := c.options.Backend.ListDocumentsByID(ctx, int32(ctx.ExpeditionID))
	if err != nil {
		return err
	}

	return ctx.OK(DocumentsType(documents))
}
