package common

import "testing"

func TestRankedModCombo(t *testing.T) {
	mods := ModSpeed125X | ModMirror

	ranked := IsModComboRanked(mods)

	if ranked {
		return
	}

	t.Fatalf("Expected ranked mod combination for %v", mods)
}

func TestUnrankedModCombo(t *testing.T) {
	mods := ModSpeed05X | ModNoLongNotes

	ranked := IsModComboRanked(mods)

	if !ranked {
		return
	}

	t.Fatalf("Expected unranked mod combination for %v", mods)
}
