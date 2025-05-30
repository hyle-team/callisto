package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

const (
	ProposalStatusInvalid = "PROPOSAL_STATUS_INVALID"
)

// DepositParams contains the data of the deposit parameters of the x/gov module
type DepositParams struct {
	MinDeposit       sdk.Coins `json:"min_deposit,omitempty" yaml:"min_deposit"`
	MaxDepositPeriod int64     `json:"max_deposit_period,omitempty" yaml:"max_deposit_period"`
}

// NewDepositParam allows to build a new DepositParams
func NewDepositParam(d *govtypesv1.DepositParams) DepositParams {
	return DepositParams{
		MinDeposit:       d.MinDeposit,
		MaxDepositPeriod: d.MaxDepositPeriod.Nanoseconds(),
	}
}

// NewGenesisDepositParam allows to build a new DepositParams
func NewGenesisDepositParam(d *govtypesv1beta1.DepositParams) DepositParams {
	return DepositParams{
		MinDeposit:       d.MinDeposit,
		MaxDepositPeriod: d.MaxDepositPeriod.Nanoseconds(),
	}
}

// VotingParams contains the voting parameters of the x/gov module
type VotingParams struct {
	VotingPeriod int64 `json:"voting_period,omitempty" yaml:"voting_period"`
}

// NewVotingParams allows to build a new VotingParams instance
func NewVotingParams(v *govtypesv1.VotingParams) VotingParams {
	return VotingParams{
		VotingPeriod: v.VotingPeriod.Nanoseconds(),
	}
}

// NewGenesisVotingParams allows to build a new VotingParams instance
func NewGenesisVotingParams(v *govtypesv1beta1.VotingParams) VotingParams {
	return VotingParams{
		VotingPeriod: v.VotingPeriod.Nanoseconds(),
	}
}

// GovParams contains the data of the x/gov module parameters
type GovParams struct {
	DepositParams DepositParams `json:"deposit_params" yaml:"deposit_params"`
	VotingParams  VotingParams  `json:"voting_params" yaml:"voting_params"`
	TallyParams   TallyParams   `json:"tally_params" yaml:"tally_params"`
	Height        int64         `json:"height" ymal:"height"`
}

// GenesisGovParams contains the data of the x/gov module parameters
type GenesisGovParams struct {
	DepositParams DepositParams      `json:"deposit_params" yaml:"deposit_params"`
	VotingParams  VotingParams       `json:"voting_params" yaml:"voting_params"`
	TallyParams   GenesisTallyParams `json:"tally_params" yaml:"tally_params"`
	Height        int64              `json:"height" ymal:"height"`
}

// TallyParams contains the tally parameters of the x/gov module
type TallyParams struct {
	Quorum               string `json:"quorum,omitempty"`
	Threshold            string `json:"threshold,omitempty"`
	VetoThreshold        string `json:"veto_threshold,omitempty" yaml:"veto_threshold"`
	DepositLockingPeriod string `json:"deposit_locking_period,omitempty" yaml:"deposit_locking_period"`
}

// GenesisTallyParams contains genesis tally parameters of the x/gov module
type GenesisTallyParams struct {
	Quorum               sdk.Dec `json:"quorum,omitempty"`
	Threshold            sdk.Dec `json:"threshold,omitempty"`
	VetoThreshold        sdk.Dec `json:"veto_threshold,omitempty" yaml:"veto_threshold"`
	DepositLockingPeriod string  `json:"deposit_locking_period,omitempty" yaml:"deposit_locking_period"`
}

// NewTallyParams allows to build a new TallyParams instance
func NewTallyParams(t *govtypesv1.TallyParams) TallyParams {
	return TallyParams{
		Quorum:               t.Quorum,
		Threshold:            t.Threshold,
		VetoThreshold:        t.VetoThreshold,
		DepositLockingPeriod: t.DepositLockingPeriod.String(),
	}
}

// NewGenesisTallyParams allows to build a new GenesisTallyParams instance
func NewGenesisTallyParams(t *govtypesv1beta1.TallyParams) GenesisTallyParams {
	return GenesisTallyParams{
		Quorum:               t.Quorum,
		Threshold:            t.Threshold,
		VetoThreshold:        t.VetoThreshold,
		DepositLockingPeriod: t.DepositLockingPeriod.String(),
	}
}

// NewGovParams allows to build a new GovParams instance
func NewGovParams(votingParams VotingParams, depositParams DepositParams, tallyParams TallyParams, height int64) *GovParams {
	return &GovParams{
		DepositParams: depositParams,
		VotingParams:  votingParams,
		TallyParams:   tallyParams,
		Height:        height,
	}
}

// NewGenesisGovParams allows to build a new GenesisGovParams instance
func NewGenesisGovParams(votingParams VotingParams, depositParams DepositParams, tallyParams GenesisTallyParams, height int64) *GenesisGovParams {
	return &GenesisGovParams{
		DepositParams: depositParams,
		VotingParams:  votingParams,
		TallyParams:   tallyParams,
		Height:        height,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// Proposal represents a single governance proposal
type Proposal struct {
	ProposalRoute       string
	ProposalType        string
	ProposalTitle       string
	ProposalDescription string
	ProposalID          uint64
	Content             string
	Status              string
	SubmitTime          time.Time
	DepositEndTime      time.Time
	VotingStartTime     time.Time
	VotingEndTime       time.Time
	Proposer            string
	Metadata            string
}

// ProposalMetadata is used to parse metadata.json from proposals
type ProposalMetadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// NewProposal return a new Proposal instance
func NewProposal(
	proposalID uint64,
	proposalRoute string,
	proposalType string,
	proposalTitle string,
	proposalDescription string,
	content string,
	status string,
	submitTime time.Time,
	depositEndTime time.Time,
	votingStartTime time.Time,
	votingEndTime time.Time,
	proposer string,
	metadata string,
) Proposal {
	return Proposal{
		Content:             content,
		ProposalTitle:       proposalTitle,
		ProposalDescription: proposalDescription,
		ProposalRoute:       proposalRoute,
		ProposalType:        proposalType,
		ProposalID:          proposalID,
		Status:              status,
		SubmitTime:          submitTime,
		DepositEndTime:      depositEndTime,
		VotingStartTime:     votingStartTime,
		VotingEndTime:       votingEndTime,
		Proposer:            proposer,
		Metadata:            metadata,
	}
}

// Equal tells whether p and other contain the same data
func (p Proposal) Equal(other Proposal) bool {
	return p.ProposalRoute == other.ProposalRoute &&
		p.ProposalType == other.ProposalType &&
		p.ProposalID == other.ProposalID &&
		p.Content == other.Content &&
		p.Status == other.Status &&
		p.SubmitTime.Equal(other.SubmitTime) &&
		p.DepositEndTime.Equal(other.DepositEndTime) &&
		p.VotingStartTime.Equal(other.VotingStartTime) &&
		p.VotingEndTime.Equal(other.VotingEndTime) &&
		p.Proposer == other.Proposer &&
		p.Metadata == other.Metadata
}

// ProposalUpdate contains the data that should be used when updating a governance proposal
type ProposalUpdate struct {
	ProposalID      uint64
	Status          string
	VotingStartTime time.Time
	VotingEndTime   time.Time
}

// NewProposalUpdate allows to build a new ProposalUpdate instance
func NewProposalUpdate(
	proposalID uint64, status string, votingStartTime, votingEndTime time.Time,
) ProposalUpdate {
	return ProposalUpdate{
		ProposalID:      proposalID,
		Status:          status,
		VotingStartTime: votingStartTime,
		VotingEndTime:   votingEndTime,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Deposit contains the data of a single deposit made towards a proposal
type Deposit struct {
	ProposalID uint64
	Depositor  string
	Amount     sdk.Coins
	Timestamp  time.Time
	Height     int64
}

// NewDeposit return a new Deposit instance
func NewDeposit(
	proposalID uint64,
	depositor string,
	amount sdk.Coins,
	timestamp time.Time,
	height int64,
) Deposit {
	return Deposit{
		ProposalID: proposalID,
		Depositor:  depositor,
		Amount:     amount,
		Timestamp:  timestamp,
		Height:     height,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// Vote contains the data of a single proposal vote
type Vote struct {
	ProposalID uint64
	Voter      string
	Option     int32
	Timestamp  time.Time
	Height     int64
}

// NewVote return a new Vote instance
func NewVote(
	proposalID uint64,
	voter string,
	option int32,
	timestamp time.Time,
	height int64,
) Vote {
	return Vote{
		ProposalID: proposalID,
		Voter:      voter,
		Option:     option,
		Timestamp:  timestamp,
		Height:     height,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// TallyResult contains the data about the final results of a proposal
type TallyResult struct {
	ProposalID uint64
	Yes        string
	Abstain    string
	No         string
	NoWithVeto string
	Height     int64
}

// NewTallyResult return a new TallyResult instance
func NewTallyResult(
	proposalID uint64,
	yes string,
	abstain string,
	no string,
	noWithVeto string,
	height int64,
) TallyResult {
	return TallyResult{
		ProposalID: proposalID,
		Yes:        yes,
		Abstain:    abstain,
		No:         no,
		NoWithVeto: noWithVeto,
		Height:     height,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// ProposalStakingPoolSnapshot contains the data about a single staking pool snapshot to be associated with a proposal
type ProposalStakingPoolSnapshot struct {
	ProposalID uint64
	Pool       *PoolSnapshot
}

// NewProposalStakingPoolSnapshot returns a new ProposalStakingPoolSnapshot instance
func NewProposalStakingPoolSnapshot(proposalID uint64, pool *PoolSnapshot) ProposalStakingPoolSnapshot {
	return ProposalStakingPoolSnapshot{
		ProposalID: proposalID,
		Pool:       pool,
	}
}

// -------------------------------------------------------------------------------------------------------------------

// ProposalValidatorStatusSnapshot represents a single snapshot of the status of a validator associated
// with a single proposal
type ProposalValidatorStatusSnapshot struct {
	ProposalID           uint64
	ValidatorConsAddress string
	ValidatorVotingPower int64
	ValidatorStatus      int
	ValidatorJailed      bool
	Height               int64
}

// NewProposalValidatorStatusSnapshot returns a new ProposalValidatorStatusSnapshot instance
func NewProposalValidatorStatusSnapshot(
	proposalID uint64,
	validatorConsAddr string,
	validatorVotingPower int64,
	validatorStatus int,
	validatorJailed bool,
	height int64,
) ProposalValidatorStatusSnapshot {
	return ProposalValidatorStatusSnapshot{
		ProposalID:           proposalID,
		ValidatorStatus:      validatorStatus,
		ValidatorConsAddress: validatorConsAddr,
		ValidatorVotingPower: validatorVotingPower,
		ValidatorJailed:      validatorJailed,
		Height:               height,
	}
}
