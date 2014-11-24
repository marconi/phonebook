package phonebook

import (
	"fmt"

	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"github.com/marconi/phonebook/services/go/contact"
)

type PhonebookServer struct {
	host             string
	handler          *ContactHandler
	processor        *contact.ContactSvcProcessor
	transport        *thrift.TServerSocket
	transportFactory thrift.TTransportFactory
	protocolFactory  *thrift.TBinaryProtocolFactory
	server           *thrift.TSimpleServer
}

func NewPhonebookServer(host string) *PhonebookServer {
	handler := NewContactHandler()
	processor := contact.NewContactSvcProcessor(handler)
	transport, err := thrift.NewTServerSocket(host)
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return &PhonebookServer{
		host:             host,
		handler:          handler,
		processor:        processor,
		transport:        transport,
		transportFactory: transportFactory,
		protocolFactory:  protocolFactory,
		server:           server,
	}
}

func (ps *PhonebookServer) Run() {
	fmt.Printf("server listening on %s\n", ps.host)
	ps.server.Serve()
}

func (ps *PhonebookServer) Stop() {
	fmt.Println("stopping server...")
	ps.server.Stop()
}
