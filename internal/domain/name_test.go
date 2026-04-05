package domain

import (
	"reflect"
	"testing"
)

func TestGeneratedName(t *testing.T) {
	race := Race{
		ID:   1,
		Slug: "human",
		Name: "Human",
	}

	parts := []string{"Fosco", "Olla Caliente"}
	full := "Fosco Olla Caliente"

	gn := GeneratedName{
		Full:  full,
		Parts: parts,
		Race:  race,
	}

	if gn.Full != full {
		t.Errorf("expected Full to be %s, got %s", full, gn.Full)
	}

	if !reflect.DeepEqual(gn.Parts, parts) {
		t.Errorf("expected Parts to be %v, got %v", parts, gn.Parts)
	}

	if gn.Race != race {
		t.Errorf("expected Race to be %v, got %v", race, gn.Race)
	}
}

func TestGenderConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      Gender
		expected string
	}{
		{"GenderMale", GenderMale, "m"},
		{"GenderFemale", GenderFemale, "f"},
		{"GenderNeutral", GenderNeutral, "n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.got) != tt.expected {
				t.Errorf("expected %s to be %q, got %q", tt.name, tt.expected, string(tt.got))
			}
		})
	}
}
