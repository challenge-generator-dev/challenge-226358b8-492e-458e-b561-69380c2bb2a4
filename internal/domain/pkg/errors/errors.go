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