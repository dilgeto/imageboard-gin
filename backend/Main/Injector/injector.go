package Injector

import (
	"context"
	"log"
	"os"

	board "github.com/dilgeto/imageboard-gin/backend/Internal/Board"
	category "github.com/dilgeto/imageboard-gin/backend/Internal/Category"
	reply "github.com/dilgeto/imageboard-gin/backend/Internal/Reply"
	thread "github.com/dilgeto/imageboard-gin/backend/Internal/Thread"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func InjectDependencies(rout *gin.Engine) {
	errData := godotenv.Load()
	if errData != nil {
		log.Fatalf("error while loading .env: - %v", errData)
	}

	db, err := ConnectPostgreSQL(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("imageboard"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	categoryRepository := &category.Repository{DB: db}
	categoryService := &category.Service{Repository: categoryRepository}
	categoryController := category.Controller{Service: categoryService}
	categoryController.LinkPaths(rout)

	boardRepository := &board.Repository{DB: db}
	boardService := &board.Service{Repository: boardRepository}
	boardController := board.Controller{Service: boardService}
	boardController.LinkPaths(rout)

	threadRepository := &thread.Repository{DB: db}
	threadService := &thread.Service{Repository: threadRepository}
	threadController := thread.Controller{Service: threadService}
	threadController.LinkPaths(rout)

	replyRepository := &reply.Repository{DB: db}
	replyService := &reply.Service{Repository: replyRepository}
	replyController := reply.Controller{Service: replyService}
	replyController.LinkPaths(rout)
}

func ConnectPostgreSQL(user string, pass string, host string, port string, dbname string) (*pgx.Conn, error) {
	dataSource := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/imageboard?sslmode=disable"

	db, err := pgx.Connect(context.Background(), dataSource)
	return db, err
}
