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

// ChannelsEditLocationRequest represents TL type `channels.editLocation#58e63f6d`.
// Edit location of geogroup
//
// See https://core.telegram.org/method/channels.editLocation for reference.
type ChannelsEditLocationRequest struct {
	// Geogroup
	Channel InputChannelClass
	// New geolocation
	GeoPoint InputGeoPointClass
	// Address string
	Address string
}

// ChannelsEditLocationRequestTypeID is TL type id of ChannelsEditLocationRequest.
const ChannelsEditLocationRequestTypeID = 0x58e63f6d

// String implements fmt.Stringer.
func (e *ChannelsEditLocationRequest) String() string {
	if e == nil {
		return "ChannelsEditLocationRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsEditLocationRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(e.Channel.String())
	sb.WriteString(",\n")
	sb.WriteString("\tGeoPoint: ")
	sb.WriteString(e.GeoPoint.String())
	sb.WriteString(",\n")
	sb.WriteString("\tAddress: ")
	sb.WriteString(fmt.Sprint(e.Address))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *ChannelsEditLocationRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editLocation#58e63f6d as nil")
	}
	b.PutID(ChannelsEditLocationRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field channel: %w", err)
	}
	if e.GeoPoint == nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field geo_point is nil")
	}
	if err := e.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editLocation#58e63f6d: field geo_point: %w", err)
	}
	b.PutString(e.Address)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditLocationRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editLocation#58e63f6d to nil")
	}
	if err := b.ConsumeID(ChannelsEditLocationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field geo_point: %w", err)
		}
		e.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editLocation#58e63f6d: field address: %w", err)
		}
		e.Address = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditLocationRequest.
var (
	_ bin.Encoder = &ChannelsEditLocationRequest{}
	_ bin.Decoder = &ChannelsEditLocationRequest{}
)

// ChannelsEditLocation invokes method channels.editLocation#58e63f6d returning error if any.
// Edit location of geogroup
//
// See https://core.telegram.org/method/channels.editLocation for reference.
func (c *Client) ChannelsEditLocation(ctx context.Context, request *ChannelsEditLocationRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
