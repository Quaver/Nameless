package common

import "fmt"

type Mode int32

const (
	ModeKeys4 Mode = iota + 1
	ModeKeys7

	// ModeKeys1 New game modes so they start counting from 3

	ModeKeys1 = iota + 1
	ModeKeys2
	ModeKeys3
	ModeKeys5
	ModeKeys6
	ModeKeys8
	ModeKeys9
	ModeKeys10
	GameModeEnumMaxValue
)

// GetModeString Returns a string version of game mode
func GetModeString(mode Mode) (string, error) {
	switch mode {
	case ModeKeys1:
		return "keys1", nil
	case ModeKeys2:
		return "keys2", nil
	case ModeKeys3:
		return "keys3", nil
	case ModeKeys4:
		return "keys4", nil
	case ModeKeys5:
		return "keys5", nil
	case ModeKeys6:
		return "keys6", nil
	case ModeKeys7:
		return "keys7", nil
	case ModeKeys8:
		return "keys8", nil
	case ModeKeys9:
		return "keys9", nil
	case ModeKeys10:
		return "keys10", nil
	default:
		return "", fmt.Errorf("%v is not a valid mode", mode)
	}
}
