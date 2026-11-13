# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener el proyecto completo y arrancable. Guarda los archivos que te devuelva en la
carpeta del repo, respetando exactamente las rutas.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Superficie de practica — NO resuelvas

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs.

- `internal/domain/infrastructure/config.go` — La fase "Configuración del Entorno y Creación de Endpoints Básicos" se implementa aqui: este archivo es el ejercicio.
- `internal/domain/pkg/errors/errors.go` — La fase "Implementación de Validaciones y Manejo de Errores" se implementa aqui: este archivo es el ejercicio.
- `internal/domain/pkg/validator/validator.go` — La fase "Implementación de Validaciones y Manejo de Errores" ya venia resuelta por el codigo base: este archivo pasa a ser el ejercicio.

## Como saber que terminaste

```bash
go mod tidy && go build ./...
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Perfil
Chapter Backend, Especialidad Desarrollador, Tecnología Go, Senior L2

### Brecha de conocimiento
Construir una API REST en Go con Gin para gestionar tareas, Implementa endpoints REST en Go con Gin para crear, listar y actualizar tareas

### Misión / candidato
Exponer la gestion de tareas por HTTP

### Reto
- Tema: api-rest-tareas-go-gin
- Seniority: senior-l2
- Tipo: practical
- Título: Desarrollo de API REST para Gestión de Tareas en Go con Gin
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Configuración del Entorno y Creación de Endpoints Básicos — objetivo: Configurar el entorno de desarrollo y crear los endpoints básicos para crear y listar tareas. — entregable (NO resolver): Código fuente con los endpoints básicos para crear y listar tareas.
- Fase 2: Implementación de Validaciones y Manejo de Errores — objetivo: Implementar validaciones adicionales y manejar correctamente los errores en los endpoints. — entregable (NO resolver): Código fuente con validaciones adicionales y manejo de errores en los endpoints.
- Fase 3: Actualización de Tareas y Optimización del Código — objetivo: Implementar un endpoint para actualizar tareas y optimizar el código existente. — entregable (NO resolver): Código fuente con el endpoint para actualizar tareas y el código optimizado.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (entrega de archivos).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Entrega de archivos
Entregá cada archivo con su ruta completa, respetando exactamente la estructura de rutas
indicada por los marcadores: un bloque por archivo, precedido por su ruta. No generes ZIPs
ni archivos adjuntos.
La entrega debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: Entregá los archivos completos al finalizar. No preguntes si el usuario
quiere recibirlos. Simplemente entregalos; No debes desplegar en el chat el resumen de lo que arreglaste, solo entregá los archivos.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del último archivo entregado
No preguntes si el usuario quiere los archivos: simplemente entregalos siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/user/project

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
)

// === ARCHIVO: main.go ===
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure"
)

func main() {
	router := gin.Default()

	repo := infrastructure.NewInMemoryTaskRepository()
	taskService := application.NewTaskService(repo)

	infrastructure.SetupRouter(router, taskService)

	log.Println("Servidor iniciando en el puerto 8080")
	if err := router.Run(); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

// === ARCHIVO: internal/app/routes.go ===
package app

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	// Sample route - will be replaced by domain code
	router.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sample endpoint",
		})
	})
}

// === ARCHIVO: internal/contracts/types.go ===
package contracts

// Sample contract types - interfaces and models go here

// === ARCHIVO: internal/domain/cmd/server/main.go ===
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	taskRepo := infrastructure.NewInMemoryTaskRepository()

	taskService := application.NewTaskService(taskRepo)

	infrastructure.SetupRouter(router, taskService)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Server starting on port %s", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server exited")
}

// === ARCHIVO: internal/domain/task.go ===
package domain

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

var validStatuses = map[TaskStatus]bool{
	TaskStatusPending:    true,
	TaskStatusInProgress: true,
	TaskStatusCompleted:  true,
}

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(id, title, description string, status TaskStatus) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}
	if !validStatuses[status] {
		return nil, errors.New("invalid status")
	}

	now := time.Now()
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (t *Task) UpdateTitle(title string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	t.Title = title
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}
	t.Description = description
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateStatus(status TaskStatus) error {
	if !validStatuses[status] {
		return errors.New("invalid status")
	}
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}

// === ARCHIVO: internal/domain/application/task_service.go ===
package application

import (
	"context"
	"errors"

	"github.com/user/project/internal/domain"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	FindByID(ctx context.Context, id string) (*domain.Task, error)
	FindAll(ctx context.Context) ([]*domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string, status domain.TaskStatus) (*domain.Task, error) {
	task, err := domain.NewTask("", title, description, status)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, task); err != nil {
		return nil, errors.New("failed to save task")
	}

	return task, nil
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, errors.New("failed to list tasks")
	}
	return tasks, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id, title, description string, status domain.TaskStatus) (*domain.Task, error) {
	task, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("task not found")
	}

	if title != "" {
		if err := task.UpdateTitle(title); err != nil {
			return nil, err
		}
	}

	if description != "" {
		if err := task.UpdateDescription(description); err != nil {
			return nil, err
		}
	}

	if status != "" {
		if err := task.UpdateStatus(status); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Update(ctx, task); err != nil {
		return nil, errors.New("failed to update task")
	}

	return task, nil
}

// === ARCHIVO: internal/domain/infrastructure/task_repository.go ===
package infrastructure

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/user/project/internal/domain"
)

type InMemoryTaskRepository struct {
	tasks map[string]*domain.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

func (r *InMemoryTaskRepository) Save(ctx context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if task.ID == "" {
		task.ID = generateID()
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (r *InMemoryTaskRepository) FindAll(ctx context.Context) ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *InMemoryTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}
	r.tasks[task.ID] = task
	return nil
}

func generateID() string {
	return "task-" + string(time.Now().UnixNano())
}

// === ARCHIVO: internal/domain/infrastructure/router.go ===
package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/application"
)

func SetupRouter(router *gin.Engine, taskService *application.TaskService) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("/", createTaskHandler(taskService))
		taskGroup.GET("/", listTasksHandler(taskService))
		taskGroup.GET("/:id", getTaskHandler(taskService))
		taskGroup.PUT("/:id", updateTaskHandler(taskService))
	}
}

func createTaskHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title       string `json:"title" binding:"required"`
			Description string `json:"description" binding:"required"`
			Status      string `json:"status" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		status := domain.TaskStatus(req.Status)
		task, err := taskService.CreateTask(c.Request.Context(), req.Title, req.Description, status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, task)
	}
}

func listTasksHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := taskService.ListTasks(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func getTaskHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		task, err := taskService.GetTask(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	}
}

func updateTaskHandler(taskService *application.TaskService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var req struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Status      string `json:"status"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		status := domain.TaskStatus(req.Status)
		task, err := taskService.UpdateTask(c.Request.Context(), id, req.Title, req.Description, status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

// === ARCHIVO: internal/domain/infrastructure/config.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package infrastructure

// TODO(reto): Implementar la carga de configuración desde variables de entorno
// o un archivo de configuración. La configuración debe incluir al menos:
// - Puerto de escucha del servidor
// - Configuración de la base de datos (si se implementa persistencia real)
//
// Requisitos:
// 1. Usar el paquete `os` para leer variables de entorno
// 2. Proporcionar valores por defecto razonables
// 3. Validar que los valores requeridos estén presentes
// 4. Manejar errores de configuración adecuadamente

type Config struct {
	Port string
	// TODO(reto): Añadir otros campos de configuración según sea necesario
}

func LoadConfig() (*Config, error) {
  panic("TODO(reto): LoadConfig")
}

// === ARCHIVO: internal/domain/pkg/validator/validator.go ===
package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/user/project/internal/domain"
)

var validate *validator.Validate

func init() {
	panic("TODO(reto): import")
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResult struct {
	Valid  bool             `json:"valid"`
	Errors []ValidationError `json:"errors,omitempty"`
}

func ValidateTask(title, description string, status domain.TaskStatus) ValidationResult {
	panic("TODO(reto): ValidateTask")
}

func isValidStatus(status domain.TaskStatus) bool {
	panic("TODO(reto): isValidStatus")
}

func ValidateTaskInput(title, description string, status domain.TaskStatus) error {
	panic("TODO(reto): ValidateTaskInput")
}

// === ARCHIVO: internal/domain/pkg/errors/errors.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package errors

// TODO(reto): Implementar manejo de errores personalizados para la aplicación.
// Se deben definir errores específicos del dominio con códigos HTTP asociados.

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (e *AppError) Error() string {
  panic("TODO(reto): Error")
}

func NewAppError(code, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}
```
