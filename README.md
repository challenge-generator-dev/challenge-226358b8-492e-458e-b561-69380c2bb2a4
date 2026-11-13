# Desarrollo de API REST para Gestión de Tareas en Go con Gin

La empresa PragmaFintech necesita una API REST para gestionar tareas en su plataforma de gestión de proyectos. La API debe permitir crear, listar y actualizar tareas. Cada tarea tiene un identificador único, un título, una descripción y un estado (pendiente, en progreso, completada). El sistema debe validar que el título y la descripción no estén vacíos y que el estado sea uno de los valores permitidos. En caso de error en la validación, la API debe retornar un mensaje descriptivo y un código de estado HTTP apropiado.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | api-rest-tareas-go-gin |
| **Nivel** | senior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y entregará cada archivo con su ruta.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y entregará cada archivo con su ruta. Guárdalos en la carpeta donde quieras trabajar, respetando las rutas.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no entrega todos los archivos al terminar el análisis, escríbele: "entrégame todos los archivos con su ruta".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Configuración del Entorno y Creación de Endpoints Básicos

**Objetivo:** Configurar el entorno de desarrollo y crear los endpoints básicos para crear y listar tareas.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Configura el entorno de desarrollo con Go y Gin.
- Crea un endpoint POST para crear una nueva tarea. La tarea debe tener un título, una descripción y un estado inicial de 'pendiente'.
- Crea un endpoint GET para listar todas las tareas existentes.

**Entregable:** Código fuente con los endpoints básicos para crear y listar tareas.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda validar los campos obligatorios en la creación de tareas.
- Utiliza el marco de trabajo del lenguaje principal para manejar las rutas y los métodos HTTP.

</details>

### Fase 2: Implementación de Validaciones y Manejo de Errores

**Objetivo:** Implementar validaciones adicionales y manejar correctamente los errores en los endpoints.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Agrega validaciones para asegurar que el título y la descripción no estén vacíos.
- Agrega validaciones para asegurar que el estado de la tarea sea uno de los valores permitidos ('pendiente', 'en progreso', 'completada').
- Maneja los errores de validación devolviendo un mensaje descriptivo y un código de estado HTTP apropiado.

**Entregable:** Código fuente con validaciones adicionales y manejo de errores en los endpoints.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que los errores de validación deben ser manejados de manera consistente y devolver información útil al cliente.
- Utiliza el marco de trabajo del lenguaje principal para manejar los errores y las respuestas HTTP.

</details>

### Fase 3: Actualización de Tareas y Optimización del Código

**Objetivo:** Implementar un endpoint para actualizar tareas y optimizar el código existente.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Crea un endpoint PUT para actualizar una tarea existente. La tarea debe poder cambiar su título, descripción y estado.
- Optimiza el código existente para asegurar que sea eficiente y fácil de mantener.
- Asegura que las validaciones y el manejo de errores funcionen correctamente para el nuevo endpoint.

**Entregable:** Código fuente con el endpoint para actualizar tareas y el código optimizado.

<details>
<summary>Pistas de conocimiento</summary>

- Recuerda que el endpoint PUT debe validar los campos actualizados de la misma manera que el endpoint POST.
- Utiliza buenas prácticas de programación para asegurar que el código sea eficiente y fácil de mantener.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una tarea en el contexto de la API REST y cuáles son sus atributos?
- **paraQueSirve**: ¿Para qué sirve cada endpoint en la API REST y cómo se relacionan con las operaciones de gestión de tareas?
- **comoSeUsa**: ¿Cómo se usa la API REST para crear, listar y actualizar tareas?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir al usar la API REST y cómo se manejan?
- **queDecisionesImplica**: ¿Qué decisiones implica la implementación de la API REST en términos de validaciones, manejo de errores y optimización del código?

## Criterios de Evaluacion

- Configuración correcta del entorno de desarrollo con Go y Gin.
- Creación de endpoints básicos para crear y listar tareas.
- Implementación de validaciones adicionales y manejo de errores en los endpoints.
- Creación de un endpoint para actualizar tareas y optimización del código existente.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT) si no tenes un agente en el IDE. Declara lo mismo que AGENTS.md; los archivos que te devuelva el chat los guardas en esta carpeta respetando las rutas.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
go mod tidy && go build ./...
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
