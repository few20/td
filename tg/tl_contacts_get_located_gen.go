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

// ContactsGetLocatedRequest represents TL type `contacts.getLocated#d348bc44`.
// Get contacts near you
//
// See https://core.telegram.org/method/contacts.getLocated for reference.
type ContactsGetLocatedRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// While the geolocation of the current user is public, clients should update it in the background every half-an-hour or so, while setting this flag. Do this only if the new location is more than 1 KM away from the previous one, or if the previous location is unknown.
	Background bool
	// Geolocation
	GeoPoint InputGeoPointClass
	// If set, the geolocation of the current user will be public for the specified number of seconds; pass 0x7fffffff to disable expiry, 0 to make the current geolocation private; if the flag isn't set, no changes will be applied.
	//
	// Use SetSelfExpires and GetSelfExpires helpers.
	SelfExpires int
}

// ContactsGetLocatedRequestTypeID is TL type id of ContactsGetLocatedRequest.
const ContactsGetLocatedRequestTypeID = 0xd348bc44

func (g *ContactsGetLocatedRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Background == false) {
		return false
	}
	if !(g.GeoPoint == nil) {
		return false
	}
	if !(g.SelfExpires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetLocatedRequest) String() string {
	if g == nil {
		return "ContactsGetLocatedRequest(nil)"
	}
	type Alias ContactsGetLocatedRequest
	return fmt.Sprintf("ContactsGetLocatedRequest%+v", Alias(*g))
}

// FillFrom fills ContactsGetLocatedRequest from given interface.
func (g *ContactsGetLocatedRequest) FillFrom(from interface {
	GetBackground() (value bool)
	GetGeoPoint() (value InputGeoPointClass)
	GetSelfExpires() (value int, ok bool)
}) {
	g.Background = from.GetBackground()
	g.GeoPoint = from.GetGeoPoint()
	if val, ok := from.GetSelfExpires(); ok {
		g.SelfExpires = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetLocatedRequest) TypeID() uint32 {
	return ContactsGetLocatedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetLocatedRequest) TypeName() string {
	return "contacts.getLocated"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetLocatedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getLocated",
		ID:   ContactsGetLocatedRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Background",
			SchemaName: "background",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "GeoPoint",
			SchemaName: "geo_point",
		},
		{
			Name:       "SelfExpires",
			SchemaName: "self_expires",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ContactsGetLocatedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getLocated#d348bc44 as nil")
	}
	b.PutID(ContactsGetLocatedRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ContactsGetLocatedRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getLocated#d348bc44 as nil")
	}
	if !(g.Background == false) {
		g.Flags.Set(1)
	}
	if !(g.SelfExpires == 0) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getLocated#d348bc44: field flags: %w", err)
	}
	if g.GeoPoint == nil {
		return fmt.Errorf("unable to encode contacts.getLocated#d348bc44: field geo_point is nil")
	}
	if err := g.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getLocated#d348bc44: field geo_point: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutInt(g.SelfExpires)
	}
	return nil
}

// SetBackground sets value of Background conditional field.
func (g *ContactsGetLocatedRequest) SetBackground(value bool) {
	if value {
		g.Flags.Set(1)
		g.Background = true
	} else {
		g.Flags.Unset(1)
		g.Background = false
	}
}

// GetBackground returns value of Background conditional field.
func (g *ContactsGetLocatedRequest) GetBackground() (value bool) {
	return g.Flags.Has(1)
}

// GetGeoPoint returns value of GeoPoint field.
func (g *ContactsGetLocatedRequest) GetGeoPoint() (value InputGeoPointClass) {
	return g.GeoPoint
}

// GetGeoPointAsNotEmpty returns mapped value of GeoPoint field.
func (g *ContactsGetLocatedRequest) GetGeoPointAsNotEmpty() (*InputGeoPoint, bool) {
	return g.GeoPoint.AsNotEmpty()
}

// SetSelfExpires sets value of SelfExpires conditional field.
func (g *ContactsGetLocatedRequest) SetSelfExpires(value int) {
	g.Flags.Set(0)
	g.SelfExpires = value
}

// GetSelfExpires returns value of SelfExpires conditional field and
// boolean which is true if field was set.
func (g *ContactsGetLocatedRequest) GetSelfExpires() (value int, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.SelfExpires, true
}

// Decode implements bin.Decoder.
func (g *ContactsGetLocatedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getLocated#d348bc44 to nil")
	}
	if err := b.ConsumeID(ContactsGetLocatedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ContactsGetLocatedRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getLocated#d348bc44 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: field flags: %w", err)
		}
	}
	g.Background = g.Flags.Has(1)
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: field geo_point: %w", err)
		}
		g.GeoPoint = value
	}
	if g.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: field self_expires: %w", err)
		}
		g.SelfExpires = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetLocatedRequest.
var (
	_ bin.Encoder     = &ContactsGetLocatedRequest{}
	_ bin.Decoder     = &ContactsGetLocatedRequest{}
	_ bin.BareEncoder = &ContactsGetLocatedRequest{}
	_ bin.BareDecoder = &ContactsGetLocatedRequest{}
)

// ContactsGetLocated invokes method contacts.getLocated#d348bc44 returning error if any.
// Get contacts near you
//
// Possible errors:
//  400 GEO_POINT_INVALID: Invalid geoposition provided
//  406 USERPIC_UPLOAD_REQUIRED: You must have a profile picture to publish your geolocation
//
// See https://core.telegram.org/method/contacts.getLocated for reference.
func (c *Client) ContactsGetLocated(ctx context.Context, request *ContactsGetLocatedRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
