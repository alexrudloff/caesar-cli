package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON prints v as indented JSON to stdout.
func JSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding output: %v\n", err)
	}
}

// Error prints an error message to stderr and exits.
func Error(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
	os.Exit(1)
}
