package addressAPI

import (
	"os"
	"path/filepath"
	"smile/address"
	"smile/repository"
	"testing"
)

var testAddressManager *address.AddressManager

func TestMain(m *testing.M) {
	absPath, _ := filepath.Abs("../../data/th_address.db")
	dbRepo, err := repository.ConnectSqlDB(absPath)
	if err != nil {
		panic(err)
	}
	testAddressManager = address.Init(dbRepo)
	runTests := m.Run()
	os.Exit(runTests)
}
