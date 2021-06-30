package common

type Mods int64

const (
	ModNoSliderVelocities Mods = 1 << iota
	ModSpeed05X
	ModSpeed06X
	ModSpeed07X
	ModSpeed08X
	ModSpeed09X
	ModSpeed11X
	ModSpeed12X
	ModSpeed13X
	ModSpeed14X
	ModSpeed15X
	ModSpeed16X
	ModSpeed17X
	ModSpeed18X
	ModSpeed19X
	ModSpeed20X
	ModStrict
	ModChill
	ModNoPause
	ModAutoplay
	ModPaused
	ModNoFail
	ModNoLongNotes
	ModRandomize
	ModSpeed055X
	ModSpeed065X
	ModSpeed075X
	ModSpeed085X
	ModSpeed095X
	ModInverse
	ModFullLN
	ModMirror
	ModCoop
	ModSpeed105X
	ModSpeed115X
	ModSpeed125X
	ModSpeed135X
	ModSpeed145X
	ModSpeed155X
	ModSpeed165X
	ModSpeed175X
	ModSpeed185X
	ModSpeed195X
	ModHealthAdjust
	ModEnumMaxValue // This is only in place for looping purposes (i < ModEnumMaxValue - 1; i++)
)

var RankedMods = []Mods{
	ModSpeed05X,
	ModSpeed055X,
	ModSpeed06X,
	ModSpeed065X,
	ModSpeed07X,
	ModSpeed075X,
	ModSpeed08X,
	ModSpeed085X,
	ModSpeed09X,
	ModSpeed095X,
	ModSpeed105X,
	ModSpeed11X,
	ModSpeed115X,
	ModSpeed12X,
	ModSpeed125X,
	ModSpeed13X,
	ModSpeed135X,
	ModSpeed14X,
	ModSpeed145X,
	ModSpeed15X,
	ModSpeed155X,
	ModSpeed16X,
	ModSpeed165X,
	ModSpeed17X,
	ModSpeed175X,
	ModSpeed18X,
	ModSpeed185X,
	ModSpeed19X,
	ModSpeed195X,
	ModSpeed20X,
	ModMirror,
}

// IsModActivated Returns if a given mod is activated in a mod combo
func IsModActivated(modCombo Mods, mod Mods) bool {
	return modCombo&mod != 0
}

// IsModRanked Returns if a particular mod is ranked
func IsModRanked(mod Mods) bool {
	if mod == 0 {
		return true
	}

	for _, rankedMod := range RankedMods {
		if rankedMod == mod {
			return true
		}
	}

	return false
}

// IsModComboRanked Returns if a combination of mods is ranked
func IsModComboRanked(modCombo Mods) bool {
	if modCombo == 0 {
		return true
	}

	// Loops over every mod in the enum and checks if the mod is active & ranked or not
	for i := 0; (1 << i) < ModEnumMaxValue-1; i++ {
		mod := Mods(1 << i)

		if !IsModActivated(modCombo, mod) {
			continue
		}

		if !IsModRanked(Mods(1 << i)) {
			return false
		}
	}

	return true
}
