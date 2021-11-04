// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"devsmake/models"
	"devsmake/persistence"
	"devsmake/restapi/handlers"
	"devsmake/restapi/operations"
	"devsmake/util"
)

//go:generate swagger generate server --target ../../devsmake --name DevsMake --spec ../spec/swagger.yaml --principal models.Principal

func configureFlags(api *operations.DevsMakeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DevsMakeAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	repo, err := persistence.NewRepository()

	if err != nil {
		panic(err)
	}

	api.GithubAuthAuth = func(token string, scopes []string) (*models.Principal, error) {
		ok, err := util.IsAuthenticated(token)
		if err != nil {
			return nil, errors.New(401, "error authenticate")
		}
		if !ok {
			return nil, errors.New(401, "invalid token")
		}

		prin := models.Principal(token)
		return &prin, nil
	}

	api.GeneralGetHealthcheckHandler = handlers.NewHealthCheckHandler(repo)

	api.AuthGetAuthCallbackHandler = handlers.NewAuthCallbackHandler(repo.AccountRepo)
	api.AuthGetAuthLoginHandler = handlers.NewAuthLoginHandler()

	api.ProfileGetProfileHandler = handlers.NewProfileHandler(repo.AccountRepo)
	api.ProfileGetProfileIDHandler = handlers.NewProfileIDHandler(repo.AccountRepo)
	api.ProfileGetProfilesHandler = handlers.NewProfilesHandler(repo.AccountRepo)

	api.IdeaPostGetIdeasHandler = handlers.NewIdeasHandler(repo.PostRepos)
	api.IdeaPostGetIdeasUUIDHandler = handlers.NewIdeasUUIDHandler(repo.PostRepos)

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			handler.ServeHTTP(w, r)
		} else {
			http.FileServer(http.Dir("./frontend/dist")).ServeHTTP(w, r)
		}
	})
}
