package types

import (
	"fmt"
	accumulatortypes "github.com/cosmos/cosmos-sdk/x/accumulator/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft/types"
	accumulatorsource "github.com/forbole/bdjuno/v4/modules/accumulator/source"
	nftsource "github.com/forbole/bdjuno/v4/modules/nft/source"
	"os"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/juno/v4/node/remote"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govtypesv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v4/node/local"

	nodeconfig "github.com/forbole/juno/v4/node/config"

	localaccumulatorsource "github.com/forbole/bdjuno/v4/modules/accumulator/source/local"
	remoteaccumulatorsource "github.com/forbole/bdjuno/v4/modules/accumulator/source/remote"
	banksource "github.com/forbole/bdjuno/v4/modules/bank/source"
	localbanksource "github.com/forbole/bdjuno/v4/modules/bank/source/local"
	remotebanksource "github.com/forbole/bdjuno/v4/modules/bank/source/remote"
	distrsource "github.com/forbole/bdjuno/v4/modules/distribution/source"
	localdistrsource "github.com/forbole/bdjuno/v4/modules/distribution/source/local"
	remotedistrsource "github.com/forbole/bdjuno/v4/modules/distribution/source/remote"
	govsource "github.com/forbole/bdjuno/v4/modules/gov/source"
	localgovsource "github.com/forbole/bdjuno/v4/modules/gov/source/local"
	remotegovsource "github.com/forbole/bdjuno/v4/modules/gov/source/remote"
	mintsource "github.com/forbole/bdjuno/v4/modules/mint/source"
	localmintsource "github.com/forbole/bdjuno/v4/modules/mint/source/local"
	remotemintsource "github.com/forbole/bdjuno/v4/modules/mint/source/remote"
	localnftsource "github.com/forbole/bdjuno/v4/modules/nft/source/local"
	remotenftsource "github.com/forbole/bdjuno/v4/modules/nft/source/remote"
	slashingsource "github.com/forbole/bdjuno/v4/modules/slashing/source"
	localslashingsource "github.com/forbole/bdjuno/v4/modules/slashing/source/local"
	remoteslashingsource "github.com/forbole/bdjuno/v4/modules/slashing/source/remote"
	stakingsource "github.com/forbole/bdjuno/v4/modules/staking/source"
	localstakingsource "github.com/forbole/bdjuno/v4/modules/staking/source/local"
	remotestakingsource "github.com/forbole/bdjuno/v4/modules/staking/source/remote"
)

type Sources struct {
	BankSource        banksource.Source
	DistrSource       distrsource.Source
	GovSource         govsource.Source
	MintSource        mintsource.Source
	SlashingSource    slashingsource.Source
	StakingSource     stakingsource.Source
	NFTSource         nftsource.Source
	AccumulatorSource accumulatorsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig *params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	app := simapp.NewSimApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, simapp.MakeTestEncodingConfig(), simapp.EmptyAppOptions{},
	)

	sources := &Sources{
		BankSource:        localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
		DistrSource:       localdistrsource.NewSource(source, distrtypes.QueryServer(app.DistrKeeper)),
		GovSource:         localgovsource.NewSource(source, govtypesv1.QueryServer(app.GovKeeper), nil),
		MintSource:        localmintsource.NewSource(source, minttypes.QueryServer(app.MintKeeper)),
		SlashingSource:    localslashingsource.NewSource(source, slashingtypes.QueryServer(app.SlashingKeeper)),
		StakingSource:     localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: *app.StakingKeeper}),
		NFTSource:         localnftsource.NewSource(source, app.NFTKeeper),
		AccumulatorSource: localaccumulatorsource.NewSource(source, app.AccumulatorKeeper),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(app, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(app, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:        remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:       remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:         remotegovsource.NewSource(source, govtypesv1.NewQueryClient(source.GrpcConn), govtypesv1beta1.NewQueryClient(source.GrpcConn)),
		MintSource:        remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource:    remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:     remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
		NFTSource:         remotenftsource.NewSource(source, nfttypes.NewQueryClient(source.GrpcConn)),
		AccumulatorSource: remoteaccumulatorsource.NewSource(source, accumulatortypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
