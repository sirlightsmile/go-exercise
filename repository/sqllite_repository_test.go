package repository

import (
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var dbRepo *SqlDB

func TestMain(m *testing.M) {
	absPath, _ := filepath.Abs("../data/th_address.db")
	repo, err := ConnectSqlDB(absPath)
	if err != nil {
		panic(err)
	}
	dbRepo = repo
	runTests := m.Run()
	os.Exit(runTests)
}

func TestQuery(t *testing.T) {
	t.Run("Test Query", func(t *testing.T) {

		queryFull := `SELECT COUNT(*) FROM provinces`

		rows, err := dbRepo.Query(queryFull)
		if err != nil {
			t.Errorf("Failed, error : %s", err)
		}

		var count int
		expectedCount := 77
		for rows.Next() {
			if err := rows.Scan(&count); err != nil {
				t.Errorf("Failed, error : %s", err)
			}
		}

		if count != expectedCount {
			t.Errorf("Failed, expected : %v reality : %v", expectedCount, count)
		}

		query := `SELECT COUNT(*) FROM provinces WHERE UPPER(province_name) = UPPER(?)`
		provinceName := "Bangkok"

		rows, err = dbRepo.Query(query, provinceName)
		if err != nil {
			t.Errorf("Failed, error : %s", err)
		}

		var count2 int
		expectedCount2 := 1

		for rows.Next() {
			if err := rows.Scan(&count2); err != nil {
				t.Errorf("Failed, error : %s", err)
			}
		}

		if count2 != expectedCount2 {
			t.Errorf("Failed, expected : %v reality : %v", expectedCount2, count2)
		}
	})
}

func TestQueryRow(t *testing.T) {
	t.Run("Test Query Row", func(t *testing.T) {
		queryFull := `SELECT province_name_eng FROM provinces WHERE UPPER(province_name_eng) = UPPER("Bangkok")`

		query := `SELECT province_name_eng FROM provinces WHERE UPPER(province_name_eng) = UPPER(?)`

		provinceName := "Bangkok"

		row := dbRepo.QueryRow(queryFull)

		var result string

		if err := row.Scan(&result); err != nil {
			t.Errorf("Failed, error : %s", err)
		}

		if result != provinceName {
			t.Errorf("Failed, expected : %s reality : %s", provinceName, result)
		}

		row = dbRepo.QueryRow(query, provinceName)

		if err := row.Scan(&result); err != nil {
			t.Errorf("Failed, error : %s", err)
		}

		if result != provinceName {
			t.Errorf("Failed, expected : %s reality : %s", provinceName, result)
		}
	})
}
