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

// MessagesDhConfigNotModified represents TL type `messages.dhConfigNotModified#c0e24635`.
// Configuring parameters did not change.
//
// See https://core.telegram.org/constructor/messages.dhConfigNotModified for reference.
type MessagesDhConfigNotModified struct {
	// Random sequence of bytes of assigned length
	Random []byte
}

// MessagesDhConfigNotModifiedTypeID is TL type id of MessagesDhConfigNotModified.
const MessagesDhConfigNotModifiedTypeID = 0xc0e24635

func (d *MessagesDhConfigNotModified) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Random == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDhConfigNotModified) String() string {
	if d == nil {
		return "MessagesDhConfigNotModified(nil)"
	}
	type Alias MessagesDhConfigNotModified
	return fmt.Sprintf("MessagesDhConfigNotModified%+v", Alias(*d))
}

// FillFrom fills MessagesDhConfigNotModified from given interface.
func (d *MessagesDhConfigNotModified) FillFrom(from interface {
	GetRandom() (value []byte)
}) {
	d.Random = from.GetRandom()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDhConfigNotModified) TypeID() uint32 {
	return MessagesDhConfigNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDhConfigNotModified) TypeName() string {
	return "messages.dhConfigNotModified"
}

// TypeInfo returns info about TL type.
func (d *MessagesDhConfigNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.dhConfigNotModified",
		ID:   MessagesDhConfigNotModifiedTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Random",
			SchemaName: "random",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDhConfigNotModified) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dhConfigNotModified#c0e24635 as nil")
	}
	b.PutID(MessagesDhConfigNotModifiedTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDhConfigNotModified) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dhConfigNotModified#c0e24635 as nil")
	}
	b.PutBytes(d.Random)
	return nil
}

// GetRandom returns value of Random field.
func (d *MessagesDhConfigNotModified) GetRandom() (value []byte) {
	return d.Random
}

// Decode implements bin.Decoder.
func (d *MessagesDhConfigNotModified) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dhConfigNotModified#c0e24635 to nil")
	}
	if err := b.ConsumeID(MessagesDhConfigNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dhConfigNotModified#c0e24635: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDhConfigNotModified) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dhConfigNotModified#c0e24635 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfigNotModified#c0e24635: field random: %w", err)
		}
		d.Random = value
	}
	return nil
}

// construct implements constructor of MessagesDhConfigClass.
func (d MessagesDhConfigNotModified) construct() MessagesDhConfigClass { return &d }

// Ensuring interfaces in compile-time for MessagesDhConfigNotModified.
var (
	_ bin.Encoder     = &MessagesDhConfigNotModified{}
	_ bin.Decoder     = &MessagesDhConfigNotModified{}
	_ bin.BareEncoder = &MessagesDhConfigNotModified{}
	_ bin.BareDecoder = &MessagesDhConfigNotModified{}

	_ MessagesDhConfigClass = &MessagesDhConfigNotModified{}
)

// MessagesDhConfig represents TL type `messages.dhConfig#2c221edd`.
// New set of configuring parameters.
//
// See https://core.telegram.org/constructor/messages.dhConfig for reference.
type MessagesDhConfig struct {
	// New value prime, see Wikipedia¹
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
	G int
	// New value primitive root, see Wikipedia¹
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
	P []byte
	// Vestion of set of parameters
	Version int
	// Random sequence of bytes of assigned length
	Random []byte
}

// MessagesDhConfigTypeID is TL type id of MessagesDhConfig.
const MessagesDhConfigTypeID = 0x2c221edd

func (d *MessagesDhConfig) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.G == 0) {
		return false
	}
	if !(d.P == nil) {
		return false
	}
	if !(d.Version == 0) {
		return false
	}
	if !(d.Random == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDhConfig) String() string {
	if d == nil {
		return "MessagesDhConfig(nil)"
	}
	type Alias MessagesDhConfig
	return fmt.Sprintf("MessagesDhConfig%+v", Alias(*d))
}

// FillFrom fills MessagesDhConfig from given interface.
func (d *MessagesDhConfig) FillFrom(from interface {
	GetG() (value int)
	GetP() (value []byte)
	GetVersion() (value int)
	GetRandom() (value []byte)
}) {
	d.G = from.GetG()
	d.P = from.GetP()
	d.Version = from.GetVersion()
	d.Random = from.GetRandom()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDhConfig) TypeID() uint32 {
	return MessagesDhConfigTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDhConfig) TypeName() string {
	return "messages.dhConfig"
}

// TypeInfo returns info about TL type.
func (d *MessagesDhConfig) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.dhConfig",
		ID:   MessagesDhConfigTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "G",
			SchemaName: "g",
		},
		{
			Name:       "P",
			SchemaName: "p",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "Random",
			SchemaName: "random",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDhConfig) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dhConfig#2c221edd as nil")
	}
	b.PutID(MessagesDhConfigTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDhConfig) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.dhConfig#2c221edd as nil")
	}
	b.PutInt(d.G)
	b.PutBytes(d.P)
	b.PutInt(d.Version)
	b.PutBytes(d.Random)
	return nil
}

// GetG returns value of G field.
func (d *MessagesDhConfig) GetG() (value int) {
	return d.G
}

// GetP returns value of P field.
func (d *MessagesDhConfig) GetP() (value []byte) {
	return d.P
}

// GetVersion returns value of Version field.
func (d *MessagesDhConfig) GetVersion() (value int) {
	return d.Version
}

// GetRandom returns value of Random field.
func (d *MessagesDhConfig) GetRandom() (value []byte) {
	return d.Random
}

// Decode implements bin.Decoder.
func (d *MessagesDhConfig) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dhConfig#2c221edd to nil")
	}
	if err := b.ConsumeID(MessagesDhConfigTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDhConfig) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.dhConfig#2c221edd to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field g: %w", err)
		}
		d.G = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field p: %w", err)
		}
		d.P = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field version: %w", err)
		}
		d.Version = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.dhConfig#2c221edd: field random: %w", err)
		}
		d.Random = value
	}
	return nil
}

// construct implements constructor of MessagesDhConfigClass.
func (d MessagesDhConfig) construct() MessagesDhConfigClass { return &d }

// Ensuring interfaces in compile-time for MessagesDhConfig.
var (
	_ bin.Encoder     = &MessagesDhConfig{}
	_ bin.Decoder     = &MessagesDhConfig{}
	_ bin.BareEncoder = &MessagesDhConfig{}
	_ bin.BareDecoder = &MessagesDhConfig{}

	_ MessagesDhConfigClass = &MessagesDhConfig{}
)

// MessagesDhConfigClass represents messages.DhConfig generic type.
//
// See https://core.telegram.org/type/messages.DhConfig for reference.
//
// Example:
//  g, err := tg.DecodeMessagesDhConfig(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesDhConfigNotModified: // messages.dhConfigNotModified#c0e24635
//  case *tg.MessagesDhConfig: // messages.dhConfig#2c221edd
//  default: panic(v)
//  }
type MessagesDhConfigClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesDhConfigClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Random sequence of bytes of assigned length
	GetRandom() (value []byte) // AsModified tries to map MessagesDhConfigClass to MessagesDhConfig.
	AsModified() (*MessagesDhConfig, bool)
}

// AsModified tries to map MessagesDhConfigNotModified to MessagesDhConfig.
func (d *MessagesDhConfigNotModified) AsModified() (*MessagesDhConfig, bool) {
	return nil, false
}

// AsModified tries to map MessagesDhConfig to MessagesDhConfig.
func (d *MessagesDhConfig) AsModified() (*MessagesDhConfig, bool) {
	return d, true
}

// DecodeMessagesDhConfig implements binary de-serialization for MessagesDhConfigClass.
func DecodeMessagesDhConfig(buf *bin.Buffer) (MessagesDhConfigClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesDhConfigNotModifiedTypeID:
		// Decoding messages.dhConfigNotModified#c0e24635.
		v := MessagesDhConfigNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDhConfigClass: %w", err)
		}
		return &v, nil
	case MessagesDhConfigTypeID:
		// Decoding messages.dhConfig#2c221edd.
		v := MessagesDhConfig{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesDhConfigClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesDhConfigClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesDhConfig boxes the MessagesDhConfigClass providing a helper.
type MessagesDhConfigBox struct {
	DhConfig MessagesDhConfigClass
}

// Decode implements bin.Decoder for MessagesDhConfigBox.
func (b *MessagesDhConfigBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesDhConfigBox to nil")
	}
	v, err := DecodeMessagesDhConfig(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DhConfig = v
	return nil
}

// Encode implements bin.Encode for MessagesDhConfigBox.
func (b *MessagesDhConfigBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DhConfig == nil {
		return fmt.Errorf("unable to encode MessagesDhConfigClass as nil")
	}
	return b.DhConfig.Encode(buf)
}

// MessagesDhConfigClassArray is adapter for slice of MessagesDhConfigClass.
type MessagesDhConfigClassArray []MessagesDhConfigClass

// Sort sorts slice of MessagesDhConfigClass.
func (s MessagesDhConfigClassArray) Sort(less func(a, b MessagesDhConfigClass) bool) MessagesDhConfigClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDhConfigClass.
func (s MessagesDhConfigClassArray) SortStable(less func(a, b MessagesDhConfigClass) bool) MessagesDhConfigClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDhConfigClass.
func (s MessagesDhConfigClassArray) Retain(keep func(x MessagesDhConfigClass) bool) MessagesDhConfigClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesDhConfigClassArray) First() (v MessagesDhConfigClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDhConfigClassArray) Last() (v MessagesDhConfigClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDhConfigClassArray) PopFirst() (v MessagesDhConfigClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDhConfigClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDhConfigClassArray) Pop() (v MessagesDhConfigClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesDhConfigNotModified returns copy with only MessagesDhConfigNotModified constructors.
func (s MessagesDhConfigClassArray) AsMessagesDhConfigNotModified() (to MessagesDhConfigNotModifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesDhConfigNotModified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesDhConfig returns copy with only MessagesDhConfig constructors.
func (s MessagesDhConfigClassArray) AsMessagesDhConfig() (to MessagesDhConfigArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesDhConfig)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesDhConfigClassArray) AppendOnlyModified(to []*MessagesDhConfig) []*MessagesDhConfig {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesDhConfigClassArray) AsModified() (to []*MessagesDhConfig) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesDhConfigClassArray) FirstAsModified() (v *MessagesDhConfig, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesDhConfigClassArray) LastAsModified() (v *MessagesDhConfig, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesDhConfigClassArray) PopFirstAsModified() (v *MessagesDhConfig, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesDhConfigClassArray) PopAsModified() (v *MessagesDhConfig, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesDhConfigNotModifiedArray is adapter for slice of MessagesDhConfigNotModified.
type MessagesDhConfigNotModifiedArray []MessagesDhConfigNotModified

// Sort sorts slice of MessagesDhConfigNotModified.
func (s MessagesDhConfigNotModifiedArray) Sort(less func(a, b MessagesDhConfigNotModified) bool) MessagesDhConfigNotModifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDhConfigNotModified.
func (s MessagesDhConfigNotModifiedArray) SortStable(less func(a, b MessagesDhConfigNotModified) bool) MessagesDhConfigNotModifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDhConfigNotModified.
func (s MessagesDhConfigNotModifiedArray) Retain(keep func(x MessagesDhConfigNotModified) bool) MessagesDhConfigNotModifiedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesDhConfigNotModifiedArray) First() (v MessagesDhConfigNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDhConfigNotModifiedArray) Last() (v MessagesDhConfigNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDhConfigNotModifiedArray) PopFirst() (v MessagesDhConfigNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDhConfigNotModified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDhConfigNotModifiedArray) Pop() (v MessagesDhConfigNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessagesDhConfigArray is adapter for slice of MessagesDhConfig.
type MessagesDhConfigArray []MessagesDhConfig

// Sort sorts slice of MessagesDhConfig.
func (s MessagesDhConfigArray) Sort(less func(a, b MessagesDhConfig) bool) MessagesDhConfigArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesDhConfig.
func (s MessagesDhConfigArray) SortStable(less func(a, b MessagesDhConfig) bool) MessagesDhConfigArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesDhConfig.
func (s MessagesDhConfigArray) Retain(keep func(x MessagesDhConfig) bool) MessagesDhConfigArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesDhConfigArray) First() (v MessagesDhConfig, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesDhConfigArray) Last() (v MessagesDhConfig, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesDhConfigArray) PopFirst() (v MessagesDhConfig, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesDhConfig
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesDhConfigArray) Pop() (v MessagesDhConfig, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
