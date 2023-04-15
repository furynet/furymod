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

// RegisterLegacyAminoCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDefineService{}, "furymod/service/MsgDefineService", nil)
	cdc.RegisterConcrete(&MsgBindService{}, "furymod/service/MsgBindService", nil)
	cdc.RegisterConcrete(&MsgUpdateServiceBinding{}, "furymod/service/MsgUpdateServiceBinding", nil)
	cdc.RegisterConcrete(&MsgSetWithdrawAddress{}, "furymod/service/MsgSetWithdrawAddress", nil)
	cdc.RegisterConcrete(&MsgDisableServiceBinding{}, "furymod/service/MsgDisableServiceBinding", nil)
	cdc.RegisterConcrete(&MsgEnableServiceBinding{}, "furymod/service/MsgEnableServiceBinding", nil)
	cdc.RegisterConcrete(&MsgRefundServiceDeposit{}, "furymod/service/MsgRefundServiceDeposit", nil)

	cdc.RegisterConcrete(&MsgCallService{}, "furymod/service/MsgCallService", nil)
	cdc.RegisterConcrete(&MsgRespondService{}, "furymod/service/MsgRespondService", nil)
	cdc.RegisterConcrete(&MsgPauseRequestContext{}, "furymod/service/MsgPauseRequestContext", nil)
	cdc.RegisterConcrete(&MsgStartRequestContext{}, "furymod/service/MsgStartRequestContext", nil)
	cdc.RegisterConcrete(&MsgKillRequestContext{}, "furymod/service/MsgKillRequestContext", nil)
	cdc.RegisterConcrete(&MsgUpdateRequestContext{}, "furymod/service/MsgUpdateRequestContext", nil)
	cdc.RegisterConcrete(&MsgWithdrawEarnedFees{}, "furymod/service/MsgWithdrawEarnedFees", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDefineService{},
		&MsgBindService{},
		&MsgUpdateServiceBinding{},
		&MsgSetWithdrawAddress{},
		&MsgDisableServiceBinding{},
		&MsgEnableServiceBinding{},
		&MsgRefundServiceDeposit{},
		&MsgCallService{},
		&MsgRespondService{},
		&MsgPauseRequestContext{},
		&MsgStartRequestContext{},
		&MsgKillRequestContext{},
		&MsgUpdateRequestContext{},
		&MsgWithdrawEarnedFees{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
