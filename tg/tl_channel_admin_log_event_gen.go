// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// ChannelAdminLogEvent represents TL type `channelAdminLogEvent#3b5a3e40`.
// Admin log event
//
// See https://core.telegram.org/constructor/channelAdminLogEvent for reference.
type ChannelAdminLogEvent struct {
	// Event ID
	ID int64
	// Date
	Date int
	// User ID
	UserID int
	// Action
	Action ChannelAdminLogEventActionClass
}

// ChannelAdminLogEventTypeID is TL type id of ChannelAdminLogEvent.
const ChannelAdminLogEventTypeID = 0x3b5a3e40

func (c *ChannelAdminLogEvent) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Action == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelAdminLogEvent) String() string {
	if c == nil {
		return "ChannelAdminLogEvent(nil)"
	}
	type Alias ChannelAdminLogEvent
	return fmt.Sprintf("ChannelAdminLogEvent%+v", Alias(*c))
}

// FillFrom fills ChannelAdminLogEvent from given interface.
func (c *ChannelAdminLogEvent) FillFrom(from interface {
	GetID() (value int64)
	GetDate() (value int)
	GetUserID() (value int)
	GetAction() (value ChannelAdminLogEventActionClass)
}) {
	c.ID = from.GetID()
	c.Date = from.GetDate()
	c.UserID = from.GetUserID()
	c.Action = from.GetAction()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelAdminLogEvent) TypeID() uint32 {
	return ChannelAdminLogEventTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelAdminLogEvent) TypeName() string {
	return "channelAdminLogEvent"
}

// TypeInfo returns info about TL type.
func (c *ChannelAdminLogEvent) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelAdminLogEvent",
		ID:   ChannelAdminLogEventTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Action",
			SchemaName: "action",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelAdminLogEvent) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelAdminLogEvent#3b5a3e40 as nil")
	}
	b.PutID(ChannelAdminLogEventTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelAdminLogEvent) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelAdminLogEvent#3b5a3e40 as nil")
	}
	b.PutLong(c.ID)
	b.PutInt(c.Date)
	b.PutInt(c.UserID)
	if c.Action == nil {
		return fmt.Errorf("unable to encode channelAdminLogEvent#3b5a3e40: field action is nil")
	}
	if err := c.Action.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelAdminLogEvent#3b5a3e40: field action: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (c *ChannelAdminLogEvent) GetID() (value int64) {
	return c.ID
}

// GetDate returns value of Date field.
func (c *ChannelAdminLogEvent) GetDate() (value int) {
	return c.Date
}

// GetUserID returns value of UserID field.
func (c *ChannelAdminLogEvent) GetUserID() (value int) {
	return c.UserID
}

// GetAction returns value of Action field.
func (c *ChannelAdminLogEvent) GetAction() (value ChannelAdminLogEventActionClass) {
	return c.Action
}

// Decode implements bin.Decoder.
func (c *ChannelAdminLogEvent) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelAdminLogEvent#3b5a3e40 to nil")
	}
	if err := b.ConsumeID(ChannelAdminLogEventTypeID); err != nil {
		return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelAdminLogEvent) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelAdminLogEvent#3b5a3e40 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := DecodeChannelAdminLogEventAction(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field action: %w", err)
		}
		c.Action = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelAdminLogEvent.
var (
	_ bin.Encoder     = &ChannelAdminLogEvent{}
	_ bin.Decoder     = &ChannelAdminLogEvent{}
	_ bin.BareEncoder = &ChannelAdminLogEvent{}
	_ bin.BareDecoder = &ChannelAdminLogEvent{}
)
