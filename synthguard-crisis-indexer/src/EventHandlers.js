const { MockLending } = require("../generated");
const { GuardianManager } = require("../generated");

// Handler for when the Manager delegates funds to a Worker (A2A Event)
GuardianManager.FundsDelegated.handler(async ({ event, context }) => {
  const entity = {
    id: event.transactionHash + event.logIndex.toString(),
    user: event.params.user,
    workerAgent: event.params.workerAgent,
    amount: event.params.amount,
    timestamp: event.block.timestamp,
  };

  context.FundsDelegated.set(entity);
});

// Handler for when a new Agent (e.g. DCA, Rescue) is approved
GuardianManager.AgentStatusUpdated.handler(async ({ event, context }) => {
  const entity = {
    id: event.params.agent.toString(), // Use agent address as ID
    agent: event.params.agent,
    isActive: event.params.isActive,
  };

  context.AgentStatus.set(entity);
});

/**
 * Health factor update
 */
MockLending.HealthFactorUpdated.handler(async ({ event, context }) => {
  try {
    await context.MockLending_HealthFactorUpdated.set({
      id: `${event.block.number}_${event.logIndex}`,
      user: event.params.user.toLowerCase(),
      newHealth: event.params.newHealth ?? 0n,
      block_number: event.block.number
    });

    context.log.info(
      `CRISIS UPDATE: user=${event.params.user} health=${event.params.newHealth}`
    );
  } catch (err) {
    context.log.error("HealthFactorUpdated handler failed", err);
  }
});

/**
 * Rescue execution
 */
MockLending.RescueExecuted.handler(async ({ event, context }) => {
  try {
    await context.MockLending_RescueExecuted.set({
      id: `${event.block.number}_${event.logIndex}`,
      user: event.params.user.toLowerCase(),
      debtRepaid: event.params.debtRepaid ?? 0n,
      block_number: event.block.number
    });

    context.log.info(
      `RESCUE EXECUTED: user=${event.params.user} repaid=${event.params.debtRepaid}`
    );
  } catch (err) {
    context.log.error("RescueExecuted handler failed", err);
  }
});
