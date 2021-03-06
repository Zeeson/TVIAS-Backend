package controllers

import (
	"bitbucket.org/staydigital/truvest-identity-management/api/middleware"
	_ "bitbucket.org/staydigital/truvest-identity-management/docs" // docs is generated by Swag CLI, you have to import it.

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) initializeRoutes() {

	// Index Page route. Sample page to demo OAuth2
	s.Router.HandleFunc("/", s.IndexPage).Methods("GET")

	// Heartbeat Route
	s.Router.HandleFunc("/heartbeat", middleware.SetMiddlewareJSON(s.Heartbeat)).Methods("GET")
	
	// SignUp Route
	s.Router.HandleFunc("/signup", middleware.SetMiddlewareJSON(s.SignUp)).Methods("POST")

	// Login Route
	s.Router.HandleFunc("/login", middleware.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/refresh", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.Refresh))).Methods("POST")
	s.Router.HandleFunc("/logout", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.Logout))).Methods("POST")

	// OAuth Route
	s.Router.HandleFunc("/auth/{provider}", middleware.SetMiddlewareJSON(s.OauthSignIn)).Methods("GET")
	s.Router.HandleFunc("/auth/{provider}/callback", middleware.SetMiddlewareJSON(s.OauthSuccessCallback)).Methods("GET")
	s.Router.HandleFunc("/logout/{provider}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.OauthLogout))).Methods("POST")

	// Users routes
	s.Router.HandleFunc("/user/me", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetLoggedInUser))).Methods("GET")
	s.Router.HandleFunc("/users", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.CreateUser))).Methods("POST")
	s.Router.HandleFunc("/users", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetUser))).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middleware.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	s.Router.HandleFunc("/users/{id}/setPassword", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.SetPassword))).Methods("POST")
	s.Router.HandleFunc("/users/forgotPassword", middleware.SetMiddlewareJSON(s.ForgotPassword)).Methods("POST")
	s.Router.HandleFunc("/users/sendMail", middleware.SetMiddlewareJSON(s.SendMail)).Methods("POST")
	s.Router.HandleFunc("/users/{id}/enableUser", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.EnableUser))).Methods("PUT")

	// Permission routes
	s.Router.HandleFunc("/permissions", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.CreatePermission))).Methods("POST")
	s.Router.HandleFunc("/permissions", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetPermissions))).Methods("GET")
	s.Router.HandleFunc("/permissions/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetPermission))).Methods("GET")
	s.Router.HandleFunc("/permissions/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdatePermission))).Methods("PUT")
	s.Router.HandleFunc("/permissions/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.DeletePermission))).Methods("DELETE")

	// Roles routes
	s.Router.HandleFunc("/roles", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.CreateRole))).Methods("POST")
	s.Router.HandleFunc("/roles", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetRoles))).Methods("GET")
	s.Router.HandleFunc("/roles/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.GetRole))).Methods("GET")
	s.Router.HandleFunc("/roles/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.UpdateRole))).Methods("PUT")
	s.Router.HandleFunc("/roles/{id}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.DeleteRole))).Methods("DELETE")

	// Map Users to Roles routes
	s.Router.HandleFunc("/roles/{id}/users", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.AddUsersToRole))).Methods("POST")
	s.Router.HandleFunc("/roles/{id1}/users/{id2}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.DeleteUsersFromRole))).Methods("DELETE")

	// Map Permissions to Roles routes
	s.Router.HandleFunc("/roles/{id}/permissions", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.AddPermissionsToRole))).Methods("POST")
	s.Router.HandleFunc("/roles/{id1}/permissions/{id2}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(s.DeletePermissionsFromRole))).Methods("DELETE")

	// Swagger
    s.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}