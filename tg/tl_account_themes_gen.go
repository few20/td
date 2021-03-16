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

// AccountThemesNotModified represents TL type `account.themesNotModified#f41eb622`.
// No new themes were installed
//
// See https://core.telegram.org/constructor/account.themesNotModified for reference.
type AccountThemesNotModified struct {
}

// AccountThemesNotModifiedTypeID is TL type id of AccountThemesNotModified.
const AccountThemesNotModifiedTypeID = 0xf41eb622

func (t *AccountThemesNotModified) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountThemesNotModified) String() string {
	if t == nil {
		return "AccountThemesNotModified(nil)"
	}
	type Alias AccountThemesNotModified
	return fmt.Sprintf("AccountThemesNotModified%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountThemesNotModified) TypeID() uint32 {
	return AccountThemesNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountThemesNotModified) TypeName() string {
	return "account.themesNotModified"
}

// TypeInfo returns info about TL type.
func (t *AccountThemesNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.themesNotModified",
		ID:   AccountThemesNotModifiedTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountThemesNotModified) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themesNotModified#f41eb622 as nil")
	}
	b.PutID(AccountThemesNotModifiedTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *AccountThemesNotModified) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themesNotModified#f41eb622 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *AccountThemesNotModified) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themesNotModified#f41eb622 to nil")
	}
	if err := b.ConsumeID(AccountThemesNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode account.themesNotModified#f41eb622: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *AccountThemesNotModified) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themesNotModified#f41eb622 to nil")
	}
	return nil
}

// construct implements constructor of AccountThemesClass.
func (t AccountThemesNotModified) construct() AccountThemesClass { return &t }

// Ensuring interfaces in compile-time for AccountThemesNotModified.
var (
	_ bin.Encoder     = &AccountThemesNotModified{}
	_ bin.Decoder     = &AccountThemesNotModified{}
	_ bin.BareEncoder = &AccountThemesNotModified{}
	_ bin.BareDecoder = &AccountThemesNotModified{}

	_ AccountThemesClass = &AccountThemesNotModified{}
)

// AccountThemes represents TL type `account.themes#7f676421`.
// Installed themes
//
// See https://core.telegram.org/constructor/account.themes for reference.
type AccountThemes struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Themes
	Themes []Theme
}

// AccountThemesTypeID is TL type id of AccountThemes.
const AccountThemesTypeID = 0x7f676421

func (t *AccountThemes) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Hash == 0) {
		return false
	}
	if !(t.Themes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountThemes) String() string {
	if t == nil {
		return "AccountThemes(nil)"
	}
	type Alias AccountThemes
	return fmt.Sprintf("AccountThemes%+v", Alias(*t))
}

// FillFrom fills AccountThemes from given interface.
func (t *AccountThemes) FillFrom(from interface {
	GetHash() (value int)
	GetThemes() (value []Theme)
}) {
	t.Hash = from.GetHash()
	t.Themes = from.GetThemes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountThemes) TypeID() uint32 {
	return AccountThemesTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountThemes) TypeName() string {
	return "account.themes"
}

// TypeInfo returns info about TL type.
func (t *AccountThemes) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.themes",
		ID:   AccountThemesTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Themes",
			SchemaName: "themes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *AccountThemes) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themes#7f676421 as nil")
	}
	b.PutID(AccountThemesTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *AccountThemes) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.themes#7f676421 as nil")
	}
	b.PutInt(t.Hash)
	b.PutVectorHeader(len(t.Themes))
	for idx, v := range t.Themes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.themes#7f676421: field themes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (t *AccountThemes) GetHash() (value int) {
	return t.Hash
}

// GetThemes returns value of Themes field.
func (t *AccountThemes) GetThemes() (value []Theme) {
	return t.Themes
}

// Decode implements bin.Decoder.
func (t *AccountThemes) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themes#7f676421 to nil")
	}
	if err := b.ConsumeID(AccountThemesTypeID); err != nil {
		return fmt.Errorf("unable to decode account.themes#7f676421: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *AccountThemes) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.themes#7f676421 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.themes#7f676421: field hash: %w", err)
		}
		t.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.themes#7f676421: field themes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Theme
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.themes#7f676421: field themes: %w", err)
			}
			t.Themes = append(t.Themes, value)
		}
	}
	return nil
}

// construct implements constructor of AccountThemesClass.
func (t AccountThemes) construct() AccountThemesClass { return &t }

// Ensuring interfaces in compile-time for AccountThemes.
var (
	_ bin.Encoder     = &AccountThemes{}
	_ bin.Decoder     = &AccountThemes{}
	_ bin.BareEncoder = &AccountThemes{}
	_ bin.BareDecoder = &AccountThemes{}

	_ AccountThemesClass = &AccountThemes{}
)

// AccountThemesClass represents account.Themes generic type.
//
// See https://core.telegram.org/type/account.Themes for reference.
//
// Example:
//  g, err := tg.DecodeAccountThemes(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AccountThemesNotModified: // account.themesNotModified#f41eb622
//  case *tg.AccountThemes: // account.themes#7f676421
//  default: panic(v)
//  }
type AccountThemesClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountThemesClass

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

	// AsModified tries to map AccountThemesClass to AccountThemes.
	AsModified() (*AccountThemes, bool)
}

// AsModified tries to map AccountThemesNotModified to AccountThemes.
func (t *AccountThemesNotModified) AsModified() (*AccountThemes, bool) {
	return nil, false
}

// AsModified tries to map AccountThemes to AccountThemes.
func (t *AccountThemes) AsModified() (*AccountThemes, bool) {
	return t, true
}

// DecodeAccountThemes implements binary de-serialization for AccountThemesClass.
func DecodeAccountThemes(buf *bin.Buffer) (AccountThemesClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountThemesNotModifiedTypeID:
		// Decoding account.themesNotModified#f41eb622.
		v := AccountThemesNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", err)
		}
		return &v, nil
	case AccountThemesTypeID:
		// Decoding account.themes#7f676421.
		v := AccountThemes{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountThemesClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountThemes boxes the AccountThemesClass providing a helper.
type AccountThemesBox struct {
	Themes AccountThemesClass
}

// Decode implements bin.Decoder for AccountThemesBox.
func (b *AccountThemesBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountThemesBox to nil")
	}
	v, err := DecodeAccountThemes(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Themes = v
	return nil
}

// Encode implements bin.Encode for AccountThemesBox.
func (b *AccountThemesBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Themes == nil {
		return fmt.Errorf("unable to encode AccountThemesClass as nil")
	}
	return b.Themes.Encode(buf)
}

// AccountThemesClassArray is adapter for slice of AccountThemesClass.
type AccountThemesClassArray []AccountThemesClass

// Sort sorts slice of AccountThemesClass.
func (s AccountThemesClassArray) Sort(less func(a, b AccountThemesClass) bool) AccountThemesClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountThemesClass.
func (s AccountThemesClassArray) SortStable(less func(a, b AccountThemesClass) bool) AccountThemesClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountThemesClass.
func (s AccountThemesClassArray) Retain(keep func(x AccountThemesClass) bool) AccountThemesClassArray {
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
func (s AccountThemesClassArray) First() (v AccountThemesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountThemesClassArray) Last() (v AccountThemesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountThemesClassArray) PopFirst() (v AccountThemesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountThemesClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountThemesClassArray) Pop() (v AccountThemesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAccountThemes returns copy with only AccountThemes constructors.
func (s AccountThemesClassArray) AsAccountThemes() (to AccountThemesArray) {
	for _, elem := range s {
		value, ok := elem.(*AccountThemes)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s AccountThemesClassArray) AppendOnlyModified(to []*AccountThemes) []*AccountThemes {
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
func (s AccountThemesClassArray) AsModified() (to []*AccountThemes) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s AccountThemesClassArray) FirstAsModified() (v *AccountThemes, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s AccountThemesClassArray) LastAsModified() (v *AccountThemes, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *AccountThemesClassArray) PopFirstAsModified() (v *AccountThemes, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *AccountThemesClassArray) PopAsModified() (v *AccountThemes, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// AccountThemesArray is adapter for slice of AccountThemes.
type AccountThemesArray []AccountThemes

// Sort sorts slice of AccountThemes.
func (s AccountThemesArray) Sort(less func(a, b AccountThemes) bool) AccountThemesArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AccountThemes.
func (s AccountThemesArray) SortStable(less func(a, b AccountThemes) bool) AccountThemesArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AccountThemes.
func (s AccountThemesArray) Retain(keep func(x AccountThemes) bool) AccountThemesArray {
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
func (s AccountThemesArray) First() (v AccountThemes, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AccountThemesArray) Last() (v AccountThemes, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AccountThemesArray) PopFirst() (v AccountThemes, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AccountThemes
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AccountThemesArray) Pop() (v AccountThemes, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
