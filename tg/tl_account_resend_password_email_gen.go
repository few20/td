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

// AccountResendPasswordEmailRequest represents TL type `account.resendPasswordEmail#7a7f2a15`.
// Resend the code to verify an email to use as 2FA recovery method.
//
// See https://core.telegram.org/method/account.resendPasswordEmail for reference.
type AccountResendPasswordEmailRequest struct {
}

// AccountResendPasswordEmailRequestTypeID is TL type id of AccountResendPasswordEmailRequest.
const AccountResendPasswordEmailRequestTypeID = 0x7a7f2a15

// String implements fmt.Stringer.
func (r *AccountResendPasswordEmailRequest) String() string {
	if r == nil {
		return "AccountResendPasswordEmailRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountResendPasswordEmailRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *AccountResendPasswordEmailRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resendPasswordEmail#7a7f2a15 as nil")
	}
	b.PutID(AccountResendPasswordEmailRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResendPasswordEmailRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resendPasswordEmail#7a7f2a15 to nil")
	}
	if err := b.ConsumeID(AccountResendPasswordEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resendPasswordEmail#7a7f2a15: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResendPasswordEmailRequest.
var (
	_ bin.Encoder = &AccountResendPasswordEmailRequest{}
	_ bin.Decoder = &AccountResendPasswordEmailRequest{}
)

// AccountResendPasswordEmail invokes method account.resendPasswordEmail#7a7f2a15 returning error if any.
// Resend the code to verify an email to use as 2FA recovery method.
//
// See https://core.telegram.org/method/account.resendPasswordEmail for reference.
func (c *Client) AccountResendPasswordEmail(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountResendPasswordEmailRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
