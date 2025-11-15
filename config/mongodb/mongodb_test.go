package mongodb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestOperationConstants tests MongoDB operation constants.
func TestOperationConstants(t *testing.T) {
	assert.Equal(t, Operation("insert"), OpInsert)
	assert.Equal(t, Operation("find"), OpFind)
	assert.Equal(t, Operation("update"), OpUpdate)
	assert.Equal(t, Operation("delete"), OpDelete)
	assert.Equal(t, Operation("count"), OpCount)
	assert.Equal(t, Operation("aggregate"), OpAggregate)
}

// TestConfigInsert tests insert operation configuration.
func TestConfigInsert(t *testing.T) {
	doc := map[string]interface{}{
		"name": "test",
		"age":  25,
	}

	config := &Config{
		URI:        "mongodb://localhost:27017",
		Database:   "testdb",
		Collection: "users",
		Operation:  OpInsert,
		Count:      100,
		Document:   doc,
	}

	assert.Equal(t, "mongodb://localhost:27017", config.URI)
	assert.Equal(t, "testdb", config.Database)
	assert.Equal(t, "users", config.Collection)
	assert.Equal(t, OpInsert, config.Operation)
	assert.Equal(t, 100, config.Count)
	assert.Equal(t, "test", config.Document["name"])
	assert.Equal(t, 25, config.Document["age"])
}

// TestConfigFind tests find operation configuration.
func TestConfigFind(t *testing.T) {
	filter := map[string]interface{}{
		"status": "active",
	}

	config := &Config{
		URI:        "mongodb://localhost:27017",
		Database:   "testdb",
		Collection: "users",
		Operation:  OpFind,
		Count:      50,
		Filter:     filter,
	}

	assert.Equal(t, OpFind, config.Operation)
	assert.Equal(t, "active", config.Filter["status"])
}

// TestConfigUpdate tests update operation configuration.
func TestConfigUpdate(t *testing.T) {
	filter := map[string]interface{}{
		"_id": 123,
	}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"status": "updated",
		},
	}

	config := &Config{
		URI:        "mongodb://localhost:27017",
		Database:   "testdb",
		Collection: "users",
		Operation:  OpUpdate,
		Count:      10,
		Filter:     filter,
		Update:     update,
	}

	assert.Equal(t, OpUpdate, config.Operation)
	assert.Equal(t, 123, config.Filter["_id"])
	assert.NotNil(t, config.Update["$set"])
}

// TestConfigDelete tests delete operation configuration.
func TestConfigDelete(t *testing.T) {
	filter := map[string]interface{}{
		"status": "deleted",
	}

	config := &Config{
		URI:        "mongodb://localhost:27017",
		Database:   "testdb",
		Collection: "users",
		Operation:  OpDelete,
		Count:      5,
		Filter:     filter,
	}

	assert.Equal(t, OpDelete, config.Operation)
	assert.Equal(t, "deleted", config.Filter["status"])
}

// TestConfigWithTimeouts tests configuration with timeouts.
func TestConfigWithTimeouts(t *testing.T) {
	timeout := 30 * time.Second
	config := &Config{
		URI:            "mongodb://localhost:27017",
		Database:       "testdb",
		Collection:     "users",
		Operation:      OpFind,
		Count:          1,
		ConnectTimeout: &timeout,
	}

	assert.NotNil(t, config.ConnectTimeout)
	assert.Equal(t, 30*time.Second, *config.ConnectTimeout)
}

// TestConfigWithPoolSize tests configuration with pool size settings.
func TestConfigWithPoolSize(t *testing.T) {
	maxPool := uint64(100)
	minPool := uint64(10)

	config := &Config{
		URI:         "mongodb://localhost:27017",
		Database:    "testdb",
		Collection:  "users",
		Operation:   OpFind,
		Count:       1,
		MaxPoolSize: &maxPool,
		MinPoolSize: &minPool,
	}

	assert.NotNil(t, config.MaxPoolSize)
	assert.NotNil(t, config.MinPoolSize)
	assert.Equal(t, uint64(100), *config.MaxPoolSize)
	assert.Equal(t, uint64(10), *config.MinPoolSize)
}

// TestConfigCount tests count operation configuration.
func TestConfigCount(t *testing.T) {
	filter := map[string]interface{}{
		"status": "active",
	}

	config := &Config{
		URI:        "mongodb://localhost:27017",
		Database:   "testdb",
		Collection: "users",
		Operation:  OpCount,
		Count:      1,
		Filter:     filter,
	}

	assert.Equal(t, OpCount, config.Operation)
	assert.Equal(t, "active", config.Filter["status"])
}

// TestConfigAggregate tests aggregate operation configuration.
func TestConfigAggregate(t *testing.T) {
	config := &Config{
		URI:        "mongodb://localhost:27017",
		Database:   "testdb",
		Collection: "users",
		Operation:  OpAggregate,
		Count:      1,
	}

	assert.Equal(t, OpAggregate, config.Operation)
}

// TestConfigEmpty tests empty configuration.
func TestConfigEmpty(t *testing.T) {
	config := &Config{}

	assert.Empty(t, config.URI)
	assert.Empty(t, config.Database)
	assert.Empty(t, config.Collection)
	assert.Empty(t, config.Operation)
	assert.Zero(t, config.Count)
	assert.Nil(t, config.Document)
	assert.Nil(t, config.Filter)
	assert.Nil(t, config.Update)
}
