package database

import (
	"encoding/json"
	"genealogy-tree/internal/debug"
	"genealogy-tree/internal/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// -------------------------------------     GET     -------------------------------------

// Test whether DB is on and if apoc is properly installed
func (dbAdapter Adapter) GetStatus() (any, error) {
	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
			CALL apoc.create.node(["test"], {}) YIELD node
			RETURN {data: "ok"}
			`, nil)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("GetStatus", "Failed to get DB status", err)
		return nil, err
	}

	return res, nil
}

func (dbAdapter Adapter) GetPerson(request models.GetPerson) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("GetPerson", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person) 
				WHERE id(n) = $id 
				
				RETURN {
					data: {
						name: n.name, 
						birth: n.birth
					}
				}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("GetPerson", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

func (dbAdapter Adapter) GetAscendants(request models.GetAscendants) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("GetAscendants", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (origin:person) 
				WHERE id(origin) = $id
				
				MATCH (root:person) 
				WHERE ((origin)-[:CHILD_OF *..]->(root)) 
				AND NOT ((root)-[:CHILD_OF]->())
				
				MATCH levels = shortestPath((root)-[:PARENT_OF *..]->(origin))
				
				CALL apoc.path.subgraphAll(root, {
						relationshipFilter: 'PARENT_OF>',
						labelFilter: "+person",
						minLevel: 0,
						maxLevel: length(levels)
				}) YIELD nodes, relationships
				
				UNWIND relationships AS relationship
				WITH nodes, relationship WHERE type(relationship) = "PARENT_OF"
				
				UNWIND nodes AS person
				WITH relationship, person ORDER BY id(person) ASC
				
				RETURN {
					data: {
						ascendants: COLLECT(DISTINCT 
								{id: id(person), name: person.name, birth: person.birth}
						), 
						relationships: COLLECT(DISTINCT 
								{parent: id(startNode(relationship)), child: id(endNode(relationship))}
						)
					}
				}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("GetAscendantsAndDescendants", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

func (dbAdapter Adapter) GetAscendantsAndDescendants(request models.GetAscendantsAndDescendants) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("GetAscendantsAndDescendants", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (origin:person) 
				WHERE id(origin) = $id
				
				MATCH (root:person)
				WHERE ((origin)-[:CHILD_OF *..]->(root))
				AND NOT ((root)-[:CHILD_OF]->())
				
				CALL apoc.path.subgraphAll(root, {
						relationshipFilter: 'PARENT_OF>',
						labelFilter: "+person",
						minLevel: 0
				}) YIELD nodes, relationships
				
				UNWIND relationships AS relationship
				WITH nodes, relationship WHERE type(relationship) = "PARENT_OF"
				
				UNWIND nodes AS person
				WITH relationship, person ORDER BY id(person) ASC
				
				RETURN {
					data: {
						ascendants_and_descendants: COLLECT(DISTINCT 
								{id: id(person), name: person.name, birth: person.birth}
						), 
						relationships: COLLECT(DISTINCT 
								{parent: id(startNode(relationship)), child: id(endNode(relationship))}
						)
					}
				}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("GetAscendants", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

func (dbAdapter Adapter) GetAscendantsAndChildren(request models.GetAscendantsAndChildren) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("GetAscendantsAndChildren", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (origin:person) 
				WHERE id(origin) = $id
				
				MATCH (root:person) 
				WHERE ((origin)-[:CHILD_OF *..]->(root)) 
				AND NOT ((root)-[:CHILD_OF]->())
				
				MATCH levels = shortestPath((root)-[:PARENT_OF *..]->(origin))
				
				CALL apoc.path.subgraphAll(root, {
								relationshipFilter: 'PARENT_OF>',
								labelFilter: "+person",
								minLevel: 0,
								maxLevel: length(levels)+1
				}) YIELD nodes, relationships
				
				UNWIND relationships AS relationship
				WITH nodes, relationship WHERE type(relationship) = "PARENT_OF"
				
				UNWIND nodes AS person
				WITH relationship, person ORDER BY id(person) ASC
				
				RETURN {
						data: {
								ascendants: COLLECT(DISTINCT 
												{id: id(person), name: person.name, birth: person.birth}
								), 
								relationships: COLLECT(DISTINCT 
												{parent: id(startNode(relationship)), child: id(endNode(relationship))}
								)
						}
				}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("GetAscendantsAndChildren", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

// -------------------------------------     POST     -------------------------------------

func (dbAdapter Adapter) PostPerson(request models.PostPerson) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("PostPerson", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MERGE (p:person {name: $name, birth: $birth}) 
				RETURN {data: {id: id(p)}}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("PostPerson", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

func (dbAdapter Adapter) PostParentRelationship(request models.PostParentRelationship) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("PostParentRelationship", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person) WHERE id(n) = $parent_id
				MATCH (m:person) WHERE id(m) = $child_id
				
				MERGE (n)-[:PARENT_OF]->(m)
				MERGE (m)-[:CHILD_OF]->(n)
				
				RETURN {data: "ok"}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("PostParentRelationship", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

// -------------------------------------     DELETE     -------------------------------------

func (dbAdapter Adapter) DelPerson(request models.DelPerson) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("DelPerson", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person) WHERE id(n) = $id
				MATCH (n)-[parent_edge:PARENT_OF]->()
				MATCH ()-[child_edge:CHILD_OF]->(n)

				DELETE n, child_edge, parent_edge

				RETURN {data: "ok"}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("DelPerson", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

func (dbAdapter Adapter) DelParentRelationship(request models.DelParentRelationship) (any, error) {
	values, err := requestToMap(request)
	if err != nil {
		debug.ShowErr("DelParentRelationship", "Failed to convert request to map", err)
		return nil, err
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (any, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person)-[child_edge:CHILD_OF]->(m:person) 
				WHERE id(n) = $child_id AND id(m) = $parent_id
				
				MATCH (m)-[parent_edge:PARENT_OF]->(n)
				
				DELETE child_edge, parent_edge

				RETURN {data: "ok"}
				`,
				values)
			if err != nil {
				return nil, err
			}

			if result.Next() {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		},
	)

	if err != nil {
		debug.ShowErr("DelParentRelationship", "Failed to execute query", err)
		return nil, err
	}

	return res, nil
}

func requestToMap[R models.Request](req R) (map[string]any, error) {
	var values map[string]any

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &values)
	if err != nil {
		return nil, err
	}

	return values, nil
}
