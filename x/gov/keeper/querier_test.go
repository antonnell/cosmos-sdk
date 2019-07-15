package keeper

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

const custom = "custom"

func getQueriedParams(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier) (types.DepositParams, types.VotingParams, types.TallyParams) {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryParams, types.ParamDeposit}, "/"),
		Data: []byte{},
	}

	bz, err := querier(ctx, []string{types.QueryParams, types.ParamDeposit}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var depositParams types.DepositParams
	require.NoError(t, cdc.UnmarshalJSON(bz, &depositParams))

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryParams, types.ParamVoting}, "/"),
		Data: []byte{},
	}

	bz, err = querier(ctx, []string{types.QueryParams, types.ParamVoting}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var votingParams types.VotingParams
	require.NoError(t, cdc.UnmarshalJSON(bz, &votingParams))

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryParams, types.ParamTallying}, "/"),
		Data: []byte{},
	}

	bz, err = querier(ctx, []string{types.QueryParams, types.ParamTallying}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var tallyParams types.TallyParams
	require.NoError(t, cdc.UnmarshalJSON(bz, &tallyParams))

	return depositParams, votingParams, tallyParams
}

func getQueriedProposal(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) types.Proposal {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryProposal}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{types.QueryProposal}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var proposal types.Proposal
	require.NoError(t, cdc.UnmarshalJSON(bz, proposal))

	return proposal
}

func getQueriedProposals(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, depositor, voter sdk.AccAddress, status types.ProposalStatus, limit uint64) []types.Proposal {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryProposals}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryProposalsParams(status, limit, voter, depositor)),
	}

	bz, err := querier(ctx, []string{types.QueryProposals}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var proposals types.Proposals
	require.NoError(t, cdc.UnmarshalJSON(bz, &proposals))

	return proposals
}

func getQueriedDeposit(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64, depositor sdk.AccAddress) types.Deposit {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryDeposit}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryDepositParams(proposalID, depositor)),
	}

	bz, err := querier(ctx, []string{types.QueryDeposit}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var deposit types.Deposit
	require.NoError(t, cdc.UnmarshalJSON(bz, &deposit))

	return deposit
}

func getQueriedDeposits(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) []types.Deposit {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryDeposits}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{types.QueryDeposits}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var deposits []types.Deposit
	require.NoError(t, cdc.UnmarshalJSON(bz, &deposits))

	return deposits
}

func getQueriedVote(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64, voter sdk.AccAddress) types.Vote {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryVote}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryVoteParams(proposalID, voter)),
	}

	bz, err := querier(ctx, []string{types.QueryVote}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var vote types.Vote
	require.NoError(t, cdc.UnmarshalJSON(bz, &vote))

	return vote
}

func getQueriedVotes(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) []types.Vote {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryVote}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{types.QueryVotes}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var votes []types.Vote
	require.NoError(t, cdc.UnmarshalJSON(bz, &votes))

	return votes
}

func getQueriedTally(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) types.TallyResult {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, types.QuerierRoute, types.QueryTally}, "/"),
		Data: cdc.MustMarshalJSON(types.NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{types.QueryTally}, query)
	require.NoError(t, err)
	require.NotNil(t, bz)

	var tally types.TallyResult
	require.NoError(t, cdc.UnmarshalJSON(bz, &tally))

	return tally
}

func TestQueries(t *testing.T) {
	ctx, _, keeper, _ := createTestInput(t, false, 1000)
	querier := NewQuerier(keeper)

	tp := TestProposal()

	depositParams, _, _ := getQueriedParams(t, ctx, keeper.cdc, querier)

	// TestAddrs[0] proposes (and deposits) proposals #1 and #2
	proposal1, err := keeper.SubmitProposal(ctx, tp)
	deposit1 := types.NewDeposit(proposal1.ProposalID, TestAddrs[0], sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)))
	keeper.SetDeposit(ctx, deposit1)

	proposal2, err := keeper.SubmitProposal(ctx, tp)
	deposit2 := types.NewDeposit(proposal2.ProposalID, TestAddrs[0], sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000000)))
	keeper.SetDeposit(ctx, deposit2)

	// TestAddrs[1] proposes (and deposits) on proposal #3
	proposal3, err := keeper.SubmitProposal(ctx, tp)
	deposit3 := types.NewDeposit(proposal3.ProposalID, TestAddrs[1], sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)))
	keeper.SetDeposit(ctx, deposit3)

	// TestAddrs[1] deposits on proposals #2 & #3
	deposit4 := types.NewDeposit(proposal2.ProposalID, TestAddrs[1], depositParams.MinDeposit)
	deposit5 := types.NewDeposit(proposal3.ProposalID, TestAddrs[1], depositParams.MinDeposit)
	keeper.SetDeposit(ctx, deposit4)
	keeper.SetDeposit(ctx, deposit5)

	// check deposits on proposal1 match individual deposits
	deposits := getQueriedDeposits(t, ctx, keeper.cdc, querier, proposal1.ProposalID)
	require.Len(t, deposits, 1)
	require.Equal(t, deposit1, deposits[0])

	deposit := getQueriedDeposit(t, ctx, keeper.cdc, querier, proposal1.ProposalID, TestAddrs[0])
	require.Equal(t, deposit1, deposit)

	// check deposits on proposal2 match individual deposits
	deposits = getQueriedDeposits(t, ctx, keeper.cdc, querier, proposal2.ProposalID)
	require.Len(t, deposits, 2)
	require.Equal(t, deposit2, deposits[0])
	require.Equal(t, deposit4, deposits[1])

	deposit = getQueriedDeposit(t, ctx, keeper.cdc, querier, proposal2.ProposalID, TestAddrs[0])
	require.Equal(t, deposit2, deposits[0])
	deposit = getQueriedDeposit(t, ctx, keeper.cdc, querier, proposal2.ProposalID, TestAddrs[1])
	require.Equal(t, deposit4, deposits[1])

	// check deposits on proposal3 match individual deposits
	deposits = getQueriedDeposits(t, ctx, keeper.cdc, querier, proposal3.ProposalID)
	require.Len(t, deposits, 1)
	require.Equal(t, deposit5, deposits[0])

	deposit = getQueriedDeposit(t, ctx, keeper.cdc, querier, proposal3.ProposalID, TestAddrs[1])
	require.Equal(t, deposi5, deposit)

	// Only proposal #1 should be in types.Deposit Period
	proposals := getQueriedProposals(t, ctx, keeper.cdc, querier, nil, nil, types.StatusDepositPeriod, 0)
	require.Len(t, proposals, 1)
	require.Equal(t, proposal1, proposals[0])

	// Only proposals #2 and #3 should be in Voting Period
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, nil, nil, types.StatusVotingPeriod, 0)
	require.Len(t, proposals, 2)
	require.Equal(t, proposal2, proposals[0])
	require.Equal(t, proposal3, proposals[1])

	// Addrs[0] votes on proposals #2 & #3
	vote1 := types.NewVote(proposal2.ProposalID, TestAddrs[0], types.OptionYes)
	vote2 := types.NewVote(proposal3.ProposalID, TestAddrs[0], types.OptionYes)
	keeper.SetVote(ctx, vote1)
	keeper.SetVote(ctx, vote2)	

	// Addrs[1] votes on proposal #3
	vote3 := types.NewVote(proposal3.ProposalID, TestAddrs[1], types.OptionYes)
	keeper.SetVote(ctx, vote3)

	// Test query voted by TestAddrs[0]
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, nil, TestAddrs[0], types.StatusNil, 0)
	require.Equal(t, proposal2, proposals[0])
	require.Equal(t, proposal3, proposals[1])

	// Test query votes on types.Proposal 2
	votes := getQueriedVotes(t, ctx, keeper.cdc, querier, proposal2.ProposalID)
	require.Len(t, votes, 1)
	require.Equal(t, vote1, votes[0])

	vote := getQueriedVote(t, ctx, keeper.cdc, querier, proposal2.ProposalID, TestAddrs[0])
	require.Equal(t, vote1, vote)

	// Test query votes on types.Proposal 3
	votes = getQueriedVotes(t, ctx, keeper.cdc, querier, proposal3.ProposalID)
	require.Len(t, votes, 2)
	require.Equal(t, vote2, votes[0])
	require.Equal(t, vote3, votes[1])

	// Test query all proposals
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, nil, nil, types.StatusNil, 0)
	require.Equal(t, proposal1, proposals[0])
	require.Equal(t, proposal2, proposals[1])
	require.Equal(t, proposal3, proposals[2])

	// Test query voted by TestAddrs[1]
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, nil, TestAddrs[1], types.StatusNil, 0)
	require.Equal(t, proposal3.ProposalID, proposals[0].ProposalID)

	// Test query deposited by TestAddrs[0]
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, TestAddrs[0], nil, types.StatusNil, 0)
	require.Equal(t, proposal1.ProposalID, proposals[0].ProposalID)

	// Test query deposited by addr2
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, TestAddrs[1], nil, types.StatusNil, 0)
	require.Equal(t, proposal2.ProposalID, proposals[0].ProposalID)
	require.Equal(t, proposal3.ProposalID, proposals[1].ProposalID)

	// Test query voted AND deposited by addr1
	proposals = getQueriedProposals(t, ctx, keeper.cdc, querier, TestAddrs[0], TestAddrs[0], types.StatusNil, 0)
	require.Equal(t, proposal2.ProposalID, proposals[0].ProposalID)
}