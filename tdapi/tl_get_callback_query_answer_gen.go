// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// GetCallbackQueryAnswerRequest represents TL type `getCallbackQueryAnswer#6ef7a5f`.
type GetCallbackQueryAnswerRequest struct {
	// Identifier of the chat with the message
	ChatID int64
	// Identifier of the message from which the query originated. The message must not be
	// scheduled
	MessageID int64
	// Query payload
	Payload CallbackQueryPayloadClass
}

// GetCallbackQueryAnswerRequestTypeID is TL type id of GetCallbackQueryAnswerRequest.
const GetCallbackQueryAnswerRequestTypeID = 0x6ef7a5f

// Ensuring interfaces in compile-time for GetCallbackQueryAnswerRequest.
var (
	_ bin.Encoder     = &GetCallbackQueryAnswerRequest{}
	_ bin.Decoder     = &GetCallbackQueryAnswerRequest{}
	_ bin.BareEncoder = &GetCallbackQueryAnswerRequest{}
	_ bin.BareDecoder = &GetCallbackQueryAnswerRequest{}
)

func (g *GetCallbackQueryAnswerRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.Payload == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCallbackQueryAnswerRequest) String() string {
	if g == nil {
		return "GetCallbackQueryAnswerRequest(nil)"
	}
	type Alias GetCallbackQueryAnswerRequest
	return fmt.Sprintf("GetCallbackQueryAnswerRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCallbackQueryAnswerRequest) TypeID() uint32 {
	return GetCallbackQueryAnswerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCallbackQueryAnswerRequest) TypeName() string {
	return "getCallbackQueryAnswer"
}

// TypeInfo returns info about TL type.
func (g *GetCallbackQueryAnswerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCallbackQueryAnswer",
		ID:   GetCallbackQueryAnswerRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "Payload",
			SchemaName: "payload",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCallbackQueryAnswerRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCallbackQueryAnswer#6ef7a5f as nil")
	}
	b.PutID(GetCallbackQueryAnswerRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCallbackQueryAnswerRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCallbackQueryAnswer#6ef7a5f as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	if g.Payload == nil {
		return fmt.Errorf("unable to encode getCallbackQueryAnswer#6ef7a5f: field payload is nil")
	}
	if err := g.Payload.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getCallbackQueryAnswer#6ef7a5f: field payload: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCallbackQueryAnswerRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCallbackQueryAnswer#6ef7a5f to nil")
	}
	if err := b.ConsumeID(GetCallbackQueryAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCallbackQueryAnswerRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCallbackQueryAnswer#6ef7a5f to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := DecodeCallbackQueryPayload(b)
		if err != nil {
			return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: field payload: %w", err)
		}
		g.Payload = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetCallbackQueryAnswerRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCallbackQueryAnswer#6ef7a5f as nil")
	}
	b.ObjStart()
	b.PutID("getCallbackQueryAnswer")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("payload")
	if g.Payload == nil {
		return fmt.Errorf("unable to encode getCallbackQueryAnswer#6ef7a5f: field payload is nil")
	}
	if err := g.Payload.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getCallbackQueryAnswer#6ef7a5f: field payload: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetCallbackQueryAnswerRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCallbackQueryAnswer#6ef7a5f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getCallbackQueryAnswer"); err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: field message_id: %w", err)
			}
			g.MessageID = value
		case "payload":
			value, err := DecodeTDLibJSONCallbackQueryPayload(b)
			if err != nil {
				return fmt.Errorf("unable to decode getCallbackQueryAnswer#6ef7a5f: field payload: %w", err)
			}
			g.Payload = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetCallbackQueryAnswerRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetCallbackQueryAnswerRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetPayload returns value of Payload field.
func (g *GetCallbackQueryAnswerRequest) GetPayload() (value CallbackQueryPayloadClass) {
	if g == nil {
		return
	}
	return g.Payload
}

// GetCallbackQueryAnswer invokes method getCallbackQueryAnswer#6ef7a5f returning error if any.
func (c *Client) GetCallbackQueryAnswer(ctx context.Context, request *GetCallbackQueryAnswerRequest) (*CallbackQueryAnswer, error) {
	var result CallbackQueryAnswer

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
