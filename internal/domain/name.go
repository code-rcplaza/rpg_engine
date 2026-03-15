package domain

// GeneratedName es el resultado final de una generación.
// Separa las partes para que el frontend pueda mostrarlas como quiera.
type GeneratedName struct {
	Full  string   // nombre completo: "Fosco Olla Caliente"
	Parts []string // partes separadas: ["Fosco", "Olla Caliente"]
	Race  Race
}