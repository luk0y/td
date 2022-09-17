// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ClearRecentEmojiStatusesRequest represents TL type `clearRecentEmojiStatuses#e671cb5e`.
type ClearRecentEmojiStatusesRequest struct {
}

// ClearRecentEmojiStatusesRequestTypeID is TL type id of ClearRecentEmojiStatusesRequest.
const ClearRecentEmojiStatusesRequestTypeID = 0xe671cb5e

// Ensuring interfaces in compile-time for ClearRecentEmojiStatusesRequest.
var (
	_ bin.Encoder     = &ClearRecentEmojiStatusesRequest{}
	_ bin.Decoder     = &ClearRecentEmojiStatusesRequest{}
	_ bin.BareEncoder = &ClearRecentEmojiStatusesRequest{}
	_ bin.BareDecoder = &ClearRecentEmojiStatusesRequest{}
)

func (c *ClearRecentEmojiStatusesRequest) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClearRecentEmojiStatusesRequest) String() string {
	if c == nil {
		return "ClearRecentEmojiStatusesRequest(nil)"
	}
	type Alias ClearRecentEmojiStatusesRequest
	return fmt.Sprintf("ClearRecentEmojiStatusesRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ClearRecentEmojiStatusesRequest) TypeID() uint32 {
	return ClearRecentEmojiStatusesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ClearRecentEmojiStatusesRequest) TypeName() string {
	return "clearRecentEmojiStatuses"
}

// TypeInfo returns info about TL type.
func (c *ClearRecentEmojiStatusesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "clearRecentEmojiStatuses",
		ID:   ClearRecentEmojiStatusesRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ClearRecentEmojiStatusesRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clearRecentEmojiStatuses#e671cb5e as nil")
	}
	b.PutID(ClearRecentEmojiStatusesRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ClearRecentEmojiStatusesRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clearRecentEmojiStatuses#e671cb5e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ClearRecentEmojiStatusesRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clearRecentEmojiStatuses#e671cb5e to nil")
	}
	if err := b.ConsumeID(ClearRecentEmojiStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode clearRecentEmojiStatuses#e671cb5e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ClearRecentEmojiStatusesRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clearRecentEmojiStatuses#e671cb5e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ClearRecentEmojiStatusesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode clearRecentEmojiStatuses#e671cb5e as nil")
	}
	b.ObjStart()
	b.PutID("clearRecentEmojiStatuses")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ClearRecentEmojiStatusesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode clearRecentEmojiStatuses#e671cb5e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("clearRecentEmojiStatuses"); err != nil {
				return fmt.Errorf("unable to decode clearRecentEmojiStatuses#e671cb5e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ClearRecentEmojiStatuses invokes method clearRecentEmojiStatuses#e671cb5e returning error if any.
func (c *Client) ClearRecentEmojiStatuses(ctx context.Context) error {
	var ok Ok

	request := &ClearRecentEmojiStatusesRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
