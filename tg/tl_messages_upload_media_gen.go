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

// MessagesUploadMediaRequest represents TL type `messages.uploadMedia#519bc2b1`.
// Upload a file and associate it to a chat (without actually sending it to the chat)
//
// See https://core.telegram.org/method/messages.uploadMedia for reference.
type MessagesUploadMediaRequest struct {
	// The chat, can be an inputPeerEmpty¹ for bots
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputPeerEmpty
	Peer InputPeerClass
	// File uploaded in chunks as described in files »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	Media InputMediaClass
}

// MessagesUploadMediaRequestTypeID is TL type id of MessagesUploadMediaRequest.
const MessagesUploadMediaRequestTypeID = 0x519bc2b1

func (u *MessagesUploadMediaRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Peer == nil) {
		return false
	}
	if !(u.Media == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUploadMediaRequest) String() string {
	if u == nil {
		return "MessagesUploadMediaRequest(nil)"
	}
	type Alias MessagesUploadMediaRequest
	return fmt.Sprintf("MessagesUploadMediaRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUploadMediaRequest from given interface.
func (u *MessagesUploadMediaRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMedia() (value InputMediaClass)
}) {
	u.Peer = from.GetPeer()
	u.Media = from.GetMedia()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUploadMediaRequest) TypeID() uint32 {
	return MessagesUploadMediaRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUploadMediaRequest) TypeName() string {
	return "messages.uploadMedia"
}

// TypeInfo returns info about TL type.
func (u *MessagesUploadMediaRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.uploadMedia",
		ID:   MessagesUploadMediaRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Media",
			SchemaName: "media",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *MessagesUploadMediaRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uploadMedia#519bc2b1 as nil")
	}
	b.PutID(MessagesUploadMediaRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *MessagesUploadMediaRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.uploadMedia#519bc2b1 as nil")
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field peer: %w", err)
	}
	if u.Media == nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field media is nil")
	}
	if err := u.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.uploadMedia#519bc2b1: field media: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (u *MessagesUploadMediaRequest) GetPeer() (value InputPeerClass) {
	return u.Peer
}

// GetMedia returns value of Media field.
func (u *MessagesUploadMediaRequest) GetMedia() (value InputMediaClass) {
	return u.Media
}

// Decode implements bin.Decoder.
func (u *MessagesUploadMediaRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uploadMedia#519bc2b1 to nil")
	}
	if err := b.ConsumeID(MessagesUploadMediaRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.uploadMedia#519bc2b1: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *MessagesUploadMediaRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.uploadMedia#519bc2b1 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uploadMedia#519bc2b1: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.uploadMedia#519bc2b1: field media: %w", err)
		}
		u.Media = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUploadMediaRequest.
var (
	_ bin.Encoder     = &MessagesUploadMediaRequest{}
	_ bin.Decoder     = &MessagesUploadMediaRequest{}
	_ bin.BareEncoder = &MessagesUploadMediaRequest{}
	_ bin.BareDecoder = &MessagesUploadMediaRequest{}
)

// MessagesUploadMedia invokes method messages.uploadMedia#519bc2b1 returning error if any.
// Upload a file and associate it to a chat (without actually sending it to the chat)
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 FILE_PARTS_INVALID: The number of file parts is invalid
//  400 IMAGE_PROCESS_FAILED: Failure while processing image
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MEDIA_INVALID: Media invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 PHOTO_EXT_INVALID: The extension of the photo is invalid
//  400 PHOTO_SAVE_FILE_INVALID: Internal issues, try again later
//  400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels
//  400 WEBPAGE_CURL_FAILED: Failure while fetching the webpage with cURL
//
// See https://core.telegram.org/method/messages.uploadMedia for reference.
// Can be used by bots.
func (c *Client) MessagesUploadMedia(ctx context.Context, request *MessagesUploadMediaRequest) (MessageMediaClass, error) {
	var result MessageMediaBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.MessageMedia, nil
}
