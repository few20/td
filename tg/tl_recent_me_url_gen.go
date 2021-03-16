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

// RecentMeURLUnknown represents TL type `recentMeUrlUnknown#46e1d13d`.
// Unknown t.me url
//
// See https://core.telegram.org/constructor/recentMeUrlUnknown for reference.
type RecentMeURLUnknown struct {
	// URL
	URL string
}

// RecentMeURLUnknownTypeID is TL type id of RecentMeURLUnknown.
const RecentMeURLUnknownTypeID = 0x46e1d13d

func (r *RecentMeURLUnknown) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLUnknown) String() string {
	if r == nil {
		return "RecentMeURLUnknown(nil)"
	}
	type Alias RecentMeURLUnknown
	return fmt.Sprintf("RecentMeURLUnknown%+v", Alias(*r))
}

// FillFrom fills RecentMeURLUnknown from given interface.
func (r *RecentMeURLUnknown) FillFrom(from interface {
	GetURL() (value string)
}) {
	r.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLUnknown) TypeID() uint32 {
	return RecentMeURLUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLUnknown) TypeName() string {
	return "recentMeUrlUnknown"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlUnknown",
		ID:   RecentMeURLUnknownTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLUnknown) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUnknown#46e1d13d as nil")
	}
	b.PutID(RecentMeURLUnknownTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLUnknown) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUnknown#46e1d13d as nil")
	}
	b.PutString(r.URL)
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLUnknown) GetURL() (value string) {
	return r.URL
}

// Decode implements bin.Decoder.
func (r *RecentMeURLUnknown) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUnknown#46e1d13d to nil")
	}
	if err := b.ConsumeID(RecentMeURLUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLUnknown) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUnknown#46e1d13d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUnknown#46e1d13d: field url: %w", err)
		}
		r.URL = value
	}
	return nil
}

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLUnknown) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLUnknown.
var (
	_ bin.Encoder     = &RecentMeURLUnknown{}
	_ bin.Decoder     = &RecentMeURLUnknown{}
	_ bin.BareEncoder = &RecentMeURLUnknown{}
	_ bin.BareDecoder = &RecentMeURLUnknown{}

	_ RecentMeURLClass = &RecentMeURLUnknown{}
)

// RecentMeURLUser represents TL type `recentMeUrlUser#8dbc3336`.
// Recent t.me link to a user
//
// See https://core.telegram.org/constructor/recentMeUrlUser for reference.
type RecentMeURLUser struct {
	// URL
	URL string
	// User ID
	UserID int
}

// RecentMeURLUserTypeID is TL type id of RecentMeURLUser.
const RecentMeURLUserTypeID = 0x8dbc3336

func (r *RecentMeURLUser) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLUser) String() string {
	if r == nil {
		return "RecentMeURLUser(nil)"
	}
	type Alias RecentMeURLUser
	return fmt.Sprintf("RecentMeURLUser%+v", Alias(*r))
}

// FillFrom fills RecentMeURLUser from given interface.
func (r *RecentMeURLUser) FillFrom(from interface {
	GetURL() (value string)
	GetUserID() (value int)
}) {
	r.URL = from.GetURL()
	r.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLUser) TypeID() uint32 {
	return RecentMeURLUserTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLUser) TypeName() string {
	return "recentMeUrlUser"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlUser",
		ID:   RecentMeURLUserTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLUser) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUser#8dbc3336 as nil")
	}
	b.PutID(RecentMeURLUserTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLUser) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlUser#8dbc3336 as nil")
	}
	b.PutString(r.URL)
	b.PutInt(r.UserID)
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLUser) GetURL() (value string) {
	return r.URL
}

// GetUserID returns value of UserID field.
func (r *RecentMeURLUser) GetUserID() (value int) {
	return r.UserID
}

// Decode implements bin.Decoder.
func (r *RecentMeURLUser) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUser#8dbc3336 to nil")
	}
	if err := b.ConsumeID(RecentMeURLUserTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLUser) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlUser#8dbc3336 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlUser#8dbc3336: field user_id: %w", err)
		}
		r.UserID = value
	}
	return nil
}

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLUser) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLUser.
var (
	_ bin.Encoder     = &RecentMeURLUser{}
	_ bin.Decoder     = &RecentMeURLUser{}
	_ bin.BareEncoder = &RecentMeURLUser{}
	_ bin.BareDecoder = &RecentMeURLUser{}

	_ RecentMeURLClass = &RecentMeURLUser{}
)

// RecentMeURLChat represents TL type `recentMeUrlChat#a01b22f9`.
// Recent t.me link to a chat
//
// See https://core.telegram.org/constructor/recentMeUrlChat for reference.
type RecentMeURLChat struct {
	// t.me URL
	URL string
	// Chat ID
	ChatID int
}

// RecentMeURLChatTypeID is TL type id of RecentMeURLChat.
const RecentMeURLChatTypeID = 0xa01b22f9

func (r *RecentMeURLChat) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLChat) String() string {
	if r == nil {
		return "RecentMeURLChat(nil)"
	}
	type Alias RecentMeURLChat
	return fmt.Sprintf("RecentMeURLChat%+v", Alias(*r))
}

// FillFrom fills RecentMeURLChat from given interface.
func (r *RecentMeURLChat) FillFrom(from interface {
	GetURL() (value string)
	GetChatID() (value int)
}) {
	r.URL = from.GetURL()
	r.ChatID = from.GetChatID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLChat) TypeID() uint32 {
	return RecentMeURLChatTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLChat) TypeName() string {
	return "recentMeUrlChat"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlChat",
		ID:   RecentMeURLChatTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLChat) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChat#a01b22f9 as nil")
	}
	b.PutID(RecentMeURLChatTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLChat) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChat#a01b22f9 as nil")
	}
	b.PutString(r.URL)
	b.PutInt(r.ChatID)
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLChat) GetURL() (value string) {
	return r.URL
}

// GetChatID returns value of ChatID field.
func (r *RecentMeURLChat) GetChatID() (value int) {
	return r.ChatID
}

// Decode implements bin.Decoder.
func (r *RecentMeURLChat) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChat#a01b22f9 to nil")
	}
	if err := b.ConsumeID(RecentMeURLChatTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLChat) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChat#a01b22f9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChat#a01b22f9: field chat_id: %w", err)
		}
		r.ChatID = value
	}
	return nil
}

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLChat) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLChat.
var (
	_ bin.Encoder     = &RecentMeURLChat{}
	_ bin.Decoder     = &RecentMeURLChat{}
	_ bin.BareEncoder = &RecentMeURLChat{}
	_ bin.BareDecoder = &RecentMeURLChat{}

	_ RecentMeURLClass = &RecentMeURLChat{}
)

// RecentMeURLChatInvite represents TL type `recentMeUrlChatInvite#eb49081d`.
// Recent t.me invite link to a chat
//
// See https://core.telegram.org/constructor/recentMeUrlChatInvite for reference.
type RecentMeURLChatInvite struct {
	// t.me URL
	URL string
	// Chat invitation
	ChatInvite ChatInviteClass
}

// RecentMeURLChatInviteTypeID is TL type id of RecentMeURLChatInvite.
const RecentMeURLChatInviteTypeID = 0xeb49081d

func (r *RecentMeURLChatInvite) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.ChatInvite == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLChatInvite) String() string {
	if r == nil {
		return "RecentMeURLChatInvite(nil)"
	}
	type Alias RecentMeURLChatInvite
	return fmt.Sprintf("RecentMeURLChatInvite%+v", Alias(*r))
}

// FillFrom fills RecentMeURLChatInvite from given interface.
func (r *RecentMeURLChatInvite) FillFrom(from interface {
	GetURL() (value string)
	GetChatInvite() (value ChatInviteClass)
}) {
	r.URL = from.GetURL()
	r.ChatInvite = from.GetChatInvite()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLChatInvite) TypeID() uint32 {
	return RecentMeURLChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLChatInvite) TypeName() string {
	return "recentMeUrlChatInvite"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLChatInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlChatInvite",
		ID:   RecentMeURLChatInviteTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "ChatInvite",
			SchemaName: "chat_invite",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLChatInvite) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChatInvite#eb49081d as nil")
	}
	b.PutID(RecentMeURLChatInviteTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLChatInvite) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlChatInvite#eb49081d as nil")
	}
	b.PutString(r.URL)
	if r.ChatInvite == nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite is nil")
	}
	if err := r.ChatInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLChatInvite) GetURL() (value string) {
	return r.URL
}

// GetChatInvite returns value of ChatInvite field.
func (r *RecentMeURLChatInvite) GetChatInvite() (value ChatInviteClass) {
	return r.ChatInvite
}

// Decode implements bin.Decoder.
func (r *RecentMeURLChatInvite) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChatInvite#eb49081d to nil")
	}
	if err := b.ConsumeID(RecentMeURLChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLChatInvite) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlChatInvite#eb49081d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeChatInvite(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlChatInvite#eb49081d: field chat_invite: %w", err)
		}
		r.ChatInvite = value
	}
	return nil
}

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLChatInvite) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLChatInvite.
var (
	_ bin.Encoder     = &RecentMeURLChatInvite{}
	_ bin.Decoder     = &RecentMeURLChatInvite{}
	_ bin.BareEncoder = &RecentMeURLChatInvite{}
	_ bin.BareDecoder = &RecentMeURLChatInvite{}

	_ RecentMeURLClass = &RecentMeURLChatInvite{}
)

// RecentMeURLStickerSet represents TL type `recentMeUrlStickerSet#bc0a57dc`.
// Recent t.me stickerset installation URL
//
// See https://core.telegram.org/constructor/recentMeUrlStickerSet for reference.
type RecentMeURLStickerSet struct {
	// t.me URL
	URL string
	// Stickerset
	Set StickerSetCoveredClass
}

// RecentMeURLStickerSetTypeID is TL type id of RecentMeURLStickerSet.
const RecentMeURLStickerSetTypeID = 0xbc0a57dc

func (r *RecentMeURLStickerSet) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.URL == "") {
		return false
	}
	if !(r.Set == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecentMeURLStickerSet) String() string {
	if r == nil {
		return "RecentMeURLStickerSet(nil)"
	}
	type Alias RecentMeURLStickerSet
	return fmt.Sprintf("RecentMeURLStickerSet%+v", Alias(*r))
}

// FillFrom fills RecentMeURLStickerSet from given interface.
func (r *RecentMeURLStickerSet) FillFrom(from interface {
	GetURL() (value string)
	GetSet() (value StickerSetCoveredClass)
}) {
	r.URL = from.GetURL()
	r.Set = from.GetSet()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecentMeURLStickerSet) TypeID() uint32 {
	return RecentMeURLStickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (*RecentMeURLStickerSet) TypeName() string {
	return "recentMeUrlStickerSet"
}

// TypeInfo returns info about TL type.
func (r *RecentMeURLStickerSet) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recentMeUrlStickerSet",
		ID:   RecentMeURLStickerSetTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Set",
			SchemaName: "set",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecentMeURLStickerSet) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlStickerSet#bc0a57dc as nil")
	}
	b.PutID(RecentMeURLStickerSetTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecentMeURLStickerSet) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recentMeUrlStickerSet#bc0a57dc as nil")
	}
	b.PutString(r.URL)
	if r.Set == nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set is nil")
	}
	if err := r.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
	}
	return nil
}

// GetURL returns value of URL field.
func (r *RecentMeURLStickerSet) GetURL() (value string) {
	return r.URL
}

// GetSet returns value of Set field.
func (r *RecentMeURLStickerSet) GetSet() (value StickerSetCoveredClass) {
	return r.Set
}

// Decode implements bin.Decoder.
func (r *RecentMeURLStickerSet) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlStickerSet#bc0a57dc to nil")
	}
	if err := b.ConsumeID(RecentMeURLStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecentMeURLStickerSet) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recentMeUrlStickerSet#bc0a57dc to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field url: %w", err)
		}
		r.URL = value
	}
	{
		value, err := DecodeStickerSetCovered(b)
		if err != nil {
			return fmt.Errorf("unable to decode recentMeUrlStickerSet#bc0a57dc: field set: %w", err)
		}
		r.Set = value
	}
	return nil
}

// construct implements constructor of RecentMeURLClass.
func (r RecentMeURLStickerSet) construct() RecentMeURLClass { return &r }

// Ensuring interfaces in compile-time for RecentMeURLStickerSet.
var (
	_ bin.Encoder     = &RecentMeURLStickerSet{}
	_ bin.Decoder     = &RecentMeURLStickerSet{}
	_ bin.BareEncoder = &RecentMeURLStickerSet{}
	_ bin.BareDecoder = &RecentMeURLStickerSet{}

	_ RecentMeURLClass = &RecentMeURLStickerSet{}
)

// RecentMeURLClass represents RecentMeUrl generic type.
//
// See https://core.telegram.org/type/RecentMeUrl for reference.
//
// Example:
//  g, err := tg.DecodeRecentMeURL(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.RecentMeURLUnknown: // recentMeUrlUnknown#46e1d13d
//  case *tg.RecentMeURLUser: // recentMeUrlUser#8dbc3336
//  case *tg.RecentMeURLChat: // recentMeUrlChat#a01b22f9
//  case *tg.RecentMeURLChatInvite: // recentMeUrlChatInvite#eb49081d
//  case *tg.RecentMeURLStickerSet: // recentMeUrlStickerSet#bc0a57dc
//  default: panic(v)
//  }
type RecentMeURLClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() RecentMeURLClass

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

	// URL
	GetURL() (value string)
}

// DecodeRecentMeURL implements binary de-serialization for RecentMeURLClass.
func DecodeRecentMeURL(buf *bin.Buffer) (RecentMeURLClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RecentMeURLUnknownTypeID:
		// Decoding recentMeUrlUnknown#46e1d13d.
		v := RecentMeURLUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLUserTypeID:
		// Decoding recentMeUrlUser#8dbc3336.
		v := RecentMeURLUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLChatTypeID:
		// Decoding recentMeUrlChat#a01b22f9.
		v := RecentMeURLChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLChatInviteTypeID:
		// Decoding recentMeUrlChatInvite#eb49081d.
		v := RecentMeURLChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	case RecentMeURLStickerSetTypeID:
		// Decoding recentMeUrlStickerSet#bc0a57dc.
		v := RecentMeURLStickerSet{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RecentMeURLClass: %w", bin.NewUnexpectedID(id))
	}
}

// RecentMeURL boxes the RecentMeURLClass providing a helper.
type RecentMeURLBox struct {
	RecentMeUrl RecentMeURLClass
}

// Decode implements bin.Decoder for RecentMeURLBox.
func (b *RecentMeURLBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RecentMeURLBox to nil")
	}
	v, err := DecodeRecentMeURL(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RecentMeUrl = v
	return nil
}

// Encode implements bin.Encode for RecentMeURLBox.
func (b *RecentMeURLBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RecentMeUrl == nil {
		return fmt.Errorf("unable to encode RecentMeURLClass as nil")
	}
	return b.RecentMeUrl.Encode(buf)
}

// RecentMeURLClassArray is adapter for slice of RecentMeURLClass.
type RecentMeURLClassArray []RecentMeURLClass

// Sort sorts slice of RecentMeURLClass.
func (s RecentMeURLClassArray) Sort(less func(a, b RecentMeURLClass) bool) RecentMeURLClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLClass.
func (s RecentMeURLClassArray) SortStable(less func(a, b RecentMeURLClass) bool) RecentMeURLClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLClass.
func (s RecentMeURLClassArray) Retain(keep func(x RecentMeURLClass) bool) RecentMeURLClassArray {
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
func (s RecentMeURLClassArray) First() (v RecentMeURLClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLClassArray) Last() (v RecentMeURLClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLClassArray) PopFirst() (v RecentMeURLClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLClassArray) Pop() (v RecentMeURLClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsRecentMeURLUnknown returns copy with only RecentMeURLUnknown constructors.
func (s RecentMeURLClassArray) AsRecentMeURLUnknown() (to RecentMeURLUnknownArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLUnknown)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLUser returns copy with only RecentMeURLUser constructors.
func (s RecentMeURLClassArray) AsRecentMeURLUser() (to RecentMeURLUserArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLUser)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLChat returns copy with only RecentMeURLChat constructors.
func (s RecentMeURLClassArray) AsRecentMeURLChat() (to RecentMeURLChatArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLChat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLChatInvite returns copy with only RecentMeURLChatInvite constructors.
func (s RecentMeURLClassArray) AsRecentMeURLChatInvite() (to RecentMeURLChatInviteArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLChatInvite)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsRecentMeURLStickerSet returns copy with only RecentMeURLStickerSet constructors.
func (s RecentMeURLClassArray) AsRecentMeURLStickerSet() (to RecentMeURLStickerSetArray) {
	for _, elem := range s {
		value, ok := elem.(*RecentMeURLStickerSet)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// RecentMeURLUnknownArray is adapter for slice of RecentMeURLUnknown.
type RecentMeURLUnknownArray []RecentMeURLUnknown

// Sort sorts slice of RecentMeURLUnknown.
func (s RecentMeURLUnknownArray) Sort(less func(a, b RecentMeURLUnknown) bool) RecentMeURLUnknownArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLUnknown.
func (s RecentMeURLUnknownArray) SortStable(less func(a, b RecentMeURLUnknown) bool) RecentMeURLUnknownArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLUnknown.
func (s RecentMeURLUnknownArray) Retain(keep func(x RecentMeURLUnknown) bool) RecentMeURLUnknownArray {
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
func (s RecentMeURLUnknownArray) First() (v RecentMeURLUnknown, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLUnknownArray) Last() (v RecentMeURLUnknown, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLUnknownArray) PopFirst() (v RecentMeURLUnknown, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLUnknown
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLUnknownArray) Pop() (v RecentMeURLUnknown, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLUserArray is adapter for slice of RecentMeURLUser.
type RecentMeURLUserArray []RecentMeURLUser

// Sort sorts slice of RecentMeURLUser.
func (s RecentMeURLUserArray) Sort(less func(a, b RecentMeURLUser) bool) RecentMeURLUserArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLUser.
func (s RecentMeURLUserArray) SortStable(less func(a, b RecentMeURLUser) bool) RecentMeURLUserArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLUser.
func (s RecentMeURLUserArray) Retain(keep func(x RecentMeURLUser) bool) RecentMeURLUserArray {
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
func (s RecentMeURLUserArray) First() (v RecentMeURLUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLUserArray) Last() (v RecentMeURLUser, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLUserArray) PopFirst() (v RecentMeURLUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLUser
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLUserArray) Pop() (v RecentMeURLUser, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLChatArray is adapter for slice of RecentMeURLChat.
type RecentMeURLChatArray []RecentMeURLChat

// Sort sorts slice of RecentMeURLChat.
func (s RecentMeURLChatArray) Sort(less func(a, b RecentMeURLChat) bool) RecentMeURLChatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLChat.
func (s RecentMeURLChatArray) SortStable(less func(a, b RecentMeURLChat) bool) RecentMeURLChatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLChat.
func (s RecentMeURLChatArray) Retain(keep func(x RecentMeURLChat) bool) RecentMeURLChatArray {
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
func (s RecentMeURLChatArray) First() (v RecentMeURLChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLChatArray) Last() (v RecentMeURLChat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLChatArray) PopFirst() (v RecentMeURLChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLChat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLChatArray) Pop() (v RecentMeURLChat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLChatInviteArray is adapter for slice of RecentMeURLChatInvite.
type RecentMeURLChatInviteArray []RecentMeURLChatInvite

// Sort sorts slice of RecentMeURLChatInvite.
func (s RecentMeURLChatInviteArray) Sort(less func(a, b RecentMeURLChatInvite) bool) RecentMeURLChatInviteArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLChatInvite.
func (s RecentMeURLChatInviteArray) SortStable(less func(a, b RecentMeURLChatInvite) bool) RecentMeURLChatInviteArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLChatInvite.
func (s RecentMeURLChatInviteArray) Retain(keep func(x RecentMeURLChatInvite) bool) RecentMeURLChatInviteArray {
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
func (s RecentMeURLChatInviteArray) First() (v RecentMeURLChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLChatInviteArray) Last() (v RecentMeURLChatInvite, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLChatInviteArray) PopFirst() (v RecentMeURLChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLChatInvite
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLChatInviteArray) Pop() (v RecentMeURLChatInvite, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// RecentMeURLStickerSetArray is adapter for slice of RecentMeURLStickerSet.
type RecentMeURLStickerSetArray []RecentMeURLStickerSet

// Sort sorts slice of RecentMeURLStickerSet.
func (s RecentMeURLStickerSetArray) Sort(less func(a, b RecentMeURLStickerSet) bool) RecentMeURLStickerSetArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RecentMeURLStickerSet.
func (s RecentMeURLStickerSetArray) SortStable(less func(a, b RecentMeURLStickerSet) bool) RecentMeURLStickerSetArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RecentMeURLStickerSet.
func (s RecentMeURLStickerSetArray) Retain(keep func(x RecentMeURLStickerSet) bool) RecentMeURLStickerSetArray {
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
func (s RecentMeURLStickerSetArray) First() (v RecentMeURLStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RecentMeURLStickerSetArray) Last() (v RecentMeURLStickerSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RecentMeURLStickerSetArray) PopFirst() (v RecentMeURLStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RecentMeURLStickerSet
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RecentMeURLStickerSetArray) Pop() (v RecentMeURLStickerSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
