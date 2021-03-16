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

// MessagesGetChatsRequest represents TL type `messages.getChats#3c6aa187`.
// Returns chat basic info on their IDs.
//
// See https://core.telegram.org/method/messages.getChats for reference.
type MessagesGetChatsRequest struct {
	// List of chat IDs
	ID []int
}

// MessagesGetChatsRequestTypeID is TL type id of MessagesGetChatsRequest.
const MessagesGetChatsRequestTypeID = 0x3c6aa187

func (g *MessagesGetChatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetChatsRequest) String() string {
	if g == nil {
		return "MessagesGetChatsRequest(nil)"
	}
	type Alias MessagesGetChatsRequest
	return fmt.Sprintf("MessagesGetChatsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetChatsRequest from given interface.
func (g *MessagesGetChatsRequest) FillFrom(from interface {
	GetID() (value []int)
}) {
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetChatsRequest) TypeID() uint32 {
	return MessagesGetChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetChatsRequest) TypeName() string {
	return "messages.getChats"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getChats",
		ID:   MessagesGetChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getChats#3c6aa187 as nil")
	}
	b.PutID(MessagesGetChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getChats#3c6aa187 as nil")
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// GetID returns value of ID field.
func (g *MessagesGetChatsRequest) GetID() (value []int) {
	return g.ID
}

// Decode implements bin.Decoder.
func (g *MessagesGetChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getChats#3c6aa187 to nil")
	}
	if err := b.ConsumeID(MessagesGetChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getChats#3c6aa187: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getChats#3c6aa187 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChats#3c6aa187: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getChats#3c6aa187: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetChatsRequest.
var (
	_ bin.Encoder     = &MessagesGetChatsRequest{}
	_ bin.Decoder     = &MessagesGetChatsRequest{}
	_ bin.BareEncoder = &MessagesGetChatsRequest{}
	_ bin.BareDecoder = &MessagesGetChatsRequest{}
)

// MessagesGetChats invokes method messages.getChats#3c6aa187 returning error if any.
// Returns chat basic info on their IDs.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getChats for reference.
// Can be used by bots.
func (c *Client) MessagesGetChats(ctx context.Context, id []int) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &MessagesGetChatsRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
