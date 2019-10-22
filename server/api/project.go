package api

import (
	"fmt"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"

	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/data"
)

type ProjectControllerOptions struct {
	Database *sqlxcache.DB
}

func ProjectType(project *data.Project) *app.Project {
	return &app.Project{
		ID:          int(project.ID),
		Name:        project.Name,
		Slug:        project.Slug,
		Description: project.Description,
		Goal:        project.Goal,
		Location:    project.Location,
		Tags:        project.Tags,
		Private:     project.Private,
		StartTime:   project.StartTime,
		EndTime:     project.EndTime,
	}
}

func ProjectsType(projects []*data.Project) *app.Projects {
	projectsCollection := make([]*app.Project, len(projects))
	for i, project := range projects {
		projectsCollection[i] = ProjectType(project)
	}

	return &app.Projects{
		Projects: projectsCollection,
	}
}

// ProjectController implements the user resource.
type ProjectController struct {
	*goa.Controller
	options ProjectControllerOptions
}

func NewProjectController(service *goa.Service, options ProjectControllerOptions) *ProjectController {
	return &ProjectController{
		Controller: service.NewController("ProjectController"),
		options:    options,
	}
}

func (c *ProjectController) Add(ctx *app.AddProjectContext) error {
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return fmt.Errorf("JWT claims error") // internal error
	}

	goal := ""
	if ctx.Payload.Goal != nil {
		goal = *ctx.Payload.Goal
	}

	location := ""
	if ctx.Payload.Location != nil {
		location = *ctx.Payload.Location
	}

	tags := ""
	if ctx.Payload.Location != nil {
		tags = *ctx.Payload.Tags
	}

	private := true
	if ctx.Payload.Private != nil {
		private = *ctx.Payload.Private
	}

	project := &data.Project{
		Name:        ctx.Payload.Name,
		Slug:        ctx.Payload.Slug,
		Description: ctx.Payload.Description,
		Goal:        goal,
		Location:    location,
		Tags:        tags,
		Private:     private,
		StartTime:   ctx.Payload.StartTime,
		EndTime:     ctx.Payload.EndTime,
	}

	if err := c.options.Database.NamedGetContext(ctx, project, "INSERT INTO fieldkit.project (name, slug, description, goal, location, tags, private, start_time, end_time) VALUES (:name, :slug, :description, :goal, :location, :tags, :private, :start_time, :end_time) RETURNING *", project); err != nil {
		return err
	}

	if _, err := c.options.Database.ExecContext(ctx, "INSERT INTO fieldkit.project_user (project_id, user_id) VALUES ($1, $2)", project.ID, claims["sub"]); err != nil {
		return err
	}

	return ctx.OK(ProjectType(project))
}

func (c *ProjectController) Update(ctx *app.UpdateProjectContext) error {
	goal := ""
	if ctx.Payload.Goal != nil {
		goal = *ctx.Payload.Goal
	}

	location := ""
	if ctx.Payload.Location != nil {
		location = *ctx.Payload.Location
	}

	tags := ""
	if ctx.Payload.Location != nil {
		tags = *ctx.Payload.Tags
	}

	private := true
	if ctx.Payload.Private != nil {
		private = *ctx.Payload.Private
	}

	project := &data.Project{
		ID:          int32(ctx.ProjectID),
		Name:        ctx.Payload.Name,
		Slug:        ctx.Payload.Slug,
		Description: ctx.Payload.Description,
		Goal:        goal,
		Location:    location,
		Tags:        tags,
		Private:     private,
		StartTime:   ctx.Payload.StartTime,
		EndTime:     ctx.Payload.EndTime,
	}

	if err := c.options.Database.NamedGetContext(ctx, project, "UPDATE fieldkit.project SET name = :name, slug = :slug, description = :description, goal = :goal, location = :location, tags = :tags, private = :private, start_time = :start_time, end_time = :end_time WHERE id = :id RETURNING *", project); err != nil {
		return err
	}

	return ctx.OK(ProjectType(project))
}

func (c *ProjectController) Get(ctx *app.GetProjectContext) error {
	project := &data.Project{}
	if err := c.options.Database.GetContext(ctx, project, "SELECT * FROM fieldkit.project WHERE slug = $1", ctx.Project); err != nil {
		return err
	}

	return ctx.OK(ProjectType(project))
}

func (c *ProjectController) GetID(ctx *app.GetIDProjectContext) error {
	project := &data.Project{}
	if err := c.options.Database.GetContext(ctx, project, "SELECT * FROM fieldkit.project WHERE id = $1", ctx.ProjectID); err != nil {
		return err
	}

	return ctx.OK(ProjectType(project))
}

func (c *ProjectController) List(ctx *app.ListProjectContext) error {
	projects := []*data.Project{}
	if err := c.options.Database.SelectContext(ctx, &projects, "SELECT * FROM fieldkit.project"); err != nil {
		return err
	}

	return ctx.OK(ProjectsType(projects))
}

func (c *ProjectController) ListCurrent(ctx *app.ListCurrentProjectContext) error {
	token := jwt.ContextJWT(ctx)
	if token == nil {
		return fmt.Errorf("JWT token is missing from context") // internal error
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return fmt.Errorf("JWT claims error") // internal error
	}

	projects := []*data.Project{}
	if err := c.options.Database.SelectContext(ctx, &projects, "SELECT p.* FROM fieldkit.project AS p JOIN fieldkit.project_user AS u ON u.project_id = p.id WHERE u.user_id = $1", claims["sub"]); err != nil {
		return err
	}

	return ctx.OK(ProjectsType(projects))
}
