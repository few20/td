// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Auth represents TL type `auth#f8bb4a38`.
//
// See https://localhost:80/doc/constructor/auth for reference.
type Auth struct {
	// Name field of Auth.
	Name string
}

// AuthTypeID is TL type id of Auth.
const AuthTypeID = 0xf8bb4a38

func (a *Auth) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *Auth) String() string {
	if a == nil {
		return "Auth(nil)"
	}
	type Alias Auth
	return fmt.Sprintf("Auth%+v", Alias(*a))
}

// FillFrom fills Auth from given interface.
func (a *Auth) FillFrom(from interface {
	GetName() (value string)
}) {
	a.Name = from.GetName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Auth) TypeID() uint32 {
	return AuthTypeID
}

// TypeName returns name of type in TL schema.
func (*Auth) TypeName() string {
	return "auth"
}

// TypeInfo returns info about TL type.
func (a *Auth) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth",
		ID:   AuthTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *Auth) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth#f8bb4a38 as nil")
	}
	b.PutID(AuthTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *Auth) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth#f8bb4a38 as nil")
	}
	b.PutString(a.Name)
	return nil
}

// GetName returns value of Name field.
func (a *Auth) GetName() (value string) {
	return a.Name
}

// Decode implements bin.Decoder.
func (a *Auth) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth#f8bb4a38 to nil")
	}
	if err := b.ConsumeID(AuthTypeID); err != nil {
		return fmt.Errorf("unable to decode auth#f8bb4a38: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *Auth) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth#f8bb4a38 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth#f8bb4a38: field name: %w", err)
		}
		a.Name = value
	}
	return nil
}

// construct implements constructor of AuthClass.
func (a Auth) construct() AuthClass { return &a }

// Ensuring interfaces in compile-time for Auth.
var (
	_ bin.Encoder     = &Auth{}
	_ bin.Decoder     = &Auth{}
	_ bin.BareEncoder = &Auth{}
	_ bin.BareDecoder = &Auth{}

	_ AuthClass = &Auth{}
)

// AuthPassword represents TL type `authPassword#29bacabb`.
//
// See https://localhost:80/doc/constructor/authPassword for reference.
type AuthPassword struct {
	// Name field of AuthPassword.
	Name string
	// Password field of AuthPassword.
	Password string
}

// AuthPasswordTypeID is TL type id of AuthPassword.
const AuthPasswordTypeID = 0x29bacabb

func (a *AuthPassword) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Name == "") {
		return false
	}
	if !(a.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthPassword) String() string {
	if a == nil {
		return "AuthPassword(nil)"
	}
	type Alias AuthPassword
	return fmt.Sprintf("AuthPassword%+v", Alias(*a))
}

// FillFrom fills AuthPassword from given interface.
func (a *AuthPassword) FillFrom(from interface {
	GetName() (value string)
	GetPassword() (value string)
}) {
	a.Name = from.GetName()
	a.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthPassword) TypeID() uint32 {
	return AuthPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthPassword) TypeName() string {
	return "authPassword"
}

// TypeInfo returns info about TL type.
func (a *AuthPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authPassword",
		ID:   AuthPasswordTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthPassword) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authPassword#29bacabb as nil")
	}
	b.PutID(AuthPasswordTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthPassword) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authPassword#29bacabb as nil")
	}
	b.PutString(a.Name)
	b.PutString(a.Password)
	return nil
}

// GetName returns value of Name field.
func (a *AuthPassword) GetName() (value string) {
	return a.Name
}

// GetPassword returns value of Password field.
func (a *AuthPassword) GetPassword() (value string) {
	return a.Password
}

// Decode implements bin.Decoder.
func (a *AuthPassword) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authPassword#29bacabb to nil")
	}
	if err := b.ConsumeID(AuthPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode authPassword#29bacabb: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthPassword) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authPassword#29bacabb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authPassword#29bacabb: field name: %w", err)
		}
		a.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authPassword#29bacabb: field password: %w", err)
		}
		a.Password = value
	}
	return nil
}

// construct implements constructor of AuthClass.
func (a AuthPassword) construct() AuthClass { return &a }

// Ensuring interfaces in compile-time for AuthPassword.
var (
	_ bin.Encoder     = &AuthPassword{}
	_ bin.Decoder     = &AuthPassword{}
	_ bin.BareEncoder = &AuthPassword{}
	_ bin.BareDecoder = &AuthPassword{}

	_ AuthClass = &AuthPassword{}
)

// AuthClass represents Auth generic type.
//
// See https://localhost:80/doc/type/Auth for reference.
//
// Example:
//  g, err := td.DecodeAuth(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *td.Auth: // auth#f8bb4a38
//  case *td.AuthPassword: // authPassword#29bacabb
//  default: panic(v)
//  }
type AuthClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthClass

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

	// Name field of Auth.
	GetName() (value string)
}

// DecodeAuth implements binary de-serialization for AuthClass.
func DecodeAuth(buf *bin.Buffer) (AuthClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthTypeID:
		// Decoding auth#f8bb4a38.
		v := Auth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthClass: %w", err)
		}
		return &v, nil
	case AuthPasswordTypeID:
		// Decoding authPassword#29bacabb.
		v := AuthPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthClass: %w", bin.NewUnexpectedID(id))
	}
}

// Auth boxes the AuthClass providing a helper.
type AuthBox struct {
	Auth AuthClass
}

// Decode implements bin.Decoder for AuthBox.
func (b *AuthBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthBox to nil")
	}
	v, err := DecodeAuth(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Auth = v
	return nil
}

// Encode implements bin.Encode for AuthBox.
func (b *AuthBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Auth == nil {
		return fmt.Errorf("unable to encode AuthClass as nil")
	}
	return b.Auth.Encode(buf)
}

// AuthClassArray is adapter for slice of AuthClass.
type AuthClassArray []AuthClass

// Sort sorts slice of AuthClass.
func (s AuthClassArray) Sort(less func(a, b AuthClass) bool) AuthClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthClass.
func (s AuthClassArray) SortStable(less func(a, b AuthClass) bool) AuthClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthClass.
func (s AuthClassArray) Retain(keep func(x AuthClass) bool) AuthClassArray {
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
func (s AuthClassArray) First() (v AuthClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthClassArray) Last() (v AuthClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthClassArray) PopFirst() (v AuthClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthClassArray) Pop() (v AuthClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsAuth returns copy with only Auth constructors.
func (s AuthClassArray) AsAuth() (to AuthArray) {
	for _, elem := range s {
		value, ok := elem.(*Auth)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsAuthPassword returns copy with only AuthPassword constructors.
func (s AuthClassArray) AsAuthPassword() (to AuthPasswordArray) {
	for _, elem := range s {
		value, ok := elem.(*AuthPassword)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AuthArray is adapter for slice of Auth.
type AuthArray []Auth

// Sort sorts slice of Auth.
func (s AuthArray) Sort(less func(a, b Auth) bool) AuthArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of Auth.
func (s AuthArray) SortStable(less func(a, b Auth) bool) AuthArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of Auth.
func (s AuthArray) Retain(keep func(x Auth) bool) AuthArray {
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
func (s AuthArray) First() (v Auth, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthArray) Last() (v Auth, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthArray) PopFirst() (v Auth, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero Auth
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthArray) Pop() (v Auth, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AuthPasswordArray is adapter for slice of AuthPassword.
type AuthPasswordArray []AuthPassword

// Sort sorts slice of AuthPassword.
func (s AuthPasswordArray) Sort(less func(a, b AuthPassword) bool) AuthPasswordArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthPassword.
func (s AuthPasswordArray) SortStable(less func(a, b AuthPassword) bool) AuthPasswordArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthPassword.
func (s AuthPasswordArray) Retain(keep func(x AuthPassword) bool) AuthPasswordArray {
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
func (s AuthPasswordArray) First() (v AuthPassword, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthPasswordArray) Last() (v AuthPassword, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthPasswordArray) PopFirst() (v AuthPassword, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthPassword
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthPasswordArray) Pop() (v AuthPassword, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
