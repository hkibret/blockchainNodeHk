"# blockchainNodeHk" 
# I'm building a blockchain node 
A BLOCKCHAIN NODE with blocks, transactions, and consensus mechanisms work together to form a secure, tamper‑resistant ledger. 
The project evolves step by step, starting from a simple append‑only chain and gradually adding realistic blockchain features.

# Features Implemented
Block structure: Each block contains an index, timestamp, transactions, previous hash, and its own hash.

Blockchain ledger: Append‑only chain with validation to detect tampering.

Transactions: Structured sender/receiver/amount records with cryptographic signatures (ECDSA).

Mempool: Pending transactions are collected before being mined into a block.

Proof of Work (PoW): Blocks require solving a computational puzzle (leading zeros in hash) before being added.

Validation: Ensures chain integrity by checking hashes and links between blocks.