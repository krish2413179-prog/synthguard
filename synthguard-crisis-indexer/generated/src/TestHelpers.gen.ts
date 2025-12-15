/* TypeScript file generated from TestHelpers.res by genType. */

/* eslint-disable */
/* tslint:disable */

const TestHelpersJS = require('./TestHelpers.res.js');

import type {MockLending_HealthFactorUpdated_event as Types_MockLending_HealthFactorUpdated_event} from './Types.gen';

import type {MockLending_RescueExecuted_event as Types_MockLending_RescueExecuted_event} from './Types.gen';

import type {t as Address_t} from 'envio/src/Address.gen';

import type {t as TestHelpers_MockDb_t} from './TestHelpers_MockDb.gen';

/** The arguements that get passed to a "processEvent" helper function */
export type EventFunctions_eventProcessorArgs<event> = {
  readonly event: event; 
  readonly mockDb: TestHelpers_MockDb_t; 
  readonly chainId?: number
};

export type EventFunctions_eventProcessor<event> = (_1:EventFunctions_eventProcessorArgs<event>) => Promise<TestHelpers_MockDb_t>;

export type EventFunctions_MockBlock_t = {
  readonly hash?: string; 
  readonly number?: number; 
  readonly timestamp?: number
};

export type EventFunctions_MockTransaction_t = {};

export type EventFunctions_mockEventData = {
  readonly chainId?: number; 
  readonly srcAddress?: Address_t; 
  readonly logIndex?: number; 
  readonly block?: EventFunctions_MockBlock_t; 
  readonly transaction?: EventFunctions_MockTransaction_t
};

export type MockLending_HealthFactorUpdated_createMockArgs = {
  readonly user?: Address_t; 
  readonly newHealth?: bigint; 
  readonly mockEventData?: EventFunctions_mockEventData
};

export type MockLending_RescueExecuted_createMockArgs = {
  readonly user?: Address_t; 
  readonly debtRepaid?: bigint; 
  readonly mockEventData?: EventFunctions_mockEventData
};

export const MockDb_createMockDb: () => TestHelpers_MockDb_t = TestHelpersJS.MockDb.createMockDb as any;

export const Addresses_mockAddresses: Address_t[] = TestHelpersJS.Addresses.mockAddresses as any;

export const Addresses_defaultAddress: Address_t = TestHelpersJS.Addresses.defaultAddress as any;

export const MockLending_HealthFactorUpdated_processEvent: EventFunctions_eventProcessor<Types_MockLending_HealthFactorUpdated_event> = TestHelpersJS.MockLending.HealthFactorUpdated.processEvent as any;

export const MockLending_HealthFactorUpdated_createMockEvent: (args:MockLending_HealthFactorUpdated_createMockArgs) => Types_MockLending_HealthFactorUpdated_event = TestHelpersJS.MockLending.HealthFactorUpdated.createMockEvent as any;

export const MockLending_RescueExecuted_processEvent: EventFunctions_eventProcessor<Types_MockLending_RescueExecuted_event> = TestHelpersJS.MockLending.RescueExecuted.processEvent as any;

export const MockLending_RescueExecuted_createMockEvent: (args:MockLending_RescueExecuted_createMockArgs) => Types_MockLending_RescueExecuted_event = TestHelpersJS.MockLending.RescueExecuted.createMockEvent as any;

export const Addresses: { mockAddresses: Address_t[]; defaultAddress: Address_t } = TestHelpersJS.Addresses as any;

export const MockLending: { RescueExecuted: { processEvent: EventFunctions_eventProcessor<Types_MockLending_RescueExecuted_event>; createMockEvent: (args:MockLending_RescueExecuted_createMockArgs) => Types_MockLending_RescueExecuted_event }; HealthFactorUpdated: { processEvent: EventFunctions_eventProcessor<Types_MockLending_HealthFactorUpdated_event>; createMockEvent: (args:MockLending_HealthFactorUpdated_createMockArgs) => Types_MockLending_HealthFactorUpdated_event } } = TestHelpersJS.MockLending as any;

export const MockDb: { createMockDb: () => TestHelpers_MockDb_t } = TestHelpersJS.MockDb as any;
