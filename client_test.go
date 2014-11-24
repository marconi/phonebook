package phonebook_test

import (
	"testing"
	"time"

	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"github.com/marconi/phonebook"
	"github.com/marconi/phonebook/services/go/contact"
	. "github.com/smartystreets/goconvey/convey"
)

const TEST_HOST string = "localhost:9191"

var (
	server *phonebook.PhonebookServer
	client *contact.ContactSvcClient
)

func init() {
	// init server
	server = phonebook.NewPhonebookServer(TEST_HOST)

	// init client
	socket, err := thrift.NewTSocket(TEST_HOST)
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport := transportFactory.GetTransport(socket)
	client = contact.NewContactSvcClientFactory(transport, protocolFactory)

	// run server
	go server.Run()

	// run client
	time.Sleep(2 * time.Second) // wait for server
	if err := client.Transport.Open(); err != nil {
		panic(err)
	}
}

func TestClientSpec(t *testing.T) {
	Convey("testing RPC client", t, func() {
		Convey("should be able to create contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			contact2, err := client.Create(contact1)
			So(err, ShouldBeNil)
			So(contact2.Id, ShouldEqual, contact1.Id)
		})

		Convey("should be able to read contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			client.Create(contact1)

			contact2, err := client.Read(contact1.Id)
			So(err, ShouldBeNil)
			So(contact1.Id, ShouldEqual, contact2.Id)
		})

		Convey("should be able to update contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			contact2, err := client.Create(contact1)
			So(err, ShouldBeNil)

			contact2.Name = "alice"
			client.Update(contact2)

			contact3, _ := client.Read(contact1.Id)
			So(contact2.Name, ShouldEqual, contact3.Name)
		})

		Convey("should be able to fetch contacts", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			contact2 := contact.NewContactInit("alice", "222-22222", "alice@wonderland.com")
			client.Create(contact1)
			client.Create(contact2)

			contacts, err := client.Fetch()
			So(err, ShouldBeNil)
			So(len(contacts), ShouldEqual, 2)
		})

		Convey("should be able to destroy contact", func() {
			contact1 := contact.NewContactInit("bob", "111-1111", "bob@wonderland.com")
			client.Create(contact1)
			client.Destroy(contact1.Id)

			contact2, err := client.Read(contact1.Id)
			So(err, ShouldNotBeNil)
			So(contact2, ShouldBeNil)
		})

		Reset(func() {
			client.Reset()
		})
	})

	server.Stop()
}
