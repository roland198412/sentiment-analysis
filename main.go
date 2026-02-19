package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jonreiter/govader"
)

var db *sql.DB

type Data struct {
	ID    int64
	Value string
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DB_HOST")
	cfg.DBName = os.Getenv("DB_NAME")

	tableName := os.Getenv("TABLE_NAME")
	tableColumnName := os.Getenv("TABLE_COLUMN_NAME")

	// Option to delete negative rows
	enableDeleteNegativeRow := os.Getenv("ENABLE_DELETE_NEGATIVE_ROW") // "true" to delete

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	data, err := getColumnData(tableName, tableColumnName)
	if err != nil {
		log.Fatal(err)
	}

	// Creating Sentiment Analysis Model
	analyzer := govader.NewSentimentIntensityAnalyzer()

	for _, row := range data {

		sentiment := analyzer.PolarityScores(row.Value)
		//fmt.Println("Compound score:", sentiment.Compound)
		//fmt.Println("Positive score:", sentiment.Positive)
		//fmt.Println("Neutral score:", sentiment.Neutral)
		if sentiment.Negative > 0.2 && sentiment.Compound < -0.05 {
			fmt.Printf("NEGATIVE: %s: %f\n", row.Value, sentiment.Negative)
			if enableDeleteNegativeRow == "true" {
				if err := deleteRowByID(tableName, row.ID); err != nil {
					log.Printf("Failed to delete row ID %d: %v", row.ID, err)
				} else {
					log.Printf("Deleted row ID %d", row.ID)
				}
			}
		}
	}

}

func getColumnData(tableName, columnName string) ([]Data, error) {
	var data []Data

	rows, err := db.Query("SELECT id, " + columnName + " FROM " + tableName)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var record Data
		if err := rows.Scan(&record.ID, &record.Value); err != nil {
			return data, err
		}

		data = append(data, record)
	}

	if err := rows.Err(); err != nil {
		return data, err
	}

	return data, nil
}

func deleteRowByID(tableName string, id int64) error {
	_, err := db.Exec("DELETE FROM "+tableName+" WHERE id = ?", id)
	return err
}
