// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// InputPhoneCall represents TL type `inputPhoneCall#1e36fded`.
type InputPhoneCall struct {
	// ID field of InputPhoneCall.
	ID int64
	// AccessHash field of InputPhoneCall.
	AccessHash int64
}

// InputPhoneCallTypeID is TL type id of InputPhoneCall.
const InputPhoneCallTypeID = 0x1e36fded

// Encode implements bin.Encoder.
func (i *InputPhoneCall) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneCall#1e36fded as nil")
	}
	b.PutID(InputPhoneCallTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhoneCall) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneCall#1e36fded to nil")
	}
	if err := b.ConsumeID(InputPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputPhoneCall.
var (
	_ bin.Encoder = &InputPhoneCall{}
	_ bin.Decoder = &InputPhoneCall{}
)