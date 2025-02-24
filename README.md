# blockchain

## Consensus-Breaking Change

### 1. What Does It Mean by Breaking Consensus?
Breaking consensus refers to introducing changes to the blockchain that make previously valid blocks or transactions invalid, or vice versa. This occurs when the blockchain's rules (protocol) are altered in a way that nodes running the old version of the software can no longer agree on the state of the blockchain with nodes running the new version. Such changes require all validators and full nodes to upgrade their software to maintain network consensus.

### 2. Why Does This Change Break Consensus?
The introduced change appends a `_v2` suffix to the `Name` field of the `Resource` during resource creation. This alters the state of the blockchain by changing how resources are stored and identified. Nodes running the old version would compute a different state for the same transaction because they do not append the `_v2` suffix, leading to a divergence in the blockchain's state and breaking consensus.

