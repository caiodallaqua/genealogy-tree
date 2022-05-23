package database

import (
	"encoding/json"
	"genealogy-tree/internal/models"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// -------------------------------------     GET     -------------------------------------

// Test if db is on and if apoc is properly installed
func (dbAdapter Adapter) GetStatus() error {
	_, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(`CALL apoc.create.node(["test"], {});`, nil)
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
		log.Println("deu ruim")
		return err
	}

	return nil
}

func (dbAdapter Adapter) GetPerson(request models.GetPerson) {
	values, err := requestToMap(request)
	if err != nil {
		return
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person) 
				WHERE id(n) = $id 
				
				RETURN {
					name: n.name, 
					birth: n.birth
				}
				`,
				values)
			if err != nil {
				return nil, err
			}

			var response []interface{}
			for result.Next() {
				response = append(response, result.Record().Values...)
			}

			if len(response) == 0 {
				return nil, err
			}

			return response, nil
		},
	)

	if err != nil {
		log.Printf("deu ruim: %v", err)
		return
	}

	log.Println(res)
}

func (dbAdapter Adapter) GetAscendants(request models.GetAscendants) {
	values, err := requestToMap(request)
	if err != nil {
		return
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
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
						ascendants: COLLECT(DISTINCT 
								{id: id(person), name: person.name, birth: person.birth}
						), 
						relationships: COLLECT(DISTINCT 
								{parent: id(startNode(relationship)), child: id(endNode(relationship))}
						)
				}
				`,
				values)
			if err != nil {
				return nil, err
			}

			var response []interface{}
			for result.Next() {
				response = append(response, result.Record().Values...)
			}

			if len(response) == 0 {
				return nil, err
			}

			return response, nil
		},
	)

	if err != nil {
		log.Printf("deu ruim: %v", err)
		return
	}

	log.Println(res)
}

// -------------------------------------     POST     -------------------------------------

func (dbAdapter Adapter) PostPerson(request models.PostPerson) {
	values, err := requestToMap(request)
	if err != nil {
		return
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				`
				MERGE (p:person {name: $name, birth: $birth}) 
				RETURN id(p)
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
		log.Println("deu ruim")
		return
	}

	log.Println(res)
}

func (dbAdapter Adapter) PostParentRelationship(request models.PostParentRelationship) {
	values, err := requestToMap(request)
	if err != nil {
		return
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person) WHERE id(n) = $parent
				MATCH (m:person) WHERE id(m) = $children
				MERGE (n)-[:PARENT_OF]->(m)
				MERGE (m)-[:CHILD_OF]->(n)
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
		log.Printf("deu ruim: %v", err)
		return
	}

	log.Println(res)
}

// -------------------------------------     DELETE     -------------------------------------

func (dbAdapter Adapter) DelPerson(request models.DelParentRelationship) {
	values, err := requestToMap(request)
	if err != nil {
		return
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person) WHERE id(n) = $id
				MATCH (n)-[parent_edge:PARENT_OF]->()
				MATCH ()-[child_edge:CHILD_OF]->(n)

				DELETE n, child_edge, parent_edge
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
		log.Printf("deu ruim: %v", err)
		return
	}

	log.Println(res)
}

func (dbAdapter Adapter) DelParentRelationship(request models.DelParentRelationship) {
	values, err := requestToMap(request)
	if err != nil {
		return
	}

	res, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				`
				MATCH (n:person)-[child_edge:CHILD_OF]->(m:person) 
				WHERE id(n) = $children AND id(m) = $parent
				
				MATCH (m)-[parent_edge:PARENT_OF]->(n)
				
				DELETE child_edge, parent_edge
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
		log.Printf("deu ruim: %v", err)
		return
	}

	log.Println(res)
}

func requestToMap(req models.Request) (map[string]interface{}, error) {
	var values map[string]interface{}
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
