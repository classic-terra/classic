package ante

import (
	ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
	channelkeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	AccountKeeper      cosmosante.AccountKeeper
	BankKeeper         BankKeeper
	FeegrantKeeper     cosmosante.FeegrantKeeper
	OracleKeeper       OracleKeeper
	TreasuryKeeper     TreasuryKeeper
	SignModeHandler    signing.SignModeHandler
	SigGasConsumer     cosmosante.SignatureVerificationGasConsumer
	IBCChannelKeeper   channelkeeper.Keeper
	DistributionKeeper distributionkeeper.Keeper
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.OracleKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "oracle keeper is required for ante builder")
	}

	if options.TreasuryKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "treasury keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = cosmosante.DefaultSigVerificationGasConsumer
	}

	return sdk.ChainAnteDecorators(
		cosmosante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		cosmosante.NewRejectExtensionOptionsDecorator(),
		NewSpammingPreventionDecorator(options.OracleKeeper), // spamming prevention
		NewTaxFeeDecorator(options.TreasuryKeeper),           // mempool gas fee validation & record tax proceeds
		cosmosante.NewValidateBasicDecorator(),
		cosmosante.NewTxTimeoutHeightDecorator(),
		cosmosante.NewValidateMemoDecorator(options.AccountKeeper),
		cosmosante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		cosmosante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		NewBurnTaxFeeDecorator(options.TreasuryKeeper, options.BankKeeper, options.DistributionKeeper), // burn tax proceeds
		cosmosante.NewSetPubKeyDecorator(options.AccountKeeper),                                        // SetPubKeyDecorator must be called before all signature verification decorators
		cosmosante.NewValidateSigCountDecorator(options.AccountKeeper),
		cosmosante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		cosmosante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(&options.IBCChannelKeeper),
	), nil
}
