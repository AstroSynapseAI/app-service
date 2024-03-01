package app

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/AstroSynapseAI/app-service/controllers"
	"github.com/AstroSynapseAI/app-service/models"
	"github.com/AstroSynapseAI/app-service/repositories"
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
	"github.com/AstroSynapseAI/app-service/sdk/crud/orms/gorm"
	"github.com/AstroSynapseAI/app-service/sdk/rest"
	"github.com/gorilla/handlers"
)

type Routes struct {
	DB *database.Database
}

var _ rest.Routes = (*Routes)(nil)

func NewRoutes(db *database.Database) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes *Routes) LoadRoutes(router *rest.Rest) {
	// CUSTOM CONTROLLERS
	api := router.API()

	// Route custom AvatarsController
	api.Route("/avatars").Controller(controllers.NewAvatarsController(routes.DB))
	// Route custom UsersController
	api.Route("/users").Controller(controllers.NewUsersController(routes.DB))
	// Route custom AgentsController
	api.Route("/agents").Controller(controllers.NewAgentsController(routes.DB))
	// Route custom DocumentsController
	api.Route("/documents").Controller(controllers.NewDocumentsController(routes.DB))
	// Route custom ToolsController
	api.Route("/tools").Controller(controllers.NewToolsController(routes.DB))
	// Route custom PluginsController
	api.Route("/plugins").Controller(controllers.NewPluginsController(routes.DB))
	// Route custom LLMSController
	api.Route("/llms").Controller(controllers.NewLLMSController(routes.DB))

	// CRUD CONTROLLERS
	// Routing CRUD Roles controller
	rolesCtrl := rest.NewCRUDController[models.Role](
		gorm.NewRepository[models.Role](routes.DB, models.Role{}),
	)
	api.Route("/roles").Controller(rolesCtrl)

	// Routing CRUD Accounts controller
	accountsCtrl := rest.NewCRUDController[models.Account](
		gorm.NewRepository[models.Account](routes.DB, models.Account{}),
	)
	api.Route("/accounts").Controller(accountsCtrl)
}

func (routes *Routes) LoadMiddlewares(router *rest.Rest) {
	// TMP API Auth Middleware
	router.Mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			publicRoutes := []string{
				"/api/users/login",
				"/api/users/register",
				"/api/users/register/invite",
				"/api/users/password_recovery",
			}

			for _, validRoute := range publicRoutes {
				if r.URL.Path == validRoute {
					next.ServeHTTP(w, r)
					return
				}
			}

			if strings.HasPrefix(r.URL.Path, "/api") {
				authHeader := r.Header.Get("Authorization")
				if authHeader != "" {
					headerParts := strings.Split(authHeader, " ")
					if len(headerParts) == 2 && strings.EqualFold(headerParts[0], "Bearer") {
						tokenString := headerParts[1]
						usersRepo := repositories.NewUsersRepository(routes.DB)
						_, err := usersRepo.GetByAPIToken(tokenString)
						if err != nil {
							w.WriteHeader(http.StatusUnauthorized)
							return
						}
					}
				}
			}

			next.ServeHTTP(w, r)
		})
	})

	if os.Getenv("ENVIRONMENT") == "LOCAL DEV" {
		// CORS Middleware
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"http://localhost:5173"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

		cors := handlers.CORS(headersOk, originsOk, methodsOk)
		router.Mux.Use(cors)

		// LOGGING MIDDLEWARE
		router.Mux.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasPrefix(r.URL.Path, "/api") {
					log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

					// Logging Request Headers
					// for name, headers := range r.Header {
					// 	for _, h := range headers {
					// 		log.Printf("%v: %v\n", name, h)
					// 	}
					// }
					if r.Method == "POST" || r.Method == "PUT" {
						bodyBytes, err := io.ReadAll(r.Body)
						if err != nil {
							log.Printf("Error reading body: %v", err)
							http.Error(w, "can't read body", http.StatusBadRequest)
							return
						}

						// After reading the body, you need to replace it for further handlers
						r.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))

						log.Printf("Body: %s\n", string(bodyBytes))
					}
				}
				next.ServeHTTP(w, r)
			})
		})
	}

}
