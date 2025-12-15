package main

import (
	"log"
	"net/http"
	"os"

	"RazdelyCheck/internal/handler"
	"RazdelyCheck/internal/middleware"
	repoimpl "RazdelyCheck/internal/repoimpl"
	"RazdelyCheck/internal/router"
	"RazdelyCheck/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dsn := "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") +
		"?sslmode=" + os.Getenv("DB_SSLMODE")

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	userRepo := repoimpl.NewUserRepo(db)
	checkResultRepo := repoimpl.NewCheckResultRepo(db)
	groupRepo := repoimpl.NewGroupRepo(db, checkResultRepo)
	itemRepo := repoimpl.NewItemRepo(db)
	checkRepo := repoimpl.NewCheckRepo(db)
	checkSourceRepo := repoimpl.NewCheckSourceRepo(db)

	userService := service.NewUserService(userRepo)
	groupService := service.NewGroupService(groupRepo, userRepo)
	checkService := service.NewCheckService(checkRepo, groupRepo)
	itemService := service.NewItemService(itemRepo, checkRepo, db)
	checkSourceService := service.NewCheckSourceService(
		checkSourceRepo,
		db,
		checkService,
		os.Getenv("CHECK_API_TOKEN"),
	)
	checkResultService := service.NewCheckResultService(checkResultRepo, db)

	userHandler := handler.NewUserHandler(userService)
	groupHandler := handler.NewGroupHandler(groupService)
	checkHandler := handler.NewCheckHandler(checkService)
	itemHandler := handler.NewItemHandler(itemService)
	checkSourceHandler := handler.NewCheckSourceHandler(checkSourceService)
	checkResultHandler := handler.NewCheckResultHandler(checkResultService)

	r := chi.NewRouter()

	r.Use(middleware.Logging)
	r.Use(middleware.Recovery)
	r.Use(middleware.Auth)

	router.NewUserRouter(r, userHandler)
	router.NewGroupRouter(r, groupHandler)
	router.NewCheckRouter(r, checkHandler)
	router.NewItemRouter(r, itemHandler)
	router.NewCheckSourceRouter(r, checkSourceHandler)
	router.NewCheckResultRouter(r, checkResultHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
