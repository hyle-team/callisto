package database

import (
	"encoding/json"
	"fmt"
	"time"

	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/forbole/bdjuno/v4/types"

	dbtypes "github.com/forbole/bdjuno/v4/database/types"

	"github.com/lib/pq"
)

// SaveGovParams saves the given x/gov parameters inside the database
func (db *Db) SaveGovParams(params *types.GovParams) error {

	depositParamsBz, err := json.Marshal(&params.DepositParams)
	if err != nil {
		return fmt.Errorf("error while marshaling deposit params: %s", err)
	}

	votingParamsBz, err := json.Marshal(&params.VotingParams)
	if err != nil {
		return fmt.Errorf("error while marshaling voting params: %s", err)
	}

	tallyingParams, err := json.Marshal(&params.TallyParams)
	if err != nil {
		return fmt.Errorf("error while marshaling tally params: %s", err)
	}

	stmt := `
INSERT INTO gov_params(deposit_params, voting_params, tally_params, height)
VALUES ($1, $2, $3, $4)
ON CONFLICT (one_row_id) DO UPDATE
	SET deposit_params = excluded.deposit_params,
  		voting_params = excluded.voting_params,
		tally_params = excluded.tally_params,
		height = excluded.height
WHERE gov_params.height <= excluded.height`
	_, err = db.SQL.Exec(stmt, string(depositParamsBz), string(votingParamsBz), string(tallyingParams), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing gov params: %s", err)
	}

	return nil
}

// SaveGenesisGovParams saves the genesis x/gov parameters inside the database
func (db *Db) SaveGenesisGovParams(params *types.GenesisGovParams) error {

	depositParamsBz, err := json.Marshal(&params.DepositParams)
	if err != nil {
		return fmt.Errorf("error while marshaling genesis deposit params: %s", err)
	}

	votingParamsBz, err := json.Marshal(&params.VotingParams)
	if err != nil {
		return fmt.Errorf("error while marshaling genesis voting params: %s", err)
	}

	tallyingParams, err := json.Marshal(&params.TallyParams)
	if err != nil {
		return fmt.Errorf("error while marshaling genesis tally params: %s", err)
	}

	stmt := `
INSERT INTO gov_params(deposit_params, voting_params, tally_params, height)
VALUES ($1, $2, $3, $4)
ON CONFLICT (one_row_id) DO UPDATE
	SET deposit_params = excluded.deposit_params,
  		voting_params = excluded.voting_params,
		tally_params = excluded.tally_params,
		height = excluded.height
WHERE gov_params.height <= excluded.height`
	_, err = db.SQL.Exec(stmt, string(depositParamsBz), string(votingParamsBz), string(tallyingParams), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing genesis gov params: %s", err)
	}

	return nil
}

// GetGovParams returns the most recent governance parameters
func (db *Db) GetGovParams() (*types.GovParams, error) {
	var rows []dbtypes.GovParamsRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM gov_params`)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	row := rows[0]

	var depositParams types.DepositParams
	err = json.Unmarshal([]byte(row.DepositParams), &depositParams)
	if err != nil {
		return nil, err
	}

	var votingParams types.VotingParams
	err = json.Unmarshal([]byte(row.VotingParams), &votingParams)
	if err != nil {
		return nil, err
	}

	var tallyParams types.TallyParams
	err = json.Unmarshal([]byte(row.TallyParams), &tallyParams)
	if err != nil {
		return nil, err
	}

	return types.NewGovParams(
		votingParams, depositParams, tallyParams,
		row.Height,
	), nil
}

// --------------------------------------------------------------------------------------------------------------------

// SaveProposals allows to save for the given height the given total amount of coins
func (db *Db) SaveProposals(proposals []types.Proposal) error {
	if len(proposals) == 0 {
		return nil
	}

	var accounts []types.Account

	proposalsQuery := `
INSERT INTO proposal(
	id, title, description, content, proposer_address, proposal_route, proposal_type, status,
    submit_time, deposit_end_time, voting_start_time, voting_end_time, metadata
) VALUES`
	var proposalsParams []interface{}

	for i, proposal := range proposals {
		// Prepare the account query
		accounts = append(accounts, types.NewAccount(proposal.Proposer))

		// Prepare the proposal query
		vi := i * 13
		proposalsQuery += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d, $%d),",
			vi+1, vi+2, vi+3, vi+4, vi+5, vi+6, vi+7, vi+8, vi+9, vi+10, vi+11, vi+12, vi+13)

		proposalsParams = append(proposalsParams,
			proposal.ProposalID,
			proposal.ProposalTitle,
			proposal.ProposalDescription,
			proposal.Content,
			proposal.Proposer,
			proposal.ProposalRoute,
			proposal.ProposalType,
			proposal.Status,
			proposal.SubmitTime,
			proposal.DepositEndTime,
			proposal.VotingStartTime,
			proposal.VotingEndTime,
			proposal.Metadata,
		)
	}

	// Store the accounts
	err := db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing proposers accounts: %s", err)
	}

	// Store the proposals
	proposalsQuery = proposalsQuery[:len(proposalsQuery)-1] // Remove trailing ","
	proposalsQuery += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(proposalsQuery, proposalsParams...)
	if err != nil {
		return fmt.Errorf("error while storing proposals: %s", err)
	}

	return nil
}

// GetProposal returns the proposal with the given id, or nil if not found
func (db *Db) GetProposal(id uint64) (types.Proposal, error) {
	var rows []*dbtypes.ProposalRow
	err := db.Sqlx.Select(&rows, `SELECT * FROM proposal WHERE id = $1`, id)
	if err != nil {
		return types.Proposal{}, err
	}

	if len(rows) == 0 {
		return types.Proposal{}, nil
	}

	row := rows[0]

	proposal := types.NewProposal(
		row.ProposalID,
		row.ProposalRoute,
		row.ProposalType,
		row.Title,
		row.Description,
		row.Content,
		row.Status,
		row.SubmitTime,
		row.DepositEndTime,
		row.VotingStartTime,
		row.VotingEndTime,
		row.Proposer,
		row.Metadata,
	)
	return proposal, nil
}

// GetOpenProposalsIds returns all the ids of the proposals that are in deposit or voting period at the given block time
func (db *Db) GetOpenProposalsIds(blockTime time.Time) ([]uint64, error) {
	var ids []uint64
	stmt := `SELECT id FROM proposal WHERE status = $1 OR status = $2`
	err := db.Sqlx.Select(&ids, stmt, govtypesv1.StatusDepositPeriod.String(), govtypesv1.StatusVotingPeriod.String())
	if err != nil {
		return ids, err
	}

	// Get also the invalid status proposals due to gRPC failure but still are in deposit period or voting period
	var idsInvalid []uint64
	stmt = `SELECT id FROM proposal WHERE status = $1 AND (voting_end_time > $2 OR deposit_end_time > $2)`
	err = db.Sqlx.Select(&idsInvalid, stmt, types.ProposalStatusInvalid, blockTime)
	ids = append(ids, idsInvalid...)

	return ids, err
}

// --------------------------------------------------------------------------------------------------------------------

// UpdateProposal updates a proposal stored inside the database
func (db *Db) UpdateProposal(update types.ProposalUpdate) error {
	query := `UPDATE proposal SET status = $1, voting_start_time = $2, voting_end_time = $3 where id = $4`
	_, err := db.SQL.Exec(query,
		update.Status,
		update.VotingStartTime,
		update.VotingEndTime,
		update.ProposalID,
	)
	if err != nil {
		return fmt.Errorf("error while updating proposal: %s", err)
	}

	return nil
}

// SaveDeposits allows to save multiple deposits
func (db *Db) SaveDeposits(deposits []types.Deposit) error {
	if len(deposits) == 0 {
		return nil
	}

	query := `INSERT INTO proposal_deposit (proposal_id, depositor_address, amount, timestamp, height) VALUES `
	var param []interface{}
	var accounts []types.Account
	for i, deposit := range deposits {
		vi := i * 5
		query += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4, vi+5)
		param = append(param, deposit.ProposalID,
			deposit.Depositor,
			pq.Array(dbtypes.NewDbCoins(deposit.Amount)),
			deposit.Timestamp,
			deposit.Height,
		)
		accounts = append(accounts, types.NewAccount(deposit.Depositor))
	}

	// Store the depositor account
	err := db.SaveAccounts(accounts)
	if err != nil {
		return fmt.Errorf("error while storing depositor account: %s", err)
	}

	query = query[:len(query)-1] // Remove trailing ","
	query += `
ON CONFLICT ON CONSTRAINT unique_deposit DO UPDATE
	SET amount = excluded.amount,
		timestamp = excluded.timestamp,
		height = excluded.height
WHERE proposal_deposit.height <= excluded.height`
	_, err = db.SQL.Exec(query, param...)
	if err != nil {
		return fmt.Errorf("error while storing deposits: %s", err)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// SaveVote allows to save for the given height and the message vote
func (db *Db) SaveVote(vote types.Vote) error {
	query := `
INSERT INTO proposal_vote (proposal_id, voter_address, option, timestamp, height)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT ON CONSTRAINT unique_vote DO UPDATE
	SET option = excluded.option,
		timestamp = excluded.timestamp,
		height = excluded.height
WHERE proposal_vote.height <= excluded.height`

	// Store the voter account
	err := db.SaveAccounts([]types.Account{types.NewAccount(vote.Voter)})
	if err != nil {
		return fmt.Errorf("error while storing voter account: %s", err)
	}

	_, err = db.SQL.Exec(query, vote.ProposalID, vote.Voter, govtypesv1.VoteOption(vote.Option).String(),
		vote.Timestamp, vote.Height)
	if err != nil {
		return fmt.Errorf("error while storing vote: %s", err)
	}

	return nil
}

// SaveTallyResults allows to save for the given height the given total amount of coins
func (db *Db) SaveTallyResults(tallys []types.TallyResult) error {
	if len(tallys) == 0 {
		return nil
	}

	query := `INSERT INTO proposal_tally_result(proposal_id, yes, abstain, no, no_with_veto, height) VALUES`

	var param []interface{}
	for i, tally := range tallys {
		vi := i * 6
		query += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4, vi+5, vi+6)
		param = append(param, tally.ProposalID,
			tally.Yes,
			tally.Abstain,
			tally.No,
			tally.NoWithVeto,
			tally.Height,
		)
	}

	query = query[:len(query)-1] // Remove trailing ","
	query += `
ON CONFLICT ON CONSTRAINT unique_tally_result DO UPDATE
	SET yes = excluded.yes,
	    abstain = excluded.abstain,
	    no = excluded.no,
	    no_with_veto = excluded.no_with_veto,
	    height = excluded.height
WHERE proposal_tally_result.height <= excluded.height`
	_, err := db.SQL.Exec(query, param...)
	if err != nil {
		return fmt.Errorf("error while storing tally result: %s", err)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// SaveProposalStakingPoolSnapshot allows to save the given snapshot of the staking pool
func (db *Db) SaveProposalStakingPoolSnapshot(snapshot types.ProposalStakingPoolSnapshot) error {
	stmt := `
INSERT INTO proposal_staking_pool_snapshot (proposal_id, bonded_tokens, not_bonded_tokens, height)
VALUES ($1, $2, $3, $4)
ON CONFLICT ON CONSTRAINT unique_staking_pool_snapshot DO UPDATE SET
	proposal_id = excluded.proposal_id,
    bonded_tokens = excluded.bonded_tokens,
	not_bonded_tokens = excluded.not_bonded_tokens,
	height = excluded.height
WHERE proposal_staking_pool_snapshot.height <= excluded.height`

	_, err := db.SQL.Exec(stmt,
		snapshot.ProposalID, snapshot.Pool.BondedTokens.String(), snapshot.Pool.NotBondedTokens.String(), snapshot.Pool.Height)
	if err != nil {
		return fmt.Errorf("error while storing proposal staking pool snapshot: %s", err)
	}

	return nil
}

// SaveProposalValidatorsStatusesSnapshots allows to save the given validator statuses snapshots
func (db *Db) SaveProposalValidatorsStatusesSnapshots(snapshots []types.ProposalValidatorStatusSnapshot) error {
	if len(snapshots) == 0 {
		return nil
	}

	stmt := `
INSERT INTO proposal_validator_status_snapshot(proposal_id, validator_address, voting_power, status, jailed, height)
VALUES `

	var args []interface{}
	for i, snapshot := range snapshots {
		si := i * 6

		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d),", si+1, si+2, si+3, si+4, si+5, si+6)
		args = append(args,
			snapshot.ProposalID, snapshot.ValidatorConsAddress, snapshot.ValidatorVotingPower,
			snapshot.ValidatorStatus, snapshot.ValidatorJailed, snapshot.Height)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += `
ON CONFLICT ON CONSTRAINT unique_validator_status_snapshot DO UPDATE
	SET proposal_id = excluded.proposal_id,
		validator_address = excluded.validator_address,
		voting_power = excluded.voting_power,
		status = excluded.status,
		jailed = excluded.jailed,
		height = excluded.height
WHERE proposal_validator_status_snapshot.height <= excluded.height`
	_, err := db.SQL.Exec(stmt, args...)
	if err != nil {
		return fmt.Errorf("error while storing proposal validator statuses snapshot: %s", err)
	}

	return nil
}

// SaveSoftwareUpgradePlan allows to save the given software upgrade plan with its proposal id
func (db *Db) SaveSoftwareUpgradePlan(proposalID uint64, plan upgradetypes.Plan, height int64) error {

	stmt := `
INSERT INTO software_upgrade_plan(proposal_id, plan_name, upgrade_height, info, height)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (proposal_id) DO UPDATE SET
	plan_name = excluded.plan_name,
	upgrade_height = excluded.upgrade_height,
	info = excluded.info,
	height = excluded.height
WHERE software_upgrade_plan.height <= excluded.height`

	_, err := db.SQL.Exec(stmt,
		proposalID, plan.Name, plan.Height, plan.Info, height)
	if err != nil {
		return fmt.Errorf("error while storing software upgrade plan: %s", err)
	}

	return nil
}

// DeleteSoftwareUpgradePlan allows to delete a SoftwareUpgradePlan with proposal ID
func (db *Db) DeleteSoftwareUpgradePlan(proposalID uint64) error {
	stmt := `DELETE FROM software_upgrade_plan WHERE proposal_id = $1`

	_, err := db.SQL.Exec(stmt, proposalID)
	if err != nil {
		return fmt.Errorf("error while deleting software upgrade plan: %s", err)
	}

	return nil
}

// CheckSoftwareUpgradePlan returns true if an upgrade is scheduled at the given height
func (db *Db) CheckSoftwareUpgradePlan(upgradeHeight int64) (bool, error) {
	var exist bool

	stmt := `SELECT EXISTS (SELECT 1 FROM software_upgrade_plan WHERE upgrade_height=$1)`
	err := db.SQL.QueryRow(stmt, upgradeHeight).Scan(&exist)
	if err != nil {
		return exist, fmt.Errorf("error while checking software upgrade plan existence: %s", err)
	}

	return exist, nil
}

// TruncateSoftwareUpgradePlan delete software upgrade plans once the upgrade height passed
func (db *Db) TruncateSoftwareUpgradePlan(height int64) error {
	stmt := `DELETE FROM software_upgrade_plan WHERE upgrade_height <= $1`

	_, err := db.SQL.Exec(stmt, height)
	if err != nil {
		return fmt.Errorf("error while deleting software upgrade plan: %s", err)
	}

	return nil
}
