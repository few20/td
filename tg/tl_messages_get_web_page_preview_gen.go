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

// MessagesGetWebPagePreviewRequest represents TL type `messages.getWebPagePreview#8b68b0cc`.
// Get preview of webpage
//
// See https://core.telegram.org/method/messages.getWebPagePreview for reference.
type MessagesGetWebPagePreviewRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Message from which to extract the preview
	Message string
	// Message entities for styled text
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// MessagesGetWebPagePreviewRequestTypeID is TL type id of MessagesGetWebPagePreviewRequest.
const MessagesGetWebPagePreviewRequestTypeID = 0x8b68b0cc

// String implements fmt.Stringer.
func (g *MessagesGetWebPagePreviewRequest) String() string {
	if g == nil {
		return "MessagesGetWebPagePreviewRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetWebPagePreviewRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(g.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(g.Message))
	sb.WriteString(",\n")
	if g.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range g.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetWebPagePreviewRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPagePreview#8b68b0cc as nil")
	}
	b.PutID(MessagesGetWebPagePreviewRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field flags: %w", err)
	}
	b.PutString(g.Message)
	if g.Flags.Has(3) {
		b.PutVectorHeader(len(g.Entities))
		for idx, v := range g.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.getWebPagePreview#8b68b0cc: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetEntities sets value of Entities conditional field.
func (g *MessagesGetWebPagePreviewRequest) SetEntities(value []MessageEntityClass) {
	g.Flags.Set(3)
	g.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (g *MessagesGetWebPagePreviewRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.Entities, true
}

// Decode implements bin.Decoder.
func (g *MessagesGetWebPagePreviewRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPagePreview#8b68b0cc to nil")
	}
	if err := b.ConsumeID(MessagesGetWebPagePreviewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field message: %w", err)
		}
		g.Message = value
	}
	if g.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getWebPagePreview#8b68b0cc: field entities: %w", err)
			}
			g.Entities = append(g.Entities, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetWebPagePreviewRequest.
var (
	_ bin.Encoder = &MessagesGetWebPagePreviewRequest{}
	_ bin.Decoder = &MessagesGetWebPagePreviewRequest{}
)

// MessagesGetWebPagePreview invokes method messages.getWebPagePreview#8b68b0cc returning error if any.
// Get preview of webpage
//
// See https://core.telegram.org/method/messages.getWebPagePreview for reference.
func (c *Client) MessagesGetWebPagePreview(ctx context.Context, request *MessagesGetWebPagePreviewRequest) (MessageMediaClass, error) {
	var result MessageMediaBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.MessageMedia, nil
}
