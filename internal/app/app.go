package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type AppConfig struct {
	Port int
}

func ConfigWithPort(port int) AppConfig {
	return AppConfig{
		Port: port,
	}
}

type App struct {
	port      int
	Mux *http.ServeMux
	DbConn    *pgx.Conn
}

func New(config AppConfig) App {
	return App{
		port:      config.Port,
		Mux: http.NewServeMux(),
	}
}

func (a *App) ConnectDB(con string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

    //TODO: make pool of connection
	conn, err := pgx.Connect(ctx, con)

	if err != nil {
		return err
	}

	a.DbConn = conn

	return nil
}

func (a *App) CloseDB() {
    ctx, cancel := context.WithTimeout(context.Background(),30 * time.Second)
    defer cancel()

    a.DbConn.Close(ctx)
}

func (a *App) Start() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", a.port),
		a.initRouter(),
	)
}

