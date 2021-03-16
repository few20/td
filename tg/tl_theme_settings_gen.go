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

// ThemeSettings represents TL type `themeSettings#9c14984a`.
// Theme settings
//
// See https://core.telegram.org/constructor/themeSettings for reference.
type ThemeSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Base theme
	BaseTheme BaseThemeClass
	// Accent color, RGB24 format
	AccentColor int
	// Message gradient color (top), RGB24 format
	//
	// Use SetMessageTopColor and GetMessageTopColor helpers.
	MessageTopColor int
	// Message gradient color (bottom), RGB24 format
	//
	// Use SetMessageBottomColor and GetMessageBottomColor helpers.
	MessageBottomColor int
	// Wallpaper
	//
	// Use SetWallpaper and GetWallpaper helpers.
	Wallpaper WallPaperClass
}

// ThemeSettingsTypeID is TL type id of ThemeSettings.
const ThemeSettingsTypeID = 0x9c14984a

func (t *ThemeSettings) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.BaseTheme == nil) {
		return false
	}
	if !(t.AccentColor == 0) {
		return false
	}
	if !(t.MessageTopColor == 0) {
		return false
	}
	if !(t.MessageBottomColor == 0) {
		return false
	}
	if !(t.Wallpaper == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThemeSettings) String() string {
	if t == nil {
		return "ThemeSettings(nil)"
	}
	type Alias ThemeSettings
	return fmt.Sprintf("ThemeSettings%+v", Alias(*t))
}

// FillFrom fills ThemeSettings from given interface.
func (t *ThemeSettings) FillFrom(from interface {
	GetBaseTheme() (value BaseThemeClass)
	GetAccentColor() (value int)
	GetMessageTopColor() (value int, ok bool)
	GetMessageBottomColor() (value int, ok bool)
	GetWallpaper() (value WallPaperClass, ok bool)
}) {
	t.BaseTheme = from.GetBaseTheme()
	t.AccentColor = from.GetAccentColor()
	if val, ok := from.GetMessageTopColor(); ok {
		t.MessageTopColor = val
	}

	if val, ok := from.GetMessageBottomColor(); ok {
		t.MessageBottomColor = val
	}

	if val, ok := from.GetWallpaper(); ok {
		t.Wallpaper = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThemeSettings) TypeID() uint32 {
	return ThemeSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ThemeSettings) TypeName() string {
	return "themeSettings"
}

// TypeInfo returns info about TL type.
func (t *ThemeSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "themeSettings",
		ID:   ThemeSettingsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BaseTheme",
			SchemaName: "base_theme",
		},
		{
			Name:       "AccentColor",
			SchemaName: "accent_color",
		},
		{
			Name:       "MessageTopColor",
			SchemaName: "message_top_color",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "MessageBottomColor",
			SchemaName: "message_bottom_color",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Wallpaper",
			SchemaName: "wallpaper",
			Null:       !t.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThemeSettings) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#9c14984a as nil")
	}
	b.PutID(ThemeSettingsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThemeSettings) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#9c14984a as nil")
	}
	if !(t.MessageTopColor == 0) {
		t.Flags.Set(0)
	}
	if !(t.MessageBottomColor == 0) {
		t.Flags.Set(0)
	}
	if !(t.Wallpaper == nil) {
		t.Flags.Set(1)
	}
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#9c14984a: field flags: %w", err)
	}
	if t.BaseTheme == nil {
		return fmt.Errorf("unable to encode themeSettings#9c14984a: field base_theme is nil")
	}
	if err := t.BaseTheme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#9c14984a: field base_theme: %w", err)
	}
	b.PutInt(t.AccentColor)
	if t.Flags.Has(0) {
		b.PutInt(t.MessageTopColor)
	}
	if t.Flags.Has(0) {
		b.PutInt(t.MessageBottomColor)
	}
	if t.Flags.Has(1) {
		if t.Wallpaper == nil {
			return fmt.Errorf("unable to encode themeSettings#9c14984a: field wallpaper is nil")
		}
		if err := t.Wallpaper.Encode(b); err != nil {
			return fmt.Errorf("unable to encode themeSettings#9c14984a: field wallpaper: %w", err)
		}
	}
	return nil
}

// GetBaseTheme returns value of BaseTheme field.
func (t *ThemeSettings) GetBaseTheme() (value BaseThemeClass) {
	return t.BaseTheme
}

// GetAccentColor returns value of AccentColor field.
func (t *ThemeSettings) GetAccentColor() (value int) {
	return t.AccentColor
}

// SetMessageTopColor sets value of MessageTopColor conditional field.
func (t *ThemeSettings) SetMessageTopColor(value int) {
	t.Flags.Set(0)
	t.MessageTopColor = value
}

// GetMessageTopColor returns value of MessageTopColor conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetMessageTopColor() (value int, ok bool) {
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.MessageTopColor, true
}

// SetMessageBottomColor sets value of MessageBottomColor conditional field.
func (t *ThemeSettings) SetMessageBottomColor(value int) {
	t.Flags.Set(0)
	t.MessageBottomColor = value
}

// GetMessageBottomColor returns value of MessageBottomColor conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetMessageBottomColor() (value int, ok bool) {
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.MessageBottomColor, true
}

// SetWallpaper sets value of Wallpaper conditional field.
func (t *ThemeSettings) SetWallpaper(value WallPaperClass) {
	t.Flags.Set(1)
	t.Wallpaper = value
}

// GetWallpaper returns value of Wallpaper conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetWallpaper() (value WallPaperClass, ok bool) {
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.Wallpaper, true
}

// Decode implements bin.Decoder.
func (t *ThemeSettings) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#9c14984a to nil")
	}
	if err := b.ConsumeID(ThemeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode themeSettings#9c14984a: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThemeSettings) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#9c14984a to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field flags: %w", err)
		}
	}
	{
		value, err := DecodeBaseTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field base_theme: %w", err)
		}
		t.BaseTheme = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field accent_color: %w", err)
		}
		t.AccentColor = value
	}
	if t.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field message_top_color: %w", err)
		}
		t.MessageTopColor = value
	}
	if t.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field message_bottom_color: %w", err)
		}
		t.MessageBottomColor = value
	}
	if t.Flags.Has(1) {
		value, err := DecodeWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field wallpaper: %w", err)
		}
		t.Wallpaper = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ThemeSettings.
var (
	_ bin.Encoder     = &ThemeSettings{}
	_ bin.Decoder     = &ThemeSettings{}
	_ bin.BareEncoder = &ThemeSettings{}
	_ bin.BareDecoder = &ThemeSettings{}
)
