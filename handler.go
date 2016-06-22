/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package natsdb

import (
	"encoding/json"

	"github.com/nats-io/nats"
)

var err error

// Handler : this struct manages all nats connections and maps them
// to the injected entity
type Handler struct {
	Nats                   *nats.Conn
	NewModel               func() Model
	NotFoundErrorMessage   []byte
	UnexpectedErrorMessage []byte
	DeletedMessage         []byte
}

// Find : Based on the fields of the input json it will
// search and return any entity matching these fields
func (h *Handler) Find(msg *nats.Msg) {
	e := h.NewModel()
	e.MapInput(msg.Data)
	entities := e.Find()
	body, err := json.Marshal(&entities)

	if err != nil {
		h.Nats.Publish(msg.Reply, h.UnexpectedErrorMessage)
		return
	}
	h.Nats.Publish(msg.Reply, body)
}

// Get : Based on a json input with an id field will
// return the client details for this id
func (h *Handler) Get(msg *nats.Msg) {
	e := h.NewModel()
	if ok := e.LoadFromInputOrFail(msg, h); ok {
		body, err := json.Marshal(e)
		if err != nil {
			h.Fail(msg)
		}
		h.Nats.Publish(msg.Reply, body)
	}
}

// Del : Based on a json input with an id field will
// delete the client details for this id
func (h *Handler) Del(msg *nats.Msg) {
	e := h.NewModel()
	if ok := e.LoadFromInputOrFail(msg, h); ok {
		e.Delete()
		h.Nats.Publish(msg.Reply, h.DeletedMessage)
	}
}

// Set : Based on a json input with an id field will
// update the client details for this id with all extra fields
// defined on the message
func (h *Handler) Set(msg *nats.Msg) {
	e := h.NewModel()
	if ok := e.LoadFromInput(msg.Data); ok {
		e.Update(msg.Data)
		body, err := json.Marshal(e)
		if err != nil {
			h.Fail(msg)
		}
		h.Nats.Publish(msg.Reply, body)
	} else {
		input := h.NewModel()
		input.MapInput(msg.Data)
		if input.HasID() == false {
			if err = e.Save(); err != nil {
				h.Nats.Publish(msg.Reply, []byte(err.Error()))
			} else {
				body, err := json.Marshal(e)
				if err != nil {
					h.Fail(msg)
				}
				h.Nats.Publish(msg.Reply, body)
			}
		} else {
			h.Nats.Publish(msg.Reply, h.NotFoundErrorMessage)
		}

	}
}

// Fail : It replies the nats message with a not found error
func (h *Handler) Fail(msg *nats.Msg) {
	h.Nats.Publish(msg.Reply, h.NotFoundErrorMessage)
}
