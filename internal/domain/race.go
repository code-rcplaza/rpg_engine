package domain

// Race represents a playable RPG race.
type Race struct {
	ID   int
	Slug string // "human", "elf", "halfling"
	Name string // "Human", "Elf", "Halfling"
}

// Gender represents the gender used for name generation.
// Enums in Go are defined as a custom type + constants.
type Gender string

const (
	GenderMale    Gender = "m"
	GenderFemale  Gender = "f"
	GenderNeutral Gender = "n" // applies to any gender
)

// NameStyle groups patterns under a naming convention.
// Used when a race has mutually exclusive styles.
// Example: tiefling has "infernal" and "virtue".
type NameStyle struct {
	ID     int
	RaceID int
	Slug   string // "infernal", "virtue", "human", "elven"
}

// NameComponent is a word or fragment that forms part of a name.
type NameComponent struct {
	ID            int
	RaceID        int
	ComponentType string // "first_name", "last_name", "clan_name", "family_name", "nickname", "virtue_name"
	Gender        Gender
	Value         string
}

// NamePattern defines the structure of how a race's name is built.
// Example human:     [first_name (order 1), last_name (order 2)]
// Example half-orc:  [first_name (order 1)]
// Example dragonborn:[clan_name (order 1), first_name (order 2)]
type NamePattern struct {
	ID            int
	RaceID        int
	StyleID       *int   // nil = always applies regardless of style
	Order         int    // position in the final name
	ComponentType string // which component type goes in this position
	Required      bool   // if false, included randomly
	MaxCount      int    // max times this component can appear (gnome nicknames = 3)
}

// CompositePart is a fragment for races with compound last names (halflings).
// "Brush" + "gather" → "Brushgather"
type CompositePart struct {
	ID       int
	RaceID   int
	Position string // "first", "second"
	Category string // "noun", "adjective", "verb"
	Value    string
}