/* TypeScript file generated from Types.res by genType. */

/* eslint-disable */
/* tslint:disable */

import type {AgentStatus_t as Entities_AgentStatus_t} from '../src/db/Entities.gen';

import type {FundsDelegated_t as Entities_FundsDelegated_t} from '../src/db/Entities.gen';

import type {HandlerContext as $$handlerContext} from './Types.ts';

import type {HandlerWithOptions as $$fnWithEventConfig} from './bindings/OpaqueTypes.ts';

import type {LoaderContext as $$loaderContext} from './Types.ts';

import type {MockLending_HealthFactorUpdated_t as Entities_MockLending_HealthFactorUpdated_t} from '../src/db/Entities.gen';

import type {MockLending_RescueExecuted_t as Entities_MockLending_RescueExecuted_t} from '../src/db/Entities.gen';

import type {SingleOrMultiple as $$SingleOrMultiple_t} from './bindings/OpaqueTypes';

import type {entityHandlerContext as Internal_entityHandlerContext} from 'envio/src/Internal.gen';

import type {eventOptions as Internal_eventOptions} from 'envio/src/Internal.gen';

import type {genericContractRegisterArgs as Internal_genericContractRegisterArgs} from 'envio/src/Internal.gen';

import type {genericContractRegister as Internal_genericContractRegister} from 'envio/src/Internal.gen';

import type {genericEvent as Internal_genericEvent} from 'envio/src/Internal.gen';

import type {genericHandlerArgs as Internal_genericHandlerArgs} from 'envio/src/Internal.gen';

import type {genericHandlerWithLoader as Internal_genericHandlerWithLoader} from 'envio/src/Internal.gen';

import type {genericHandler as Internal_genericHandler} from 'envio/src/Internal.gen';

import type {genericLoaderArgs as Internal_genericLoaderArgs} from 'envio/src/Internal.gen';

import type {genericLoader as Internal_genericLoader} from 'envio/src/Internal.gen';

import type {logger as Envio_logger} from 'envio/src/Envio.gen';

import type {t as Address_t} from 'envio/src/Address.gen';

export type id = string;
export type Id = id;

export type contractRegistrations = {
  readonly log: Envio_logger; 
  readonly addGuardianManager: (_1:Address_t) => void; 
  readonly addMockLending: (_1:Address_t) => void
};

export type entityLoaderContext<entity,indexedFieldOperations> = {
  readonly get: (_1:id) => Promise<(undefined | entity)>; 
  readonly getOrThrow: (_1:id, message:(undefined | string)) => Promise<entity>; 
  readonly getWhere: indexedFieldOperations; 
  readonly getOrCreate: (_1:entity) => Promise<entity>; 
  readonly set: (_1:entity) => void; 
  readonly deleteUnsafe: (_1:id) => void
};

export type loaderContext = $$loaderContext;

export type entityHandlerContext<entity> = Internal_entityHandlerContext<entity>;

export type handlerContext = $$handlerContext;

export type agentStatus = Entities_AgentStatus_t;
export type AgentStatus = agentStatus;

export type fundsDelegated = Entities_FundsDelegated_t;
export type FundsDelegated = fundsDelegated;

export type mockLending_HealthFactorUpdated = Entities_MockLending_HealthFactorUpdated_t;
export type MockLending_HealthFactorUpdated = mockLending_HealthFactorUpdated;

export type mockLending_RescueExecuted = Entities_MockLending_RescueExecuted_t;
export type MockLending_RescueExecuted = mockLending_RescueExecuted;

export type Transaction_t = {};

export type Block_t = {
  readonly number: number; 
  readonly timestamp: number; 
  readonly hash: string
};

export type AggregatedBlock_t = {
  readonly hash: string; 
  readonly number: number; 
  readonly timestamp: number
};

export type AggregatedTransaction_t = {};

export type eventLog<params> = Internal_genericEvent<params,Block_t,Transaction_t>;
export type EventLog<params> = eventLog<params>;

export type SingleOrMultiple_t<a> = $$SingleOrMultiple_t<a>;

export type HandlerTypes_args<eventArgs,context> = { readonly event: eventLog<eventArgs>; readonly context: context };

export type HandlerTypes_contractRegisterArgs<eventArgs> = Internal_genericContractRegisterArgs<eventLog<eventArgs>,contractRegistrations>;

export type HandlerTypes_contractRegister<eventArgs> = Internal_genericContractRegister<HandlerTypes_contractRegisterArgs<eventArgs>>;

export type HandlerTypes_loaderArgs<eventArgs> = Internal_genericLoaderArgs<eventLog<eventArgs>,loaderContext>;

export type HandlerTypes_loader<eventArgs,loaderReturn> = Internal_genericLoader<HandlerTypes_loaderArgs<eventArgs>,loaderReturn>;

export type HandlerTypes_handlerArgs<eventArgs,loaderReturn> = Internal_genericHandlerArgs<eventLog<eventArgs>,handlerContext,loaderReturn>;

export type HandlerTypes_handler<eventArgs,loaderReturn> = Internal_genericHandler<HandlerTypes_handlerArgs<eventArgs,loaderReturn>>;

export type HandlerTypes_loaderHandler<eventArgs,loaderReturn,eventFilters> = Internal_genericHandlerWithLoader<HandlerTypes_loader<eventArgs,loaderReturn>,HandlerTypes_handler<eventArgs,loaderReturn>,eventFilters>;

export type HandlerTypes_eventConfig<eventFilters> = Internal_eventOptions<eventFilters>;

export type fnWithEventConfig<fn,eventConfig> = $$fnWithEventConfig<fn,eventConfig>;

export type handlerWithOptions<eventArgs,loaderReturn,eventFilters> = fnWithEventConfig<HandlerTypes_handler<eventArgs,loaderReturn>,HandlerTypes_eventConfig<eventFilters>>;

export type contractRegisterWithOptions<eventArgs,eventFilters> = fnWithEventConfig<HandlerTypes_contractRegister<eventArgs>,HandlerTypes_eventConfig<eventFilters>>;

export type GuardianManager_chainId = 84532;

export type GuardianManager_FundsDelegated_eventArgs = {
  readonly user: Address_t; 
  readonly workerAgent: Address_t; 
  readonly amount: bigint
};

export type GuardianManager_FundsDelegated_block = Block_t;

export type GuardianManager_FundsDelegated_transaction = Transaction_t;

export type GuardianManager_FundsDelegated_event = {
  /** The parameters or arguments associated with this event. */
  readonly params: GuardianManager_FundsDelegated_eventArgs; 
  /** The unique identifier of the blockchain network where this event occurred. */
  readonly chainId: GuardianManager_chainId; 
  /** The address of the contract that emitted this event. */
  readonly srcAddress: Address_t; 
  /** The index of this event's log within the block. */
  readonly logIndex: number; 
  /** The transaction that triggered this event. Configurable in `config.yaml` via the `field_selection` option. */
  readonly transaction: GuardianManager_FundsDelegated_transaction; 
  /** The block in which this event was recorded. Configurable in `config.yaml` via the `field_selection` option. */
  readonly block: GuardianManager_FundsDelegated_block
};

export type GuardianManager_FundsDelegated_loaderArgs = Internal_genericLoaderArgs<GuardianManager_FundsDelegated_event,loaderContext>;

export type GuardianManager_FundsDelegated_loader<loaderReturn> = Internal_genericLoader<GuardianManager_FundsDelegated_loaderArgs,loaderReturn>;

export type GuardianManager_FundsDelegated_handlerArgs<loaderReturn> = Internal_genericHandlerArgs<GuardianManager_FundsDelegated_event,handlerContext,loaderReturn>;

export type GuardianManager_FundsDelegated_handler<loaderReturn> = Internal_genericHandler<GuardianManager_FundsDelegated_handlerArgs<loaderReturn>>;

export type GuardianManager_FundsDelegated_contractRegister = Internal_genericContractRegister<Internal_genericContractRegisterArgs<GuardianManager_FundsDelegated_event,contractRegistrations>>;

export type GuardianManager_FundsDelegated_eventFilter = { readonly user?: SingleOrMultiple_t<Address_t>; readonly workerAgent?: SingleOrMultiple_t<Address_t> };

export type GuardianManager_FundsDelegated_eventFiltersArgs = { 
/** The unique identifier of the blockchain network where this event occurred. */
readonly chainId: GuardianManager_chainId; 
/** Addresses of the contracts indexing the event. */
readonly addresses: Address_t[] };

export type GuardianManager_FundsDelegated_eventFiltersDefinition = 
    GuardianManager_FundsDelegated_eventFilter
  | GuardianManager_FundsDelegated_eventFilter[];

export type GuardianManager_FundsDelegated_eventFilters = 
    GuardianManager_FundsDelegated_eventFilter
  | GuardianManager_FundsDelegated_eventFilter[]
  | ((_1:GuardianManager_FundsDelegated_eventFiltersArgs) => GuardianManager_FundsDelegated_eventFiltersDefinition);

export type GuardianManager_AgentStatusUpdated_eventArgs = { readonly agent: Address_t; readonly isActive: boolean };

export type GuardianManager_AgentStatusUpdated_block = Block_t;

export type GuardianManager_AgentStatusUpdated_transaction = Transaction_t;

export type GuardianManager_AgentStatusUpdated_event = {
  /** The parameters or arguments associated with this event. */
  readonly params: GuardianManager_AgentStatusUpdated_eventArgs; 
  /** The unique identifier of the blockchain network where this event occurred. */
  readonly chainId: GuardianManager_chainId; 
  /** The address of the contract that emitted this event. */
  readonly srcAddress: Address_t; 
  /** The index of this event's log within the block. */
  readonly logIndex: number; 
  /** The transaction that triggered this event. Configurable in `config.yaml` via the `field_selection` option. */
  readonly transaction: GuardianManager_AgentStatusUpdated_transaction; 
  /** The block in which this event was recorded. Configurable in `config.yaml` via the `field_selection` option. */
  readonly block: GuardianManager_AgentStatusUpdated_block
};

export type GuardianManager_AgentStatusUpdated_loaderArgs = Internal_genericLoaderArgs<GuardianManager_AgentStatusUpdated_event,loaderContext>;

export type GuardianManager_AgentStatusUpdated_loader<loaderReturn> = Internal_genericLoader<GuardianManager_AgentStatusUpdated_loaderArgs,loaderReturn>;

export type GuardianManager_AgentStatusUpdated_handlerArgs<loaderReturn> = Internal_genericHandlerArgs<GuardianManager_AgentStatusUpdated_event,handlerContext,loaderReturn>;

export type GuardianManager_AgentStatusUpdated_handler<loaderReturn> = Internal_genericHandler<GuardianManager_AgentStatusUpdated_handlerArgs<loaderReturn>>;

export type GuardianManager_AgentStatusUpdated_contractRegister = Internal_genericContractRegister<Internal_genericContractRegisterArgs<GuardianManager_AgentStatusUpdated_event,contractRegistrations>>;

export type GuardianManager_AgentStatusUpdated_eventFilter = { readonly agent?: SingleOrMultiple_t<Address_t> };

export type GuardianManager_AgentStatusUpdated_eventFiltersArgs = { 
/** The unique identifier of the blockchain network where this event occurred. */
readonly chainId: GuardianManager_chainId; 
/** Addresses of the contracts indexing the event. */
readonly addresses: Address_t[] };

export type GuardianManager_AgentStatusUpdated_eventFiltersDefinition = 
    GuardianManager_AgentStatusUpdated_eventFilter
  | GuardianManager_AgentStatusUpdated_eventFilter[];

export type GuardianManager_AgentStatusUpdated_eventFilters = 
    GuardianManager_AgentStatusUpdated_eventFilter
  | GuardianManager_AgentStatusUpdated_eventFilter[]
  | ((_1:GuardianManager_AgentStatusUpdated_eventFiltersArgs) => GuardianManager_AgentStatusUpdated_eventFiltersDefinition);

export type MockLending_chainId = 84532;

export type MockLending_HealthFactorUpdated_eventArgs = { readonly user: Address_t; readonly newHealth: bigint };

export type MockLending_HealthFactorUpdated_block = Block_t;

export type MockLending_HealthFactorUpdated_transaction = Transaction_t;

export type MockLending_HealthFactorUpdated_event = {
  /** The parameters or arguments associated with this event. */
  readonly params: MockLending_HealthFactorUpdated_eventArgs; 
  /** The unique identifier of the blockchain network where this event occurred. */
  readonly chainId: MockLending_chainId; 
  /** The address of the contract that emitted this event. */
  readonly srcAddress: Address_t; 
  /** The index of this event's log within the block. */
  readonly logIndex: number; 
  /** The transaction that triggered this event. Configurable in `config.yaml` via the `field_selection` option. */
  readonly transaction: MockLending_HealthFactorUpdated_transaction; 
  /** The block in which this event was recorded. Configurable in `config.yaml` via the `field_selection` option. */
  readonly block: MockLending_HealthFactorUpdated_block
};

export type MockLending_HealthFactorUpdated_loaderArgs = Internal_genericLoaderArgs<MockLending_HealthFactorUpdated_event,loaderContext>;

export type MockLending_HealthFactorUpdated_loader<loaderReturn> = Internal_genericLoader<MockLending_HealthFactorUpdated_loaderArgs,loaderReturn>;

export type MockLending_HealthFactorUpdated_handlerArgs<loaderReturn> = Internal_genericHandlerArgs<MockLending_HealthFactorUpdated_event,handlerContext,loaderReturn>;

export type MockLending_HealthFactorUpdated_handler<loaderReturn> = Internal_genericHandler<MockLending_HealthFactorUpdated_handlerArgs<loaderReturn>>;

export type MockLending_HealthFactorUpdated_contractRegister = Internal_genericContractRegister<Internal_genericContractRegisterArgs<MockLending_HealthFactorUpdated_event,contractRegistrations>>;

export type MockLending_HealthFactorUpdated_eventFilter = { readonly user?: SingleOrMultiple_t<Address_t> };

export type MockLending_HealthFactorUpdated_eventFiltersArgs = { 
/** The unique identifier of the blockchain network where this event occurred. */
readonly chainId: MockLending_chainId; 
/** Addresses of the contracts indexing the event. */
readonly addresses: Address_t[] };

export type MockLending_HealthFactorUpdated_eventFiltersDefinition = 
    MockLending_HealthFactorUpdated_eventFilter
  | MockLending_HealthFactorUpdated_eventFilter[];

export type MockLending_HealthFactorUpdated_eventFilters = 
    MockLending_HealthFactorUpdated_eventFilter
  | MockLending_HealthFactorUpdated_eventFilter[]
  | ((_1:MockLending_HealthFactorUpdated_eventFiltersArgs) => MockLending_HealthFactorUpdated_eventFiltersDefinition);

export type MockLending_RescueExecuted_eventArgs = { readonly user: Address_t; readonly debtRepaid: bigint };

export type MockLending_RescueExecuted_block = Block_t;

export type MockLending_RescueExecuted_transaction = Transaction_t;

export type MockLending_RescueExecuted_event = {
  /** The parameters or arguments associated with this event. */
  readonly params: MockLending_RescueExecuted_eventArgs; 
  /** The unique identifier of the blockchain network where this event occurred. */
  readonly chainId: MockLending_chainId; 
  /** The address of the contract that emitted this event. */
  readonly srcAddress: Address_t; 
  /** The index of this event's log within the block. */
  readonly logIndex: number; 
  /** The transaction that triggered this event. Configurable in `config.yaml` via the `field_selection` option. */
  readonly transaction: MockLending_RescueExecuted_transaction; 
  /** The block in which this event was recorded. Configurable in `config.yaml` via the `field_selection` option. */
  readonly block: MockLending_RescueExecuted_block
};

export type MockLending_RescueExecuted_loaderArgs = Internal_genericLoaderArgs<MockLending_RescueExecuted_event,loaderContext>;

export type MockLending_RescueExecuted_loader<loaderReturn> = Internal_genericLoader<MockLending_RescueExecuted_loaderArgs,loaderReturn>;

export type MockLending_RescueExecuted_handlerArgs<loaderReturn> = Internal_genericHandlerArgs<MockLending_RescueExecuted_event,handlerContext,loaderReturn>;

export type MockLending_RescueExecuted_handler<loaderReturn> = Internal_genericHandler<MockLending_RescueExecuted_handlerArgs<loaderReturn>>;

export type MockLending_RescueExecuted_contractRegister = Internal_genericContractRegister<Internal_genericContractRegisterArgs<MockLending_RescueExecuted_event,contractRegistrations>>;

export type MockLending_RescueExecuted_eventFilter = { readonly user?: SingleOrMultiple_t<Address_t> };

export type MockLending_RescueExecuted_eventFiltersArgs = { 
/** The unique identifier of the blockchain network where this event occurred. */
readonly chainId: MockLending_chainId; 
/** Addresses of the contracts indexing the event. */
readonly addresses: Address_t[] };

export type MockLending_RescueExecuted_eventFiltersDefinition = 
    MockLending_RescueExecuted_eventFilter
  | MockLending_RescueExecuted_eventFilter[];

export type MockLending_RescueExecuted_eventFilters = 
    MockLending_RescueExecuted_eventFilter
  | MockLending_RescueExecuted_eventFilter[]
  | ((_1:MockLending_RescueExecuted_eventFiltersArgs) => MockLending_RescueExecuted_eventFiltersDefinition);

export type chainId = number;

export type chain = 84532;
