# Chapter 3: Write a Log Package
## The Log is a powerful Tool
The ext filesystems, for example, log changes to a journal instead of directly changing the 's data filesystems
Logging to the journal is simple and fast, so there's little chance to losing data. When crashes before ext had finished updating the disk files, then in the next boot, the file would process the data and complete updates.

In Database, they records the change to the log, called a write-ahead log (WAL), and later process the WAL to the changes to the db data files.

DB developer also use the WAL for replication. Raft, a consensus algorithm, uses the same idea to get distributed services to agree on a cluster-wide state. Each node in a Raft cluster runs a state machine with a log as its input. The leader of the Raft cluster appends changes to its followers' logs.

Databases often provide a way to restore their state, referred to as point-in-time recovery.
