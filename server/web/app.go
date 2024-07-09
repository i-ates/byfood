package web

import (
	"log"
	"net/http"
	"task4/routes"
)

// App represents the web application configuration.
type App struct {
	handlers   map[string]http.HandlerFunc
	userRouter routes.UserRouterInterface
}

// NewApp creates a new App instance with the specified configuration.
func NewApp(cors bool, userRouter routes.UserRouterInterface) App {
	app := App{
		handlers:   make(map[string]http.HandlerFunc),
		userRouter: userRouter,
	}

	// Define HTTP request handlers for user-related routes
	getUserHandler := app.userRouter.GetUsers
	addUserHandler := app.userRouter.AddUser
	updateUserHandler := app.userRouter.UpdateUser
	deleteUserHandler := app.userRouter.DeleteUser

	// Enable or disable CORS (Cross-Origin Resource Sharing) based on the 'cors' flag
	if !cors {
		getUserHandler = disableCors(getUserHandler)
		addUserHandler = disableCors(addUserHandler)
		updateUserHandler = disableCors(updateUserHandler)
		deleteUserHandler = disableCors(deleteUserHandler)
	}

	// Register the handlers for specific routes
	app.handlers["/api/get-users"] = getUserHandler
	app.handlers["/api/add-user"] = addUserHandler
	app.handlers["/api/update-user"] = updateUserHandler
	app.handlers["/api/delete-user"] = deleteUserHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP // Serve static files for the web app
	return app
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}

// Serve starts the web server and listens on port 8080.
func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}
