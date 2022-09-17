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

// PremiumState represents TL type `premiumState#6d4d8ec0`.
type PremiumState struct {
	// Text description of the state of the current Premium subscription; may be empty if the
	// current user has no Telegram Premium subscription
	State FormattedText
	// The list of available options for buying Telegram Premium
	PaymentOptions []PremiumPaymentOption
	// The list of available promotion animations for Premium features
	Animations []PremiumFeaturePromotionAnimation
}

// PremiumStateTypeID is TL type id of PremiumState.
const PremiumStateTypeID = 0x6d4d8ec0

// Ensuring interfaces in compile-time for PremiumState.
var (
	_ bin.Encoder     = &PremiumState{}
	_ bin.Decoder     = &PremiumState{}
	_ bin.BareEncoder = &PremiumState{}
	_ bin.BareDecoder = &PremiumState{}
)

func (p *PremiumState) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.State.Zero()) {
		return false
	}
	if !(p.PaymentOptions == nil) {
		return false
	}
	if !(p.Animations == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumState) String() string {
	if p == nil {
		return "PremiumState(nil)"
	}
	type Alias PremiumState
	return fmt.Sprintf("PremiumState%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumState) TypeID() uint32 {
	return PremiumStateTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumState) TypeName() string {
	return "premiumState"
}

// TypeInfo returns info about TL type.
func (p *PremiumState) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumState",
		ID:   PremiumStateTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "PaymentOptions",
			SchemaName: "payment_options",
		},
		{
			Name:       "Animations",
			SchemaName: "animations",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumState) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumState#6d4d8ec0 as nil")
	}
	b.PutID(PremiumStateTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumState) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumState#6d4d8ec0 as nil")
	}
	if err := p.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumState#6d4d8ec0: field state: %w", err)
	}
	b.PutInt(len(p.PaymentOptions))
	for idx, v := range p.PaymentOptions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare premiumState#6d4d8ec0: field payment_options element with index %d: %w", idx, err)
		}
	}
	b.PutInt(len(p.Animations))
	for idx, v := range p.Animations {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare premiumState#6d4d8ec0: field animations element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumState) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumState#6d4d8ec0 to nil")
	}
	if err := b.ConsumeID(PremiumStateTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumState#6d4d8ec0: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumState) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumState#6d4d8ec0 to nil")
	}
	{
		if err := p.State.Decode(b); err != nil {
			return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field state: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field payment_options: %w", err)
		}

		if headerLen > 0 {
			p.PaymentOptions = make([]PremiumPaymentOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PremiumPaymentOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare premiumState#6d4d8ec0: field payment_options: %w", err)
			}
			p.PaymentOptions = append(p.PaymentOptions, value)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field animations: %w", err)
		}

		if headerLen > 0 {
			p.Animations = make([]PremiumFeaturePromotionAnimation, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PremiumFeaturePromotionAnimation
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare premiumState#6d4d8ec0: field animations: %w", err)
			}
			p.Animations = append(p.Animations, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumState) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumState#6d4d8ec0 as nil")
	}
	b.ObjStart()
	b.PutID("premiumState")
	b.Comma()
	b.FieldStart("state")
	if err := p.State.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode premiumState#6d4d8ec0: field state: %w", err)
	}
	b.Comma()
	b.FieldStart("payment_options")
	b.ArrStart()
	for idx, v := range p.PaymentOptions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode premiumState#6d4d8ec0: field payment_options element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("animations")
	b.ArrStart()
	for idx, v := range p.Animations {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode premiumState#6d4d8ec0: field animations element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumState) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumState#6d4d8ec0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumState"); err != nil {
				return fmt.Errorf("unable to decode premiumState#6d4d8ec0: %w", err)
			}
		case "state":
			if err := p.State.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field state: %w", err)
			}
		case "payment_options":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value PremiumPaymentOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field payment_options: %w", err)
				}
				p.PaymentOptions = append(p.PaymentOptions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field payment_options: %w", err)
			}
		case "animations":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value PremiumFeaturePromotionAnimation
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field animations: %w", err)
				}
				p.Animations = append(p.Animations, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode premiumState#6d4d8ec0: field animations: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetState returns value of State field.
func (p *PremiumState) GetState() (value FormattedText) {
	if p == nil {
		return
	}
	return p.State
}

// GetPaymentOptions returns value of PaymentOptions field.
func (p *PremiumState) GetPaymentOptions() (value []PremiumPaymentOption) {
	if p == nil {
		return
	}
	return p.PaymentOptions
}

// GetAnimations returns value of Animations field.
func (p *PremiumState) GetAnimations() (value []PremiumFeaturePromotionAnimation) {
	if p == nil {
		return
	}
	return p.Animations
}
