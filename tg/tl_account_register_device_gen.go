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

// AccountRegisterDeviceRequest represents TL type `account.registerDevice#68976c6f`.
// Register device to receive PUSH notifications
//
// See https://core.telegram.org/method/account.registerDevice for reference.
type AccountRegisterDeviceRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Avoid receiving (silent and invisible background) notifications. Useful to save battery.
	NoMuted bool
	// Device token type.Possible values:1 - APNS (device token for apple push)2 - FCM (firebase token for google firebase)3 - MPNS (channel URI for microsoft push)4 - Simple push (endpoint for firefox's simple push API)5 - Ubuntu phone (token for ubuntu push)6 - Blackberry (token for blackberry push)7 - Unused8 - WNS (windows push)9 - APNS VoIP (token for apple push VoIP)10 - Web push (web push, see below)11 - MPNS VoIP (token for microsoft push VoIP)12 - Tizen (token for tizen push)For 10 web push, the token must be a JSON-encoded object containing the keys described in PUSH updates
	TokenType int
	// Device token
	Token string
	// If (boolTrue) is transmitted, a sandbox-certificate will be used during transmission.
	AppSandbox bool
	// For FCM and APNS VoIP, optional encryption key used to encrypt push notifications
	Secret []byte
	// List of user identifiers of other users currently using the client
	OtherUids []int
}

// AccountRegisterDeviceRequestTypeID is TL type id of AccountRegisterDeviceRequest.
const AccountRegisterDeviceRequestTypeID = 0x68976c6f

// String implements fmt.Stringer.
func (r *AccountRegisterDeviceRequest) String() string {
	if r == nil {
		return "AccountRegisterDeviceRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountRegisterDeviceRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(r.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tTokenType: ")
	sb.WriteString(fmt.Sprint(r.TokenType))
	sb.WriteString(",\n")
	sb.WriteString("\tToken: ")
	sb.WriteString(fmt.Sprint(r.Token))
	sb.WriteString(",\n")
	sb.WriteString("\tAppSandbox: ")
	sb.WriteString(fmt.Sprint(r.AppSandbox))
	sb.WriteString(",\n")
	sb.WriteString("\tSecret: ")
	sb.WriteString(fmt.Sprint(r.Secret))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range r.OtherUids {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *AccountRegisterDeviceRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.registerDevice#68976c6f as nil")
	}
	b.PutID(AccountRegisterDeviceRequestTypeID)
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.registerDevice#68976c6f: field flags: %w", err)
	}
	b.PutInt(r.TokenType)
	b.PutString(r.Token)
	b.PutBool(r.AppSandbox)
	b.PutBytes(r.Secret)
	b.PutVectorHeader(len(r.OtherUids))
	for _, v := range r.OtherUids {
		b.PutInt(v)
	}
	return nil
}

// SetNoMuted sets value of NoMuted conditional field.
func (r *AccountRegisterDeviceRequest) SetNoMuted(value bool) {
	if value {
		r.Flags.Set(0)
	} else {
		r.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (r *AccountRegisterDeviceRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.registerDevice#68976c6f to nil")
	}
	if err := b.ConsumeID(AccountRegisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.registerDevice#68976c6f: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field flags: %w", err)
		}
	}
	r.NoMuted = r.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field token_type: %w", err)
		}
		r.TokenType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field token: %w", err)
		}
		r.Token = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field app_sandbox: %w", err)
		}
		r.AppSandbox = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field secret: %w", err)
		}
		r.Secret = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field other_uids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode account.registerDevice#68976c6f: field other_uids: %w", err)
			}
			r.OtherUids = append(r.OtherUids, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountRegisterDeviceRequest.
var (
	_ bin.Encoder = &AccountRegisterDeviceRequest{}
	_ bin.Decoder = &AccountRegisterDeviceRequest{}
)

// AccountRegisterDevice invokes method account.registerDevice#68976c6f returning error if any.
// Register device to receive PUSH notifications
//
// See https://core.telegram.org/method/account.registerDevice for reference.
func (c *Client) AccountRegisterDevice(ctx context.Context, request *AccountRegisterDeviceRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
