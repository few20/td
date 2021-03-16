// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// RPCError represents TL type `rpc_error#2144ca19`.
type RPCError struct {
	// ErrorCode field of RPCError.
	ErrorCode int
	// ErrorMessage field of RPCError.
	ErrorMessage string
}

// RPCErrorTypeID is TL type id of RPCError.
const RPCErrorTypeID = 0x2144ca19

func (r *RPCError) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ErrorCode == 0) {
		return false
	}
	if !(r.ErrorMessage == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCError) String() string {
	if r == nil {
		return "RPCError(nil)"
	}
	type Alias RPCError
	return fmt.Sprintf("RPCError%+v", Alias(*r))
}

// FillFrom fills RPCError from given interface.
func (r *RPCError) FillFrom(from interface {
	GetErrorCode() (value int)
	GetErrorMessage() (value string)
}) {
	r.ErrorCode = from.GetErrorCode()
	r.ErrorMessage = from.GetErrorMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCError) TypeID() uint32 {
	return RPCErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCError) TypeName() string {
	return "rpc_error"
}

// TypeInfo returns info about TL type.
func (r *RPCError) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_error",
		ID:   RPCErrorTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ErrorCode",
			SchemaName: "error_code",
		},
		{
			Name:       "ErrorMessage",
			SchemaName: "error_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCError) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_error#2144ca19 as nil")
	}
	b.PutID(RPCErrorTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RPCError) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_error#2144ca19 as nil")
	}
	b.PutInt(r.ErrorCode)
	b.PutString(r.ErrorMessage)
	return nil
}

// GetErrorCode returns value of ErrorCode field.
func (r *RPCError) GetErrorCode() (value int) {
	return r.ErrorCode
}

// GetErrorMessage returns value of ErrorMessage field.
func (r *RPCError) GetErrorMessage() (value string) {
	return r.ErrorMessage
}

// Decode implements bin.Decoder.
func (r *RPCError) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_error#2144ca19 to nil")
	}
	if err := b.ConsumeID(RPCErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_error#2144ca19: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RPCError) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_error#2144ca19 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_error#2144ca19: field error_code: %w", err)
		}
		r.ErrorCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_error#2144ca19: field error_message: %w", err)
		}
		r.ErrorMessage = value
	}
	return nil
}

// Ensuring interfaces in compile-time for RPCError.
var (
	_ bin.Encoder     = &RPCError{}
	_ bin.Decoder     = &RPCError{}
	_ bin.BareEncoder = &RPCError{}
	_ bin.BareDecoder = &RPCError{}
)
