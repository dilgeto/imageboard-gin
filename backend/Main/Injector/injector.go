package injector

import {
	"context"
	"log"

	category "github.com/dilgeto/imageboard-gin/backend/Internal/Category"
	board "github.com/dilgeto/imageboard-gin/backend/Internal/Board"
	thread "github.com/dilgeto/imageboard-gin/backend/Internal/Thread"
	reply "github.com/dilgeto/imageboard-gin/backend/Internal/Reply"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
}

func InjectDependencies(rout *gin.Engine) {
	db, err := ConnectPostgreSQL("postgres", "postgres", "localhost", "5432", "imageboard")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	categoryRepository := &category.Repository{DB: db}
	categoryService := &category.Service{Repository: categoryRepository}
	categoryController := &category.Controller{Service: categoryService}
	categoryController.LinkPaths(rout)

	boardRepository := &board.Repository{DB: db}
	boardService := &board.Service{Repository: boardRepository}
	boardController := &board.Controller{Service: boardService}
	boardController.LinkPaths(rout)

	threadRepository := &thread.Repository{DB: db}
	threadService := &thread.Service{Repository: threadRepository}
	threadController := &thread.Controller{Service: threadService}
	threadController.LinkPaths(rout)

	replyRepository := &reply.Repository{DB: db}
	replyService := &reply.Service{Repository: replyRepository}
	replyController := &reply.Controller{Service: replyService}
	replyController.LinkPaths(rout)
}

func ConnectPostgreSQL(user string, pass string, host string, port string, dbname string) (*pgx.Conn, err) {
	dataSource := "postgres://" + user + ":" + pass + "@" + host + ":" + port + "/imageboard?sslmode=disable"

	db, err := pgx.Connect(context.Background(), dataSource)
	return db, err
}