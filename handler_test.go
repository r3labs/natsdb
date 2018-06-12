/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package natsdb

import (
	"os"
	"testing"
	"time"

	"github.com/nats-io/go-nats"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetHandler(t *testing.T) {
	var natsURL = os.Getenv("NATS_URI")
	n, _ := nats.Connect(natsURL)

	handler := Handler{
		NotFoundErrorMessage:   []byte(`{"error":"not found"}`),
		UnexpectedErrorMessage: []byte(`{"error":"unexpected"}`),
		DeletedMessage:         []byte(`{"deleted"}`),
		Nats:                   n,
		NewModel: func() Model {
			return &Entity{}
		},
	}

	Convey("Scenario: getting a client", t, func() {
		n.Subscribe("client.get", handler.Get)
		Convey("Given the client does not exist on the database", func() {
			msg, err := n.Request("client.get", []byte(`{"id":"32"}`), time.Second)
			So(string(msg.Data), ShouldEqual, string(`{"id":0,"name":""}`))
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Scenario: deleting a client", t, func() {
		n.Subscribe("client.del", handler.Del)
		Convey("Given the client does not exist on the database", func() {
			msg, err := n.Request("client.del", []byte(`{"id":"32"}`), time.Second)
			So(string(msg.Data), ShouldEqual, string(handler.DeletedMessage))
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Scenario: setting a client", t, func() {
		n.Subscribe("client.set", handler.Set)
		Convey("Given the client does not exist on the database", func() {
			msg, err := n.Request("client.set", []byte(`{"id":"32"}`), time.Second)
			So(string(msg.Data), ShouldEqual, string(`{"id":0,"name":""}`))
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Scenario: setting a client", t, func() {
		n.Subscribe("client.find", handler.Set)
		Convey("Given the client does not exist on the database", func() {
			msg, err := n.Request("client.find", []byte(`{"id":"32"}`), time.Second)
			So(string(msg.Data), ShouldEqual, string(`{"id":0,"name":""}`))
			So(err, ShouldEqual, nil)
		})
	})
}
