// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// ContactsAddContactRequest represents TL type `contacts.addContact#e8f463d0`.
// Add an existing telegram user as contact.
// Use contacts.importContacts to add contacts by phone number, without knowing their Telegram ID.
//
// See https://core.telegram.org/method/contacts.addContact for reference.
type ContactsAddContactRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Allow the other user to see our phone number?
	AddPhonePrivacyException bool
	// Telegram ID of the other user
	ID InputUserClass
	// First name
	FirstName string
	// Last name
	LastName string
	// User's phone number
	Phone string
}

// ContactsAddContactRequestTypeID is TL type id of ContactsAddContactRequest.
const ContactsAddContactRequestTypeID = 0xe8f463d0

// String implements fmt.Stringer.
func (a *ContactsAddContactRequest) String() string {
	if a == nil {
		return "ContactsAddContactRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsAddContactRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(a.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(a.ID.String())
	sb.WriteString(",\n")
	sb.WriteString("\tFirstName: ")
	sb.WriteString(fmt.Sprint(a.FirstName))
	sb.WriteString(",\n")
	sb.WriteString("\tLastName: ")
	sb.WriteString(fmt.Sprint(a.LastName))
	sb.WriteString(",\n")
	sb.WriteString("\tPhone: ")
	sb.WriteString(fmt.Sprint(a.Phone))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *ContactsAddContactRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode contacts.addContact#e8f463d0 as nil")
	}
	b.PutID(ContactsAddContactRequestTypeID)
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field flags: %w", err)
	}
	if a.ID == nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field id is nil")
	}
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.addContact#e8f463d0: field id: %w", err)
	}
	b.PutString(a.FirstName)
	b.PutString(a.LastName)
	b.PutString(a.Phone)
	return nil
}

// SetAddPhonePrivacyException sets value of AddPhonePrivacyException conditional field.
func (a *ContactsAddContactRequest) SetAddPhonePrivacyException(value bool) {
	if value {
		a.Flags.Set(0)
	} else {
		a.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (a *ContactsAddContactRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode contacts.addContact#e8f463d0 to nil")
	}
	if err := b.ConsumeID(ContactsAddContactRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field flags: %w", err)
		}
	}
	a.AddPhonePrivacyException = a.Flags.Has(0)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field first_name: %w", err)
		}
		a.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field last_name: %w", err)
		}
		a.LastName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.addContact#e8f463d0: field phone: %w", err)
		}
		a.Phone = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsAddContactRequest.
var (
	_ bin.Encoder = &ContactsAddContactRequest{}
	_ bin.Decoder = &ContactsAddContactRequest{}
)

// ContactsAddContact invokes method contacts.addContact#e8f463d0 returning error if any.
// Add an existing telegram user as contact.
// Use contacts.importContacts to add contacts by phone number, without knowing their Telegram ID.
//
// See https://core.telegram.org/method/contacts.addContact for reference.
func (c *Client) ContactsAddContact(ctx context.Context, request *ContactsAddContactRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
