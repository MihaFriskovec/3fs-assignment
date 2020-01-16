// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/MihaFriskovec/3fs-assignment/restapi/operations/groups"
	"github.com/MihaFriskovec/3fs-assignment/restapi/operations/users"
)

// NewAssignmentAPI creates a new Assignment instance
func NewAssignmentAPI(spec *loads.Document) *AssignmentAPI {
	return &AssignmentAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		GroupsCreateGroupHandler: groups.CreateGroupHandlerFunc(func(params groups.CreateGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation GroupsCreateGroup has not yet been implemented")
		}),
		UsersCreateUserHandler: users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UsersCreateUser has not yet been implemented")
		}),
		GroupsDeleteGroupHandler: groups.DeleteGroupHandlerFunc(func(params groups.DeleteGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation GroupsDeleteGroup has not yet been implemented")
		}),
		UsersDeleteUserHandler: users.DeleteUserHandlerFunc(func(params users.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UsersDeleteUser has not yet been implemented")
		}),
		GroupsListGroupsHandler: groups.ListGroupsHandlerFunc(func(params groups.ListGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation GroupsListGroups has not yet been implemented")
		}),
		UsersListUsersHandler: users.ListUsersHandlerFunc(func(params users.ListUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation UsersListUsers has not yet been implemented")
		}),
		GroupsReadGroupHandler: groups.ReadGroupHandlerFunc(func(params groups.ReadGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation GroupsReadGroup has not yet been implemented")
		}),
		UsersReadUserHandler: users.ReadUserHandlerFunc(func(params users.ReadUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UsersReadUser has not yet been implemented")
		}),
		GroupsUpdateGroupHandler: groups.UpdateGroupHandlerFunc(func(params groups.UpdateGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation GroupsUpdateGroup has not yet been implemented")
		}),
		UsersUpdateUserHandler: users.UpdateUserHandlerFunc(func(params users.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation UsersUpdateUser has not yet been implemented")
		}),
	}
}

/*AssignmentAPI First assignment */
type AssignmentAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// GroupsCreateGroupHandler sets the operation handler for the create group operation
	GroupsCreateGroupHandler groups.CreateGroupHandler
	// UsersCreateUserHandler sets the operation handler for the create user operation
	UsersCreateUserHandler users.CreateUserHandler
	// GroupsDeleteGroupHandler sets the operation handler for the delete group operation
	GroupsDeleteGroupHandler groups.DeleteGroupHandler
	// UsersDeleteUserHandler sets the operation handler for the delete user operation
	UsersDeleteUserHandler users.DeleteUserHandler
	// GroupsListGroupsHandler sets the operation handler for the list groups operation
	GroupsListGroupsHandler groups.ListGroupsHandler
	// UsersListUsersHandler sets the operation handler for the list users operation
	UsersListUsersHandler users.ListUsersHandler
	// GroupsReadGroupHandler sets the operation handler for the read group operation
	GroupsReadGroupHandler groups.ReadGroupHandler
	// UsersReadUserHandler sets the operation handler for the read user operation
	UsersReadUserHandler users.ReadUserHandler
	// GroupsUpdateGroupHandler sets the operation handler for the update group operation
	GroupsUpdateGroupHandler groups.UpdateGroupHandler
	// UsersUpdateUserHandler sets the operation handler for the update user operation
	UsersUpdateUserHandler users.UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AssignmentAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AssignmentAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AssignmentAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AssignmentAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AssignmentAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AssignmentAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AssignmentAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AssignmentAPI
func (o *AssignmentAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GroupsCreateGroupHandler == nil {
		unregistered = append(unregistered, "groups.CreateGroupHandler")
	}

	if o.UsersCreateUserHandler == nil {
		unregistered = append(unregistered, "users.CreateUserHandler")
	}

	if o.GroupsDeleteGroupHandler == nil {
		unregistered = append(unregistered, "groups.DeleteGroupHandler")
	}

	if o.UsersDeleteUserHandler == nil {
		unregistered = append(unregistered, "users.DeleteUserHandler")
	}

	if o.GroupsListGroupsHandler == nil {
		unregistered = append(unregistered, "groups.ListGroupsHandler")
	}

	if o.UsersListUsersHandler == nil {
		unregistered = append(unregistered, "users.ListUsersHandler")
	}

	if o.GroupsReadGroupHandler == nil {
		unregistered = append(unregistered, "groups.ReadGroupHandler")
	}

	if o.UsersReadUserHandler == nil {
		unregistered = append(unregistered, "users.ReadUserHandler")
	}

	if o.GroupsUpdateGroupHandler == nil {
		unregistered = append(unregistered, "groups.UpdateGroupHandler")
	}

	if o.UsersUpdateUserHandler == nil {
		unregistered = append(unregistered, "users.UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AssignmentAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AssignmentAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *AssignmentAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *AssignmentAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AssignmentAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AssignmentAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the assignment API
func (o *AssignmentAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AssignmentAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/groups"] = groups.NewCreateGroup(o.context, o.GroupsCreateGroupHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = users.NewCreateUser(o.context, o.UsersCreateUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/groups/{id}"] = groups.NewDeleteGroup(o.context, o.GroupsDeleteGroupHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{id}"] = users.NewDeleteUser(o.context, o.UsersDeleteUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/groups"] = groups.NewListGroups(o.context, o.GroupsListGroupsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = users.NewListUsers(o.context, o.UsersListUsersHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/groups/{id}"] = groups.NewReadGroup(o.context, o.GroupsReadGroupHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{id}"] = users.NewReadUser(o.context, o.UsersReadUserHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/groups/{id}"] = groups.NewUpdateGroup(o.context, o.GroupsUpdateGroupHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{id}"] = users.NewUpdateUser(o.context, o.UsersUpdateUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AssignmentAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AssignmentAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AssignmentAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AssignmentAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
