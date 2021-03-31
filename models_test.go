package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/asdine/storm"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	tmpfile, err := ioutil.TempFile("", "shorturl")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	db, err = storm.Open(tmpfile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	os.Exit(m.Run())
}

func TestZeroURL(t *testing.T) {
	assert := assert.New(t)

	u := URL{}
	assert.Equal(u.ID, "")
	assert.Equal(u.URL, "")
	assert.Equal(u.Name, "")
	assert.Equal(u.CreatedAt, time.Time{})
	assert.Equal(u.UpdatedAt, time.Time{})
}

func TestNewURL(t *testing.T) {
	assert := assert.New(t)

	u, err := NewURL("https://www.google.com", 5)
	assert.Nil(err, nil)

	assert.NotEqual(u.ID, "")
	assert.Equal(u.URL, "https://www.google.com")
	assert.Equal(u.Name, "")
}
