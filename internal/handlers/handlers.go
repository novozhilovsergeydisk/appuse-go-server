package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

// Home - это название функции, которая обрабатывает маршрут /
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFound(w, r)
		return
	}

	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "home", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}

	// Парсинг шаблонов с функциями
	tmpl := template.Must(template.New("index.html").
		Funcs(template.FuncMap{
			"safeHTML": func(s string) template.HTML { return template.HTML(s) },
		}).
		ParseFiles(indexPath, navbarPath, footerPath))

	// Данные для шаблона
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Welcome to Go Server",
		Message: "Hello from Go HTTP Server!",
	}

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Test это название функции, которая обрабатывает маршрут /test
func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Page"))
}

// ServerInstallation это название функции, которая обрабатывает маршрут /server_installation
func ServerInstallation(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "server_installation", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}

	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UsingPsql это название функции, которая обрабатывает маршрут /using_psql
func UsingPsql(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "using_psql", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}

	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ServerConfiguration это название функции, которая обрабатывает маршрут /server_configuration
func ServerConfiguration(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "server_configuration", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}

	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// PostgreSQLQuiz это название функции, которая обрабатывает маршрут /postgres_quiz
func PostgreSQLQuiz(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "postgres_quiz", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "PostgreSQL Quiz template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}
	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ConcurrentAccessManagement это название функции, которая обрабатывает маршрут /concurrent_access_management
func ConcurrentAccessManagement(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "concurrent_access_management", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Concurrent Access Management template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}
	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Vacuum - это название функции, которая обрабатывает маршрут /vacuum
func Vacuum(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "vacuum", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}
	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// BufferCacheLog - это название функции, которая обрабатывает маршрут /buffer_cache_log
func BufferCacheLog(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "buffer_cache_log", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}
	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DatabasesSchemas - это название функции, которая обрабатывает маршрут /databases_schemas
func DatabasesSchemas(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "databases_schemas", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}
	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SystemCatalog - это название функции, которая обрабатывает маршрут /system_catalog
func SystemCatalog(w http.ResponseWriter, r *http.Request) {
	// Полные пути к шаблонам
	indexPath := filepath.Join("templates", "system_catalog", "index.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	footerPath := filepath.Join("templates", "footer.html")

	// Проверка существования файлов
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "Index template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(navbarPath); os.IsNotExist(err) {
		http.Error(w, "Navbar template not found", http.StatusInternalServerError)
		return
	}
	if _, err := os.Stat(footerPath); os.IsNotExist(err) {
		http.Error(w, "Footer template not found", http.StatusInternalServerError)
		return
	}
	// Парсинг шаблонов
	tmpl := template.Must(template.ParseFiles(indexPath, navbarPath, footerPath))

	// Рендеринг
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NotFound обработчик 404 ошибки
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	tmplPath := filepath.Join("templates", "404.html")
	footerPath := filepath.Join("templates", "footer.html")
	navbarPath := filepath.Join("templates", "navbar.html")

	if _, err := os.Stat(tmplPath); err == nil {
		// Проверяем наличие футера и навбара
		footerExists := false
		navbarExists := false

		if _, err := os.Stat(footerPath); err == nil {
			footerExists = true
		}

		if _, err := os.Stat(navbarPath); err == nil {
			navbarExists = true
		}

		if footerExists && navbarExists {
			// Если существуют все шаблоны
			tmpl := template.Must(template.ParseFiles(tmplPath, navbarPath, footerPath))
			tmpl.Execute(w, nil)
			return
		} else {
			// Если есть только 404.html
			tmpl := template.Must(template.ParseFiles(tmplPath))
			tmpl.Execute(w, nil)
			return
		}
	}

	http.Error(w, "404 page not found", http.StatusNotFound)
}
