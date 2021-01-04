package types

import (
	"database/sql"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// ValidatorData contains all the data of a single validator.
// It implements types.Validator interface
type ValidatorData struct {
	ConsAddress         string `db:"consensus_address"`
	ValAddress          string `db:"operator_address"`
	ConsPubKey          string `db:"consensus_pubkey"`
	SelfDelegateAddress string `db:"self_delegate_address"`
	MaxRate             string `db:"max_rate"`
	MaxChangeRate       string `db:"max_change_rate"`
}

// NewValidatorData allows to build a new ValidatorData
func NewValidatorData(consAddress, valAddress, consPubKey, selfDelegateAddress, maxRate, maxChangeRate string) ValidatorData {
	return ValidatorData{
		ConsAddress:         consAddress,
		ValAddress:          valAddress,
		ConsPubKey:          consPubKey,
		SelfDelegateAddress: selfDelegateAddress,
		MaxRate:             maxRate,
		MaxChangeRate:       maxChangeRate,
	}
}

// GetConsAddr implements types.Validator
func (v ValidatorData) GetConsAddr() sdk.ConsAddress {
	addr, err := sdk.ConsAddressFromBech32(v.ConsAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetConsPubKey implements types.Validator
func (v ValidatorData) GetConsPubKey() crypto.PubKey {
	return sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, v.ConsPubKey)
}

// GetOperator implements types.Validator
func (v ValidatorData) GetOperator() sdk.ValAddress {
	addr, err := sdk.ValAddressFromBech32(v.ValAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetSelfDelegateAddress implements types.Validator
func (v ValidatorData) GetSelfDelegateAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(v.SelfDelegateAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

// GetMaxChangeRate implements types.Validator
func (v ValidatorData) GetMaxChangeRate() *sdk.Dec {
	n, err := strconv.ParseInt(v.MaxChangeRate, 10, 64)
	if err != nil {
		panic(err)
	}
	result := sdk.NewDec(n)
	return &result
}

// GetMaxRate implements types.Validator
func (v ValidatorData) GetMaxRate() *sdk.Dec {
	n, err := strconv.ParseInt(v.MaxRate, 10, 64)
	if err != nil {
		panic(err)
	}
	result := sdk.NewDec(n)
	return &result
}

// ________________________________________________

// ValidatorRow represents a single row of the validator table
type ValidatorRow struct {
	ConsAddress string `db:"consensus_address"`
	ConsPubKey  string `db:"consensus_pubkey"`
}

// NewValidatorRow returns a new ValidatorRow
func NewValidatorRow(consAddress, consPubKey string) ValidatorRow {
	return ValidatorRow{
		ConsAddress: consAddress,
		ConsPubKey:  consPubKey,
	}
}

// Equal tells whether v and w contain the same data
func (v ValidatorRow) Equal(w ValidatorRow) bool {
	return v.ConsAddress == w.ConsAddress &&
		v.ConsPubKey == w.ConsPubKey
}

// ________________________________________________

// ValidatorInfoRow represents a single row of the validator_info table
type ValidatorInfoRow struct {
	ConsAddress         string `db:"consensus_address"`
	ValAddress          string `db:"operator_address"`
	SelfDelegateAddress string `db:"self_delegate_address"`
	MaxChangeRate       string `db:"max_change_rate"`
	MaxRate             string `db:"max_rate"`
}

// NewValidatorInfoRow allows to build a new ValidatorInfoRow
func NewValidatorInfoRow(consAddress, valAddress, selfDelegateAddress, maxChangeRate, maxRate string) ValidatorInfoRow {
	return ValidatorInfoRow{
		ConsAddress:         consAddress,
		ValAddress:          valAddress,
		SelfDelegateAddress: selfDelegateAddress,
		MaxChangeRate:       maxChangeRate,
		MaxRate:             maxRate,
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorInfoRow) Equal(w ValidatorInfoRow) bool {
	return v.ConsAddress == w.ConsAddress &&
		v.ValAddress == w.ValAddress &&
		v.SelfDelegateAddress == w.SelfDelegateAddress &&
		v.MaxRate == w.MaxRate &&
		v.MaxChangeRate == w.MaxChangeRate
}

//________________________________________________________________

// ValidatorDescriptionRow represent a row in validator_description
type ValidatorDescriptionRow struct {
	ValAddress      string         `db:"validator_address"`
	Moniker         sql.NullString `db:"moniker"`
	Identity        sql.NullString `db:"identity"`
	Website         sql.NullString `db:"website"`
	SecurityContact sql.NullString `db:"security_contact"`
	Details         sql.NullString `db:"details"`
}

// NewValidatorDescriptionRow return a row representing data structure in validator_description
func NewValidatorDescriptionRow(valAddress, moniker, identity, website, securityContact, details string) ValidatorDescriptionRow {
	return ValidatorDescriptionRow{
		ValAddress:      valAddress,
		Moniker:         ToNullString(moniker),
		Identity:        ToNullString(identity),
		Website:         ToNullString(website),
		SecurityContact: ToNullString(securityContact),
		Details:         ToNullString(details),
	}
}

// Equals return true if two ValidatorDescriptionRow are equal
func (w ValidatorDescriptionRow) Equals(v ValidatorDescriptionRow) bool {
	return v.ValAddress == w.ValAddress &&
		v.Moniker == w.Moniker &&
		v.Identity == w.Identity &&
		v.Website == w.Website &&
		v.SecurityContact == w.SecurityContact &&
		v.Details == w.Details
}

// ValidatorDescriptionHistoryRow represents a single row inside the validator_description_history table
type ValidatorDescriptionHistoryRow struct {
	ValAddress      string         `db:"operator_address"`
	Moniker         sql.NullString `db:"moniker"`
	Identity        sql.NullString `db:"identity"`
	Website         sql.NullString `db:"website"`
	SecurityContact sql.NullString `db:"security_contact"`
	Details         sql.NullString `db:"details"`
	Height          int64          `db:"height"`
	Timestamp       time.Time      `db:"timestamp"`
}

// NewValidatorDescriptionHistoryRow represents a single row inside the validator_description_history table
func NewValidatorDescriptionHistoryRow(
	valAddress, moniker, identity, website, securityContact, details string,
	height int64, timestamp time.Time,
) ValidatorDescriptionHistoryRow {
	return ValidatorDescriptionHistoryRow{
		ValAddress:      valAddress,
		Moniker:         sql.NullString{String: moniker, Valid: true},
		Identity:        sql.NullString{String: identity, Valid: true},
		Website:         sql.NullString{String: website, Valid: true},
		SecurityContact: sql.NullString{String: securityContact, Valid: true},
		Details:         sql.NullString{String: details, Valid: true},
		Height:          height,
		Timestamp:       timestamp,
	}
}

// Equals return true if two ValidatorDescriptionRow are equal
func (w ValidatorDescriptionHistoryRow) Equals(v ValidatorDescriptionHistoryRow) bool {
	return v.ValAddress == w.ValAddress &&
		v.Moniker == w.Moniker &&
		v.Identity == w.Identity &&
		v.Website == w.Website &&
		v.SecurityContact == w.SecurityContact &&
		v.Details == w.Details &&
		v.Height == w.Height &&
		v.Timestamp.Equal(w.Timestamp)
}

// ________________________________________________

// ValidatorCommissionRow represents a single row of the validator_commission database table
type ValidatorCommissionRow struct {
	OperatorAddress   string         `db:"validator_address"`
	Commission        sql.NullString `db:"commission"`
	MinSelfDelegation sql.NullString `db:"min_self_delegation"`
}

// NewValidatorCommissionRow allows to easily build a new ValidatorCommissionRow instance
func NewValidatorCommissionRow(operatorAddress string, commission string, minSelfDelegation string) ValidatorCommissionRow {
	return ValidatorCommissionRow{
		OperatorAddress:   operatorAddress,
		Commission:        sql.NullString{String: commission, Valid: true},
		MinSelfDelegation: sql.NullString{String: minSelfDelegation, Valid: true},
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorCommissionRow) Equal(w ValidatorCommissionRow) bool {
	return v.OperatorAddress == w.OperatorAddress &&
		v.Commission == w.Commission &&
		v.MinSelfDelegation == w.MinSelfDelegation
}

// ValidatorCommissionHistoryRow represents a single row of the validator_commission_history table
type ValidatorCommissionHistoryRow struct {
	CommissionID int64     `db:"commission_id"`
	Height       int64     `db:"height"`
	Timestamp    time.Time `db:"timestamp"`
}

// NewValidatorCommissionHistoryRow allows to easily build a new ValidatorCommissionHistoryRow instance
func NewValidatorCommissionHistoryRow(
	commissionID int64, height int64, timestamp time.Time,
) ValidatorCommissionHistoryRow {
	return ValidatorCommissionHistoryRow{
		CommissionID: commissionID,
		Height:       height,
		Timestamp:    timestamp,
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorCommissionHistoryRow) Equal(w ValidatorCommissionHistoryRow) bool {
	return v.CommissionID == w.CommissionID &&
		v.Height == w.Height &&
		v.Timestamp.Equal(w.Timestamp)
}

// ________________________________________________

// ValidatorVotingPowerRow represents a single row of the validator_voting_power database table
type ValidatorVotingPowerRow struct {
	ValidatorAddress string `db:"validator_address"`
	VotingPower      int64  `db:"voting_power"`
}

// NewValidatorVotingPowerRow allows to easily build a new ValidatorVotingPowerRow instance
func NewValidatorVotingPowerRow(address string, votingPower int64) ValidatorVotingPowerRow {
	return ValidatorVotingPowerRow{
		ValidatorAddress: address,
		VotingPower:      votingPower,
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorVotingPowerRow) Equal(w ValidatorVotingPowerRow) bool {
	return v.ValidatorAddress == w.ValidatorAddress &&
		v.VotingPower == w.VotingPower
}

// ValidatorVotingPowerHistoryRow represents a single row of the validator_voting_power_history database table
type ValidatorVotingPowerHistoryRow struct {
	ValidatorAddress string    `db:"validator_address"`
	VotingPower      int64     `db:"voting_power"`
	Height           int64     `db:"height"`
	Timestamp        time.Time `db:"timestamp"`
}

// NewValidatorVotingPowerHistoryRow allows to easily build a new ValidatorVotingPowerHistoryRow instance
func NewValidatorVotingPowerHistoryRow(
	address string, votingPower int64, height int64, timestamp time.Time,
) ValidatorVotingPowerHistoryRow {
	return ValidatorVotingPowerHistoryRow{
		ValidatorAddress: address,
		VotingPower:      votingPower,
		Height:           height,
		Timestamp:        timestamp,
	}
}

// Equal tells whether v and w represent the same rows
func (v ValidatorVotingPowerHistoryRow) Equal(w ValidatorVotingPowerHistoryRow) bool {
	return v.ValidatorAddress == w.ValidatorAddress &&
		v.VotingPower == w.VotingPower &&
		v.Height == w.Height &&
		v.Timestamp.Equal(w.Timestamp)
}

// ________________________________________________

// ValidatorUptimeRow represents a single row of the validator_uptime table
type ValidatorUptimeRow struct {
	ID                 int64  `db:"id"`
	ConsAddr           string `db:"validator_address"`
	SignedBlockWindow  int64  `db:"signed_blocks_window"`
	MissedBlockCounter int64  `db:"missed_blocks_counter"`
}

// NewValidatorUptimeRow allows to build a new ValidatorUptimeRow
func NewValidatorUptimeRow(consAddr string, signedBlocWindow, missedBlocksCounter int64) ValidatorUptimeRow {
	return ValidatorUptimeRow{
		ConsAddr:           consAddr,
		SignedBlockWindow:  signedBlocWindow,
		MissedBlockCounter: missedBlocksCounter,
	}
}

// Equal tells whether v and w contain the same data
func (v ValidatorUptimeRow) Equal(w ValidatorUptimeRow) bool {
	return v.ConsAddr == w.ConsAddr &&
		v.SignedBlockWindow == w.SignedBlockWindow &&
		v.MissedBlockCounter == w.MissedBlockCounter
}

// ValidatorUptimeHistoryRow represents a single row of the validator_uptime_history table
type ValidatorUptimeHistoryRow struct {
	UptimeID  int64     `db:"uptime_id"`
	Height    int64     `db:"height"`
	Timestamp time.Time `db:"timestamp"`
}

// NewValidatorUptimesHistoryRow builds a new ValidatorUptimeHistoryRow
func NewValidatorUptimesHistoryRow(
	uptimeID int64, height int64, timestamp time.Time,
) ValidatorUptimeHistoryRow {
	return ValidatorUptimeHistoryRow{
		UptimeID:  uptimeID,
		Height:    height,
		Timestamp: timestamp,
	}
}

// Equal tells whether v and w contain the same data
func (v ValidatorUptimeHistoryRow) Equal(w ValidatorUptimeHistoryRow) bool {
	return v.UptimeID == w.UptimeID &&
		v.Height == w.Height &&
		v.Timestamp.Equal(w.Timestamp)
}

//------------------------------------------------------------
// ValidatorStatus represents a single row of the validator_status table
type ValidatorStatusRow struct {
	Status      int    `db:"status"`
	Jailed      bool   `db:"jailed"`
	ConsAddress string `db:"validator_address"`
}

// NewValidatorUptimesHistoryRow builds a new ValidatorUptimeHistoryRow
func NewValidatorStatusRow(
	status int, jailed bool, consAddess string,
) ValidatorStatusRow {
	return ValidatorStatusRow{
		Status:      status,
		Jailed:      jailed,
		ConsAddress: consAddess,
	}
}

// Equal tells whether v and w contain the same data
func (v ValidatorStatusRow) Equal(w ValidatorStatusRow) bool {
	return v.Status == w.Status &&
		v.Jailed == w.Jailed &&
		v.ConsAddress == w.ConsAddress
}

//-------------------------------------------------------------

// ValidatorStatusHistoryRow represents a single row of the validator_status_history table
type ValidatorStatusHistoryRow struct {
	Status      int       `db:"status"`
	Jailed      bool      `db:"jailed"`
	Height      int64     `db:"height"`
	Timestamp   time.Time `db:"timestamp"`
	ConsAddress string    `db:"validator_address"`
}

// NewValidatorStatusHistoryRow builds a new ValidatorUptimeHistoryRow
func NewValidatorStatusHistoryRow(
	status int, jailed bool, height int64, timestamp time.Time, consAddress string,
) ValidatorStatusHistoryRow {
	return ValidatorStatusHistoryRow{
		Status:      status,
		Jailed:      jailed,
		Height:      height,
		Timestamp:   timestamp,
		ConsAddress: consAddress,
	}
}

// Equal tells whether v and w contain the same data
func (v ValidatorStatusHistoryRow) Equal(w ValidatorStatusHistoryRow) bool {
	return v.Status == w.Status &&
		v.Jailed == w.Jailed &&
		v.Height == w.Height &&
		v.Timestamp.Equal(w.Timestamp)
}

//--------------------------------------------------------

// DoubleSignVoteRow represents a single row of the double_sign_vote table
type DoubleSignVoteRow struct {
	ID               int64     `db:"id"`
	VoteType         int       `db:"type"`
	Height           int64     `db:"height"`
	Round            int       `db:"round"`
	BlockID          string    `db:"block_id"`
	Timestamp        time.Time `db:"timestamp"`
	ValidatorAddress string    `db:"validator_address"`
	ValidatorIndex   int       `db:"validator_index"`
	Signature        string    `db:"signature"`
}

// NewDoubleSignVoteRow allows to build a new NewDoubleSignVoteRow
func NewDoubleSignVoteRow(
	id int64,
	voteType int,
	height int64,
	round int,
	blockID string,
	timestamp time.Time,
	validatorAddress string,
	validatorIndex int,
	signature string,
) DoubleSignVoteRow {
	return DoubleSignVoteRow{
		ID:               id,
		VoteType:         voteType,
		Height:           height,
		Round:            round,
		BlockID:          blockID,
		Timestamp:        timestamp,
		ValidatorAddress: validatorAddress,
		ValidatorIndex:   validatorIndex,
		Signature:        signature,
	}
}

// Equal tells whether v and w represent the same rows
func (v DoubleSignVoteRow) Equal(w DoubleSignVoteRow) bool {
	return v.ID == w.ID &&
		v.VoteType == w.VoteType &&
		v.Height == w.Height &&
		v.Round == w.Round &&
		v.BlockID == w.BlockID &&
		v.Timestamp.Equal(w.Timestamp) &&
		v.ValidatorAddress == w.ValidatorAddress &&
		v.ValidatorIndex == w.ValidatorIndex &&
		v.Signature == w.Signature
}

//--------------------------------------------------------

// DoubleSignEvidenceRow represents a single row of the double_sign_evidence table
type DoubleSignEvidenceRow struct {
	PublicKey string `db:"public_key"`
	VoteAID   int64  `db:"vote_a_id"`
	VoteBID   int64  `db:"vote_b_id"`
}

// NewDoubleSignEvidenceRow allows to build a new NewDoubleSignEvidenceRow
func NewDoubleSignEvidenceRow(
	publicKey string,
	voteAID int64,
	voteBID int64,
) DoubleSignEvidenceRow {
	return DoubleSignEvidenceRow{
		PublicKey: publicKey,
		VoteAID:   voteAID,
		VoteBID:   voteBID,
	}
}

// Equal tells whether v and w represent the same rows
func (v DoubleSignEvidenceRow) Equal(w DoubleSignEvidenceRow) bool {
	return v.PublicKey == w.PublicKey &&
		v.VoteAID == w.VoteAID &&
		v.VoteBID == w.VoteBID
}