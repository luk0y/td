// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesTranslateTextRequest represents TL type `messages.translateText#24ce6dee`.
// Translate a given text
//
// See https://core.telegram.org/method/messages.translateText for reference.
type MessagesTranslateTextRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If the text is a chat message, the peer ID
	//
	// Use SetPeer and GetPeer helpers.
	Peer InputPeerClass
	// If the text is a chat message, the message ID
	//
	// Use SetMsgID and GetMsgID helpers.
	MsgID int
	// The text to translate
	//
	// Use SetText and GetText helpers.
	Text string
	// Two-letter ISO 639-1 language code of the language from which the message is
	// translated, if not set will be autodetected
	//
	// Use SetFromLang and GetFromLang helpers.
	FromLang string
	// Two-letter ISO 639-1 language code of the language to which the message is translated
	ToLang string
}

// MessagesTranslateTextRequestTypeID is TL type id of MessagesTranslateTextRequest.
const MessagesTranslateTextRequestTypeID = 0x24ce6dee

// Ensuring interfaces in compile-time for MessagesTranslateTextRequest.
var (
	_ bin.Encoder     = &MessagesTranslateTextRequest{}
	_ bin.Decoder     = &MessagesTranslateTextRequest{}
	_ bin.BareEncoder = &MessagesTranslateTextRequest{}
	_ bin.BareDecoder = &MessagesTranslateTextRequest{}
)

func (t *MessagesTranslateTextRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.MsgID == 0) {
		return false
	}
	if !(t.Text == "") {
		return false
	}
	if !(t.FromLang == "") {
		return false
	}
	if !(t.ToLang == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesTranslateTextRequest) String() string {
	if t == nil {
		return "MessagesTranslateTextRequest(nil)"
	}
	type Alias MessagesTranslateTextRequest
	return fmt.Sprintf("MessagesTranslateTextRequest%+v", Alias(*t))
}

// FillFrom fills MessagesTranslateTextRequest from given interface.
func (t *MessagesTranslateTextRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass, ok bool)
	GetMsgID() (value int, ok bool)
	GetText() (value string, ok bool)
	GetFromLang() (value string, ok bool)
	GetToLang() (value string)
}) {
	if val, ok := from.GetPeer(); ok {
		t.Peer = val
	}

	if val, ok := from.GetMsgID(); ok {
		t.MsgID = val
	}

	if val, ok := from.GetText(); ok {
		t.Text = val
	}

	if val, ok := from.GetFromLang(); ok {
		t.FromLang = val
	}

	t.ToLang = from.GetToLang()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesTranslateTextRequest) TypeID() uint32 {
	return MessagesTranslateTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesTranslateTextRequest) TypeName() string {
	return "messages.translateText"
}

// TypeInfo returns info about TL type.
func (t *MessagesTranslateTextRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.translateText",
		ID:   MessagesTranslateTextRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Text",
			SchemaName: "text",
			Null:       !t.Flags.Has(1),
		},
		{
			Name:       "FromLang",
			SchemaName: "from_lang",
			Null:       !t.Flags.Has(2),
		},
		{
			Name:       "ToLang",
			SchemaName: "to_lang",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesTranslateTextRequest) SetFlags() {
	if !(t.Peer == nil) {
		t.Flags.Set(0)
	}
	if !(t.MsgID == 0) {
		t.Flags.Set(0)
	}
	if !(t.Text == "") {
		t.Flags.Set(1)
	}
	if !(t.FromLang == "") {
		t.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesTranslateTextRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateText#24ce6dee as nil")
	}
	b.PutID(MessagesTranslateTextRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesTranslateTextRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.translateText#24ce6dee as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.translateText#24ce6dee: field flags: %w", err)
	}
	if t.Flags.Has(0) {
		if t.Peer == nil {
			return fmt.Errorf("unable to encode messages.translateText#24ce6dee: field peer is nil")
		}
		if err := t.Peer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.translateText#24ce6dee: field peer: %w", err)
		}
	}
	if t.Flags.Has(0) {
		b.PutInt(t.MsgID)
	}
	if t.Flags.Has(1) {
		b.PutString(t.Text)
	}
	if t.Flags.Has(2) {
		b.PutString(t.FromLang)
	}
	b.PutString(t.ToLang)
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesTranslateTextRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateText#24ce6dee to nil")
	}
	if err := b.ConsumeID(MessagesTranslateTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.translateText#24ce6dee: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesTranslateTextRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.translateText#24ce6dee to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.translateText#24ce6dee: field flags: %w", err)
		}
	}
	if t.Flags.Has(0) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#24ce6dee: field peer: %w", err)
		}
		t.Peer = value
	}
	if t.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#24ce6dee: field msg_id: %w", err)
		}
		t.MsgID = value
	}
	if t.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#24ce6dee: field text: %w", err)
		}
		t.Text = value
	}
	if t.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#24ce6dee: field from_lang: %w", err)
		}
		t.FromLang = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.translateText#24ce6dee: field to_lang: %w", err)
		}
		t.ToLang = value
	}
	return nil
}

// SetPeer sets value of Peer conditional field.
func (t *MessagesTranslateTextRequest) SetPeer(value InputPeerClass) {
	t.Flags.Set(0)
	t.Peer = value
}

// GetPeer returns value of Peer conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetPeer() (value InputPeerClass, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.Peer, true
}

// SetMsgID sets value of MsgID conditional field.
func (t *MessagesTranslateTextRequest) SetMsgID(value int) {
	t.Flags.Set(0)
	t.MsgID = value
}

// GetMsgID returns value of MsgID conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetMsgID() (value int, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.MsgID, true
}

// SetText sets value of Text conditional field.
func (t *MessagesTranslateTextRequest) SetText(value string) {
	t.Flags.Set(1)
	t.Text = value
}

// GetText returns value of Text conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetText() (value string, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.Text, true
}

// SetFromLang sets value of FromLang conditional field.
func (t *MessagesTranslateTextRequest) SetFromLang(value string) {
	t.Flags.Set(2)
	t.FromLang = value
}

// GetFromLang returns value of FromLang conditional field and
// boolean which is true if field was set.
func (t *MessagesTranslateTextRequest) GetFromLang() (value string, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(2) {
		return value, false
	}
	return t.FromLang, true
}

// GetToLang returns value of ToLang field.
func (t *MessagesTranslateTextRequest) GetToLang() (value string) {
	if t == nil {
		return
	}
	return t.ToLang
}

// MessagesTranslateText invokes method messages.translateText#24ce6dee returning error if any.
// Translate a given text
//
// Possible errors:
//
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 TO_LANG_INVALID: The specified destination language is invalid.
//
// See https://core.telegram.org/method/messages.translateText for reference.
func (c *Client) MessagesTranslateText(ctx context.Context, request *MessagesTranslateTextRequest) (MessagesTranslatedTextClass, error) {
	var result MessagesTranslatedTextBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TranslatedText, nil
}
