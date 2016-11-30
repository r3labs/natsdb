/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package natsdb

import (
	"encoding/json"
	"log"
)

// ErrorMessage : Struct representing an error message
type ErrorMessage struct {
	Message string `json:"_error"`
	Code    string `json:"_code"`
}

var (
	// NotFound : Error message for not found errors
	NotFound = ErrorMessage{Message: "Not found", Code: "404"}
	// Unexpected : Error message for unexpected errors
	Unexpected = ErrorMessage{Message: "Unexpected error", Code: "500"}
)

func (e *ErrorMessage) encoded() []byte {
	var err error
	str := []byte("")
	if str, err = json.Marshal(e); err != nil {
		log.Println("Couldn't marshal error message")
	}
	return str
}
