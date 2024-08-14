// Package twofer determines what to say in a cookie offer statement from a bakery
package twofer

import (
	"fmt"
)

// ShareWith offers two cookies for the price one to your specified name or to a generic you.
func ShareWith(name string) string {
	var offer string = "One for you, one for me."
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	} else {
		return offer
	}
	
}
