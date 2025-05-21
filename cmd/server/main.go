package main

import (
	"http-server/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Создаем мультиплексор
	mux := http.NewServeMux()

	// Регистрируем обработчики
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/test", handlers.Test)
	mux.HandleFunc("/server_configuration", handlers.ServerConfiguration)
	mux.HandleFunc("/using_psql", handlers.UsingPsql)
	mux.HandleFunc("/server_installation", handlers.ServerInstallation)
	mux.HandleFunc("/postgres_quiz", handlers.PostgreSQLQuiz)

	// Настройка обработки статических файлов
	staticDir := "static"
	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Настройка обработки 404
	mux.HandleFunc("/not-found", handlers.NotFound)

	// Конфигурация сервера
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Запуск сервера
	log.Println("🚀 Server started on http://localhost:8080")
	log.Printf("📁 Serving static files from /%s/", staticDir)
	log.Fatal(server.ListenAndServe())
}
