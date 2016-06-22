/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package natsdb

import (
	"github.com/nats-io/nats"
)

// Model : This is the interface you need to implement in order
// to inject a db manager to the Handler
type Model interface {
	MapInput(body []byte)
	LoadFromInput(msg []byte) bool
	LoadFromInputOrFail(msg *nats.Msg, h *Handler) bool
	Update(body []byte) error
	Delete() error
	Save() error
	Find() []interface{}
	HasID() bool
}
