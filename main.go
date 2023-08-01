package main

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	NumberOfPools        = 200
	NumberOfDependencies = 100
)

func main() {
	ctx := context.Background()
	// create a new Neo4j driver
	driver, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "test1234", ""))
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer func(driver neo4j.DriverWithContext, ctx context.Context) {
		err := driver.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(driver, ctx)

	pools := make([]*Pool, 0, NumberOfPools)
	// Generate NumberOfPools pools
	for i := 0; i < NumberOfPools; i++ {
		// create a new pool
		pool, err := createPool(ctx, driver, GenerateAddress(), GeneratePoolType())
		if err != nil {
			log.Fatalf("Failed to create pool: %v", err)
		}
		log.Printf("Created pool: %+v\n", pool)

		pools = append(pools, pool)
	}

	dependencies := make([]*Dependency, 0, NumberOfDependencies)
	// Generate NumberOfDependencies dependencies
	for i := 0; i < NumberOfDependencies; i++ {
		// create a new dependency
		dependency, err := createDependency(ctx, driver, GenerateAddress())
		if err != nil {
			log.Fatalf("Failed to create dependency: %v", err)
		}
		log.Printf("Created dependency: %+v\n", dependency)

		dependencies = append(dependencies, dependency)
	}

	for _, p := range pools {
		// choose a random dependency
		dependency := pickRandomDependency(dependencies)

		// create a relationship between the pool and the dependency
		err = createRelationship(ctx, driver, p.ID, dependency.ID)
		if err != nil {
			log.Fatalf("Failed to create relationship: %v", err)
		}
	}

	//// get the person by name
	//personByName, err := getPersonByName(ctx, driver, "Alice")
	//if err != nil {
	//	log.Fatalf("Failed to get person by name: %v", err)
	//}
	//log.Printf("Found person by name: %+v\n", personByName)
	//
	//// get the person by ID
	//personByID, err := getPersonByID(ctx, driver, personByName.ID)
	//if err != nil {
	//	log.Fatalf("Failed to get person by ID: %v", err)
	//}
	//log.Printf("Found person by ID: %+v\n", personByID)
	//
	//// update the person's age
	//updatedPerson, err := updatePersonAge(ctx, driver, personByName.ID, 35)
	//if err != nil {
	//	log.Fatalf("Failed to update person's age: %v", err)
	//}
	//log.Printf("Updated person: %+v\n", updatedPerson)

	// delete the person
	//err = deletePerson(ctx, driver, personByName.ID)
	//if err != nil {
	//	log.Fatalf("Failed to delete person: %v", err)
	//}
	//log.Printf("Deleted person with ID %d\n", personByName.ID)
}
