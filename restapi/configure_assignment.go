// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/MihaFriskovec/3fs-assignment/models"
	"github.com/MihaFriskovec/3fs-assignment/restapi/operations"
	"github.com/MihaFriskovec/3fs-assignment/restapi/operations/groups"
	"github.com/MihaFriskovec/3fs-assignment/restapi/operations/users"
	user "github.com/MihaFriskovec/3fs-assignment/users/api"

	group "github.com/MihaFriskovec/3fs-assignment/groups/api"
)

//go:generate swagger generate server --target ../../3fs-assignment --name Assignment --spec ../swagger.yml

func configureFlags(api *operations.AssignmentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AssignmentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GroupsCreateGroupHandler = groups.CreateGroupHandlerFunc(func(params groups.CreateGroupParams) middleware.Responder {
		err, _ := group.Create(params.Body)

		if err != nil {
			return groups.NewCreateGroupDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		return groups.NewCreateGroupCreated().WithPayload(params.Body)
	})

	api.GroupsListGroupsHandler = groups.ListGroupsHandlerFunc(func(params groups.ListGroupsParams) middleware.Responder {
		var page int64 = 1
		var limit int64 = 1
		var sort string = ""
		var project string = ""

		if params.Page != nil {
			page = *params.Page
		}

		if params.Limit != nil {
			limit = *params.Limit
		}

		if params.Sort != nil {
			sort = *params.Sort
		}

		if params.Select != nil {
			project = *params.Select
		}

		err, groupList := group.List(page, limit, sort, project)

		if err != nil {
			return groups.NewListGroupsBadRequest().WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		return groups.NewListGroupsOK().WithPayload(groupList)
	})

	api.GroupsReadGroupHandler = groups.ReadGroupHandlerFunc(func(params groups.ReadGroupParams) middleware.Responder {
		groupId, err := primitive.ObjectIDFromHex(params.ID)

		if err != nil {
			return groups.NewCreateGroupBadRequest().WithPayload(&models.Error{Type: "User error", Message: "Invalid ObjectId", Status: 400})
		}

		err, group := group.Read(groupId)

		if err != nil {
			return groups.NewReadGroupDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		return groups.NewReadGroupOK().WithPayload(group)
	})

	api.GroupsDeleteGroupHandler = groups.DeleteGroupHandlerFunc(func(params groups.DeleteGroupParams) middleware.Responder {
		groupID, err := primitive.ObjectIDFromHex(params.ID)

		if err != nil {
			return groups.NewDeleteGroupBadRequest().WithPayload(&models.Error{Type: "User error", Message: "Invalid ObjectId", Status: 400})
		}

		err, group := group.Delete(groupID)

		if err != nil {
			return groups.NewReadGroupDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		if group.DeletedCount > 0 {
			return groups.NewDeleteGroupNoContent()
		}

		return groups.NewDeleteGroupBadRequest().WithPayload(&models.Error{Type: "Application error", Message: "Group with given id not found", Status: 500})
	})

	api.GroupsUpdateGroupHandler = groups.UpdateGroupHandlerFunc(func(params groups.UpdateGroupParams) middleware.Responder {
		groupID, err := primitive.ObjectIDFromHex(params.ID)

		if err != nil {
			return groups.NewUpdateGroupBadRequest().WithPayload(&models.Error{Type: "User error", Message: "Invalid ObjectId", Status: 400})
		}

		err, group := group.Update(params.Body, groupID)

		if err != nil {
			return groups.NewReadGroupDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		if group.ModifiedCount > 0 {
			return groups.NewUpdateGroupOK().WithPayload(params.Body)
		}

		return groups.NewUpdateGroupDefault(500).WithPayload(&models.Error{Type: "Application error", Message: "Group with given id not found", Status: 500})
	})

	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
		err, _ := user.Create(params.Body)

		if err != nil {
			return users.NewCreateUserDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		return users.NewCreateUserCreated().WithPayload(params.Body)
	})

	api.UsersListUsersHandler = users.ListUsersHandlerFunc(func(params users.ListUsersParams) middleware.Responder {
		var page int64 = 1
		var limit int64 = 1
		var sort string = ""
		var project string = ""

		if params.Page != nil {
			page = *params.Page
		}

		if params.Limit != nil {
			limit = *params.Limit
		}

		if params.Sort != nil {
			sort = *params.Sort
		}

		if params.Select != nil {
			project = *params.Select
		}

		err, userList := user.List(page, limit, sort, project)

		if err != nil {
			return users.NewListUsersDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		return users.NewListUsersOK().WithPayload(userList)
	})

	api.UsersReadUserHandler = users.ReadUserHandlerFunc(func(params users.ReadUserParams) middleware.Responder {
		userID, err := primitive.ObjectIDFromHex(params.ID)

		if err != nil {
			return users.NewReadUserBadRequest().WithPayload(&models.Error{Type: "User error", Message: "Invalid ObjectId", Status: 400})
		}

		err, user := user.Read(userID)

		if err != nil {
			return users.NewReadUserDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		return users.NewReadUserOK().WithPayload(user)
	})

	api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(func(params users.DeleteUserParams) middleware.Responder {
		userID, err := primitive.ObjectIDFromHex(params.ID)

		if err != nil {
			return users.NewDeleteUserBadRequest().WithPayload(&models.Error{Type: "User error", Message: "Invalid ObjectId", Status: 400})
		}

		err, user := user.Delete(userID)

		if err != nil {
			return users.NewDeleteUserDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		if user.DeletedCount > 0 {
			return users.NewDeleteUserNoContent()
		}

		return users.NewDeleteUserBadRequest().WithPayload(&models.Error{Type: "Application error", Message: "User with given id not found", Status: 500})
	})

	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams) middleware.Responder {
		userID, err := primitive.ObjectIDFromHex(params.ID)

		if err != nil {
			return users.NewUpdateUserBadRequest().WithPayload(&models.Error{Type: "User error", Message: "Invalid ObjectId", Status: 400})
		}

		err, user := user.Update(params.Body, userID)

		if err != nil {
			return users.NewUpdateUserDefault(500).WithPayload(&models.Error{Type: "Database error", Message: err.Error(), Status: 500})
		}

		if user.MatchedCount > 0 {
			return users.NewUpdateUserOK().WithPayload(params.Body)
		}

		return users.NewUpdateUserDefault(500).WithPayload(&models.Error{Type: "Application error", Message: "User with given id not found", Status: 500})
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
