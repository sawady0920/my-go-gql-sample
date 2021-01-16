package main

import (
	"fmt"
	"log"
	"my-go-gql-sample/graph"
	"my-go-gql-sample/graph/generated"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	protocol := os.Getenv("MYSQL_PROTOCOL")
	dbname := os.Getenv("MYSQL_DATABASE")
	dsn := user + ":" + pass + "@" + protocol + "/" + dbname + "?parseTime=true&charset=utf8"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Printf("DB Open Error :%v", err)
		panic(err.Error())
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Printf("DB Close Error :%v", err)
			panic(err.Error())
		}
		sqlDB.Close()
	}()

	fmt.Println(dsn)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: db,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
