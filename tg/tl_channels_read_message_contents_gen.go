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

// ChannelsReadMessageContentsRequest represents TL type `channels.readMessageContents#eab5dc38`.
// Mark channel/supergroup message contents as read
//
// See https://core.telegram.org/method/channels.readMessageContents for reference.
type ChannelsReadMessageContentsRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// IDs of messages whose contents should be marked as read
	ID []int
}

// ChannelsReadMessageContentsRequestTypeID is TL type id of ChannelsReadMessageContentsRequest.
const ChannelsReadMessageContentsRequestTypeID = 0xeab5dc38

// String implements fmt.Stringer.
func (r *ChannelsReadMessageContentsRequest) String() string {
	if r == nil {
		return "ChannelsReadMessageContentsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsReadMessageContentsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(r.Channel.String())
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range r.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *ChannelsReadMessageContentsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode channels.readMessageContents#eab5dc38 as nil")
	}
	b.PutID(ChannelsReadMessageContentsRequestTypeID)
	if r.Channel == nil {
		return fmt.Errorf("unable to encode channels.readMessageContents#eab5dc38: field channel is nil")
	}
	if err := r.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.readMessageContents#eab5dc38: field channel: %w", err)
	}
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ChannelsReadMessageContentsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode channels.readMessageContents#eab5dc38 to nil")
	}
	if err := b.ConsumeID(ChannelsReadMessageContentsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: field channel: %w", err)
		}
		r.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.readMessageContents#eab5dc38: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsReadMessageContentsRequest.
var (
	_ bin.Encoder = &ChannelsReadMessageContentsRequest{}
	_ bin.Decoder = &ChannelsReadMessageContentsRequest{}
)

// ChannelsReadMessageContents invokes method channels.readMessageContents#eab5dc38 returning error if any.
// Mark channel/supergroup message contents as read
//
// See https://core.telegram.org/method/channels.readMessageContents for reference.
func (c *Client) ChannelsReadMessageContents(ctx context.Context, request *ChannelsReadMessageContentsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
