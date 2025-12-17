module ContractType = {
  @genType
  type t = 
    | @as("GuardianManager") GuardianManager
    | @as("MockLending") MockLending

  let name = "CONTRACT_TYPE"
  let variants = [
    GuardianManager,
    MockLending,
  ]
  let config = Internal.makeEnumConfig(~name, ~variants)
}

module EntityType = {
  @genType
  type t = 
    | @as("AgentStatus") AgentStatus
    | @as("FundsDelegated") FundsDelegated
    | @as("MockLending_HealthFactorUpdated") MockLending_HealthFactorUpdated
    | @as("MockLending_RescueExecuted") MockLending_RescueExecuted
    | @as("dynamic_contract_registry") DynamicContractRegistry

  let name = "ENTITY_TYPE"
  let variants = [
    AgentStatus,
    FundsDelegated,
    MockLending_HealthFactorUpdated,
    MockLending_RescueExecuted,
    DynamicContractRegistry,
  ]
  let config = Internal.makeEnumConfig(~name, ~variants)
}

let allEnums = ([
  ContractType.config->Internal.fromGenericEnumConfig,
  EntityType.config->Internal.fromGenericEnumConfig,
])
