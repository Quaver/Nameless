package common

import "fmt"

type Mode int32

const (
	ModeKeys4 Mode = iota + 1
	ModeKeys7
)

// GetModeString Returns a string version of game mode
func GetModeString(mode Mode) (string, error) {
	switch mode {
	case ModeKeys4:
		return "keys4", nil
	case ModeKeys7:
		return "keys7", nil
	default:
		return "", fmt.Errorf("%v is not a valid mode", mode)
	}
}