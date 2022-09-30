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

func TestIncompatibleModifiers(t *testing.T) {
	if !HasIncompatibleModifiers(ModSpeed055X | ModSpeed14X) {
		t.Fatal("Expected 0.55x and 1.4x to be incompatible")
	}

	if !HasIncompatibleModifiers(ModFullLN | ModInverse) {
		t.Fatal("Expected FLN and INV to be incompatible")
	}
	
	if HasIncompatibleModifiers(ModSpeed12X | ModMirror) {
		t.Fatal("Expected 1.2x and Mirror to be compatible")
	}
	
	if HasIncompatibleModifiers(ModFullLN) {
		t.Fatal("Expected FLN to be compatible by itself")
	}
}
