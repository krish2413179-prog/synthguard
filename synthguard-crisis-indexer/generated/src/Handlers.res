  @genType
module MockLending = {
  module HealthFactorUpdated = Types.MakeRegister(Types.MockLending.HealthFactorUpdated)
  module RescueExecuted = Types.MakeRegister(Types.MockLending.RescueExecuted)
}

@genType /** Register a Block Handler. It'll be called for every block by default. */
let onBlock: (
  Envio.onBlockOptions<Types.chain>,
  Envio.onBlockArgs<Types.handlerContext> => promise<unit>,
) => unit = (
  EventRegister.onBlock: (unknown, Internal.onBlockArgs => promise<unit>) => unit
)->Utils.magic
