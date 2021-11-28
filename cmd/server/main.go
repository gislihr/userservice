package main

import (
	"fmt"
	"net"

	userserviceGrpc "github.com/gislihr/userservice/grpc"
	"github.com/gislihr/userservice/postgres"
	"github.com/gislihr/userservice/proto"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	Port       int    `envconfig:"PORT" required:"true"`
	PgHost     string `envconfig:"PG_HOST" required:"true"`
	PgUser     string `envconfig:"PG_USER" required:"true"`
	PgPassword string `envconfig:"PG_PASSWORD" required:"true"`
	PgDatabase string `envconfig:"PG_DATABASE" required:"true"`
	PgSSLMode  string `envconfig:"PG_SSL_MODE" required:"true"`
}

func getDatabase(conf config) (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=%s",
			conf.PgHost,
			conf.PgUser,
			conf.PgPassword,
			conf.PgDatabase,
			conf.PgSSLMode),
	)
}

func main() {
	var conf config
	envconfig.MustProcess("", &conf)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Port))
	if err != nil {
		log.WithError(err).Fatal("failed to listen")
	}

	db, err := getDatabase(conf)
	if err != nil {
		log.WithError(err).Fatal("error connecting to database")
	}

	postgres.NewStore(db)
	grpcServer := grpc.NewServer()
	s := userserviceGrpc.New(nil)
	proto.RegisterUserServiceServer(grpcServer, s)

	log.WithField("port", conf.Port).Info("Listening...")
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
