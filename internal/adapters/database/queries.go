package database

import (
	"encoding/json"
	"genealogy-tree/internal/models"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (dbAdapter Adapter) AddPerson(request models.AddPerson) {
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

func (dbAdapter Adapter) AddParentRelationship(request models.AddParentRelationship) {
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
				WHERE id(origin) = $person
				
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
				
				UNWIND relationships AS edge
				WITH edge WHERE type(edge) = "PARENT_OF"
				
				MATCH (parent:person)-[edge]->(child:person)
				
				RETURN {
						parent: {
								id: id(parent),
								name: parent.name,
								birth: parent.birth
						}, 
						child: {
								id: id(child),
								name: child.name,
								birth: child.birth
						}
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

	if res == nil {
		log.Println("fazer outra query kkkj")
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
