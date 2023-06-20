package cprotect

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHddId(t *testing.T) {
	id, err := getDiskDriveId(false)
	log.Printf("hdd serial number: %s", id)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}

func TestMotherboardId(t *testing.T) {
	id, err := getMotherboardId(false)
	log.Printf("mother board serial number: %s", id)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}
