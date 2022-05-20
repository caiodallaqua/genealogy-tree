package database

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Adapter struct {
	addr    string
	user    string
	pwd     string
	driver  neo4j.Driver
	session neo4j.Session
}

func NewAdapter(addr, user, pwd string) (Adapter, error) {
	driver, err := neo4j.NewDriver(addr, neo4j.BasicAuth(user, pwd, ""))
	if err != nil {
		return Adapter{}, err
	}
	//defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	//defer session.Close()

	var adapter = Adapter{
		addr:    addr,
		user:    user,
		pwd:     pwd,
		driver:  driver,
		session: session,
	}

	return adapter, nil
}

func (dbAdapter Adapter) Test() {
	greeting, err := dbAdapter.session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			result, err := transaction.Run(
				"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
				map[string]interface{}{"message": "hello, world"})
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

	log.Println(greeting.(string))
}
