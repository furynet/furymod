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

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*TokenI)(nil), nil)

	cdc.RegisterConcrete(&Token{}, "furymod/token/Token", nil)

	cdc.RegisterConcrete(&MsgIssueToken{}, "furymod/token/MsgIssueToken", nil)
	cdc.RegisterConcrete(&MsgEditToken{}, "furymod/token/MsgEditToken", nil)
	cdc.RegisterConcrete(&MsgMintToken{}, "furymod/token/MsgMintToken", nil)
	cdc.RegisterConcrete(&MsgBurnToken{}, "furymod/token/MsgBurnToken", nil)
	cdc.RegisterConcrete(&MsgTransferTokenOwner{}, "furymod/token/MsgTransferTokenOwner", nil)
	cdc.RegisterConcrete(&MsgSwapFeeToken{}, "furymod/token/MsgSwapFeeToken", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueToken{},
		&MsgEditToken{},
		&MsgMintToken{},
		&MsgBurnToken{},
		&MsgTransferTokenOwner{},
		&MsgSwapFeeToken{},
	)
	registry.RegisterInterface(
		"furymod.token.TokenI",
		(*TokenI)(nil),
		&Token{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
