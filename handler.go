package phonebook

import (
	"fmt"
	"sync"

	"github.com/marconi/phonebook/services/go/contact"
)

type ContactHandler struct {
	contacts map[string]*contact.Contact
	sync.RWMutex
}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{
		contacts: make(map[string]*contact.Contact),
	}
}

func (ch *ContactHandler) Create(contact *contact.Contact) (*contact.Contact, error) {
	ch.Lock()
	defer ch.Unlock()
	ch.contacts[contact.Id] = contact
	return contact, nil
}

func (ch *ContactHandler) Read(contactId string) (*contact.Contact, error) {
	contact, ok := ch.contacts[contactId]
	if !ok {
		return nil, fmt.Errorf("Contact with ID '%s' does not exist", contactId)
	}
	return contact, nil
}

func (ch *ContactHandler) Update(contact *contact.Contact) (*contact.Contact, error) {
	ch.Lock()
	defer ch.Unlock()
	ch.contacts[contact.Id] = contact
	return contact, nil
}

func (ch *ContactHandler) Destroy(contactId string) error {
	if _, ok := ch.contacts[contactId]; ok {
		ch.Lock()
		defer ch.Unlock()
		delete(ch.contacts, contactId)
	}
	return nil
}

func (ch *ContactHandler) Fetch() ([]*contact.Contact, error) {
	var contacts []*contact.Contact
	for _, contact := range ch.contacts {
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (ch *ContactHandler) Reset() error {
	ch.Lock()
	defer ch.Unlock()
	ch.contacts = make(map[string]*contact.Contact)
	return nil
}
