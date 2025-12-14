/*
 * Please refer to https://docs.envio.dev for a thorough guide on all Envio indexer features
 */
import {
  USDT,
  USDT_Approval,
  USDT_Transfer,
} from "generated";

USDT.Approval.handler(async ({ event, context }) => {
  const entity: USDT_Approval = {
    id: `${event.chainId}_${event.block.number}_${event.logIndex}`,
    owner: event.params.owner,
    spender: event.params.spender,
    value: event.params.value,
  };

  context.USDT_Approval.set(entity);
});

USDT.Transfer.handler(async ({ event, context }) => {
  const entity: USDT_Transfer = {
    id: `${event.chainId}_${event.block.number}_${event.logIndex}`,
    from: event.params.from,
    to: event.params.to,
    value: event.params.value,
  };

  context.USDT_Transfer.set(entity);
});
