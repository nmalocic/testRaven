package main

import (
	"fmt"
	"github.com/ravendb/ravendb-go-client"
)

var (
	dbName        = "testdapr"
	serverNodeURL = "http://127.0.0.1:8080"
)

type Item struct {
	ID    string
	Value string
}

const (
	key       string = "items/1"
	secondKey string = "items/2"
)

func main() {
	//testFirstWrite()
	//testDeleteWithEtag()
	transaction()
}

func testFirstWrite() {
	store, session, err := openSession(dbName)
	if err != nil {
		panic("error")
	}
	defer store.Close()
	defer session.Close()

	item := &Item{ID: key, Value: "original"}
	err = session.Store(item)
	if err != nil {
		panic(err)
	}

	err = session.SaveChanges()
	if err != nil {
		panic("error")
	}
	updateItem := &Item{ID: key, Value: "updated"}
	store1, session1, err1 := openSession(dbName)
	if err1 != nil {
		panic("error")
	}
	defer store1.Close()
	defer session1.Close()

	err = session1.StoreWithChangeVectorAndID(updateItem, "Non-ExistingKey", key)
	if err != nil {
		panic("error")
	}
	err = session1.SaveChanges()
	if err != nil {
		panic(err)
	}
	panic("first write: this should fail")
}

func testDeleteWithEtag() {
	store, session, err := openSession(dbName)
	if err != nil {
		panic("error")
	}
	defer store.Close()
	defer session.Close()

	item := &Item{ID: key, Value: "original"}
	item2 := &Item{ID: secondKey, Value: "updated"}
	err = session.Store(item)
	if err != nil {
		panic("error")
	}
	err = session.Store(item2)
	if err != nil {
		panic("error")
	}
	err = session.SaveChanges()
	if err != nil {
		panic("error")
	}
	store1, session1, err1 := openSession(dbName)
	if err1 != nil {
		panic("error")
	}
	defer store1.Close()
	defer session1.Close()
	err = session1.DeleteByID(key, "NotCorrectChangeVectorForSure#$")
	if err != nil {
		panic("this is ok?")
	}
	err = session.SaveChanges()
	if err != nil {
		panic("this is ok also?")
	}

	panic("delete with etag: this should fail")
}

func transaction() {
	store, session, err := openSession(dbName)
	if err != nil {
		panic("error")
	}
	defer store.Close()
	defer session.Close()
	item1 := &Item{ID: key, Value: "value1"}

	item2 := &Item{ID: secondKey, Value: "value2"}

	err = session.Store(item1)
	if err != nil {
		panic(err)
	}

	err = session.Store(item2)
	if err != nil {
		panic(err)
	}

	err = session.DeleteByID(key, "")
	if err != nil {
		panic(err)
	}

	err = session.SaveChanges()
	if err != nil {
		panic(err)
	}

	fmt.Println("finished")
}

func getDocumentStore(databaseName string) (*ravendb.DocumentStore, error) {
	serverNodes := []string{serverNodeURL}
	store := ravendb.NewDocumentStore(serverNodes, databaseName)
	if err := store.Initialize(); err != nil {
		return nil, err
	}
	return store, nil
}

func openSession(databaseName string) (*ravendb.DocumentStore, *ravendb.DocumentSession, error) {
	store, err := getDocumentStore(dbName)
	if err != nil {
		return nil, nil, fmt.Errorf("getDocumentStore() failed with %s", err)
	}

	session, err := store.OpenSession("")
	if err != nil {
		return nil, nil, fmt.Errorf("store.OpenSession() failed with %s", err)
	}
	return store, session, nil
}
