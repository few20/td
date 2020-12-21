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

// MessagesSendMediaRequest represents TL type `messages.sendMedia#3491eba9`.
// Send a media
//
// See https://core.telegram.org/method/messages.sendMedia for reference.
type MessagesSendMediaRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Send message silently (no notification should be triggered)
	Silent bool
	// Send message in background
	Background bool
	// Clear the draft
	ClearDraft bool
	// Destination
	Peer InputPeerClass
	// Message ID to which this message should reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// Attached media
	Media InputMediaClass
	// Caption
	Message string
	// Random ID to avoid resending the same message
	RandomID int64
	// Reply markup for bot keyboards
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

// MessagesSendMediaRequestTypeID is TL type id of MessagesSendMediaRequest.
const MessagesSendMediaRequestTypeID = 0x3491eba9

// String implements fmt.Stringer.
func (s *MessagesSendMediaRequest) String() string {
	if s == nil {
		return "MessagesSendMediaRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSendMediaRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(s.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(s.Peer.String())
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tReplyToMsgID: ")
		sb.WriteString(fmt.Sprint(s.ReplyToMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tMedia: ")
	sb.WriteString(s.Media.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(s.Message))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(s.RandomID))
	sb.WriteString(",\n")
	if s.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(s.ReplyMarkup.String())
		sb.WriteString(",\n")
	}
	if s.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range s.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if s.Flags.Has(10) {
		sb.WriteString("\tScheduleDate: ")
		sb.WriteString(fmt.Sprint(s.ScheduleDate))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSendMediaRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendMedia#3491eba9 as nil")
	}
	b.PutID(MessagesSendMediaRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReplyToMsgID)
	}
	if s.Media == nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field media is nil")
	}
	if err := s.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field media: %w", err)
	}
	b.PutString(s.Message)
	b.PutLong(s.RandomID)
	if s.Flags.Has(2) {
		if s.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field reply_markup is nil")
		}
		if err := s.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field reply_markup: %w", err)
		}
	}
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.sendMedia#3491eba9: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(10) {
		b.PutInt(s.ScheduleDate)
	}
	return nil
}

// SetSilent sets value of Silent conditional field.
func (s *MessagesSendMediaRequest) SetSilent(value bool) {
	if value {
		s.Flags.Set(5)
	} else {
		s.Flags.Unset(5)
	}
}

// SetBackground sets value of Background conditional field.
func (s *MessagesSendMediaRequest) SetBackground(value bool) {
	if value {
		s.Flags.Set(6)
	} else {
		s.Flags.Unset(6)
	}
}

// SetClearDraft sets value of ClearDraft conditional field.
func (s *MessagesSendMediaRequest) SetClearDraft(value bool) {
	if value {
		s.Flags.Set(7)
	} else {
		s.Flags.Unset(7)
	}
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (s *MessagesSendMediaRequest) SetReplyToMsgID(value int) {
	s.Flags.Set(0)
	s.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetReplyToMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReplyToMsgID, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (s *MessagesSendMediaRequest) SetReplyMarkup(value ReplyMarkupClass) {
	s.Flags.Set(2)
	s.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (s *MessagesSendMediaRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(3)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Entities, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (s *MessagesSendMediaRequest) SetScheduleDate(value int) {
	s.Flags.Set(10)
	s.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (s *MessagesSendMediaRequest) GetScheduleDate() (value int, ok bool) {
	if !s.Flags.Has(10) {
		return value, false
	}
	return s.ScheduleDate, true
}

// Decode implements bin.Decoder.
func (s *MessagesSendMediaRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendMedia#3491eba9 to nil")
	}
	if err := b.ConsumeID(MessagesSendMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field flags: %w", err)
		}
	}
	s.Silent = s.Flags.Has(5)
	s.Background = s.Flags.Has(6)
	s.ClearDraft = s.Flags.Has(7)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field media: %w", err)
		}
		s.Media = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field message: %w", err)
		}
		s.Message = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field random_id: %w", err)
		}
		s.RandomID = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field reply_markup: %w", err)
		}
		s.ReplyMarkup = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendMedia#3491eba9: field schedule_date: %w", err)
		}
		s.ScheduleDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendMediaRequest.
var (
	_ bin.Encoder = &MessagesSendMediaRequest{}
	_ bin.Decoder = &MessagesSendMediaRequest{}
)

// MessagesSendMedia invokes method messages.sendMedia#3491eba9 returning error if any.
// Send a media
//
// See https://core.telegram.org/method/messages.sendMedia for reference.
func (c *Client) MessagesSendMedia(ctx context.Context, request *MessagesSendMediaRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
