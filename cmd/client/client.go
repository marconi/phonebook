package main

import (
	"fmt"

	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"github.com/marconi/phonebook/services/go/contact"
)

func main() {
	socket, err := thrift.NewTSocket("localhost:9090")
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport := transportFactory.GetTransport(socket)
	client := contact.NewContactSvcClientFactory(transport, protocolFactory)
	defer client.Transport.Close()
	if err := client.Transport.Open(); err != nil {
		panic(err)
	}

	c1 := contact.NewContactInit("Bob", "111-1111", "bob@wonderland.com")
	c1, err = client.Create(c1)
	if err != nil {
		panic(err)
	}

	c2, err := client.Read(c1.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(c2)
}
