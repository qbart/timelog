package cli

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

// AreYouSure displays user prompt for confirmation.
func AreYouSure(msg string, yes func(), no func()) {
	red := color.New(color.FgRed)
	red.Print(msg, " y/N: ")

	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)
	if len(s) == 1 && (s[0] == 'y' || s[0] == 'Y') {
		yes()
	} else {
		no()
	}
}
