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

// String implements fmt.Stringer.
func (c *ChannelAdminLogEvent) String() string {
	if c == nil {
		return "ChannelAdminLogEvent(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelAdminLogEvent")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(c.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(c.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(c.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tAction: ")
	sb.WriteString(c.Action.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChannelAdminLogEvent) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelAdminLogEvent#3b5a3e40 as nil")
	}
	b.PutID(ChannelAdminLogEventTypeID)
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

// Decode implements bin.Decoder.
func (c *ChannelAdminLogEvent) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelAdminLogEvent#3b5a3e40 to nil")
	}
	if err := b.ConsumeID(ChannelAdminLogEventTypeID); err != nil {
		return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: %w", err)
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
	_ bin.Encoder = &ChannelAdminLogEvent{}
	_ bin.Decoder = &ChannelAdminLogEvent{}
)
