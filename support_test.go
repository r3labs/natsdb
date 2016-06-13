/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package natsdb

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

// Entity : a mocked model to b used on testing
type Entity struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// Find : mocked find method for testing
func (e *Entity) Find() []interface{} {
	strings := []string{}
	list := make([]interface{}, len(strings))
	for i, s := range strings {
		list[i] = s
	}

	return list
}

// MapInput : mocked MapInput method for testing
func (e *Entity) MapInput(body []byte) {
	json.Unmarshal(body, &e)
}

// HasID : mocked HasID method for testing
func (e *Entity) HasID() bool {
	return e.ID != 0
}

// LoadFromInput : mocked LoadFromInput method for testing
func (e *Entity) LoadFromInput(msg []byte) bool {
	return true
}

// LoadFromInputOrFail : mocked LoadFromInputOrFail method for testing
func (e *Entity) LoadFromInputOrFail(msg *nats.Msg, h *Handler) bool {
	return true
}

// Update : mocked Update method for testing
func (e *Entity) Update(body []byte) {
	stored := Entity{
		ID:   22,
		Name: "UPDATED",
	}
	e = &stored
}

// Delete : mocked Delete method for testing
func (e *Entity) Delete() {
}

// Save : mocked Save method for testing
func (e *Entity) Save() {
}
