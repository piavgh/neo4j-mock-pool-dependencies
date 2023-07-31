package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func createPool(ctx context.Context, driver neo4j.DriverWithContext, address string, poolType string) (*Contract, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MERGE (p:Contract {address: $address, type: $type}) RETURN id(p)",
		map[string]interface{}{"address": address, "type": poolType},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, err
	}

	id, ok := record.Values[0].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}

	return &Contract{ID: id, Address: address}, nil
}

func createDependency(ctx context.Context, driver neo4j.DriverWithContext, address string) (*Contract, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MERGE (d:Contract {address: $address}) RETURN id(d)",
		map[string]interface{}{"address": address},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, err
	}

	id, ok := record.Values[0].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}

	return &Contract{ID: id, Address: address}, nil
}

func createRelationship(ctx context.Context, driver neo4j.DriverWithContext, sourceId int64, targetId int64) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MATCH\n  (a:Contract),\n  (b:Contract)\nWHERE id(a) = $sourceId AND id(b) = $targetId\nMERGE (a)-[r:DEPEND_ON {name: a.address + '->' + b.address}]->(b)\nRETURN type(r), r.name",
		map[string]interface{}{"sourceId": sourceId, "targetId": targetId},
	)
	if err != nil {
		return err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return err
	}

	relationshipType, ok := record.Values[0].(string)
	if !ok {
		return fmt.Errorf("invalid relationship type")
	}

	relationshipName, ok := record.Values[1].(string)
	if !ok {
		return fmt.Errorf("invalid relationship name type")
	}

	fmt.Println("relationshipType = ", relationshipType)
	fmt.Println("relationshipName = ", relationshipName)

	return nil
}

func getPoolByAddress(ctx context.Context, driver neo4j.DriverWithContext, address string) (*Contract, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MATCH (p:Contract) WHERE p.address = $address RETURN id(p) LIMIT 1",
		map[string]interface{}{"address": address},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, err
	}

	id, ok := record.Values[0].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}

	return &Contract{ID: id, Address: address}, nil
}

func getPoolByID(ctx context.Context, driver neo4j.DriverWithContext, id int64) (*Contract, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MATCH (p:Contract) WHERE id(p) = $id RETURN p.address",
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, err
	}

	address, ok := record.Values[0].(string)
	if !ok {
		return nil, fmt.Errorf("invalid address type")
	}

	return &Contract{ID: id, Address: address}, nil
}

func updatePoolType(ctx context.Context, driver neo4j.DriverWithContext, id int64, poolType int) (*Contract, error) {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	result, err := session.Run(
		ctx,
		"MATCH (p:Contract) WHERE id(p) = $id SET p.type = $type RETURN p.address, p.type",
		map[string]interface{}{"id": id, "type": poolType},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single(ctx)
	if err != nil {
		return nil, err
	}

	address, ok := record.Values[0].(string)
	if !ok {
		return nil, fmt.Errorf("invalid address type")
	}

	newType, ok := record.Values[1].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid poolType type")
	}

	return &Contract{ID: id, Address: address, Type: strconv.FormatInt(newType, 10)}, nil
}

func deletePool(ctx context.Context, driver neo4j.DriverWithContext, id int64) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	_, err := session.Run(
		ctx,
		"MATCH (p:Contract) WHERE id(p) = $id DELETE p",
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return err
	}

	return nil
}
