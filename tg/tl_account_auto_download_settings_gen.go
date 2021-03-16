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

// AccountAutoDownloadSettings represents TL type `account.autoDownloadSettings#63cacf26`.
// Media autodownload settings
//
// See https://core.telegram.org/constructor/account.autoDownloadSettings for reference.
type AccountAutoDownloadSettings struct {
	// Low data usage preset
	Low AutoDownloadSettings
	// Medium data usage preset
	Medium AutoDownloadSettings
	// High data usage preset
	High AutoDownloadSettings
}

// AccountAutoDownloadSettingsTypeID is TL type id of AccountAutoDownloadSettings.
const AccountAutoDownloadSettingsTypeID = 0x63cacf26

func (a *AccountAutoDownloadSettings) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Low.Zero()) {
		return false
	}
	if !(a.Medium.Zero()) {
		return false
	}
	if !(a.High.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AccountAutoDownloadSettings) String() string {
	if a == nil {
		return "AccountAutoDownloadSettings(nil)"
	}
	type Alias AccountAutoDownloadSettings
	return fmt.Sprintf("AccountAutoDownloadSettings%+v", Alias(*a))
}

// FillFrom fills AccountAutoDownloadSettings from given interface.
func (a *AccountAutoDownloadSettings) FillFrom(from interface {
	GetLow() (value AutoDownloadSettings)
	GetMedium() (value AutoDownloadSettings)
	GetHigh() (value AutoDownloadSettings)
}) {
	a.Low = from.GetLow()
	a.Medium = from.GetMedium()
	a.High = from.GetHigh()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountAutoDownloadSettings) TypeID() uint32 {
	return AccountAutoDownloadSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountAutoDownloadSettings) TypeName() string {
	return "account.autoDownloadSettings"
}

// TypeInfo returns info about TL type.
func (a *AccountAutoDownloadSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.autoDownloadSettings",
		ID:   AccountAutoDownloadSettingsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Low",
			SchemaName: "low",
		},
		{
			Name:       "Medium",
			SchemaName: "medium",
		},
		{
			Name:       "High",
			SchemaName: "high",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AccountAutoDownloadSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.autoDownloadSettings#63cacf26 as nil")
	}
	b.PutID(AccountAutoDownloadSettingsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AccountAutoDownloadSettings) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.autoDownloadSettings#63cacf26 as nil")
	}
	if err := a.Low.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.autoDownloadSettings#63cacf26: field low: %w", err)
	}
	if err := a.Medium.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.autoDownloadSettings#63cacf26: field medium: %w", err)
	}
	if err := a.High.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.autoDownloadSettings#63cacf26: field high: %w", err)
	}
	return nil
}

// GetLow returns value of Low field.
func (a *AccountAutoDownloadSettings) GetLow() (value AutoDownloadSettings) {
	return a.Low
}

// GetMedium returns value of Medium field.
func (a *AccountAutoDownloadSettings) GetMedium() (value AutoDownloadSettings) {
	return a.Medium
}

// GetHigh returns value of High field.
func (a *AccountAutoDownloadSettings) GetHigh() (value AutoDownloadSettings) {
	return a.High
}

// Decode implements bin.Decoder.
func (a *AccountAutoDownloadSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.autoDownloadSettings#63cacf26 to nil")
	}
	if err := b.ConsumeID(AccountAutoDownloadSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AccountAutoDownloadSettings) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.autoDownloadSettings#63cacf26 to nil")
	}
	{
		if err := a.Low.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: field low: %w", err)
		}
	}
	{
		if err := a.Medium.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: field medium: %w", err)
		}
	}
	{
		if err := a.High.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.autoDownloadSettings#63cacf26: field high: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountAutoDownloadSettings.
var (
	_ bin.Encoder     = &AccountAutoDownloadSettings{}
	_ bin.Decoder     = &AccountAutoDownloadSettings{}
	_ bin.BareEncoder = &AccountAutoDownloadSettings{}
	_ bin.BareDecoder = &AccountAutoDownloadSettings{}
)
