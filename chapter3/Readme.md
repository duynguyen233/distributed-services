# Chapter 3: Write a Log Package
## The Log is a powerful Tool
The ext filesystems, for example, log changes to a journal instead of directly changing the 's data filesystems
Logging to the journal is simple and fast, so there's little chance to losing data. When crashes before ext had finished updating the disk files, then in the next boot, the file would process the data and complete updates.

In Database, they records the change to the log, called a write-ahead log (WAL), and later process the WAL to the changes to the db data files.

DB developer also use the WAL for replication. Raft, a consensus algorithm, uses the same idea to get distributed services to agree on a cluster-wide state. Each node in a Raft cluster runs a state machine with a log as its input. The leader of the Raft cluster appends changes to its followers' logs.

Databases often provide a way to restore their state, referred to as point-in-time recovery.

## How logs work
- Logs append-only sequence of records
- Log assigns a record a unique and sequential offset number
- A log is like table that always orders the records by time and indexes each record by its offset and time created
- We don't having disks with infinite space -> split the log into a list of segments. When the logs grows too big, we free up disk space by deleting old segments.
- This cleaning up run in a background process while our service can still produce to the active segment and consume from other segments with no, or at least fewer, conflicts where goroutines access the same data
- Each segment's have a store file and index file. The segment's store file is where we store the record data; we continually append records to this file.
- The segment's index file is where we index each record in the store file -> speed up the process
- Two-step process: first you get the entry from the index for the record, which tells you the position of the record in the store file, and then you read the record at that position in the store file
