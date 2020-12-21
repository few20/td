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

// InputCheckPasswordEmpty represents TL type `inputCheckPasswordEmpty#9880f658`.
// There is no password
//
// See https://core.telegram.org/constructor/inputCheckPasswordEmpty for reference.
type InputCheckPasswordEmpty struct {
}

// InputCheckPasswordEmptyTypeID is TL type id of InputCheckPasswordEmpty.
const InputCheckPasswordEmptyTypeID = 0x9880f658

// String implements fmt.Stringer.
func (i *InputCheckPasswordEmpty) String() string {
	if i == nil {
		return "InputCheckPasswordEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputCheckPasswordEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputCheckPasswordEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCheckPasswordEmpty#9880f658 as nil")
	}
	b.PutID(InputCheckPasswordEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCheckPasswordEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCheckPasswordEmpty#9880f658 to nil")
	}
	if err := b.ConsumeID(InputCheckPasswordEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCheckPasswordEmpty#9880f658: %w", err)
	}
	return nil
}

// construct implements constructor of InputCheckPasswordSRPClass.
func (i InputCheckPasswordEmpty) construct() InputCheckPasswordSRPClass { return &i }

// Ensuring interfaces in compile-time for InputCheckPasswordEmpty.
var (
	_ bin.Encoder = &InputCheckPasswordEmpty{}
	_ bin.Decoder = &InputCheckPasswordEmpty{}

	_ InputCheckPasswordSRPClass = &InputCheckPasswordEmpty{}
)

// InputCheckPasswordSRP represents TL type `inputCheckPasswordSRP#d27ff082`.
// Constructor for checking the validity of a 2FA SRP password (see SRP)
//
// See https://core.telegram.org/constructor/inputCheckPasswordSRP for reference.
type InputCheckPasswordSRP struct {
	// SRP ID
	SrpID int64
	// A parameter (see SRP)
	A []byte
	// M1 parameter (see SRP)
	M1 []byte
}

// InputCheckPasswordSRPTypeID is TL type id of InputCheckPasswordSRP.
const InputCheckPasswordSRPTypeID = 0xd27ff082

// String implements fmt.Stringer.
func (i *InputCheckPasswordSRP) String() string {
	if i == nil {
		return "InputCheckPasswordSRP(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputCheckPasswordSRP")
	sb.WriteString("{\n")
	sb.WriteString("\tSrpID: ")
	sb.WriteString(fmt.Sprint(i.SrpID))
	sb.WriteString(",\n")
	sb.WriteString("\tA: ")
	sb.WriteString(fmt.Sprint(i.A))
	sb.WriteString(",\n")
	sb.WriteString("\tM1: ")
	sb.WriteString(fmt.Sprint(i.M1))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputCheckPasswordSRP) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputCheckPasswordSRP#d27ff082 as nil")
	}
	b.PutID(InputCheckPasswordSRPTypeID)
	b.PutLong(i.SrpID)
	b.PutBytes(i.A)
	b.PutBytes(i.M1)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputCheckPasswordSRP) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputCheckPasswordSRP#d27ff082 to nil")
	}
	if err := b.ConsumeID(InputCheckPasswordSRPTypeID); err != nil {
		return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: field srp_id: %w", err)
		}
		i.SrpID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: field A: %w", err)
		}
		i.A = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputCheckPasswordSRP#d27ff082: field M1: %w", err)
		}
		i.M1 = value
	}
	return nil
}

// construct implements constructor of InputCheckPasswordSRPClass.
func (i InputCheckPasswordSRP) construct() InputCheckPasswordSRPClass { return &i }

// Ensuring interfaces in compile-time for InputCheckPasswordSRP.
var (
	_ bin.Encoder = &InputCheckPasswordSRP{}
	_ bin.Decoder = &InputCheckPasswordSRP{}

	_ InputCheckPasswordSRPClass = &InputCheckPasswordSRP{}
)

// InputCheckPasswordSRPClass represents InputCheckPasswordSRP generic type.
//
// See https://core.telegram.org/type/InputCheckPasswordSRP for reference.
//
// Example:
//  g, err := DecodeInputCheckPasswordSRP(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputCheckPasswordEmpty: // inputCheckPasswordEmpty#9880f658
//  case *InputCheckPasswordSRP: // inputCheckPasswordSRP#d27ff082
//  default: panic(v)
//  }
type InputCheckPasswordSRPClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputCheckPasswordSRPClass
	fmt.Stringer
}

// DecodeInputCheckPasswordSRP implements binary de-serialization for InputCheckPasswordSRPClass.
func DecodeInputCheckPasswordSRP(buf *bin.Buffer) (InputCheckPasswordSRPClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputCheckPasswordEmptyTypeID:
		// Decoding inputCheckPasswordEmpty#9880f658.
		v := InputCheckPasswordEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCheckPasswordSRPClass: %w", err)
		}
		return &v, nil
	case InputCheckPasswordSRPTypeID:
		// Decoding inputCheckPasswordSRP#d27ff082.
		v := InputCheckPasswordSRP{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputCheckPasswordSRPClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputCheckPasswordSRPClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputCheckPasswordSRP boxes the InputCheckPasswordSRPClass providing a helper.
type InputCheckPasswordSRPBox struct {
	InputCheckPasswordSRP InputCheckPasswordSRPClass
}

// Decode implements bin.Decoder for InputCheckPasswordSRPBox.
func (b *InputCheckPasswordSRPBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputCheckPasswordSRPBox to nil")
	}
	v, err := DecodeInputCheckPasswordSRP(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputCheckPasswordSRP = v
	return nil
}

// Encode implements bin.Encode for InputCheckPasswordSRPBox.
func (b *InputCheckPasswordSRPBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputCheckPasswordSRP == nil {
		return fmt.Errorf("unable to encode InputCheckPasswordSRPClass as nil")
	}
	return b.InputCheckPasswordSRP.Encode(buf)
}
