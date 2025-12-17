
type hyperSyncConfig = {endpointUrl: string}
type hyperFuelConfig = {endpointUrl: string}

@genType.opaque
type rpcConfig = {
  syncConfig: InternalConfig.sourceSync,
}

@genType
type syncSource = HyperSync(hyperSyncConfig) | HyperFuel(hyperFuelConfig) | Rpc(rpcConfig)

@genType.opaque
type aliasAbi = Ethers.abi

type eventName = string

type contract = {
  name: string,
  abi: aliasAbi,
  addresses: array<string>,
  events: array<eventName>,
}

type configYaml = {
  syncSource,
  startBlock: int,
  confirmedBlockThreshold: int,
  contracts: dict<contract>,
  lowercaseAddresses: bool,
}

let publicConfig = ChainMap.fromArrayUnsafe([
  {
    let contracts = Js.Dict.fromArray([
      (
        "MockLending",
        {
          name: "MockLending",
          abi: Types.MockLending.abi,
          addresses: [
            "0x7871A1eE915f3085126f1676daEf75A24626EcBa",
          ],
          events: [
            Types.MockLending.HealthFactorUpdated.name,
            Types.MockLending.RescueExecuted.name,
          ],
        }
      ),
      (
        "GuardianManager",
        {
          name: "GuardianManager",
          abi: Types.GuardianManager.abi,
          addresses: [
            "0xc56f35202CD3c79F70351b213E8Fb130B678Fe7C",
          ],
          events: [
            Types.GuardianManager.FundsDelegated.name,
            Types.GuardianManager.AgentStatusUpdated.name,
          ],
        }
      ),
    ])
    let chain = ChainMap.Chain.makeUnsafe(~chainId=84532)
    (
      chain,
      {
        confirmedBlockThreshold: 200,
        syncSource: Rpc({syncConfig: Config.getSyncConfig({})}),
        startBlock: 34900000,
        contracts,
        lowercaseAddresses: false
      }
    )
  },
])

@genType
let getGeneratedByChainId: int => configYaml = chainId => {
  let chain = ChainMap.Chain.makeUnsafe(~chainId)
  if !(publicConfig->ChainMap.has(chain)) {
    Js.Exn.raiseError(
      "No chain with id " ++ chain->ChainMap.Chain.toString ++ " found in config.yaml",
    )
  }
  publicConfig->ChainMap.get(chain)
}
