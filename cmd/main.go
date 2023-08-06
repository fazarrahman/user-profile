package main

import (
	"os"

	"github.com/fazarrahman/user-profile/generated"
	"github.com/fazarrahman/user-profile/handler"
	"github.com/fazarrahman/user-profile/repository"
	"github.com/fazarrahman/user-profile/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	var svc service.ServiceInterface = service.NewService(service.NewServiceOptions{Repository: repo})
	return handler.NewServer(svc)
}
