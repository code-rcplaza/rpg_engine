package domain

import "fmt"

// NotFoundError es un error tipado para cuando algo no existe en la DB.
// Tener errores propios permite al caller hacer distinción:
//
//	if errors.As(err, &domain.NotFoundError{}) { ... }
type NotFoundError struct {
	Entity string // "raza", "componente"
	ID     any    // el identificador que no se encontró
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s no encontrado: %v", e.Entity, e.ID)
}