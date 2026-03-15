package domain

// Race representa una raza jugable del RPG.
// Los campos en mayúscula son públicos (exportados).
type Race struct {
	ID   int
	Slug string // "humano", "elfo", "mediano"
	Name string // "Humano", "Elfo", "Mediano"
}

// Gender representa el género para la generación de nombres.
// En Go los "enums" se hacen con un tipo propio + constantes.
type Gender string

const (
	GenderMale    Gender = "m"
	GenderFemale  Gender = "f"
	GenderNeutral Gender = "n" // sin distinción de género
)

// NameComponent es una palabra o fragmento que forma parte de un nombre.
// Por ejemplo: "Juan" (nombre_pila humano masculino)
type NameComponent struct {
	ID            int
	RaceID        int
	ComponentType string // "nombre_pila", "apellido", "clan", "apodo_compuesto"
	Gender        Gender // puede ser GenderNeutral si aplica a cualquiera
	Value         string // el valor real: "Juan", "García", "Olla"
}

// NamePattern define la estructura de cómo se arma el nombre de una raza.
// Ejemplo humano: [nombre_pila (orden 1), apellido (orden 2)]
// Ejemplo semi-orco: [nombre_pila (orden 1)]
type NamePattern struct {
	ID            int
	RaceID        int
	Order         int    // posición en el nombre final
	ComponentType string // qué tipo de componente va en esta posición
	Required      bool   // si es false, se incluye aleatoriamente
}

// CompositePart es un fragmento para razas con nombres compuestos.
// Exclusivo para razas como el Mediano: "Olla" + "Caliente"
type CompositePart struct {
	ID       int
	RaceID   int
	Position string // "primera", "segunda"
	Category string // "sustantivo", "adjetivo"
	Value    string // "Olla", "Barril", "Caliente", "Viejo"
}