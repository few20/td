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

// ChannelsDeleteHistoryRequest represents TL type `channels.deleteHistory#af369d42`.
// Delete the history of a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteHistory for reference.
type ChannelsDeleteHistoryRequest struct {
	// Supergroup¹ whose history must be deleted
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// ID of message up to which the history must be deleted
	MaxID int
}

// ChannelsDeleteHistoryRequestTypeID is TL type id of ChannelsDeleteHistoryRequest.
const ChannelsDeleteHistoryRequestTypeID = 0xaf369d42

func (d *ChannelsDeleteHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteHistoryRequest) String() string {
	if d == nil {
		return "ChannelsDeleteHistoryRequest(nil)"
	}
	type Alias ChannelsDeleteHistoryRequest
	return fmt.Sprintf("ChannelsDeleteHistoryRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteHistoryRequest from given interface.
func (d *ChannelsDeleteHistoryRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetMaxID() (value int)
}) {
	d.Channel = from.GetChannel()
	d.MaxID = from.GetMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteHistoryRequest) TypeID() uint32 {
	return ChannelsDeleteHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteHistoryRequest) TypeName() string {
	return "channels.deleteHistory"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteHistory",
		ID:   ChannelsDeleteHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteHistory#af369d42 as nil")
	}
	b.PutID(ChannelsDeleteHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteHistory#af369d42 as nil")
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#af369d42: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#af369d42: field channel: %w", err)
	}
	b.PutInt(d.MaxID)
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteHistoryRequest) GetChannel() (value InputChannelClass) {
	return d.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteHistoryRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// GetMaxID returns value of MaxID field.
func (d *ChannelsDeleteHistoryRequest) GetMaxID() (value int) {
	return d.MaxID
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteHistory#af369d42 to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteHistory#af369d42: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteHistory#af369d42 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#af369d42: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#af369d42: field max_id: %w", err)
		}
		d.MaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteHistoryRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteHistoryRequest{}
	_ bin.Decoder     = &ChannelsDeleteHistoryRequest{}
	_ bin.BareEncoder = &ChannelsDeleteHistoryRequest{}
	_ bin.BareDecoder = &ChannelsDeleteHistoryRequest{}
)

// ChannelsDeleteHistory invokes method channels.deleteHistory#af369d42 returning error if any.
// Delete the history of a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//
// See https://core.telegram.org/method/channels.deleteHistory for reference.
func (c *Client) ChannelsDeleteHistory(ctx context.Context, request *ChannelsDeleteHistoryRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
