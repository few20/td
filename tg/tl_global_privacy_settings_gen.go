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

// GlobalPrivacySettings represents TL type `globalPrivacySettings#bea2f424`.
// Global privacy settings
//
// See https://core.telegram.org/constructor/globalPrivacySettings for reference.
type GlobalPrivacySettings struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether to archive and mute new chats from non-contacts
	ArchiveAndMuteNewNoncontactPeers bool
}

// GlobalPrivacySettingsTypeID is TL type id of GlobalPrivacySettings.
const GlobalPrivacySettingsTypeID = 0xbea2f424

// String implements fmt.Stringer.
func (g *GlobalPrivacySettings) String() string {
	if g == nil {
		return "GlobalPrivacySettings(nil)"
	}
	var sb strings.Builder
	sb.WriteString("GlobalPrivacySettings")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(g.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *GlobalPrivacySettings) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#bea2f424 as nil")
	}
	b.PutID(GlobalPrivacySettingsTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode globalPrivacySettings#bea2f424: field flags: %w", err)
	}
	return nil
}

// SetArchiveAndMuteNewNoncontactPeers sets value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) SetArchiveAndMuteNewNoncontactPeers(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (g *GlobalPrivacySettings) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#bea2f424 to nil")
	}
	if err := b.ConsumeID(GlobalPrivacySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: field flags: %w", err)
		}
	}
	g.ArchiveAndMuteNewNoncontactPeers = g.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for GlobalPrivacySettings.
var (
	_ bin.Encoder = &GlobalPrivacySettings{}
	_ bin.Decoder = &GlobalPrivacySettings{}
)
