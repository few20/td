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

// StickersSetStickerSetThumbRequest represents TL type `stickers.setStickerSetThumb#9a364e30`.
// Set stickerset thumbnail
//
// See https://core.telegram.org/method/stickers.setStickerSetThumb for reference.
type StickersSetStickerSetThumbRequest struct {
	// Stickerset
	Stickerset InputStickerSetClass
	// Thumbnail
	Thumb InputDocumentClass
}

// StickersSetStickerSetThumbRequestTypeID is TL type id of StickersSetStickerSetThumbRequest.
const StickersSetStickerSetThumbRequestTypeID = 0x9a364e30

func (s *StickersSetStickerSetThumbRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Stickerset == nil) {
		return false
	}
	if !(s.Thumb == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickersSetStickerSetThumbRequest) String() string {
	if s == nil {
		return "StickersSetStickerSetThumbRequest(nil)"
	}
	type Alias StickersSetStickerSetThumbRequest
	return fmt.Sprintf("StickersSetStickerSetThumbRequest%+v", Alias(*s))
}

// FillFrom fills StickersSetStickerSetThumbRequest from given interface.
func (s *StickersSetStickerSetThumbRequest) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
	GetThumb() (value InputDocumentClass)
}) {
	s.Stickerset = from.GetStickerset()
	s.Thumb = from.GetThumb()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersSetStickerSetThumbRequest) TypeID() uint32 {
	return StickersSetStickerSetThumbRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersSetStickerSetThumbRequest) TypeName() string {
	return "stickers.setStickerSetThumb"
}

// TypeInfo returns info about TL type.
func (s *StickersSetStickerSetThumbRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.setStickerSetThumb",
		ID:   StickersSetStickerSetThumbRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
		{
			Name:       "Thumb",
			SchemaName: "thumb",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickersSetStickerSetThumbRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.setStickerSetThumb#9a364e30 as nil")
	}
	b.PutID(StickersSetStickerSetThumbRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickersSetStickerSetThumbRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers.setStickerSetThumb#9a364e30 as nil")
	}
	if s.Stickerset == nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field stickerset is nil")
	}
	if err := s.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field stickerset: %w", err)
	}
	if s.Thumb == nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field thumb is nil")
	}
	if err := s.Thumb.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.setStickerSetThumb#9a364e30: field thumb: %w", err)
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (s *StickersSetStickerSetThumbRequest) GetStickerset() (value InputStickerSetClass) {
	return s.Stickerset
}

// GetThumb returns value of Thumb field.
func (s *StickersSetStickerSetThumbRequest) GetThumb() (value InputDocumentClass) {
	return s.Thumb
}

// GetThumbAsNotEmpty returns mapped value of Thumb field.
func (s *StickersSetStickerSetThumbRequest) GetThumbAsNotEmpty() (*InputDocument, bool) {
	return s.Thumb.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (s *StickersSetStickerSetThumbRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.setStickerSetThumb#9a364e30 to nil")
	}
	if err := b.ConsumeID(StickersSetStickerSetThumbRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.setStickerSetThumb#9a364e30: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickersSetStickerSetThumbRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers.setStickerSetThumb#9a364e30 to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.setStickerSetThumb#9a364e30: field stickerset: %w", err)
		}
		s.Stickerset = value
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.setStickerSetThumb#9a364e30: field thumb: %w", err)
		}
		s.Thumb = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersSetStickerSetThumbRequest.
var (
	_ bin.Encoder     = &StickersSetStickerSetThumbRequest{}
	_ bin.Decoder     = &StickersSetStickerSetThumbRequest{}
	_ bin.BareEncoder = &StickersSetStickerSetThumbRequest{}
	_ bin.BareDecoder = &StickersSetStickerSetThumbRequest{}
)

// StickersSetStickerSetThumb invokes method stickers.setStickerSetThumb#9a364e30 returning error if any.
// Set stickerset thumbnail
//
// Possible errors:
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/stickers.setStickerSetThumb for reference.
// Can be used by bots.
func (c *Client) StickersSetStickerSetThumb(ctx context.Context, request *StickersSetStickerSetThumbRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
