package shared

import (
	"encoding/json"
	"fmt"
)

// MustShowJSON attempts to marshal the given value and panics if an error
// occurs.
func MustShowJSON(e interface{}) {
	bin, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bin))
}
