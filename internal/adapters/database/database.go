package database

import (
	"genealogy-tree/internal/debug"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Adapter struct {
	addr string
	user string
	pwd  string

	// Both kept in the adapter to be deferred later on
	driver  neo4j.Driver
	session neo4j.Session
}

func NewAdapter(addr, user, pwd string) (Adapter, error) {
	driver, err := neo4j.NewDriver(addr, neo4j.BasicAuth(user, pwd, ""))
	if err != nil {
		debug.ShowErr("NewAdapter", "Failed to connect to DB", err)
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
