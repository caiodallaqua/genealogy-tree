# Genealogy Tree API

Backed by Neo4J, the Genealogy Tree API allows for Person entities manipulation and their parental relationships for managing genealogy trees.

## Build

To run the project, just execute `docker compose up -d` in the project's root folder. This will build three services:

- Genealogy Tree: REST API, available for requests on [localhost:8998](http://localhost:8998);
- Neo4J: NoSQL Graph Database powered by Cypher (a query language for graphs);
- Swagger UI: A web interface for documentation at [localhost:8080](http://localhost:8080), where you can see the available endpoints, their examples and more.

## Data Model

```json
{
  "person": {
    "id": "(unique, auto-generated) positive integer",
    "name": "(required) string",
    "birth": "(optional) 2020-12-09T16:09:53+00:00"
  }
}
```

To traverse the graph through parental relationships, two types of edges were defined:

```json
{
  "edge": {
    "type": "PARENT_OF"
  },

  "edge": {
    "type": "CHILD_OF"
  }
}
```

Both connecting exactly two nodes (a start node and an end one) as usual in standard graph theory.

**Example**

Let E be Eduardo Suplicy and S be Supla. Eduardo is Supla's father. their parental relationship is given by:

```(E:person)-[:PARENT_OF]->(S:person) AND (S:person)-[:CHILD_OF]->(E:person)```

As you might guess, the edges are redundant. Throughout the queries it is assumed that whenever one exists the other will be there too. Combining them in that way provides a good flexibility.

## Project Architecture

Pattern: Ports & Adapters.

```sh
.
├── cmd
├── docs
├── env
└── internal
    ├── adapters
    │   ├── api
    │   ├── database
    │   └── rest
    ├── debug
    ├── models
    └── ports
```

**Description**

- cmd: runs the application, main file
- docs: swagger auto-generated documentation
- env: all env vars setup happens here, then they're injected through main
- internal
  - adapters
    - api: internal api that serves database to rest so that it depends only on contracts and not on implementations
    - database: driven component responsible for contact with DB (Neo4j)
    - rest: entrypoint for requests and responsible for responses, it is the driver component of the application
  - debug: helper functions for printing to console when "dev" mode is set
  - models: request and response models of the application
  - ports: ports for all components (defined as interfaces)

## Contributing

**Code Quality**

- Dependencies are injected in `cmd/main.go`. Please avoid importing components directly. When dependency is needed, make sure it happens over abstractions rather than implementations so that the components remain relatively uncoupled.
- Do not read environment variables directly inside `internal`. This should be done in the `env` package and injected in `cmd/main.go`. One file to rule them all is as good as it gets.

**Average CCN (Cyclomatic Complexity Number)**: 2.3

CCN reference:
- 1-4: good
- 5-7: ok
- 8-10: bad
- 11-11+: really bad

To generate the complete report, install [lizard](https://github.com/terryyin/lizard). Then run `lizard -l go --CCN 4 --exclude ./docs .`. 
Here we exclude `./docs` because it is auto-generated by swagger.

**Commit messages**

```sh
git commit -m "feat: add magical unicorn"
               ^--^  ^-----------------^
               |     |
               |     +-> Summary in present tense, lowercase.
               |
               +-------> Type: arch, build, chore, doc, feat, fix, refactor, style or test.
```

Types:

- arch: (project architecture, organization changes)

- build: (dockerfiles, dependencies, etc)

- chore: (ci/cd scripts; no production code change)

- doc: (documentation)

- feat: (new feature)

- fix: (bug fix)

- refactor: (refactoring production code, eg. renaming variables)

- style: (commenting, formatting, etc; no production code change)

- test: (missing tests, refactoring tests; no production code change)