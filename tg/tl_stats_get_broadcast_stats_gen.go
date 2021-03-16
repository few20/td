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

// StatsGetBroadcastStatsRequest represents TL type `stats.getBroadcastStats#ab42441a`.
// Get channel statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/method/stats.getBroadcastStats for reference.
type StatsGetBroadcastStatsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to enable dark theme for graph colors
	Dark bool
	// The channel
	Channel InputChannelClass
}

// StatsGetBroadcastStatsRequestTypeID is TL type id of StatsGetBroadcastStatsRequest.
const StatsGetBroadcastStatsRequestTypeID = 0xab42441a

func (g *StatsGetBroadcastStatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Dark == false) {
		return false
	}
	if !(g.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetBroadcastStatsRequest) String() string {
	if g == nil {
		return "StatsGetBroadcastStatsRequest(nil)"
	}
	type Alias StatsGetBroadcastStatsRequest
	return fmt.Sprintf("StatsGetBroadcastStatsRequest%+v", Alias(*g))
}

// FillFrom fills StatsGetBroadcastStatsRequest from given interface.
func (g *StatsGetBroadcastStatsRequest) FillFrom(from interface {
	GetDark() (value bool)
	GetChannel() (value InputChannelClass)
}) {
	g.Dark = from.GetDark()
	g.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetBroadcastStatsRequest) TypeID() uint32 {
	return StatsGetBroadcastStatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetBroadcastStatsRequest) TypeName() string {
	return "stats.getBroadcastStats"
}

// TypeInfo returns info about TL type.
func (g *StatsGetBroadcastStatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getBroadcastStats",
		ID:   StatsGetBroadcastStatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StatsGetBroadcastStatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastStats#ab42441a as nil")
	}
	b.PutID(StatsGetBroadcastStatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetBroadcastStatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastStats#ab42441a as nil")
	}
	if !(g.Dark == false) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastStats#ab42441a: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getBroadcastStats#ab42441a: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastStats#ab42441a: field channel: %w", err)
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *StatsGetBroadcastStatsRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
		g.Dark = true
	} else {
		g.Flags.Unset(0)
		g.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (g *StatsGetBroadcastStatsRequest) GetDark() (value bool) {
	return g.Flags.Has(0)
}

// GetChannel returns value of Channel field.
func (g *StatsGetBroadcastStatsRequest) GetChannel() (value InputChannelClass) {
	return g.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *StatsGetBroadcastStatsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (g *StatsGetBroadcastStatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastStats#ab42441a to nil")
	}
	if err := b.ConsumeID(StatsGetBroadcastStatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getBroadcastStats#ab42441a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetBroadcastStatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastStats#ab42441a to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastStats#ab42441a: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastStats#ab42441a: field channel: %w", err)
		}
		g.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGetBroadcastStatsRequest.
var (
	_ bin.Encoder     = &StatsGetBroadcastStatsRequest{}
	_ bin.Decoder     = &StatsGetBroadcastStatsRequest{}
	_ bin.BareEncoder = &StatsGetBroadcastStatsRequest{}
	_ bin.BareDecoder = &StatsGetBroadcastStatsRequest{}
)

// StatsGetBroadcastStats invokes method stats.getBroadcastStats#ab42441a returning error if any.
// Get channel statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// Possible errors:
//  400 BROADCAST_REQUIRED: This method can only be called on a channel, please use stats.getMegagroupStats for supergroups
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//
// See https://core.telegram.org/method/stats.getBroadcastStats for reference.
func (c *Client) StatsGetBroadcastStats(ctx context.Context, request *StatsGetBroadcastStatsRequest) (*StatsBroadcastStats, error) {
	var result StatsBroadcastStats

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
