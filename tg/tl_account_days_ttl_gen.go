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

// AccountDaysTTL represents TL type `accountDaysTTL#b8d0afdf`.
// Time to live in days of the current account
//
// See https://core.telegram.org/constructor/accountDaysTTL for reference.
type AccountDaysTTL struct {
	// This account will self-destruct in the specified number of days
	Days int
}

// AccountDaysTTLTypeID is TL type id of AccountDaysTTL.
const AccountDaysTTLTypeID = 0xb8d0afdf

// String implements fmt.Stringer.
func (a *AccountDaysTTL) String() string {
	if a == nil {
		return "AccountDaysTTL(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountDaysTTL")
	sb.WriteString("{\n")
	sb.WriteString("\tDays: ")
	sb.WriteString(fmt.Sprint(a.Days))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *AccountDaysTTL) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode accountDaysTTL#b8d0afdf as nil")
	}
	b.PutID(AccountDaysTTLTypeID)
	b.PutInt(a.Days)
	return nil
}

// Decode implements bin.Decoder.
func (a *AccountDaysTTL) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode accountDaysTTL#b8d0afdf to nil")
	}
	if err := b.ConsumeID(AccountDaysTTLTypeID); err != nil {
		return fmt.Errorf("unable to decode accountDaysTTL#b8d0afdf: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode accountDaysTTL#b8d0afdf: field days: %w", err)
		}
		a.Days = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountDaysTTL.
var (
	_ bin.Encoder = &AccountDaysTTL{}
	_ bin.Decoder = &AccountDaysTTL{}
)
