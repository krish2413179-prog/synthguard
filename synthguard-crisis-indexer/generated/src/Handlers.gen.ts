/* TypeScript file generated from Handlers.res by genType. */

/* eslint-disable */
/* tslint:disable */

const HandlersJS = require('./Handlers.res.js');

import type {GuardianManager_AgentStatusUpdated_eventFilters as Types_GuardianManager_AgentStatusUpdated_eventFilters} from './Types.gen';

import type {GuardianManager_AgentStatusUpdated_event as Types_GuardianManager_AgentStatusUpdated_event} from './Types.gen';

import type {GuardianManager_FundsDelegated_eventFilters as Types_GuardianManager_FundsDelegated_eventFilters} from './Types.gen';

import type {GuardianManager_FundsDelegated_event as Types_GuardianManager_FundsDelegated_event} from './Types.gen';

import type {HandlerTypes_eventConfig as Types_HandlerTypes_eventConfig} from './Types.gen';

import type {MockLending_HealthFactorUpdated_eventFilters as Types_MockLending_HealthFactorUpdated_eventFilters} from './Types.gen';

import type {MockLending_HealthFactorUpdated_event as Types_MockLending_HealthFactorUpdated_event} from './Types.gen';

import type {MockLending_RescueExecuted_eventFilters as Types_MockLending_RescueExecuted_eventFilters} from './Types.gen';

import type {MockLending_RescueExecuted_event as Types_MockLending_RescueExecuted_event} from './Types.gen';

import type {chain as Types_chain} from './Types.gen';

import type {contractRegistrations as Types_contractRegistrations} from './Types.gen';

import type {fnWithEventConfig as Types_fnWithEventConfig} from './Types.gen';

import type {genericContractRegisterArgs as Internal_genericContractRegisterArgs} from 'envio/src/Internal.gen';

import type {genericContractRegister as Internal_genericContractRegister} from 'envio/src/Internal.gen';

import type {genericHandlerArgs as Internal_genericHandlerArgs} from 'envio/src/Internal.gen';

import type {genericHandlerWithLoader as Internal_genericHandlerWithLoader} from 'envio/src/Internal.gen';

import type {genericHandler as Internal_genericHandler} from 'envio/src/Internal.gen';

import type {genericLoaderArgs as Internal_genericLoaderArgs} from 'envio/src/Internal.gen';

import type {genericLoader as Internal_genericLoader} from 'envio/src/Internal.gen';

import type {handlerContext as Types_handlerContext} from './Types.gen';

import type {loaderContext as Types_loaderContext} from './Types.gen';

import type {onBlockArgs as Envio_onBlockArgs} from 'envio/src/Envio.gen';

import type {onBlockOptions as Envio_onBlockOptions} from 'envio/src/Envio.gen';

export const GuardianManager_FundsDelegated_contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_GuardianManager_FundsDelegated_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_FundsDelegated_eventFilters>> = HandlersJS.GuardianManager.FundsDelegated.contractRegister as any;

export const GuardianManager_FundsDelegated_handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_FundsDelegated_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_FundsDelegated_eventFilters>> = HandlersJS.GuardianManager.FundsDelegated.handler as any;

export const GuardianManager_FundsDelegated_handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_GuardianManager_FundsDelegated_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_FundsDelegated_event,Types_handlerContext,loaderReturn>>,Types_GuardianManager_FundsDelegated_eventFilters>) => void = HandlersJS.GuardianManager.FundsDelegated.handlerWithLoader as any;

export const GuardianManager_AgentStatusUpdated_contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_AgentStatusUpdated_eventFilters>> = HandlersJS.GuardianManager.AgentStatusUpdated.contractRegister as any;

export const GuardianManager_AgentStatusUpdated_handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_AgentStatusUpdated_eventFilters>> = HandlersJS.GuardianManager.AgentStatusUpdated.handler as any;

export const GuardianManager_AgentStatusUpdated_handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_handlerContext,loaderReturn>>,Types_GuardianManager_AgentStatusUpdated_eventFilters>) => void = HandlersJS.GuardianManager.AgentStatusUpdated.handlerWithLoader as any;

export const MockLending_HealthFactorUpdated_contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_MockLending_HealthFactorUpdated_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_MockLending_HealthFactorUpdated_eventFilters>> = HandlersJS.MockLending.HealthFactorUpdated.contractRegister as any;

export const MockLending_HealthFactorUpdated_handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_HealthFactorUpdated_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_MockLending_HealthFactorUpdated_eventFilters>> = HandlersJS.MockLending.HealthFactorUpdated.handler as any;

export const MockLending_HealthFactorUpdated_handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_MockLending_HealthFactorUpdated_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_HealthFactorUpdated_event,Types_handlerContext,loaderReturn>>,Types_MockLending_HealthFactorUpdated_eventFilters>) => void = HandlersJS.MockLending.HealthFactorUpdated.handlerWithLoader as any;

export const MockLending_RescueExecuted_contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_MockLending_RescueExecuted_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_MockLending_RescueExecuted_eventFilters>> = HandlersJS.MockLending.RescueExecuted.contractRegister as any;

export const MockLending_RescueExecuted_handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_RescueExecuted_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_MockLending_RescueExecuted_eventFilters>> = HandlersJS.MockLending.RescueExecuted.handler as any;

export const MockLending_RescueExecuted_handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_MockLending_RescueExecuted_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_RescueExecuted_event,Types_handlerContext,loaderReturn>>,Types_MockLending_RescueExecuted_eventFilters>) => void = HandlersJS.MockLending.RescueExecuted.handlerWithLoader as any;

/** Register a Block Handler. It'll be called for every block by default. */
export const onBlock: (_1:Envio_onBlockOptions<Types_chain>, _2:((_1:Envio_onBlockArgs<Types_handlerContext>) => Promise<void>)) => void = HandlersJS.onBlock as any;

export const MockLending: { RescueExecuted: {
  handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_MockLending_RescueExecuted_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_RescueExecuted_event,Types_handlerContext,loaderReturn>>,Types_MockLending_RescueExecuted_eventFilters>) => void; 
  handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_RescueExecuted_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_MockLending_RescueExecuted_eventFilters>>; 
  contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_MockLending_RescueExecuted_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_MockLending_RescueExecuted_eventFilters>>
}; HealthFactorUpdated: {
  handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_MockLending_HealthFactorUpdated_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_HealthFactorUpdated_event,Types_handlerContext,loaderReturn>>,Types_MockLending_HealthFactorUpdated_eventFilters>) => void; 
  handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_MockLending_HealthFactorUpdated_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_MockLending_HealthFactorUpdated_eventFilters>>; 
  contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_MockLending_HealthFactorUpdated_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_MockLending_HealthFactorUpdated_eventFilters>>
} } = HandlersJS.MockLending as any;

export const GuardianManager: { FundsDelegated: {
  handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_GuardianManager_FundsDelegated_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_FundsDelegated_event,Types_handlerContext,loaderReturn>>,Types_GuardianManager_FundsDelegated_eventFilters>) => void; 
  handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_FundsDelegated_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_FundsDelegated_eventFilters>>; 
  contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_GuardianManager_FundsDelegated_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_FundsDelegated_eventFilters>>
}; AgentStatusUpdated: {
  handlerWithLoader: <loaderReturn>(_1:Internal_genericHandlerWithLoader<Internal_genericLoader<Internal_genericLoaderArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_loaderContext>,loaderReturn>,Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_handlerContext,loaderReturn>>,Types_GuardianManager_AgentStatusUpdated_eventFilters>) => void; 
  handler: Types_fnWithEventConfig<Internal_genericHandler<Internal_genericHandlerArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_handlerContext,void>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_AgentStatusUpdated_eventFilters>>; 
  contractRegister: Types_fnWithEventConfig<Internal_genericContractRegister<Internal_genericContractRegisterArgs<Types_GuardianManager_AgentStatusUpdated_event,Types_contractRegistrations>>,Types_HandlerTypes_eventConfig<Types_GuardianManager_AgentStatusUpdated_eventFilters>>
} } = HandlersJS.GuardianManager as any;
