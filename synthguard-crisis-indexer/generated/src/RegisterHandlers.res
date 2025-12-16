@val external require: string => unit = "require"

let registerContractHandlers = (
  ~contractName,
  ~handlerPathRelativeToRoot,
  ~handlerPathRelativeToConfig,
) => {
  try {
    require(`../${Path.relativePathToRootFromGenerated}/${handlerPathRelativeToRoot}`)
  } catch {
  | exn =>
    let params = {
      "Contract Name": contractName,
      "Expected Handler Path": handlerPathRelativeToConfig,
      "Code": "EE500",
    }
    let logger = Logging.createChild(~params)

    let errHandler = exn->ErrorHandling.make(~msg="Failed to import handler file", ~logger)
    errHandler->ErrorHandling.log
    errHandler->ErrorHandling.raiseExn
  }
}

let makeGeneratedConfig = () => {
  let chains = [
    {
      let contracts = [
        {
          InternalConfig.name: "MockLending",
          abi: Types.MockLending.abi,
          addresses: [
            "0xa0f95A73BA2c1395E9F4B95e6F6b7faF3E07A447"->Address.Evm.fromStringOrThrow
,
          ],
          events: [
            (Types.MockLending.HealthFactorUpdated.register() :> Internal.eventConfig),
            (Types.MockLending.RescueExecuted.register() :> Internal.eventConfig),
          ],
          startBlock: Some(35051000),
        },
      ]
      let chain = ChainMap.Chain.makeUnsafe(~chainId=84532)
      {
        InternalConfig.maxReorgDepth: 200,
        startBlock: 35051000,
        id: 84532,
        contracts,
        sources: NetworkSources.evm(~chain, ~contracts=[{name: "MockLending",events: [Types.MockLending.HealthFactorUpdated.register(), Types.MockLending.RescueExecuted.register()],abi: Types.MockLending.abi}], ~hyperSync=None, ~allEventSignatures=[Types.MockLending.eventSignatures]->Belt.Array.concatMany, ~shouldUseHypersyncClientDecoder=true, ~rpcs=[{url: "https://sepolia.base.org", sourceFor: Sync, syncConfig: {}}], ~lowercaseAddresses=false)
      }
    },
  ]

  Config.make(
    ~shouldRollbackOnReorg=true,
    ~shouldSaveFullHistory=false,
    ~isUnorderedMultichainMode=false,
    ~chains,
    ~enableRawEvents=false,
    ~batchSize=?Env.batchSize,
    ~preloadHandlers=false,
    ~lowercaseAddresses=false,
    ~shouldUseHypersyncClientDecoder=true,
  )
}

%%private(
  let config: ref<option<Config.t>> = ref(None)
)

let registerAllHandlers = () => {
  let configWithoutRegistrations = makeGeneratedConfig()
  EventRegister.startRegistration(
    ~ecosystem=configWithoutRegistrations.ecosystem,
    ~multichain=configWithoutRegistrations.multichain,
    ~preloadHandlers=configWithoutRegistrations.preloadHandlers,
  )

  registerContractHandlers(
    ~contractName="MockLending",
    ~handlerPathRelativeToRoot="src/EventHandlers.ts",
    ~handlerPathRelativeToConfig="src/EventHandlers.ts",
  )

  let generatedConfig = {
    // Need to recreate initial config one more time,
    // since configWithoutRegistrations called register for event
    // before they were ready
    ...makeGeneratedConfig(),
    registrations: Some(EventRegister.finishRegistration()),
  }
  config := Some(generatedConfig)
  generatedConfig
}

let getConfig = () => {
  switch config.contents {
  | Some(config) => config
  | None => registerAllHandlers()
  }
}

let getConfigWithoutRegistrations = makeGeneratedConfig
