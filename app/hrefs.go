//************************************************************************//
// API "mara": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/codeclysm/mara-api/design
// --out=$(GOPATH)/src/github.com/codeclysm/mara-api
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// CalendarHref returns the resource href.
func CalendarHref(id interface{}) string {
	return fmt.Sprintf("/appointments/%v", id)
}
