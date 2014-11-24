package phonebook_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/marconi/phonebook"
	"github.com/marconi/phonebook/services/go/contact"
)

func TestHandlerSpec(t *testing.T) {
	handler := phonebook.NewContactHandler()

	Convey("testing handler", t, func() {
		Convey("should be able to create contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			contact2, err := handler.Create(contact1)
			So(err, ShouldBeNil)
			So(contact1.Id, ShouldEqual, contact2.Id)

			contacts, err := handler.Fetch()
			So(err, ShouldBeNil)
			So(len(contacts), ShouldEqual, 1)
		})

		Convey("should be able to read contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			handler.Create(contact1)

			contact2, err := handler.Read(contact1.Id)
			So(err, ShouldBeNil)
			So(contact1.Id, ShouldEqual, contact2.Id)
		})

		Convey("should be able to update contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			contact2, _ := handler.Create(contact1)

			contact2.Name = "alice"
			handler.Update(contact2)

			contact3, _ := handler.Read(contact1.Id)
			So(contact2.Name, ShouldEqual, contact3.Name)
		})

		Convey("should be able to destroy contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			handler.Create(contact1)
			handler.Destroy(contact1.Id)

			contact2, err := handler.Read(contact1.Id)
			So(err, ShouldNotBeNil)
			So(contact2, ShouldBeNil)
		})

		Convey("should be able to fetch contacts", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			contact2 := contact.NewContactInit("alice", "222-22222", "alice@wonderland.com")
			handler.Create(contact1)
			handler.Create(contact2)

			contacts, err := handler.Fetch()
			So(err, ShouldBeNil)
			So(len(contacts), ShouldEqual, 2)
		})

		Reset(func() {
			handler.Reset()
		})
	})
}
