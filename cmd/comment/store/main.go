package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/beeblogit/app_go_interaction/internal/comment"
	"github.com/beeblogit/app_go_interaction/pkg/bootstrap"
	"github.com/beeblogit/app_go_interaction/pkg/handler"
	"gorm.io/gorm"
	"os"
	//"context"
	"github.com/go-kit/kit/transport/awslambda"
)

var db *gorm.DB
var h *awslambda.Handler

func init() {

	var err error

	//_ = godotenv.Load()
	l := bootstrap.InitLogger()

	db, _ = bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	pagLimDef := os.Getenv("PAGINATOR_LIMIT_DEFAULT")
	if pagLimDef == "" {
		l.Fatal("paginator limit default is required")
	}

	//ctx := context.Background()
	repo := comment.NewRepo(db, l)
	service := comment.NewService(l, repo)

	endpoint := comment.MakeEndpoints(service, comment.Config{LimPageDef: pagLimDef})
	h = handler.NewLambdaCommentStore(endpoint)

}

func main() {
	lambda.Start(h)
}
