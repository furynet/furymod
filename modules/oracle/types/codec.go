package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided Amino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateFeed{}, "furymod/oracle/MsgCreateFeed", nil)
	cdc.RegisterConcrete(&MsgStartFeed{}, "furymod/oracle/MsgStartFeed", nil)
	cdc.RegisterConcrete(&MsgPauseFeed{}, "furymod/oracle/MsgPauseFeed", nil)
	cdc.RegisterConcrete(&MsgEditFeed{}, "furymod/oracle/MsgEditFeed", nil)

	cdc.RegisterConcrete(&Feed{}, "furymod/oracle/Feed", nil)
	cdc.RegisterConcrete(&FeedContext{}, "furymod/oracle/FeedContext", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateFeed{},
		&MsgStartFeed{},
		&MsgPauseFeed{},
		&MsgEditFeed{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
