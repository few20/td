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

// HelpTermsOfServiceUpdateEmpty represents TL type `help.termsOfServiceUpdateEmpty#e3309f7f`.
// No changes were made to telegram's terms of service
//
// See https://core.telegram.org/constructor/help.termsOfServiceUpdateEmpty for reference.
type HelpTermsOfServiceUpdateEmpty struct {
	// New TOS updates will have to be queried using help.getTermsOfServiceUpdate¹ in expires seconds
	//
	// Links:
	//  1) https://core.telegram.org/method/help.getTermsOfServiceUpdate
	Expires int
}

// HelpTermsOfServiceUpdateEmptyTypeID is TL type id of HelpTermsOfServiceUpdateEmpty.
const HelpTermsOfServiceUpdateEmptyTypeID = 0xe3309f7f

func (t *HelpTermsOfServiceUpdateEmpty) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Expires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *HelpTermsOfServiceUpdateEmpty) String() string {
	if t == nil {
		return "HelpTermsOfServiceUpdateEmpty(nil)"
	}
	type Alias HelpTermsOfServiceUpdateEmpty
	return fmt.Sprintf("HelpTermsOfServiceUpdateEmpty%+v", Alias(*t))
}

// FillFrom fills HelpTermsOfServiceUpdateEmpty from given interface.
func (t *HelpTermsOfServiceUpdateEmpty) FillFrom(from interface {
	GetExpires() (value int)
}) {
	t.Expires = from.GetExpires()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpTermsOfServiceUpdateEmpty) TypeID() uint32 {
	return HelpTermsOfServiceUpdateEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpTermsOfServiceUpdateEmpty) TypeName() string {
	return "help.termsOfServiceUpdateEmpty"
}

// TypeInfo returns info about TL type.
func (t *HelpTermsOfServiceUpdateEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.termsOfServiceUpdateEmpty",
		ID:   HelpTermsOfServiceUpdateEmptyTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *HelpTermsOfServiceUpdateEmpty) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfServiceUpdateEmpty#e3309f7f as nil")
	}
	b.PutID(HelpTermsOfServiceUpdateEmptyTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *HelpTermsOfServiceUpdateEmpty) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfServiceUpdateEmpty#e3309f7f as nil")
	}
	b.PutInt(t.Expires)
	return nil
}

// GetExpires returns value of Expires field.
func (t *HelpTermsOfServiceUpdateEmpty) GetExpires() (value int) {
	return t.Expires
}

// Decode implements bin.Decoder.
func (t *HelpTermsOfServiceUpdateEmpty) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfServiceUpdateEmpty#e3309f7f to nil")
	}
	if err := b.ConsumeID(HelpTermsOfServiceUpdateEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.termsOfServiceUpdateEmpty#e3309f7f: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *HelpTermsOfServiceUpdateEmpty) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfServiceUpdateEmpty#e3309f7f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfServiceUpdateEmpty#e3309f7f: field expires: %w", err)
		}
		t.Expires = value
	}
	return nil
}

// construct implements constructor of HelpTermsOfServiceUpdateClass.
func (t HelpTermsOfServiceUpdateEmpty) construct() HelpTermsOfServiceUpdateClass { return &t }

// Ensuring interfaces in compile-time for HelpTermsOfServiceUpdateEmpty.
var (
	_ bin.Encoder     = &HelpTermsOfServiceUpdateEmpty{}
	_ bin.Decoder     = &HelpTermsOfServiceUpdateEmpty{}
	_ bin.BareEncoder = &HelpTermsOfServiceUpdateEmpty{}
	_ bin.BareDecoder = &HelpTermsOfServiceUpdateEmpty{}

	_ HelpTermsOfServiceUpdateClass = &HelpTermsOfServiceUpdateEmpty{}
)

// HelpTermsOfServiceUpdate represents TL type `help.termsOfServiceUpdate#28ecf961`.
// Info about an update of telegram's terms of service. If the terms of service are declined, then the account.deleteAccount¹ method should be called with the reason "Decline ToS update"
//
// Links:
//  1) https://core.telegram.org/method/account.deleteAccount
//
// See https://core.telegram.org/constructor/help.termsOfServiceUpdate for reference.
type HelpTermsOfServiceUpdate struct {
	// New TOS updates will have to be queried using help.getTermsOfServiceUpdate¹ in expires seconds
	//
	// Links:
	//  1) https://core.telegram.org/method/help.getTermsOfServiceUpdate
	Expires int
	// New terms of service
	TermsOfService HelpTermsOfService
}

// HelpTermsOfServiceUpdateTypeID is TL type id of HelpTermsOfServiceUpdate.
const HelpTermsOfServiceUpdateTypeID = 0x28ecf961

func (t *HelpTermsOfServiceUpdate) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Expires == 0) {
		return false
	}
	if !(t.TermsOfService.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *HelpTermsOfServiceUpdate) String() string {
	if t == nil {
		return "HelpTermsOfServiceUpdate(nil)"
	}
	type Alias HelpTermsOfServiceUpdate
	return fmt.Sprintf("HelpTermsOfServiceUpdate%+v", Alias(*t))
}

// FillFrom fills HelpTermsOfServiceUpdate from given interface.
func (t *HelpTermsOfServiceUpdate) FillFrom(from interface {
	GetExpires() (value int)
	GetTermsOfService() (value HelpTermsOfService)
}) {
	t.Expires = from.GetExpires()
	t.TermsOfService = from.GetTermsOfService()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpTermsOfServiceUpdate) TypeID() uint32 {
	return HelpTermsOfServiceUpdateTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpTermsOfServiceUpdate) TypeName() string {
	return "help.termsOfServiceUpdate"
}

// TypeInfo returns info about TL type.
func (t *HelpTermsOfServiceUpdate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.termsOfServiceUpdate",
		ID:   HelpTermsOfServiceUpdateTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
		{
			Name:       "TermsOfService",
			SchemaName: "terms_of_service",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *HelpTermsOfServiceUpdate) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfServiceUpdate#28ecf961 as nil")
	}
	b.PutID(HelpTermsOfServiceUpdateTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *HelpTermsOfServiceUpdate) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode help.termsOfServiceUpdate#28ecf961 as nil")
	}
	b.PutInt(t.Expires)
	if err := t.TermsOfService.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.termsOfServiceUpdate#28ecf961: field terms_of_service: %w", err)
	}
	return nil
}

// GetExpires returns value of Expires field.
func (t *HelpTermsOfServiceUpdate) GetExpires() (value int) {
	return t.Expires
}

// GetTermsOfService returns value of TermsOfService field.
func (t *HelpTermsOfServiceUpdate) GetTermsOfService() (value HelpTermsOfService) {
	return t.TermsOfService
}

// Decode implements bin.Decoder.
func (t *HelpTermsOfServiceUpdate) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfServiceUpdate#28ecf961 to nil")
	}
	if err := b.ConsumeID(HelpTermsOfServiceUpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode help.termsOfServiceUpdate#28ecf961: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *HelpTermsOfServiceUpdate) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode help.termsOfServiceUpdate#28ecf961 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.termsOfServiceUpdate#28ecf961: field expires: %w", err)
		}
		t.Expires = value
	}
	{
		if err := t.TermsOfService.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.termsOfServiceUpdate#28ecf961: field terms_of_service: %w", err)
		}
	}
	return nil
}

// construct implements constructor of HelpTermsOfServiceUpdateClass.
func (t HelpTermsOfServiceUpdate) construct() HelpTermsOfServiceUpdateClass { return &t }

// Ensuring interfaces in compile-time for HelpTermsOfServiceUpdate.
var (
	_ bin.Encoder     = &HelpTermsOfServiceUpdate{}
	_ bin.Decoder     = &HelpTermsOfServiceUpdate{}
	_ bin.BareEncoder = &HelpTermsOfServiceUpdate{}
	_ bin.BareDecoder = &HelpTermsOfServiceUpdate{}

	_ HelpTermsOfServiceUpdateClass = &HelpTermsOfServiceUpdate{}
)

// HelpTermsOfServiceUpdateClass represents help.TermsOfServiceUpdate generic type.
//
// See https://core.telegram.org/type/help.TermsOfServiceUpdate for reference.
//
// Example:
//  g, err := tg.DecodeHelpTermsOfServiceUpdate(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpTermsOfServiceUpdateEmpty: // help.termsOfServiceUpdateEmpty#e3309f7f
//  case *tg.HelpTermsOfServiceUpdate: // help.termsOfServiceUpdate#28ecf961
//  default: panic(v)
//  }
type HelpTermsOfServiceUpdateClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() HelpTermsOfServiceUpdateClass

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

	// New TOS updates will have to be queried using help.getTermsOfServiceUpdate¹ in expires seconds
	//
	// Links:
	//  1) https://core.telegram.org/method/help.getTermsOfServiceUpdate
	GetExpires() (value int)
	// AsNotEmpty tries to map HelpTermsOfServiceUpdateClass to HelpTermsOfServiceUpdate.
	AsNotEmpty() (*HelpTermsOfServiceUpdate, bool)
}

// AsNotEmpty tries to map HelpTermsOfServiceUpdateEmpty to HelpTermsOfServiceUpdate.
func (t *HelpTermsOfServiceUpdateEmpty) AsNotEmpty() (*HelpTermsOfServiceUpdate, bool) {
	return nil, false
}

// AsNotEmpty tries to map HelpTermsOfServiceUpdate to HelpTermsOfServiceUpdate.
func (t *HelpTermsOfServiceUpdate) AsNotEmpty() (*HelpTermsOfServiceUpdate, bool) {
	return t, true
}

// DecodeHelpTermsOfServiceUpdate implements binary de-serialization for HelpTermsOfServiceUpdateClass.
func DecodeHelpTermsOfServiceUpdate(buf *bin.Buffer) (HelpTermsOfServiceUpdateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpTermsOfServiceUpdateEmptyTypeID:
		// Decoding help.termsOfServiceUpdateEmpty#e3309f7f.
		v := HelpTermsOfServiceUpdateEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpTermsOfServiceUpdateClass: %w", err)
		}
		return &v, nil
	case HelpTermsOfServiceUpdateTypeID:
		// Decoding help.termsOfServiceUpdate#28ecf961.
		v := HelpTermsOfServiceUpdate{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpTermsOfServiceUpdateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpTermsOfServiceUpdateClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpTermsOfServiceUpdate boxes the HelpTermsOfServiceUpdateClass providing a helper.
type HelpTermsOfServiceUpdateBox struct {
	TermsOfServiceUpdate HelpTermsOfServiceUpdateClass
}

// Decode implements bin.Decoder for HelpTermsOfServiceUpdateBox.
func (b *HelpTermsOfServiceUpdateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpTermsOfServiceUpdateBox to nil")
	}
	v, err := DecodeHelpTermsOfServiceUpdate(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.TermsOfServiceUpdate = v
	return nil
}

// Encode implements bin.Encode for HelpTermsOfServiceUpdateBox.
func (b *HelpTermsOfServiceUpdateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.TermsOfServiceUpdate == nil {
		return fmt.Errorf("unable to encode HelpTermsOfServiceUpdateClass as nil")
	}
	return b.TermsOfServiceUpdate.Encode(buf)
}

// HelpTermsOfServiceUpdateClassArray is adapter for slice of HelpTermsOfServiceUpdateClass.
type HelpTermsOfServiceUpdateClassArray []HelpTermsOfServiceUpdateClass

// Sort sorts slice of HelpTermsOfServiceUpdateClass.
func (s HelpTermsOfServiceUpdateClassArray) Sort(less func(a, b HelpTermsOfServiceUpdateClass) bool) HelpTermsOfServiceUpdateClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpTermsOfServiceUpdateClass.
func (s HelpTermsOfServiceUpdateClassArray) SortStable(less func(a, b HelpTermsOfServiceUpdateClass) bool) HelpTermsOfServiceUpdateClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpTermsOfServiceUpdateClass.
func (s HelpTermsOfServiceUpdateClassArray) Retain(keep func(x HelpTermsOfServiceUpdateClass) bool) HelpTermsOfServiceUpdateClassArray {
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
func (s HelpTermsOfServiceUpdateClassArray) First() (v HelpTermsOfServiceUpdateClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpTermsOfServiceUpdateClassArray) Last() (v HelpTermsOfServiceUpdateClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpTermsOfServiceUpdateClassArray) PopFirst() (v HelpTermsOfServiceUpdateClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpTermsOfServiceUpdateClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpTermsOfServiceUpdateClassArray) Pop() (v HelpTermsOfServiceUpdateClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpTermsOfServiceUpdateEmpty returns copy with only HelpTermsOfServiceUpdateEmpty constructors.
func (s HelpTermsOfServiceUpdateClassArray) AsHelpTermsOfServiceUpdateEmpty() (to HelpTermsOfServiceUpdateEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpTermsOfServiceUpdateEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsHelpTermsOfServiceUpdate returns copy with only HelpTermsOfServiceUpdate constructors.
func (s HelpTermsOfServiceUpdateClassArray) AsHelpTermsOfServiceUpdate() (to HelpTermsOfServiceUpdateArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpTermsOfServiceUpdate)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s HelpTermsOfServiceUpdateClassArray) AppendOnlyNotEmpty(to []*HelpTermsOfServiceUpdate) []*HelpTermsOfServiceUpdate {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s HelpTermsOfServiceUpdateClassArray) AsNotEmpty() (to []*HelpTermsOfServiceUpdate) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s HelpTermsOfServiceUpdateClassArray) FirstAsNotEmpty() (v *HelpTermsOfServiceUpdate, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s HelpTermsOfServiceUpdateClassArray) LastAsNotEmpty() (v *HelpTermsOfServiceUpdate, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *HelpTermsOfServiceUpdateClassArray) PopFirstAsNotEmpty() (v *HelpTermsOfServiceUpdate, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *HelpTermsOfServiceUpdateClassArray) PopAsNotEmpty() (v *HelpTermsOfServiceUpdate, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// HelpTermsOfServiceUpdateEmptyArray is adapter for slice of HelpTermsOfServiceUpdateEmpty.
type HelpTermsOfServiceUpdateEmptyArray []HelpTermsOfServiceUpdateEmpty

// Sort sorts slice of HelpTermsOfServiceUpdateEmpty.
func (s HelpTermsOfServiceUpdateEmptyArray) Sort(less func(a, b HelpTermsOfServiceUpdateEmpty) bool) HelpTermsOfServiceUpdateEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpTermsOfServiceUpdateEmpty.
func (s HelpTermsOfServiceUpdateEmptyArray) SortStable(less func(a, b HelpTermsOfServiceUpdateEmpty) bool) HelpTermsOfServiceUpdateEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpTermsOfServiceUpdateEmpty.
func (s HelpTermsOfServiceUpdateEmptyArray) Retain(keep func(x HelpTermsOfServiceUpdateEmpty) bool) HelpTermsOfServiceUpdateEmptyArray {
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
func (s HelpTermsOfServiceUpdateEmptyArray) First() (v HelpTermsOfServiceUpdateEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpTermsOfServiceUpdateEmptyArray) Last() (v HelpTermsOfServiceUpdateEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpTermsOfServiceUpdateEmptyArray) PopFirst() (v HelpTermsOfServiceUpdateEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpTermsOfServiceUpdateEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpTermsOfServiceUpdateEmptyArray) Pop() (v HelpTermsOfServiceUpdateEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// HelpTermsOfServiceUpdateArray is adapter for slice of HelpTermsOfServiceUpdate.
type HelpTermsOfServiceUpdateArray []HelpTermsOfServiceUpdate

// Sort sorts slice of HelpTermsOfServiceUpdate.
func (s HelpTermsOfServiceUpdateArray) Sort(less func(a, b HelpTermsOfServiceUpdate) bool) HelpTermsOfServiceUpdateArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpTermsOfServiceUpdate.
func (s HelpTermsOfServiceUpdateArray) SortStable(less func(a, b HelpTermsOfServiceUpdate) bool) HelpTermsOfServiceUpdateArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpTermsOfServiceUpdate.
func (s HelpTermsOfServiceUpdateArray) Retain(keep func(x HelpTermsOfServiceUpdate) bool) HelpTermsOfServiceUpdateArray {
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
func (s HelpTermsOfServiceUpdateArray) First() (v HelpTermsOfServiceUpdate, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpTermsOfServiceUpdateArray) Last() (v HelpTermsOfServiceUpdate, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpTermsOfServiceUpdateArray) PopFirst() (v HelpTermsOfServiceUpdate, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpTermsOfServiceUpdate
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpTermsOfServiceUpdateArray) Pop() (v HelpTermsOfServiceUpdate, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
