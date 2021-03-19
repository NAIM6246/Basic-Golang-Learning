package system

import (
	"Golang/config"
	"Golang/conn"
	"Golang/handler"
	"Golang/repository"
	"Golang/services"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()
	//
	container.Provide(config.NewDBConfig)
	container.Provide(conn.ConnectDB)

	//user
	container.Provide(repository.NewUserRepository)
	container.Provide(services.NewUserService)

	//Article
	container.Provide(repository.NewArticleRepository)
	container.Provide(services.NewArticleService)

	//Handlers
	container.Provide(handler.NewUserHandler)
	container.Provide(handler.NewArticleHandler)

	//server
	container.Provide(NewServer)

	return container
}

//System	:
type System struct {
}

//NewSystem	:
func NewSystem() {
	container := buildContainer()
	err := container.Invoke(func(server *Server) {
		server.run()
	})
	if err != nil {
		panic(err)
	}
}

//Server	:
type Server struct {
	userHandler    handler.IUserHandler
	articleHandler handler.IArticleHandler
	router         *chi.Mux
	dbContext      *conn.DB
}

//Constructor of server	:
func NewServer(
	userHandler handler.IUserHandler,
	articleHandler handler.IArticleHandler,
	dbContext *conn.DB) *Server {

	return &Server{
		userHandler:    userHandler,
		articleHandler: articleHandler,
		router:         chi.NewRouter(),
		dbContext:      dbContext,
	}
}

const port = ":3000"

func (s *Server) run() {
	s.setMiddlewares()
	s.mapHandlers()
	defer s.dispose()
	http.ListenAndServe(port, s.router)
}

func (s *Server) setMiddlewares() {
	s.dbContext.Migration()
}

func (s *Server) mapHandlers() {
	s.router.Route("/articles", s.articleHandler.Handle)
	s.router.Route("/users", s.userHandler.Handle)
}

func (s *Server) dispose() {
	s.dbContext.Close()
}
