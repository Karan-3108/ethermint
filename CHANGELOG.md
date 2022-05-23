
<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## Unreleased

### Improvements

* (cli) [Karan-3108#1086](https://github.com/Karan-3108/ethermint/pull/1086) Add rollback command.

### Bug Fixes

* (rpc) [Karan-3108#1082](https://github.com/Karan-3108/ethermint/pull/1082) fix gas price returned in getTransaction api.
* (evm) [Karan-3108#1088](https://github.com/Karan-3108/ethermint/pull/1088) Fix ability to append log in tx post processing.

## [v0.15.0] - 2022-05-09

### State Machine Breaking

* (ante) [Karan-3108#1060](https://github.com/Karan-3108/ethermint/pull/1060) Check `EnableCreate`/`EnableCall` in `AnteHandler` to short-circuit EVM transactions.

### API Breaking

* (rpc) [Karan-3108#1070](https://github.com/Karan-3108/ethermint/pull/1070) Refactor `rpc/` package:
  * `Backend` interface is now `BackendI`, which implements `EVMBackend` (for Ethereum namespaces) and `CosmosBackend` (for Cosmos namespaces)
  * Previous `EVMBackend` type is now `Backend`, which is the concrete implementation of `BackendI`
  * Move `rpc/ethereum/types` -> `rpc/types`
  * Move `rpc/ethereum/backend` -> `rpc/backend`
  * Move `rpc/ethereum/namespaces` -> `rpc/namespaces/ethereum`
* (rpc) [Karan-3108#1068](https://github.com/Karan-3108/ethermint/pull/1068) Fix London hard-fork check logic in JSON-RPC APIs.

### Improvements

* (ci, evm) [Karan-3108#1063](https://github.com/Karan-3108/ethermint/pull/1063) Run simulations on CI.

### Bug Fixes

* (rpc) [Karan-3108#1059](https://github.com/Karan-3108/ethermint/pull/1059) Remove unnecessary event filtering logic on the `eth_baseFee` JSON-RPC endpoint.

## [v0.14.0] - 2022-04-19

### API Breaking

* (evm) [Karan-3108#1051](https://github.com/Karan-3108/ethermint/pull/1051) Context block height fix on TraceTx. Removes `tx_index` on `QueryTraceTxRequest` proto type.

### Improvements

* (deps) [Karan-3108#1046](https://github.com/Karan-3108/ethermint/pull/1046) Bump Cosmos SDK version to [`v0.45.3`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.3)
* (rpc) [Karan-3108#1056](https://github.com/Karan-3108/ethermint/pull/1056) Make json-rpc namespaces extensible

### Bug Fixes

* (rpc) [Karan-3108#1050](https://github.com/Karan-3108/ethermint/pull/1050) `eth_getBlockByNumber` fix on batch transactions
* (app) [Karan-3108#658](https://github.com/Karan-3108/ethermint/issues/658) Support simulations for the EVM.

## [v0.13.0] - 2022-04-05

### API Breaking

* (evm) [Karan-3108#1027](https://github.com/Karan-3108/ethermint/pull/1027) Change the `PostTxProcessing` hook interface to include the full message data.
* (feemarket) [Karan-3108#1026](https://github.com/Karan-3108/ethermint/pull/1026) Fix REST endpoints to use `/ethermint/feemarket/*` instead of `/feemarket/evm/*`.

### Improvements

* (deps) [Karan-3108#1029](https://github.com/Karan-3108/ethermint/pull/1029) Bump Cosmos SDK version to [`v0.45.2`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.2)
* (evm) [Karan-3108#1025](https://github.com/Karan-3108/ethermint/pull/1025) Allow to append logs after a post processing hook.

## [v0.12.2] - 2022-03-30

### Bug Fixes

* (feemarket) [Karan-3108#1021](https://github.com/Karan-3108/ethermint/pull/1021) Fix fee market migration.

## [v0.12.1] - 2022-03-29

### Bug Fixes

* (evm) [Karan-3108#1016](https://github.com/Karan-3108/ethermint/pull/1016) Update validate basic check for storage state.

## [v0.12.0] - 2022-03-24

### Bug Fixes

* (rpc) [Karan-3108#1012](https://github.com/Karan-3108/ethermint/pull/1012) fix the tx hash in filter entries created by `eth_newPendingTransactionFilter`.
* (rpc) [Karan-3108#1006](https://github.com/Karan-3108/ethermint/pull/1006) Use `string` as the parameters type to correct ambiguous results.
* (ante) [Karan-3108#1004](https://github.com/Karan-3108/ethermint/pull/1004) Make `MaxTxGasWanted` configurable.
* (ante) [Karan-3108#991](https://github.com/Karan-3108/ethermint/pull/991) Set an upper bound to gasWanted to prevent DoS attack.
* (rpc) [Karan-3108#990](https://github.com/Karan-3108/ethermint/pull/990) Calculate reward values from all `MsgEthereumTx` from a block in `eth_feeHistory`.

## [v0.11.0] - 2022-03-06

### State Machine Breaking

* (ante) [Karan-3108#964](https://github.com/Karan-3108/ethermint/pull/964) add NewInfiniteGasMeterWithLimit for storing the user provided gas limit. Fixes block's consumed gas calculation in the block creation phase.

### Bug Fixes

* (rpc) [Karan-3108#975](https://github.com/Karan-3108/ethermint/pull/975) Fix unexpected `nil` values for `reward`, returned by `EffectiveGasTipValue(blockBaseFee)` in the `eth_feeHistory` RPC method.

### Improvements

- (rpc) [Karan-3108#979](https://github.com/Karan-3108/ethermint/pull/979) Add configurable timeouts to http server
- (rpc) [Karan-3108#988](https://github.com/Karan-3108/ethermint/pull/988) json-rpc server always use local rpc client

## [v0.10.1] - 2022-03-04

### Bug Fixes

* (rpc) [Karan-3108#970](https://github.com/Karan-3108/ethermint/pull/970) Fix unexpected nil reward values on `eth_feeHistory` response
* (evm) [Karan-3108#529](https://github.com/Karan-3108/ethermint/issues/529) Add support return value on trace tx response.

### Improvements

* (rpc) [Karan-3108#968](https://github.com/Karan-3108/ethermint/pull/968) Add some buffer to returned gas price to provide better default UX for client.

## [v0.10.0] - 2022-02-26

### API Breaking

* (ante) [Karan-3108#866](https://github.com/Karan-3108/ethermint/pull/866) `NewAnteHandler` constructor now receives a `HandlerOptions` field.
* (evm) [Karan-3108#849](https://github.com/Karan-3108/ethermint/pull/849) `PostTxProcessing` hook now takes an Ethereum tx `Receipt` and a `from` `Address` as arguments.
* (ante) [Karan-3108#916](https://github.com/Karan-3108/ethermint/pull/916) Don't check min-gas-price for eth tx if london hardfork enabled and feemarket enabled.

### State Machine Breaking

* (deps) [Karan-3108#912](https://github.com/Karan-3108/ethermint/pull/912) Bump Cosmos SDK version to [`v0.45.1`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.1)
* (evm) [Karan-3108#840](https://github.com/Karan-3108/ethermint/pull/840) Store empty topics as empty array rather than nil.
* (feemarket) [Karan-3108#822](https://github.com/Karan-3108/ethermint/pull/822) Update EIP1559 base fee in `BeginBlock`.
* (evm) [Karan-3108#817](https://github.com/Karan-3108/ethermint/pull/817) Use `effectiveGasPrice` in ante handler, add `effectiveGasPrice` to tx receipt.
* (evm) [Karan-3108#808](https://github.com/Karan-3108/ethermint/issues/808) increase nonce in ante handler for contract creation transaction.
* (evm) [Karan-3108#851](https://github.com/Karan-3108/ethermint/pull/851) fix contract address used in EVM, this issue is caused by [Karan-3108#808](https://github.com/Karan-3108/ethermint/issues/808).
* (evm)  Reject invalid `MsgEthereumTx` wrapping tx
* (evm)  Fix `SelfDestruct` opcode by deleting account code and state.
* (feemarket) [Karan-3108#855](https://github.com/Karan-3108/ethermint/pull/855) Consistent `BaseFee` check logic.
* (evm) [Karan-3108#729](https://github.com/Karan-3108/ethermint/pull/729) Refactor EVM `StateDB` implementation.
* (evm) [Karan-3108#945](https://github.com/Karan-3108/ethermint/pull/945) Bumb Go-ethereum version to [`v1.10.16`](https://github.com/ethereum/go-ethereum/releases/tag/v1.10.16)

### Features

* (ante) [Karan-3108#950](https://github.com/Karan-3108/ethermint/pull/950) Add support for EIP712 signed Cosmos transactions

### Improvements

* (types) [Karan-3108#884](https://github.com/Karan-3108/ethermint/pull/884) Introduce a new `EthAccountI` interface for EVM-compatible account types.
* (types) [Karan-3108#849](https://github.com/Karan-3108/ethermint/pull/849) Add `Type` function to distinguish EOAs from Contract accounts.
* (evm) [Karan-3108#826](https://github.com/Karan-3108/ethermint/issues/826) Improve allocation of bytes of `tx.To` address.
* (evm) [Karan-3108#827](https://github.com/Karan-3108/ethermint/issues/827) Speed up creation of event logs by using the slice insertion idiom with indices.
* (ante) [Karan-3108#819](https://github.com/Karan-3108/ethermint/pull/819) Remove redundant ante handlers
* (app) [Karan-3108#873](https://github.com/Karan-3108/ethermint/pull/873) Validate code hash in GenesisAccount
* (evm) [Karan-3108#901](https://github.com/Karan-3108/ethermint/pull/901) Support multiple `MsgEthereumTx` in single tx.
* (config) [Karan-3108#908](https://github.com/Karan-3108/ethermint/pull/908) Add `api.enable` flag for Cosmos SDK Rest server
* (feemarket) [Karan-3108#919](https://github.com/Karan-3108/ethermint/pull/919) Initialize baseFee in default genesis state.
* (feemarket) [Karan-3108#943](https://github.com/Karan-3108/ethermint/pull/943) Store the base fee as a module param instead of using state storage.

### Bug Fixes

* (rpc) [Karan-3108#955](https://github.com/Karan-3108/ethermint/pull/955) Fix websocket server push duplicated messages to subscriber.
* (rpc) [Karan-3108#953](https://github.com/Karan-3108/ethermint/pull/953) Add `eth_signTypedData` api support.
* (log) [Karan-3108#948](https://github.com/Karan-3108/ethermint/pull/948) Redirect go-ethereum's logs to cosmos-sdk logger.
* (evm) [Karan-3108#884](https://github.com/Karan-3108/ethermint/pull/884) Support multiple account types on the EVM `StateDB`.
* (rpc) [Karan-3108#831](https://github.com/Karan-3108/ethermint/pull/831) Fix BaseFee value when height is specified.
* (evm) [Karan-3108#838](https://github.com/Karan-3108/ethermint/pull/838) Fix splitting of trace.Memory into 32 chunks.
* (rpc) [Karan-3108#860](https://github.com/Karan-3108/ethermint/pull/860) Fix `eth_getLogs` when specify blockHash without address/topics, and limit the response size.
* (rpc) [Karan-3108#865](https://github.com/Karan-3108/ethermint/pull/865) Fix RPC Filter parameters being ignored
* (evm) [Karan-3108#871](https://github.com/Karan-3108/ethermint/pull/871) Set correct nonce in `EthCall` and `EstimateGas` grpc query.
* (rpc) [Karan-3108#878](https://github.com/Karan-3108/ethermint/pull/878) Workaround to make GetBlock RPC api report correct block gas used.
* (rpc) [Karan-3108#900](https://github.com/Karan-3108/ethermint/pull/900) `newPendingTransactions` filter return ethereum tx hash.
* (rpc) [Karan-3108#933](https://github.com/Karan-3108/ethermint/pull/933) Fix `newPendingTransactions` subscription deadlock when a Websocket client exits without unsubscribing and the node errors.
* (evm) [Karan-3108#932](https://github.com/Karan-3108/ethermint/pull/932) Fix base fee check logic in state transition.

## [v0.9.0] - 2021-12-01

### State Machine Breaking

* (evm) [Karan-3108#802](https://github.com/Karan-3108/ethermint/pull/802) Clear access list for each transaction

### Improvements

* (app) [Karan-3108#794](https://github.com/Karan-3108/ethermint/pull/794) Setup in-place store migrators.
* (ci) [Karan-3108#784](https://github.com/Karan-3108/ethermint/pull/784) Enable automatic backport of PRs.
* (rpc) [Karan-3108#786](https://github.com/Karan-3108/ethermint/pull/786) Improve error message of `SendTransaction`/`SendRawTransaction` JSON-RPC APIs.
* (rpc) [Karan-3108#810](https://github.com/Karan-3108/ethermint/pull/810) Optimize tx index lookup in web3 rpc

### Bug Fixes

* (license) [Karan-3108#800](https://github.com/Karan-3108/ethermint/pull/800) Re-license project to [LGPLv3](https://choosealicense.com/licenses/lgpl-3.0/#) to comply with go-ethereum.
* (evm) [Karan-3108#794](https://github.com/Karan-3108/ethermint/pull/794) Register EVM gRPC `Msg` server.
* (rpc) [Karan-3108#781](https://github.com/Karan-3108/ethermint/pull/781) Fix get block invalid transactions filter.
* (rpc) [Karan-3108#782](https://github.com/Karan-3108/ethermint/pull/782) Fix wrong block gas limit returned by JSON-RPC.
* (evm) [Karan-3108#798](https://github.com/Karan-3108/ethermint/pull/798) Fix the semantic of `ForEachStorage` callback's return value

## [v0.8.1] - 2021-11-23

### Bug Fixes

* (feemarket) [Karan-3108#770](https://github.com/Karan-3108/ethermint/pull/770) Enable fee market (EIP1559) by default.
* (rpc) [Karan-3108#769](https://github.com/Karan-3108/ethermint/pull/769) Fix default Ethereum signer for JSON-RPC.

## [v0.8.0] - 2021-11-17

### State Machine Breaking

* (evm, ante) [Karan-3108#620](https://github.com/Karan-3108/ethermint/pull/620) Add fee market field to EVM `Keeper` and `AnteHandler`.
* (all) [Karan-3108#231](https://github.com/Karan-3108/ethermint/pull/231) Bump go-ethereum version to [`v1.10.9`](https://github.com/ethereum/go-ethereum/releases/tag/v1.10.9)
* (ante) [Karan-3108#703](https://github.com/Karan-3108/ethermint/pull/703) Fix some fields in transaction are not authenticated by signature.
* (evm) [Karan-3108#751](https://github.com/Karan-3108/ethermint/pull/751) don't revert gas refund logic when transaction reverted

### Features

* (rpc, evm) [Karan-3108#673](https://github.com/Karan-3108/ethermint/pull/673) Use tendermint events to store fee market basefee.
* (rpc) [Karan-3108#624](https://github.com/Karan-3108/ethermint/pull/624) Implement new JSON-RPC endpoints from latest geth version
* (evm) [Karan-3108#662](https://github.com/Karan-3108/ethermint/pull/662) Disable basefee for non london blocks
* (cmd) [Karan-3108#712](https://github.com/Karan-3108/ethermint/pull/712) add tx cli to build evm transaction
* (rpc) [Karan-3108#733](https://github.com/Karan-3108/ethermint/pull/733) add JSON_RPC endpoint `personal_unpair`
* (rpc) [Karan-3108#734](https://github.com/Karan-3108/ethermint/pull/734) add JSON_RPC endpoint `eth_feeHistory`
* (rpc) [Karan-3108#740](https://github.com/Karan-3108/ethermint/pull/740) add JSON_RPC endpoint `personal_initializeWallet`
* (rpc) [Karan-3108#743](https://github.com/Karan-3108/ethermint/pull/743) add JSON_RPC endpoint `debug_traceBlockByHash`
* (rpc) [Karan-3108#748](https://github.com/Karan-3108/ethermint/pull/748) add JSON_RPC endpoint `personal_listWallets`
* (rpc) [Karan-3108#754](https://github.com/Karan-3108/ethermint/pull/754) add JSON_RPC endpoint `debug_intermediateRoots`

### Bug Fixes

* (evm) [Karan-3108#746](https://github.com/Karan-3108/ethermint/pull/746) Set EVM debugging based on tracer configuration.
* (app,cli) [Karan-3108#725](https://github.com/Karan-3108/ethermint/pull/725) Fix cli-config for  `keys` command.
* (rpc) [Karan-3108#727](https://github.com/Karan-3108/ethermint/pull/727) Decode raw transaction using RLP.
* (rpc) [Karan-3108#661](https://github.com/Karan-3108/ethermint/pull/661) Fix OOM bug when creating too many filters using JSON-RPC.
* (evm) [Karan-3108#660](https://github.com/Karan-3108/ethermint/pull/660) Fix `nil` pointer panic in `ApplyNativeMessage`.
* (evm, test) [Karan-3108#649](https://github.com/Karan-3108/ethermint/pull/649) Test DynamicFeeTx.
* (evm) [Karan-3108#702](https://github.com/Karan-3108/ethermint/pull/702) Fix panic in web3 RPC handlers
* (rpc) [Karan-3108#720](https://github.com/Karan-3108/ethermint/pull/720) Fix `debug_traceTransaction` failure
* (rpc) [Karan-3108#741](https://github.com/Karan-3108/ethermint/pull/741) Fix `eth_getBlockByNumberAndHash` return with non eth txs
* (rpc) [Karan-3108#743](https://github.com/Karan-3108/ethermint/pull/743) Fix debug JSON RPC handler crash on non-existing block

### Improvements

* (tests) [Karan-3108#704](https://github.com/Karan-3108/ethermint/pull/704) Introduce E2E testing framework for clients
* (deps) [Karan-3108#737](https://github.com/Karan-3108/ethermint/pull/737) Bump ibc-go to [`v2.0.0`](https://github.com/cosmos/ibc-go/releases/tag/v2.0.0)
* (rpc) [Karan-3108#671](https://github.com/Karan-3108/ethermint/pull/671) Don't pass base fee externally for `EthCall`/`EthEstimateGas` apis.
* (evm) [Karan-3108#674](https://github.com/Karan-3108/ethermint/pull/674) Refactor `ApplyMessage`, remove
  `ApplyNativeMessage`.
* (rpc) [Karan-3108#714](https://github.com/Karan-3108/ethermint/pull/714) remove `MsgEthereumTx` support in `TxConfig`

## [v0.7.2] - 2021-10-24

### Improvements

* (deps) [Karan-3108#692](https://github.com/Karan-3108/ethermint/pull/692) Bump Cosmos SDK version to [`v0.44.3`](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.44.3).
* (rpc) [Karan-3108#679](https://github.com/Karan-3108/ethermint/pull/679) Fix file close handle.
* (deps) [Karan-3108#668](https://github.com/Karan-3108/ethermint/pull/668) Bump Tendermint version to [`v0.34.14`](https://github.com/tendermint/tendermint/releases/tag/v0.34.14).

### Bug Fixes

* (rpc) [Karan-3108#667](https://github.com/Karan-3108/ethermint/issues/667) Fix `ExpandHome` restrictions bypass

## [v0.7.1] - 2021-10-08

### Bug Fixes

* (evm) [Karan-3108#650](https://github.com/Karan-3108/ethermint/pull/650) Fix panic when flattening the cache context in case transaction is reverted.
* (rpc, test) [Karan-3108#608](https://github.com/Karan-3108/ethermint/pull/608) Fix rpc test.

## [v0.7.0] - 2021-10-07

### API Breaking

* (rpc) [Karan-3108#400](https://github.com/Karan-3108/ethermint/issues/400) Restructure JSON-RPC directory and rename server config

### Improvements

* (deps) [Karan-3108#621](https://github.com/Karan-3108/ethermint/pull/621) Bump IBC-go to [`v1.2.1`](https://github.com/cosmos/ibc-go/releases/tag/v1.2.1)
* (evm) [Karan-3108#613](https://github.com/Karan-3108/ethermint/pull/613) Refactor `traceTx`
* (deps) [Karan-3108#610](https://github.com/Karan-3108/ethermint/pull/610) Bump Cosmos SDK to [v0.44.1](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.44.1).

### Bug Fixes

* (rpc) [Karan-3108#642](https://github.com/Karan-3108/ethermint/issues/642) Fix `eth_getLogs` when string is specified in filter's from or to fields
* (evm) [Karan-3108#616](https://github.com/Karan-3108/ethermint/issues/616) Fix halt on deeply nested stack of cache context. Stack is now flattened before iterating over the tx logs.
* (rpc, evm) [Karan-3108#614](https://github.com/Karan-3108/ethermint/issues/614) Use JSON for (un)marshaling tx `Log`s from events.
* (rpc) [Karan-3108#611](https://github.com/Karan-3108/ethermint/pull/611) Fix panic on JSON-RPC when querying for an invalid block height.
* (cmd) [Karan-3108#483](https://github.com/Karan-3108/ethermint/pull/483) Use config values on genesis accounts.

## [v0.6.0] - 2021-09-29

### State Machine Breaking

* (app) [Karan-3108#476](https://github.com/Karan-3108/ethermint/pull/476) Update Bech32 HRP to `ethm`.
* (evm) [Karan-3108#556](https://github.com/Karan-3108/ethermint/pull/556) Remove tx logs and block bloom from chain state
* (evm) [Karan-3108#590](https://github.com/Karan-3108/ethermint/pull/590) Contract storage key is not hashed anymore

### API Breaking

* (evm) [Karan-3108#469](https://github.com/Karan-3108/ethermint/pull/469) Deprecate `YoloV3Block` and `EWASMBlock` from `ChainConfig`

### Features

* (evm) [Karan-3108#469](https://github.com/Karan-3108/ethermint/pull/469) Support [EIP-1559](https://eips.ethereum.org/EIPS/eip-1559)
* (evm) [Karan-3108#417](https://github.com/Karan-3108/ethermint/pull/417) Add `EvmHooks` for tx post-processing
* (rpc) [Karan-3108#506](https://github.com/Karan-3108/ethermint/pull/506) Support for `debug_traceTransaction` RPC endpoint
* (rpc) [Karan-3108#555](https://github.com/Karan-3108/ethermint/pull/555) Support for `debug_traceBlockByNumber` RPC endpoint

### Bug Fixes

* (rpc, server) [Karan-3108#600](https://github.com/Karan-3108/ethermint/pull/600) Add TLS configuration for websocket API
* (rpc) [Karan-3108#598](https://github.com/Karan-3108/ethermint/pull/598) Check truncation when creating a `BlockNumber` from `big.Int`
* (evm) [Karan-3108#597](https://github.com/Karan-3108/ethermint/pull/597) Check for `uint64` -> `int64` block height overflow on `GetHashFn`
* (evm) [Karan-3108#579](https://github.com/Karan-3108/ethermint/pull/579) Update `DeriveChainID` function to handle `v` signature values `< 35`.
* (encoding) [Karan-3108#478](https://github.com/Karan-3108/ethermint/pull/478) Register `Evidence` to amino codec.
* (rpc) [Karan-3108#478](https://github.com/Karan-3108/ethermint/pull/481) Getting the node configuration when calling the `miner` rpc methods.
* (cli) [Karan-3108#561](https://github.com/Karan-3108/ethermint/pull/561) `Export` and `Start` commands now use the same home directory.

### Improvements

* (evm) [Karan-3108#461](https://github.com/Karan-3108/ethermint/pull/461) Increase performance of `StateDB` transaction log storage (r/w).
* (evm) [Karan-3108#566](https://github.com/Karan-3108/ethermint/pull/566) Introduce `stateErr` store in `StateDB` to avoid meaningless operations if any error happened before
* (rpc, evm) [Karan-3108#587](https://github.com/Karan-3108/ethermint/pull/587) Apply bloom filter when query ethlogs with range of blocks
* (evm) [Karan-3108#586](https://github.com/Karan-3108/ethermint/pull/586) Benchmark evm keeper

## [v0.5.0] - 2021-08-20

### State Machine Breaking

* (app, rpc) [Karan-3108#447](https://github.com/Karan-3108/ethermint/pull/447) Chain ID format has been changed from `<identifier>-<epoch>` to `<identifier>_<EIP155_number>-<epoch>`
in order to clearly distinguish permanent vs impermanent components.
* (app, evm) [Karan-3108#434](https://github.com/Karan-3108/ethermint/pull/434) EVM `Keeper` struct and `NewEVM` function now have a new `trace` field to define
the Tracer type used to collect execution traces from the EVM transaction execution.
* (evm) [Karan-3108#175](https://github.com/Karan-3108/ethermint/issues/175) The msg `TxData` field is now represented as a `*proto.Any`.
* (evm) [Karan-3108#84](https://github.com/Karan-3108/ethermint/pull/84) Remove `journal`, `CommitStateDB` and `stateObjects`.
* (rpc, evm) [Karan-3108#81](https://github.com/Karan-3108/ethermint/pull/81) Remove tx `Receipt` from store and replace it with fields obtained from the Tendermint RPC client.
* (evm) [Karan-3108#72](https://github.com/Karan-3108/ethermint/issues/72) Update `AccessList` to use `TransientStore` instead of map.
* (evm) [Karan-3108#68](https://github.com/Karan-3108/ethermint/issues/68) Replace block hash storage map to use staking `HistoricalInfo`.
* (evm) [Karan-3108#276](https://github.com/Karan-3108/ethermint/pull/276) Vm errors don't result in cosmos tx failure, just
  different tx state and events.
* (evm) [Karan-3108#342](https://github.com/Karan-3108/ethermint/issues/342) Don't clear balance when resetting the account.
* (evm) [Karan-3108#334](https://github.com/Karan-3108/ethermint/pull/334) Log index changed to the index in block rather than
  tx.
* (evm) [Karan-3108#399](https://github.com/Karan-3108/ethermint/pull/399) Exception in sub-message call reverts the call if it's not propagated.

### API Breaking

* (proto) [Karan-3108#448](https://github.com/Karan-3108/ethermint/pull/448) Bump version for all Ethermint messages to `v1`
* (server) [Karan-3108#434](https://github.com/Karan-3108/ethermint/pull/434) `evm-rpc` flags and app config have been renamed to `json-rpc`.
* (proto, evm) [Karan-3108#207](https://github.com/Karan-3108/ethermint/issues/207) Replace `big.Int` in favor of `sdk.Int` for `TxData` fields
* (proto, evm) [Karan-3108#81](https://github.com/Karan-3108/ethermint/pull/81) gRPC Query and Tx service changes:
  * The `TxReceipt`, `TxReceiptsByBlockHeight` endpoints have been removed from the Query service.
  * The `ContractAddress`, `Bloom` have been removed from the `MsgEthereumTxResponse` and the
    response now contains the ethereum-formatted `Hash` in hex format.
* (eth) [Karan-3108#845](https://github.com/cosmos/ethermint/pull/845) The `eth` namespace must be included in the list of API's as default to run the rpc server without error.
* (evm) [Karan-3108#202](https://github.com/Karan-3108/ethermint/pull/202) Web3 api `SendTransaction`/`SendRawTransaction` returns ethereum compatible transaction hash, and query api `GetTransaction*` also accept that.
* (rpc) [Karan-3108#258](https://github.com/Karan-3108/ethermint/pull/258) Return empty `BloomFilter` instead of throwing an error when it cannot be found (`nil` or empty).
* (rpc) [Karan-3108#277](https://github.com/Karan-3108/ethermint/pull/321) Fix `BloomFilter` response.

### Improvements

* (client) [Karan-3108#450](https://github.com/Karan-3108/ethermint/issues/450) Add EIP55 hex address support on `debug addr` command.
* (server) [Karan-3108#343](https://github.com/Karan-3108/ethermint/pull/343) Define a wrap tendermint logger `Handler` go-ethereum's `root` logger.
* (rpc) [Karan-3108#457](https://github.com/Karan-3108/ethermint/pull/457) Configure RPC gas cap through app config.
* (evm) [Karan-3108#434](https://github.com/Karan-3108/ethermint/pull/434) Support different `Tracer` types for the EVM.
* (deps) [Karan-3108#427](https://github.com/Karan-3108/ethermint/pull/427) Bump ibc-go to [`v1.0.0`](https://github.com/cosmos/ibc-go/releases/tag/v1.0.0)
* (gRPC) [Karan-3108#239](https://github.com/Karan-3108/ethermint/pull/239) Query `ChainConfig` via gRPC.
* (rpc) [Karan-3108#181](https://github.com/Karan-3108/ethermint/pull/181) Use evm denomination for params on tx fee.
* (deps) [Karan-3108#423](https://github.com/Karan-3108/ethermint/pull/423) Bump Cosmos SDK and Tendermint versions to [v0.43.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.43.0) and [v0.34.11](https://github.com/tendermint/tendermint/releases/tag/v0.34.11), respectively.
* (evm) [Karan-3108#66](https://github.com/Karan-3108/ethermint/issues/66) Support legacy transaction types for signing.
* (evm) [Karan-3108#24](https://github.com/Karan-3108/ethermint/pull/24) Implement metrics for `MsgEthereumTx`, state transitions, `BeginBlock` and `EndBlock`.
* (rpc)  [Karan-3108#124](https://github.com/Karan-3108/ethermint/issues/124) Implement `txpool_content`, `txpool_inspect` and `txpool_status` RPC methods
* (rpc) [Karan-3108#112](https://github.com/Karan-3108/ethermint/pull/153) Fix `eth_coinbase` to return the ethereum address of the validator
* (rpc) [Karan-3108#176](https://github.com/Karan-3108/ethermint/issues/176) Support fetching pending nonce
* (rpc) [Karan-3108#272](https://github.com/Karan-3108/ethermint/pull/272) do binary search to estimate gas accurately
* (rpc) [Karan-3108#313](https://github.com/Karan-3108/ethermint/pull/313) Implement internal debug namespace (Not including logger functions nor traces).
* (rpc) [Karan-3108#349](https://github.com/Karan-3108/ethermint/pull/349) Implement configurable JSON-RPC APIs to manage enabled namespaces.
* (rpc) [Karan-3108#377](https://github.com/Karan-3108/ethermint/pull/377) Implement `miner_` namespace. `miner_setEtherbase` and `miner_setGasPrice` are working as intended. All the other calls are not applicable and return `unsupported`.
* (eth) [Karan-3108#460](https://github.com/Karan-3108/ethermint/issues/460) Add support for EIP-1898.

### Bug Fixes

* (keys) [Karan-3108#346](https://github.com/Karan-3108/ethermint/pull/346) Fix `keys add` command with `--ledger` flag for the `secp256k1` signing algorithm.
* (evm) [Karan-3108#291](https://github.com/Karan-3108/ethermint/pull/291) Use block proposer address (validator operator) for `COINBASE` opcode.
* (rpc) [Karan-3108#81](https://github.com/Karan-3108/ethermint/pull/81) Fix transaction hashing and decoding on `eth_sendTransaction`.
* (rpc) [Karan-3108#45](https://github.com/Karan-3108/ethermint/pull/45) Use `EmptyUncleHash` and `EmptyRootHash` for empty ethereum `Header` fields.

## [v0.4.1] - 2021-03-01

### API Breaking

* (faucet) [Karan-3108#678](https://github.com/cosmos/ethermint/pull/678) Faucet module has been removed in favor of client libraries such as [`@cosmjs/faucet`](https://github.com/cosmos/cosmjs/tree/master/packages/faucet).
* (evm) [Karan-3108#670](https://github.com/cosmos/ethermint/pull/670) Migrate types to the ones defined by the protobuf messages, which are required for the stargate release.

### Bug Fixes

* (evm) [Karan-3108#799](https://github.com/cosmos/ethermint/issues/799) Fix wrong precision in calculation of gas fee.
* (evm) [Karan-3108#760](https://github.com/cosmos/ethermint/issues/760) Fix Failed to call function EstimateGas.
* (evm) [Karan-3108#767](https://github.com/cosmos/ethermint/issues/767) Fix error of timeout when using Truffle to deploy contract.
* (evm) [Karan-3108#751](https://github.com/cosmos/ethermint/issues/751) Fix misused method to calculate block hash in evm related function.
* (evm) [Karan-3108#721](https://github.com/cosmos/ethermint/issues/721) Fix mismatch block hash in rpc response when use eht.getBlock.
* (evm) [Karan-3108#730](https://github.com/cosmos/ethermint/issues/730) Fix 'EIP2028' not open when Istanbul version has been enabled.
* (evm) [Karan-3108#749](https://github.com/cosmos/ethermint/issues/749) Fix panic in `AnteHandler` when gas price larger than 100000
* (evm) [Karan-3108#747](https://github.com/cosmos/ethermint/issues/747) Fix format errors in String() of QueryETHLogs
* (evm) [Karan-3108#742](https://github.com/cosmos/ethermint/issues/742) Add parameter check for evm query func.
* (evm) [Karan-3108#687](https://github.com/cosmos/ethermint/issues/687) Fix nonce check to explicitly check for the correct nonce, rather than a simple 'greater than' comparison.
* (api) [Karan-3108#687](https://github.com/cosmos/ethermint/issues/687) Returns error for a transaction with an incorrect nonce.
* (evm) [Karan-3108#674](https://github.com/cosmos/ethermint/issues/674) Reset all cache after account data has been committed in `EndBlock` to make sure every node state consistent.
* (evm) [Karan-3108#672](https://github.com/cosmos/ethermint/issues/672) Fix panic of `wrong Block.Header.AppHash` when restart a node with snapshot.
* (evm) [Karan-3108#775](https://github.com/cosmos/ethermint/issues/775) MisUse of headHash as blockHash when create EVM context.

### Features
* (api) [Karan-3108#821](https://github.com/cosmos/ethermint/pull/821) Individually enable the api modules. Will be implemented in the latest version of ethermint with the upcoming stargate upgrade.

### Features
* (api) [Karan-3108#825](https://github.com/cosmos/ethermint/pull/825) Individually enable the api modules. Will be implemented in the latest version of ethermint with the upcoming stargate upgrade.

## [v0.4.0] - 2020-12-15

### API Breaking

* (evm) [Karan-3108#661](https://github.com/cosmos/ethermint/pull/661) `Balance` field has been removed from the evm module's `GenesisState`.

### Features

* (rpc) [Karan-3108#571](https://github.com/cosmos/ethermint/pull/571) Add pending queries to JSON-RPC calls. This allows for the querying of pending transactions and other relevant information that pertains to the pending state:
  * `eth_getBalance`
  * `eth_getTransactionCount`
  * `eth_getBlockTransactionCountByNumber`
  * `eth_getBlockByNumber`
  * `eth_getTransactionByHash`
  * `eth_getTransactionByBlockNumberAndIndex`
  * `eth_sendTransaction` - the nonce will automatically update to its pending nonce (when none is explicitly provided)

### Improvements

* (evm) [Karan-3108#661](https://github.com/cosmos/ethermint/pull/661) Add invariant check for account balance and account nonce.
* (deps) [Karan-3108#654](https://github.com/cosmos/ethermint/pull/654) Bump go-ethereum version to [v1.9.25](https://github.com/ethereum/go-ethereum/releases/tag/v1.9.25)
* (evm) [Karan-3108#627](https://github.com/cosmos/ethermint/issues/627) Add extra EIPs parameter to apply custom EVM jump tables.

### Bug Fixes

* (evm) [Karan-3108#661](https://github.com/cosmos/ethermint/pull/661) Set nonce to the EVM account on genesis initialization.
* (rpc) [Karan-3108#648](https://github.com/cosmos/ethermint/issues/648) Fix block cumulative gas used value.
* (evm) [Karan-3108#621](https://github.com/cosmos/ethermint/issues/621) EVM `GenesisAccount` fields now share the same format as the auth module `Account`.
* (evm) [Karan-3108#618](https://github.com/cosmos/ethermint/issues/618) Add missing EVM `Context` `GetHash` field that retrieves a the header hash from a given block height.
* (app) [Karan-3108#617](https://github.com/cosmos/ethermint/issues/617) Fix genesis export functionality.
* (rpc) [Karan-3108#574](https://github.com/cosmos/ethermint/issues/574) Fix outdated version from `eth_protocolVersion`.

## [v0.3.1] - 2020-11-24

### Improvements

* (deps) [Karan-3108#615](https://github.com/cosmos/ethermint/pull/615) Bump Cosmos SDK version to [v0.39.2](https://github.com/cosmos/cosmos-sdk/tag/v0.39.2)
* (deps) [Karan-3108#610](https://github.com/cosmos/ethermint/pull/610) Update Go dependency to 1.15+.
* (evm) [Karan-3108#603](https://github.com/cosmos/ethermint/pull/603) Add state transition params that enable or disable the EVM `Call` and `Create` operations.
* (deps) [Karan-3108#602](https://github.com/cosmos/ethermint/pull/602) Bump tendermint version to [v0.33.9](https://github.com/tendermint/tendermint/releases/tag/v0.33.9)

### Bug Fixes

* (rpc) [Karan-3108#613](https://github.com/cosmos/ethermint/issues/613) Fix potential deadlock caused if the keyring `List` returned an error.

## [v0.3.0] - 2020-11-16

### API Breaking

* (crypto) [Karan-3108#559](https://github.com/cosmos/ethermint/pull/559) Refactored crypto package in preparation for the SDK's Stargate release:
  * `crypto.PubKeySecp256k1` and `crypto.PrivKeySecp256k1` are now `ethsecp256k1.PubKey` and `ethsecp256k1.PrivKey`, respectively
  * Moved SDK `SigningAlgo` implementation for Ethermint's Secp256k1 key to `crypto/hd` package.
* (rpc) [Karan-3108#588](https://github.com/cosmos/ethermint/pull/588) The `rpc` package has been refactored to account for the separation of each
corresponding Ethereum API namespace:
  * `rpc/namespaces/eth`: `eth` namespace. Exposes the `PublicEthereumAPI` and the `PublicFilterAPI`.
  * `rpc/namespaces/personal`: `personal` namespace. Exposes the `PrivateAccountAPI`.
  * `rpc/namespaces/net`: `net` namespace. Exposes the `PublicNetAPI`.
  * `rpc/namespaces/web3`: `web3` namespace. Exposes the `PublicWeb3API`.
* (evm) [Karan-3108#588](https://github.com/cosmos/ethermint/pull/588) The EVM transaction CLI has been removed in favor of the JSON-RPC.

### Improvements

* (deps) [Karan-3108#594](https://github.com/cosmos/ethermint/pull/594) Bump go-ethereum version to [v1.9.24](https://github.com/ethereum/go-ethereum/releases/tag/v1.9.24)

### Bug Fixes

* (ante) [Karan-3108#597](https://github.com/cosmos/ethermint/pull/597) Fix incorrect fee check on `AnteHandler`.
* (evm) [Karan-3108#583](https://github.com/cosmos/ethermint/pull/583) Fixes incorrect resetting of tx count and block bloom during `BeginBlock`, as well as gas consumption.
* (crypto) [Karan-3108#577](https://github.com/cosmos/ethermint/pull/577) Fix `BIP44HDPath` that did not prepend `m/` to the path. This now uses the `DefaultBaseDerivationPath` variable from go-ethereum to ensure addresses are consistent.

## [v0.2.1] - 2020-09-30

### Features

* (rpc) [Karan-3108#552](https://github.com/cosmos/ethermint/pull/552) Implement Eth Personal namespace `personal_importRawKey`.

### Bug fixes

* (keys) [Karan-3108#554](https://github.com/cosmos/ethermint/pull/554) Fix private key derivation.
* (app/ante) [Karan-3108#550](https://github.com/cosmos/ethermint/pull/550) Update ante handler nonce verification to accept any nonce greater than or equal to the expected nonce to allow to successive transactions.

## [v0.2.0] - 2020-09-24

### State Machine Breaking

* (app) [Karan-3108#540](https://github.com/cosmos/ethermint/issues/540) Chain identifier's format has been changed to match the Cosmos `chainID` [standard](https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-5.md), which is required for IBC. The epoch number of the ID is used as the EVM `chainID`.

### API Breaking

* (types) [Karan-3108#503](https://github.com/cosmos/ethermint/pull/503) The `types.DenomDefault` constant for `"aphoton"` has been renamed to `types.AttoPhoton`.

### Improvements

* (types) [Karan-3108#504](https://github.com/cosmos/ethermint/pull/504) Unmarshal a JSON `EthAccount` using an Ethereum hex address in addition to Bech32.
* (types) [Karan-3108#503](https://github.com/cosmos/ethermint/pull/503) Add `--coin-denom` flag to testnet command that sets the given coin denomination to SDK and Ethermint parameters.
* (types) [Karan-3108#502](https://github.com/cosmos/ethermint/pull/502) `EthAccount` now also exposes the Ethereum hex address in `string` format to clients.
* (types) [Karan-3108#494](https://github.com/cosmos/ethermint/pull/494) Update `EthAccount` public key JSON type to `string`.
* (app) [Karan-3108#471](https://github.com/cosmos/ethermint/pull/471) Add `x/upgrade` module for managing software updates.
* (evm) [Karan-3108#458](https://github.com/cosmos/ethermint/pull/458) Define parameter for token denomination used for the EVM module.
* (evm) [Karan-3108#443](https://github.com/cosmos/ethermint/issues/443) Support custom Ethereum `ChainConfig` params.
* (types) [Karan-3108#434](https://github.com/cosmos/ethermint/issues/434) Update default denomination to Atto Photon (`aphoton`).
* (types) [Karan-3108#515](https://github.com/cosmos/ethermint/pull/515) Update minimum gas price to be 1.

### Bug Fixes

* (ante) [Karan-3108#525](https://github.com/cosmos/ethermint/pull/525) Add message validation decorator to `AnteHandler` for `MsgEthereumTx`.
* (types) [Karan-3108#507](https://github.com/cosmos/ethermint/pull/507) Fix hardcoded `aphoton` on `EthAccount` balance getter and setter.
* (types) [Karan-3108#501](https://github.com/cosmos/ethermint/pull/501) Fix bech32 encoding error by using the compressed ethereum secp256k1 public key.
* (evm) [Karan-3108#496](https://github.com/cosmos/ethermint/pull/496) Fix bugs on `journal.revert` and `CommitStateDB.Copy`.
* (types) [Karan-3108#480](https://github.com/cosmos/ethermint/pull/480) Update [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) coin type to `60` to satisfy [EIP84](https://github.com/ethereum/EIPs/issues/84).
* (types) [Karan-3108#513](https://github.com/cosmos/ethermint/pull/513) Fix simulated transaction bug that was causing a consensus error by unintentionally affecting the state.

## [v0.1.0] - 2020-08-23

### Improvements

* (sdk) [Karan-3108#386](https://github.com/cosmos/ethermint/pull/386) Bump Cosmos SDK version to [v0.39.1](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.1)
* (evm) [Karan-3108#181](https://github.com/cosmos/ethermint/issues/181) Updated EVM module to the recommended module structure.
* (app) [Karan-3108#188](https://github.com/cosmos/ethermint/issues/186)  Misc cleanup:
  * (evm) Rename `EthereumTxMsg` --> `MsgEthereumTx` and `EmintMsg` --> `MsgEthermint` for consistency with SDK standards
  * Updated integration and unit tests to use `EthermintApp` as testing suite
  * Use expected `Keeper` interface for `AccountKeeper`
  * Replaced `count` type in keeper with `int`
  * Add SDK events for transactions
* [Karan-3108#236](https://github.com/cosmos/ethermint/pull/236) Changes from upgrade:
  * (`app/ante`) Moved `AnteHandler` implementation to `app/ante`
  * (keys) Marked `ExportEthKeyCommand` as **UNSAFE**
  * (evm) Moved `BeginBlock` and `EndBlock` to `x/evm/abci.go`
* (evm) [Karan-3108#255](https://github.com/cosmos/ethermint/pull/255) Add missing `GenesisState` fields and support `ExportGenesis` functionality.
* [Karan-3108#272](https://github.com/cosmos/ethermint/pull/272) Add `Logger` for evm module.
* [Karan-3108#317](https://github.com/cosmos/ethermint/pull/317) `GenesisAccount` validation.
* (evm) [Karan-3108#319](https://github.com/cosmos/ethermint/pull/319) Various evm improvements:
  * Add transaction `[]*ethtypes.Logs` to evm's `GenesisState` to persist logs after an upgrade.
  * Remove evm `CodeKey` and `BlockKey`in favor of a prefix `Store`.
  * Set `BlockBloom` during `EndBlock` instead of `BeginBlock`.
  * `Commit` state object and `Finalize` storage after `InitGenesis` setup.
* (rpc) [Karan-3108#325](https://github.com/cosmos/ethermint/pull/325) `eth_coinbase` JSON-RPC query now returns the node's validator address.

### Features

* (build) [Karan-3108#378](https://github.com/cosmos/ethermint/pull/378) Create multi-node, local, automated testnet setup with `make localnet-start`.
* (rpc) [Karan-3108#330](https://github.com/cosmos/ethermint/issues/330) Implement `PublicFilterAPI`'s `EventSystem` which subscribes to Tendermint events upon `Filter` creation.
* (rpc) [Karan-3108#231](https://github.com/cosmos/ethermint/issues/231) Implement `NewBlockFilter` in rpc/filters.go which instantiates a polling block filter
  * Polls for new blocks via `BlockNumber` rpc call; if block number changes, it requests the new block via `GetBlockByNumber` rpc call and adds it to its internal list of blocks
  * Update `uninstallFilter` and `getFilterChanges` accordingly
  * `uninstallFilter` stops the polling goroutine
  * `getFilterChanges` returns the filter's internal list of block hashes and resets it
* (rpc) [Karan-3108#54](https://github.com/cosmos/ethermint/issues/54), [Karan-3108#55](https://github.com/cosmos/ethermint/issues/55)
  Implement `eth_getFilterLogs` and `eth_getLogs`:
  * For a given filter, look through each block for transactions. If there are transactions in the block, get the logs from it, and filter using the filterLogs method
  * `eth_getLogs` and `eth_getFilterChanges` for log filters use the same underlying method as `eth_getFilterLogs`
  * update `HandleMsgEthereumTx` to store logs using the ethereum hash
* (app) [Karan-3108#187](https://github.com/cosmos/ethermint/issues/187) Add support for simulations.

### Bug Fixes

* (evm) [Karan-3108#767](https://github.com/cosmos/ethermint/issues/767) Fix error of timeout when using Truffle to deploy contract.
* (evm) [Karan-3108#751](https://github.com/cosmos/ethermint/issues/751) Fix misused method to calculate block hash in evm related function.
* (evm) [Karan-3108#721](https://github.com/cosmos/ethermint/issues/721) Fix mismatch block hash in rpc response when use eth.getBlock.
* (evm) [Karan-3108#730](https://github.com/cosmos/ethermint/issues/730) Fix 'EIP2028' not open when Istanbul version has been enabled.
* (app) [Karan-3108#749](https://github.com/cosmos/ethermint/issues/749) Fix panic in `AnteHandler` when gas price larger than 100000
* (rpc) [Karan-3108#305](https://github.com/cosmos/ethermint/issues/305) Update `eth_getTransactionCount` to check for account existence before getting sequence and return 0 as the nonce if it doesn't exist.
* (evm) [Karan-3108#319](https://github.com/cosmos/ethermint/pull/319) Fix `SetBlockHash` that was setting the incorrect height during `BeginBlock`.
* (evm) [Karan-3108#176](https://github.com/cosmos/ethermint/issues/176) Updated Web3 transaction hash from using RLP hash. Now all transaction hashes exposed are amino hashes:
  * Removes `Hash()` (RLP) function from `MsgEthereumTx` to avoid confusion or misuse in future.
