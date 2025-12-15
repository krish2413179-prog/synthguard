const { MockLending } = require("../generated");

/**
 * Health factor update
 */
MockLending.HealthFactorUpdated.handler(async ({ event, context }) => {
  try {
    const user = event.params.user;
    // CRITICAL: Ensure BigInt is converted to string for database and apply safety check
    const newHealth = (event.params.newHealth || 0n).toString(); 
    
    // CRITICAL: Safely access block number from event.block object
    const blockNum = (event.block.number || 0).toString(); 

    context.log.info(
      `CRISIS UPDATE: User ${user} health is now ${newHealth}`
    );

    await context.MockLending_HealthFactorUpdated.set({
      // ID FIX: Use event.block.number for maximum stability
      id: `${blockNum}_${event.logIndex}`, 
      user,
      newHealth,
      block_number: blockNum
    });
  } catch (err) {
    context.log.error("HealthFactorUpdated handler failed", err);
  }
});

/**
 * Rescue execution
 */
MockLending.RescueExecuted.handler(async ({ event, context }) => {
  try {
    const user = event.params.user;
    // CRITICAL: Ensure BigInt is converted to string for database and apply safety check
    const debtRepaid = (event.params.debtRepaid || 0n).toString();

    // CRITICAL: Safely access block number from event.block object
    const blockNum = (event.block.number || 0).toString();

    context.log.info(
      `RESCUE EXECUTED: Paid off ${debtRepaid} for user ${user}`
    );

    await context.MockLending_RescueExecuted.set({
      // ID FIX: Use event.block.number for maximum stability
      id: `${blockNum}_${event.logIndex}`,
      user,
      debtRepaid,
      block_number: blockNum
    });
  } catch (err) {
    context.log.error("RescueExecuted handler failed", err);
  }
});