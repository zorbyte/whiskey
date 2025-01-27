package args

import "strings"

// Argument for commands.
type Argument struct {
	ID       string
	name     string
	required bool
	greedy   bool
	keyed    bool
	info     string

	parser Parser
}

// New creates an argument.
func New(ID string, parser Parser, info ...string) *Argument {
	name, required, greedy, keyed, err := parseID(ID)
	if err != nil {
		panic(err)
	}

	finalInfo := getInfo(info)
	userFriendlyID := makeUserFriendlyID(ID, finalInfo)

	return &Argument{userFriendlyID, name, required, greedy, keyed, finalInfo, parser}
}

func getInfo(info []string) (finalInfo string) {
	switch len(info) {
	case 0:
		break
	case 1:
		finalInfo = info[0]
	default:
		finalInfo = strings.Join(info, ", ")
	}

	return
}

func makeUserFriendlyID(originalID, finalInfo string) (ID string) {
	if finalInfo != "" {
		ID = string(originalID[:len(originalID)-2]) + "(" + finalInfo + ")" + originalID[len(originalID)-1:]
	} else {
		ID = originalID
	}

	return
}
