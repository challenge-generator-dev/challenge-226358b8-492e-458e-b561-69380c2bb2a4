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