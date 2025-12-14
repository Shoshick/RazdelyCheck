package main

import (
	"log"
	"net/http"
	"os"

	"RazdelyCheck/internal/handler"
	"RazdelyCheck/internal/middleware"
	"RazdelyCheck/internal/repoimpl"
	"RazdelyCheck/internal/router"
	"RazdelyCheck/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	// Загружаем переменные окружения из .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	dsn := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSL

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Репозитории
	userRepo := repo_impl.NewUserRepo(db)
	checkResultRepo := repo_impl.NewCheckResultRepo(db)
	groupRepo := repo_impl.NewGroupRepo(db, checkResultRepo)
	itemRepo := repo_impl.NewItemRepo(db)
	checkRepo := repo_impl.NewCheckRepo(db)
	checkSourceRepo := repo_impl.NewCheckSourceRepo(db)
	// новый репо

	// Сервисы
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
	checkResultService := service.NewCheckResultService(checkResultRepo, db) // новый сервис

	// Хендлеры
	userHandler := handler.NewUserHandler(userService)
	groupHandler := handler.NewGroupHandler(groupService)
	checkHandler := handler.NewCheckHandler(checkService)
	itemHandler := handler.NewItemHandler(itemService)
	checkSourceHandler := handler.NewCheckSourceHandler(checkSourceService)
	checkResultHandler := handler.NewCheckResultHandler(checkResultService) // новый хендлер

	// Роутер
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recovery)
	r.Use(middleware.Auth)

	// Подключаем роутеры
	r.Mount("/", router.NewUserRouter(userHandler))
	r.Mount("/", router.NewGroupRouter(groupHandler))
	r.Mount("/", router.NewCheckRouter(checkHandler))
	r.Mount("/", router.NewItemRouter(itemHandler))
	r.Mount("/", router.NewCheckSourceRouter(checkSourceHandler))
	r.Mount("/", router.NewCheckResultRouter(checkResultHandler)) // новый роутер

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started at :%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
