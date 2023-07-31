# neo4j-mock-pool-dependencies

Mock pools, dependencies and relationships between them

## Quick Start

1. Start a neo4j instance:
```bash
docker compose up -d
```

2. Generate mock data:
```bash
go run *.go
```

3. Query data using Cypher (Neo4j Desktop app)

Query all data:
```bash
MATCH (n) RETURN (n)
```

Query all pools from the list of dependencies:
```bash
MATCH (p:Contract)-[r:DEPEND_ON]->(d:Contract)
WHERE d.address IN [<list of dependency addresses>]
RETURN p.address
```
