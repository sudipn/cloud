package api

import (
	"context"
	"fmt"
	"net/http"

	goa "goa.design/goa/v3/pkg"

	jwt "github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"

	goahttp "goa.design/goa/v3/http"
	"goa.design/goa/v3/middleware"

	"github.com/fieldkit/cloud/server/logging"

	testSvr "github.com/fieldkit/cloud/server/api/gen/http/test/server"
	test "github.com/fieldkit/cloud/server/api/gen/test"

	tasksSvr "github.com/fieldkit/cloud/server/api/gen/http/tasks/server"
	tasks "github.com/fieldkit/cloud/server/api/gen/tasks"

	modulesSvr "github.com/fieldkit/cloud/server/api/gen/http/modules/server"
	modules "github.com/fieldkit/cloud/server/api/gen/modules"

	following "github.com/fieldkit/cloud/server/api/gen/following"
	followingSvr "github.com/fieldkit/cloud/server/api/gen/http/following/server"

	activity "github.com/fieldkit/cloud/server/api/gen/activity"
	activitySvr "github.com/fieldkit/cloud/server/api/gen/http/activity/server"

	projectSvr "github.com/fieldkit/cloud/server/api/gen/http/project/server"
	project "github.com/fieldkit/cloud/server/api/gen/project"

	stationSvr "github.com/fieldkit/cloud/server/api/gen/http/station/server"
	station "github.com/fieldkit/cloud/server/api/gen/station"

	userSvr "github.com/fieldkit/cloud/server/api/gen/http/user/server"
	user "github.com/fieldkit/cloud/server/api/gen/user"
)

func LogErrors() func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = e(ctx, request)
			if err != nil {
				id := newErrorID()
				log := logging.Logger(ctx).Sugar()
				log.Errorw("error", "error", err, "error_id", id)
			}
			return response, err
		}
	}
}

func CreateGoaV3Handler(ctx context.Context, options *ControllerOptions) (http.Handler, error) {
	testSvc := NewTestSevice(ctx, options)
	testEndpoints := test.NewEndpoints(testSvc)

	tasksSvc := NewTasksService(ctx, options)
	tasksEndpoints := tasks.NewEndpoints(tasksSvc)

	modulesSvc := NewModulesService(ctx, options)
	modulesEndpoints := modules.NewEndpoints(modulesSvc)

	followingSvc := NewFollowingService(ctx, options)
	followingEndpoints := following.NewEndpoints(followingSvc)

	activitySvc := NewActivityService(ctx, options)
	activityEndpoints := activity.NewEndpoints(activitySvc)

	projectSvc := NewProjectService(ctx, options)
	projectEndpoints := project.NewEndpoints(projectSvc)

	stationSvc := NewStationService(ctx, options)
	stationEndpoints := station.NewEndpoints(stationSvc)

	userSvc := NewUserService(ctx, options)
	userEndpoints := user.NewEndpoints(userSvc)

	logErrors := LogErrors()

	modulesEndpoints.Use(logErrors)
	tasksEndpoints.Use(logErrors)
	testEndpoints.Use(logErrors)
	followingEndpoints.Use(logErrors)
	activityEndpoints.Use(logErrors)
	projectEndpoints.Use(logErrors)
	stationEndpoints.Use(logErrors)
	userEndpoints.Use(logErrors)

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	mux := goahttp.NewMuxer()

	eh := errorHandler()

	tasksServer := tasksSvr.New(tasksEndpoints, mux, dec, enc, eh, nil)
	testServer := testSvr.New(testEndpoints, mux, dec, enc, eh, nil)
	modulesServer := modulesSvr.New(modulesEndpoints, mux, dec, enc, eh, nil)
	followingServer := followingSvr.New(followingEndpoints, mux, dec, enc, eh, nil)
	activityServer := activitySvr.New(activityEndpoints, mux, dec, enc, eh, nil)
	projectServer := projectSvr.New(projectEndpoints, mux, dec, enc, eh, nil)
	stationServer := stationSvr.New(stationEndpoints, mux, dec, enc, eh, nil)
	userServer := userSvr.New(userEndpoints, mux, dec, enc, eh, nil)

	tasksSvr.Mount(mux, tasksServer)
	testSvr.Mount(mux, testServer)
	modulesSvr.Mount(mux, modulesServer)
	followingSvr.Mount(mux, followingServer)
	activitySvr.Mount(mux, activityServer)
	projectSvr.Mount(mux, projectServer)
	stationSvr.Mount(mux, stationServer)
	userSvr.Mount(mux, userServer)

	log := Logger(ctx).Sugar()

	for _, m := range testServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range tasksServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range modulesServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range followingServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range activityServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range projectServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range stationServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range userServer.Mounts {
		log.Infow("mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}

	return mux, nil
}

func errorHandler() func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		log := Logger(ctx).Sugar()
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		log.Errorw("fatal", "id", id, "message", err.Error())
	}
}

type GenerateError func(string) error

type AuthAttempt struct {
	Token        string
	Scheme       *security.JWTScheme
	Key          []byte
	Unauthorized GenerateError
	NotFound     GenerateError
}

func Authenticate(ctx context.Context, a AuthAttempt) (context.Context, error) {
	claims := make(jwt.MapClaims)
	_, err := jwt.ParseWithClaims(a.Token, claims, func(t *jwt.Token) (interface{}, error) {
		return a.Key, nil
	})
	if err != nil {
		return ctx, a.Unauthorized("invalid token")
	}

	if claims["scopes"] == nil {
		return ctx, a.Unauthorized("invalid scopes")
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, a.Unauthorized("invalid scopes")
	}

	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := a.Scheme.Validate(scopesInToken); err != nil {
		return ctx, a.Unauthorized("invalid scopes")
	}

	withClaims := addClaimsToContext(ctx, claims)
	withAttempt := addAuthAttemptToContext(withClaims, &a)
	withLogging := logging.WithUserID(withAttempt, fmt.Sprintf("%v", claims["sub"]))

	return withLogging, nil
}