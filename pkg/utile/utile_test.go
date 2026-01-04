package utile

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStrPtr(t *testing.T) {
	input := "test"
	result := StrPtr(input)

	assert.NotNil(t, result)
	assert.Equal(t, input, *result)
}

func TestFileExists_Basic(t *testing.T) {
	nonExistent := "/tmp/this-file-should-not-exist-12345"
	res := FileExists(nonExistent)
	assert.False(t, res)

	tmpfile, err := os.CreateTemp("", "testfile")
	assert.Nil(t, err)
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	res = FileExists(tmpfile.Name())
	assert.True(t, res)
}

func TestTimePtr_Basic(t *testing.T) {
	timeNow := time.Now()

	testTime := timeNow.Add(5 * time.Hour)
	testTimePtr := TimePtr(testTime)

	assert.Equal(t, testTime, *testTimePtr)
}


