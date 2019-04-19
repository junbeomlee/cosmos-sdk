//nolint
package gov

import "github.com/cosmos/cosmos-sdk/x/gov/types"

const (
	StatusNil           = types.StatusNil
	StatusDepositPeriod = types.StatusDepositPeriod
	StatusVotingPeriod  = types.StatusVotingPeriod
	StatusPassed        = types.StatusPassed
	StatusRejected      = types.StatusRejected

	DefaultCodespace  = types.DefaultCodespace
	DefaultParamspace = types.DefaultParamspace
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute

	ProposalTypeText            = types.ProposalTypeText
	ProposalTypeSoftwareUpgrade = types.ProposalTypeSoftwareUpgrade

	OptionEmpty      = types.OptionEmpty
	OptionYes        = types.OptionYes
	OptionAbstain    = types.OptionAbstain
	OptionNo         = types.OptionNo
	OptionNoWithVeto = types.OptionNoWithVeto
)

type (
	Content   = types.Content
	Handler   = types.Handler
	Proposal  = types.Proposal
	Proposals = types.Proposals

	Deposit    = types.Deposit
	Vote       = types.Vote
	VoteOption = types.VoteOption

	TextProposal            = types.TextProposal
	SoftwareUpgradeProposal = types.SoftwareUpgradeProposal
	ProposalStatus          = types.ProposalStatus

	MsgSubmitProposal = types.MsgSubmitProposal
	MsgDeposit        = types.MsgDeposit
	MsgVote           = types.MsgVote

	TallyResult = types.TallyResult
)

var (
	ErrUnknownProposal         = types.ErrUnknownProposal
	ErrInactiveProposal        = types.ErrInactiveProposal
	ErrAlreadyActiveProposal   = types.ErrAlreadyActiveProposal
	ErrAlreadyFinishedProposal = types.ErrAlreadyFinishedProposal
	ErrAddressNotStaked        = types.ErrAddressNotStaked
	ErrInvalidTitle            = types.ErrInvalidTitle
	ErrInvalidDescription      = types.ErrInvalidDescription
	ErrInvalidProposalType     = types.ErrInvalidProposalType
	ErrInvalidVote             = types.ErrInvalidVote
	ErrInvalidGenesis          = types.ErrInvalidGenesis
	ErrNoProposalHandlerExists = types.ErrNoProposalHandlerExists

	NewProposal         = types.NewProposal
	ProposalHandler     = types.ProposalHandler
	ValidVoteOption     = types.ValidVoteOption
	ValidProposalStatus = types.ValidProposalStatus

	ValidateAbstract = types.ValidateAbstract
	EmptyTallyResult = types.EmptyTallyResult

	RegisterCodec = types.RegisterCodec

	KeyDelimiter                = types.KeyDelimiter
	KeyNextProposalID           = types.KeyNextProposalID
	PrefixActiveProposalQueue   = types.PrefixActiveProposalQueue
	PrefixInactiveProposalQueue = types.PrefixInactiveProposalQueue

	KeyProposal                      = types.KeyProposal
	KeyDeposit                       = types.KeyDeposit
	KeyVote                          = types.KeyVote
	KeyDepositsSubspace              = types.KeyDepositsSubspace
	KeyVotesSubspace                 = types.KeyVotesSubspace
	PrefixActiveProposalQueueTime    = types.PrefixActiveProposalQueueTime
	KeyActiveProposalQueueProposal   = types.KeyActiveProposalQueueProposal
	PrefixInactiveProposalQueueTime  = types.PrefixInactiveProposalQueueTime
	KeyInactiveProposalQueueProposal = types.KeyInactiveProposalQueueProposal

	NewMsgSubmitProposal  = types.NewMsgSubmitProposal
	NewMsgDeposit         = types.NewMsgDeposit
	NewMsgVote            = types.NewMsgVote
	NewTextProposal       = types.NewTextProposal
	NewTallyResultFromMap = types.NewTallyResultFromMap

	ContentFromProposalType = types.ContentFromProposalType
)
