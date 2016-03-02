package main

import "fmt"

var currentId int

var databaseConnections DatabaseConnections

// Give us some seed data
func init() {
	RepoCreateDatabaseConnection(DatabaseConnection{Name: "Write presentation"})
	RepoCreateDatabaseConnection(DatabaseConnection{Name: "Host meetup"})
}

func RepoFindDatabaseConnection(id int) DatabaseConnection {
	for _, t := range databaseConnections {
		if t.Id == id {
			return t
		}
	}
	// return empty DatabaseConnection if not found
	return DatabaseConnection{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateDatabaseConnection(t DatabaseConnection) DatabaseConnection {
	currentId += 1
	t.Id = currentId
	databaseConnections = append(databaseConnections, t)
	return t
}

func RepoDestroyDatabaseConnection(id int) error {
	for i, t := range databaseConnections {
		if t.Id == id {
			databaseConnections = append(databaseConnections[:i], databaseConnections[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find DatabaseConnection with id of %d to delete", id)
}
