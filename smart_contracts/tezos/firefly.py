import smartpy as sp

@sp.module
def main():
    class Firefly(sp.Contract):
        def __init__(self):
            pass
    
        @sp.entry_point
        def pinBatchData(self, uuids, batchHash):
            uuids, batchHash, payloadRef, contexts = sp.unpack(
                data, (sp.bytes, sp.bytes, sp.string, sp.list[sp.bytes])
            )
            self.data_batch_pin(
                sp.sender,
                sp.timestamp,
                "firefly:contract_invoke_pin",
                uuids,
                batchHash,
                payloadRef,
                contexts,
            )
    
        @sp.entry_point
        def pinBatch(self, uuids, batchHash, payloadRef, contexts):
            self.data_batch_pin(
                sp.sender,
                sp.timestamp,
                "firefly:batch_pin",
                uuids,
                batchHash,
                payloadRef,
                contexts,
            )
    
        @sp.entry_point
        def networkAction(self, action, payload):
            self.data_batch_pin(
                sp.sender,
                sp.timestamp,
                action,
                sp.bytes(0),
                sp.bytes(0),
                payload,
                sp.list[sp.bytes],
            )
    
        @sp.entry_point
        def networkVersion(self):
            sp.result(2)
