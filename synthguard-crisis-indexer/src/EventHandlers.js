const { MockLending } = require("../generated");

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
