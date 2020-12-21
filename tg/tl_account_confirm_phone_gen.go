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

// AccountConfirmPhoneRequest represents TL type `account.confirmPhone#5f2178c3`.
// Confirm a phone number to cancel account deletion, for more info click here »
//
// See https://core.telegram.org/method/account.confirmPhone for reference.
type AccountConfirmPhoneRequest struct {
	// Phone code hash, for more info click here »
	PhoneCodeHash string
	// SMS code, for more info click here »
	PhoneCode string
}

// AccountConfirmPhoneRequestTypeID is TL type id of AccountConfirmPhoneRequest.
const AccountConfirmPhoneRequestTypeID = 0x5f2178c3

// String implements fmt.Stringer.
func (c *AccountConfirmPhoneRequest) String() string {
	if c == nil {
		return "AccountConfirmPhoneRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountConfirmPhoneRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPhoneCodeHash: ")
	sb.WriteString(fmt.Sprint(c.PhoneCodeHash))
	sb.WriteString(",\n")
	sb.WriteString("\tPhoneCode: ")
	sb.WriteString(fmt.Sprint(c.PhoneCode))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *AccountConfirmPhoneRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.confirmPhone#5f2178c3 as nil")
	}
	b.PutID(AccountConfirmPhoneRequestTypeID)
	b.PutString(c.PhoneCodeHash)
	b.PutString(c.PhoneCode)
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountConfirmPhoneRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.confirmPhone#5f2178c3 to nil")
	}
	if err := b.ConsumeID(AccountConfirmPhoneRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.confirmPhone#5f2178c3: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPhone#5f2178c3: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.confirmPhone#5f2178c3: field phone_code: %w", err)
		}
		c.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountConfirmPhoneRequest.
var (
	_ bin.Encoder = &AccountConfirmPhoneRequest{}
	_ bin.Decoder = &AccountConfirmPhoneRequest{}
)

// AccountConfirmPhone invokes method account.confirmPhone#5f2178c3 returning error if any.
// Confirm a phone number to cancel account deletion, for more info click here »
//
// See https://core.telegram.org/method/account.confirmPhone for reference.
func (c *Client) AccountConfirmPhone(ctx context.Context, request *AccountConfirmPhoneRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
