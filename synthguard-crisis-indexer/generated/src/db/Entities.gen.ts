/* TypeScript file generated from Entities.res by genType. */

/* eslint-disable */
/* tslint:disable */

export type id = string;

export type whereOperations<entity,fieldType> = { readonly eq: (_1:fieldType) => Promise<entity[]>; readonly gt: (_1:fieldType) => Promise<entity[]> };

export type MockLending_HealthFactorUpdated_t = {
  readonly block_number: bigint; 
  readonly id: id; 
  readonly newHealth: bigint; 
  readonly user: string
};

export type MockLending_HealthFactorUpdated_indexedFieldOperations = {};

export type MockLending_RescueExecuted_t = {
  readonly block_number: bigint; 
  readonly debtRepaid: bigint; 
  readonly id: id; 
  readonly user: string
};

export type MockLending_RescueExecuted_indexedFieldOperations = {};
