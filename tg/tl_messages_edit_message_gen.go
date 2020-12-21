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

// MessagesEditMessageRequest represents TL type `messages.editMessage#48f71778`.
// Edit message
//
// See https://core.telegram.org/method/messages.editMessage for reference.
type MessagesEditMessageRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Disable webpage preview
	NoWebpage bool
	// Where was the message sent
	Peer InputPeerClass
	// ID of the message to edit
	ID int
	// New message
	//
	// Use SetMessage and GetMessage helpers.
	Message string
	// New attached media
	//
	// Use SetMedia and GetMedia helpers.
	Media InputMediaClass
	// Reply markup for inline keyboards
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
	// Message entities for styled text
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Scheduled message date for scheduled messages
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
}

// MessagesEditMessageRequestTypeID is TL type id of MessagesEditMessageRequest.
const MessagesEditMessageRequestTypeID = 0x48f71778

// String implements fmt.Stringer.
func (e *MessagesEditMessageRequest) String() string {
	if e == nil {
		return "MessagesEditMessageRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesEditMessageRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(e.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(e.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	if e.Flags.Has(11) {
		sb.WriteString("\tMessage: ")
		sb.WriteString(fmt.Sprint(e.Message))
		sb.WriteString(",\n")
	}
	if e.Flags.Has(14) {
		sb.WriteString("\tMedia: ")
		sb.WriteString(e.Media.String())
		sb.WriteString(",\n")
	}
	if e.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(e.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	if e.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range e.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if e.Flags.Has(15) {
		sb.WriteString("\tScheduleDate: ")
		sb.WriteString(fmt.Sprint(e.ScheduleDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *MessagesEditMessageRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editMessage#48f71778 as nil")
	}
	b.PutID(MessagesEditMessageRequestTypeID)
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editMessage#48f71778: field flags: %w", err)
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.editMessage#48f71778: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editMessage#48f71778: field peer: %w", err)
	}
	b.PutInt(e.ID)
	if e.Flags.Has(11) {
		b.PutString(e.Message)
	}
	if e.Flags.Has(14) {
		if e.Media == nil {
			return fmt.Errorf("unable to encode messages.editMessage#48f71778: field media is nil")
		}
		if err := e.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.editMessage#48f71778: field media: %w", err)
		}
	}
	if e.Flags.Has(2) {
		if e.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.editMessage#48f71778: field reply_markup is nil")
		}
		if err := e.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.editMessage#48f71778: field reply_markup: %w", err)
		}
	}
	if e.Flags.Has(3) {
		b.PutVectorHeader(len(e.Entities))
		for idx, v := range e.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.editMessage#48f71778: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.editMessage#48f71778: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if e.Flags.Has(15) {
		b.PutInt(e.ScheduleDate)
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (e *MessagesEditMessageRequest) SetNoWebpage(value bool) {
	if value {
		e.Flags.Set(1)
	} else {
		e.Flags.Unset(1)
	}
}

// SetMessage sets value of Message conditional field.
func (e *MessagesEditMessageRequest) SetMessage(value string) {
	e.Flags.Set(11)
	e.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (e *MessagesEditMessageRequest) GetMessage() (value string, ok bool) {
	if !e.Flags.Has(11) {
		return value, false
	}
	return e.Message, true
}

// SetMedia sets value of Media conditional field.
func (e *MessagesEditMessageRequest) SetMedia(value InputMediaClass) {
	e.Flags.Set(14)
	e.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (e *MessagesEditMessageRequest) GetMedia() (value InputMediaClass, ok bool) {
	if !e.Flags.Has(14) {
		return value, false
	}
	return e.Media, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (e *MessagesEditMessageRequest) SetReplyMarkup(value ReplyMarkupClass) {
	e.Flags.Set(2)
	e.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (e *MessagesEditMessageRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !e.Flags.Has(2) {
		return value, false
	}
	return e.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (e *MessagesEditMessageRequest) SetEntities(value []MessageEntityClass) {
	e.Flags.Set(3)
	e.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (e *MessagesEditMessageRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !e.Flags.Has(3) {
		return value, false
	}
	return e.Entities, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (e *MessagesEditMessageRequest) SetScheduleDate(value int) {
	e.Flags.Set(15)
	e.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (e *MessagesEditMessageRequest) GetScheduleDate() (value int, ok bool) {
	if !e.Flags.Has(15) {
		return value, false
	}
	return e.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (e *MessagesEditMessageRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editMessage#48f71778 to nil")
	}
	if err := b.ConsumeID(MessagesEditMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editMessage#48f71778: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field flags: %w", err)
		}
	}
	e.NoWebpage = e.Flags.Has(1)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field id: %w", err)
		}
		e.ID = value
	}
	if e.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field message: %w", err)
		}
		e.Message = value
	}
	if e.Flags.Has(14) {
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field media: %w", err)
		}
		e.Media = value
	}
	if e.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	if e.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.editMessage#48f71778: field entities: %w", err)
			}
			e.Entities = append(e.Entities, value)
		}
	}
	if e.Flags.Has(15) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editMessage#48f71778: field schedule_date: %w", err)
		}
		e.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditMessageRequest.
var (
	_ bin.Encoder = &MessagesEditMessageRequest{}
	_ bin.Decoder = &MessagesEditMessageRequest{}
)

// MessagesEditMessage invokes method messages.editMessage#48f71778 returning error if any.
// Edit message
//
// See https://core.telegram.org/method/messages.editMessage for reference.
func (c *Client) MessagesEditMessage(ctx context.Context, request *MessagesEditMessageRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
