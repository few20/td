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

// MessagesGetAllStickersRequest represents TL type `messages.getAllStickers#1c9618b1`.
// Get all installed stickers
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
type MessagesGetAllStickersRequest struct {
	// Hash for pagination, for more info click here
	Hash int
}

// MessagesGetAllStickersRequestTypeID is TL type id of MessagesGetAllStickersRequest.
const MessagesGetAllStickersRequestTypeID = 0x1c9618b1

// String implements fmt.Stringer.
func (g *MessagesGetAllStickersRequest) String() string {
	if g == nil {
		return "MessagesGetAllStickersRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetAllStickersRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(g.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetAllStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllStickers#1c9618b1 as nil")
	}
	b.PutID(MessagesGetAllStickersRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllStickers#1c9618b1 to nil")
	}
	if err := b.ConsumeID(MessagesGetAllStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAllStickers#1c9618b1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAllStickers#1c9618b1: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllStickersRequest.
var (
	_ bin.Encoder = &MessagesGetAllStickersRequest{}
	_ bin.Decoder = &MessagesGetAllStickersRequest{}
)

// MessagesGetAllStickers invokes method messages.getAllStickers#1c9618b1 returning error if any.
// Get all installed stickers
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
func (c *Client) MessagesGetAllStickers(ctx context.Context, hash int) (MessagesAllStickersClass, error) {
	var result MessagesAllStickersBox

	request := &MessagesGetAllStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStickers, nil
}
