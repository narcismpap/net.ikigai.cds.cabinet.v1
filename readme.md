# cds.cabinet
This is a prototype database developed against Apple's NewSQL open source database at https://www.foundationdb.org/

While this v1 is an early prototype of the system, subsequent versions have been successfully running large data clusters with nearly infinite horizontal scalability for our customers. This repository is kept for educational purposes.

This project is all about creating an infinitely horizontally scalable database based on very simple mechanism. Under the hood, we utilize the capabilities of FoundationDB, alongside extremely simple data structures which permit us to store complex data sets at scale.

## Documentation
You can read the technology specs for this prototype over at [V1 Final Docs](net.ikigai.cds.cabinet.v1.pdf).

## Cabinet V1 Implementation
Ikigai Cabinet is a globally-distributed, high-performance, horizontally scalable data store capable of performing a set of core operations designated as building blocks for any modern application.
The approach taken with this product is heavily inspired by the modern service-mesh platform design, where multiple self-contained layers can work in unison to build a sophisticated end- products, without a single point of failure.

This datastore implements learnings from existing classic relational databases (such as PostgreSQL), simple NoSQL document stores (MongoDB) and alternative representation data stores (Graph, Neo4J) and attempts to create a simple design which can be extended by multiple layers within the application code.

Cabinet does not intent to ever replace a traditional monolithic database and it will never attempt to achieve parity or compatibility with the functionalities provided by these systems. You can view the Cabinet as more of a highly-distributed, fast file system with additional abilities, while all the responsibilities of maintaining indexes, meta, edges and ensuring updates are always done in a single transaction will fall down to the application and associated developers.

Moving the complexity away from the database layer ensures a very simple product that is easy to maintain and poses very few constraints on the application development process. This does however slightly increase the engineering workload as they are fully responsible for the data stored: Ikigai Cabinet makes very few assumptions on what it right or wrong and does not perform high-level constraint, foreign key or schema integrity validations on its own.


## Elements

Core:
1. Nodes (document stores)
2. Edges (QUAD stores for relations, similar to a Graph database)
3. Indexes (views with searchable arbitrary properties)
4. Meta (meta values/property list)
5. Counts (distributed real-time atomical counts)
6. Sequential (sequential number generation in a distributed cluster)

Composite Elements:
1. Transactions (complex ReadCheck and Mutation requests processed in an ACID
environment)


## DB Design

1. Node: /n/{NODE_ID} = {protobuf}
2. Edge: /e/{SUBJECT_ID}/{PREDICATE}/{TARGET_ID} = {protobuf} Index: i/n/{INDEX_ID}/{VALUE}/{NODE_ID} = null
3. Meta: /m/{n/e/i}/{ID}/0xY = {binaryVal}
4. Count: /c/{n/e/i}/CountName/{ID}/0xY/{0x0 - OxF} = atomic<INT> Sequences: /s/{SEQ_NAME}/{SEQ_ID} = {NODE_ID}

### Compiling protobuf

    protoc --proto_path=./cabinet --go_out=plugins=grpc:./rpc cabinet/*.proto
    
### fdb

To build packages which use the go bindings, you will need to
set the following environment variables:

       CGO_CPPFLAGS="-I/Users/michael/Documents/Code/ikigai/net.ikigai.cds.cabinet.v1/src/github.com/apple/foundationdb/bindings/c"
       CGO_CFLAGS="-g -O2"
       CGO_LDFLAGS="-L/usr/local/lib"


Author: Narcis M. Pap, Singapore.
