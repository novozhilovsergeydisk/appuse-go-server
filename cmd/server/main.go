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
	mux.HandleFunc("/concurrent_access_management", handlers.ConcurrentAccessManagement)
	mux.HandleFunc("/vacuum", handlers.Vacuum)
	mux.HandleFunc("/buffer_cache_log", handlers.BufferCacheLog)
	mux.HandleFunc("/databases_schemas", handlers.DatabasesSchemas)
	mux.HandleFunc("/system_catalog", handlers.SystemCatalog)
	mux.HandleFunc("/tablespaces", handlers.Tablespaces)
	mux.HandleFunc("/low_level", handlers.LowLevel)
	mux.HandleFunc("/monitoring", handlers.Monitoring)

	// Настройка обработки статических файлов
	staticDir := "static"
	fileServer := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Обработка favicon.ico
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/images/favicon.ico")
	})

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
