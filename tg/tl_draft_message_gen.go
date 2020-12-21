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

// DraftMessageEmpty represents TL type `draftMessageEmpty#1b0c841a`.
// Empty draft
//
// See https://core.telegram.org/constructor/draftMessageEmpty for reference.
type DraftMessageEmpty struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// When was the draft last updated
	//
	// Use SetDate and GetDate helpers.
	Date int
}

// DraftMessageEmptyTypeID is TL type id of DraftMessageEmpty.
const DraftMessageEmptyTypeID = 0x1b0c841a

// String implements fmt.Stringer.
func (d *DraftMessageEmpty) String() string {
	if d == nil {
		return "DraftMessageEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DraftMessageEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(d.Flags.String())
	sb.WriteString(",\n")
	if d.Flags.Has(0) {
		sb.WriteString("\tDate: ")
		sb.WriteString(fmt.Sprint(d.Date))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DraftMessageEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode draftMessageEmpty#1b0c841a as nil")
	}
	b.PutID(DraftMessageEmptyTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode draftMessageEmpty#1b0c841a: field flags: %w", err)
	}
	if d.Flags.Has(0) {
		b.PutInt(d.Date)
	}
	return nil
}

// SetDate sets value of Date conditional field.
func (d *DraftMessageEmpty) SetDate(value int) {
	d.Flags.Set(0)
	d.Date = value
}

// GetDate returns value of Date conditional field and
// boolean which is true if field was set.
func (d *DraftMessageEmpty) GetDate() (value int, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Date, true
}

// Decode implements bin.Decoder.
func (d *DraftMessageEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode draftMessageEmpty#1b0c841a to nil")
	}
	if err := b.ConsumeID(DraftMessageEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode draftMessageEmpty#1b0c841a: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode draftMessageEmpty#1b0c841a: field flags: %w", err)
		}
	}
	if d.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessageEmpty#1b0c841a: field date: %w", err)
		}
		d.Date = value
	}
	return nil
}

// construct implements constructor of DraftMessageClass.
func (d DraftMessageEmpty) construct() DraftMessageClass { return &d }

// Ensuring interfaces in compile-time for DraftMessageEmpty.
var (
	_ bin.Encoder = &DraftMessageEmpty{}
	_ bin.Decoder = &DraftMessageEmpty{}

	_ DraftMessageClass = &DraftMessageEmpty{}
)

// DraftMessage represents TL type `draftMessage#fd8e711f`.
// Represents a message draft.
//
// See https://core.telegram.org/constructor/draftMessage for reference.
type DraftMessage struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether no webpage preview will be generated
	NoWebpage bool
	// The message this message will reply to
	//
	// Use SetReplyToMsgID and GetReplyToMsgID helpers.
	ReplyToMsgID int
	// The draft
	Message string
	// Message entities for styled text.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Date of last update of the draft.
	Date int
}

// DraftMessageTypeID is TL type id of DraftMessage.
const DraftMessageTypeID = 0xfd8e711f

// String implements fmt.Stringer.
func (d *DraftMessage) String() string {
	if d == nil {
		return "DraftMessage(nil)"
	}
	var sb strings.Builder
	sb.WriteString("DraftMessage")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(d.Flags.String())
	sb.WriteString(",\n")
	if d.Flags.Has(0) {
		sb.WriteString("\tReplyToMsgID: ")
		sb.WriteString(fmt.Sprint(d.ReplyToMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(d.Message))
	sb.WriteString(",\n")
	if d.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range d.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(d.Date))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *DraftMessage) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode draftMessage#fd8e711f as nil")
	}
	b.PutID(DraftMessageTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode draftMessage#fd8e711f: field flags: %w", err)
	}
	if d.Flags.Has(0) {
		b.PutInt(d.ReplyToMsgID)
	}
	b.PutString(d.Message)
	if d.Flags.Has(3) {
		b.PutVectorHeader(len(d.Entities))
		for idx, v := range d.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode draftMessage#fd8e711f: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode draftMessage#fd8e711f: field entities element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(d.Date)
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (d *DraftMessage) SetNoWebpage(value bool) {
	if value {
		d.Flags.Set(1)
	} else {
		d.Flags.Unset(1)
	}
}

// SetReplyToMsgID sets value of ReplyToMsgID conditional field.
func (d *DraftMessage) SetReplyToMsgID(value int) {
	d.Flags.Set(0)
	d.ReplyToMsgID = value
}

// GetReplyToMsgID returns value of ReplyToMsgID conditional field and
// boolean which is true if field was set.
func (d *DraftMessage) GetReplyToMsgID() (value int, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.ReplyToMsgID, true
}

// SetEntities sets value of Entities conditional field.
func (d *DraftMessage) SetEntities(value []MessageEntityClass) {
	d.Flags.Set(3)
	d.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (d *DraftMessage) GetEntities() (value []MessageEntityClass, ok bool) {
	if !d.Flags.Has(3) {
		return value, false
	}
	return d.Entities, true
}

// Decode implements bin.Decoder.
func (d *DraftMessage) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode draftMessage#fd8e711f to nil")
	}
	if err := b.ConsumeID(DraftMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode draftMessage#fd8e711f: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode draftMessage#fd8e711f: field flags: %w", err)
		}
	}
	d.NoWebpage = d.Flags.Has(1)
	if d.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#fd8e711f: field reply_to_msg_id: %w", err)
		}
		d.ReplyToMsgID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#fd8e711f: field message: %w", err)
		}
		d.Message = value
	}
	if d.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#fd8e711f: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode draftMessage#fd8e711f: field entities: %w", err)
			}
			d.Entities = append(d.Entities, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode draftMessage#fd8e711f: field date: %w", err)
		}
		d.Date = value
	}
	return nil
}

// construct implements constructor of DraftMessageClass.
func (d DraftMessage) construct() DraftMessageClass { return &d }

// Ensuring interfaces in compile-time for DraftMessage.
var (
	_ bin.Encoder = &DraftMessage{}
	_ bin.Decoder = &DraftMessage{}

	_ DraftMessageClass = &DraftMessage{}
)

// DraftMessageClass represents DraftMessage generic type.
//
// See https://core.telegram.org/type/DraftMessage for reference.
//
// Example:
//  g, err := DecodeDraftMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *DraftMessageEmpty: // draftMessageEmpty#1b0c841a
//  case *DraftMessage: // draftMessage#fd8e711f
//  default: panic(v)
//  }
type DraftMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() DraftMessageClass
	fmt.Stringer
}

// DecodeDraftMessage implements binary de-serialization for DraftMessageClass.
func DecodeDraftMessage(buf *bin.Buffer) (DraftMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DraftMessageEmptyTypeID:
		// Decoding draftMessageEmpty#1b0c841a.
		v := DraftMessageEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DraftMessageClass: %w", err)
		}
		return &v, nil
	case DraftMessageTypeID:
		// Decoding draftMessage#fd8e711f.
		v := DraftMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DraftMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DraftMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// DraftMessage boxes the DraftMessageClass providing a helper.
type DraftMessageBox struct {
	DraftMessage DraftMessageClass
}

// Decode implements bin.Decoder for DraftMessageBox.
func (b *DraftMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DraftMessageBox to nil")
	}
	v, err := DecodeDraftMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DraftMessage = v
	return nil
}

// Encode implements bin.Encode for DraftMessageBox.
func (b *DraftMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DraftMessage == nil {
		return fmt.Errorf("unable to encode DraftMessageClass as nil")
	}
	return b.DraftMessage.Encode(buf)
}
