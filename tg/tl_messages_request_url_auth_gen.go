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

// MessagesRequestUrlAuthRequest represents TL type `messages.requestUrlAuth#e33f5613`.
// Get more info about a Seamless Telegram Login authorization request, for more info click here »
//
// See https://core.telegram.org/method/messages.requestUrlAuth for reference.
type MessagesRequestUrlAuthRequest struct {
	// Peer where the message is located
	Peer InputPeerClass
	// The message
	MsgID int
	// The ID of the button with the authorization request
	ButtonID int
}

// MessagesRequestUrlAuthRequestTypeID is TL type id of MessagesRequestUrlAuthRequest.
const MessagesRequestUrlAuthRequestTypeID = 0xe33f5613

// String implements fmt.Stringer.
func (r *MessagesRequestUrlAuthRequest) String() string {
	if r == nil {
		return "MessagesRequestUrlAuthRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesRequestUrlAuthRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(r.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(r.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tButtonID: ")
	sb.WriteString(fmt.Sprint(r.ButtonID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *MessagesRequestUrlAuthRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestUrlAuth#e33f5613 as nil")
	}
	b.PutID(MessagesRequestUrlAuthRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.requestUrlAuth#e33f5613: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestUrlAuth#e33f5613: field peer: %w", err)
	}
	b.PutInt(r.MsgID)
	b.PutInt(r.ButtonID)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestUrlAuthRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestUrlAuth#e33f5613 to nil")
	}
	if err := b.ConsumeID(MessagesRequestUrlAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestUrlAuth#e33f5613: field button_id: %w", err)
		}
		r.ButtonID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesRequestUrlAuthRequest.
var (
	_ bin.Encoder = &MessagesRequestUrlAuthRequest{}
	_ bin.Decoder = &MessagesRequestUrlAuthRequest{}
)

// MessagesRequestUrlAuth invokes method messages.requestUrlAuth#e33f5613 returning error if any.
// Get more info about a Seamless Telegram Login authorization request, for more info click here »
//
// See https://core.telegram.org/method/messages.requestUrlAuth for reference.
func (c *Client) MessagesRequestUrlAuth(ctx context.Context, request *MessagesRequestUrlAuthRequest) (UrlAuthResultClass, error) {
	var result UrlAuthResultBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UrlAuthResult, nil
}
