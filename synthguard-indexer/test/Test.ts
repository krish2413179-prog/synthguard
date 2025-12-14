import assert from "assert";
import { 
  TestHelpers,
  USDT_Approval
} from "generated";
const { MockDb, USDT } = TestHelpers;

describe("USDT contract Approval event tests", () => {
  // Create mock db
  const mockDb = MockDb.createMockDb();

  // Creating mock for USDT contract Approval event
  const event = USDT.Approval.createMockEvent({/* It mocks event fields with default values. You can overwrite them if you need */});

  it("USDT_Approval is created correctly", async () => {
    // Processing the event
    const mockDbUpdated = await USDT.Approval.processEvent({
      event,
      mockDb,
    });

    // Getting the actual entity from the mock database
    let actualUSDTApproval = mockDbUpdated.entities.USDT_Approval.get(
      `${event.chainId}_${event.block.number}_${event.logIndex}`
    );

    // Creating the expected entity
    const expectedUSDTApproval: USDT_Approval = {
      id: `${event.chainId}_${event.block.number}_${event.logIndex}`,
      owner: event.params.owner,
      spender: event.params.spender,
      value: event.params.value,
    };
    // Asserting that the entity in the mock database is the same as the expected entity
    assert.deepEqual(actualUSDTApproval, expectedUSDTApproval, "Actual USDTApproval should be the same as the expectedUSDTApproval");
  });
});
