package database

import (
  "fmt"
  "os"
  "log"
  "github.com/jedib0t/go-pretty/v6/table"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

// Create database file if doesn't exist
func CreateDatabase(db string) {
  _, err := os.Stat(db)
  if os.IsNotExist(err) {
    fmt.Println("Database doesn't exist. Creating...")
    file, err := os.Create(db)
    if err != nil {
        log.Fatal(err)
    }
    file.Close()
  }
}

// Creates table in database if doesn't exist
func CreateTable(db *sql.DB) {
  shotsTable := `CREATE TABLE IF NOT EXISTS shots(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "Date" TEXT,
        "Gun" TEXT,
        "Shots" TEXT,
        "Notes" INT);`
  query, err := db.Prepare(shotsTable)
  if err != nil {
      log.Fatal(err)
  }
  query.Exec()
}

// Adds a record to database
func AddRecord(db *sql.DB, Date string, Gun string, Shots int, Notes string) {
  records := "INSERT INTO shots(Date, Gun, Shots, Notes) VALUES (?, ?, ?, ?)"
  query, err := db.Prepare(records)
  if err != nil {
    log.Fatal(err)
  }
  _, err = query.Exec(Date, Gun, Shots, Notes)
  if err != nil {
    log.Fatal(err)
  }
}

// Deletes a record from database
func DeleteRecord(db *sql.DB, str string) {
  query, err := db.Prepare(str)
  if err != nil {
    log.Fatal(err)
  }
  _, err = query.Exec()
  if err != nil {
    log.Fatal(err)
  }
}

// Fetches a record from database
// Uses github.com/jedib0t/go-pretty/v6/table tables to print it out pretty.
func FetchRecord(db *sql.DB, record *sql.Rows, err error) {
  if err != nil {
    log.Fatal(err)
  }
  defer record.Close()

  //totalSlice := []int{}
  var (
    totalSlice []int
    total int
    id int
    Date string
    Gun string
    Shots int
    Notes string
    recordCount int
  )

  t := table.NewWriter()
  t.SetOutputMirror(os.Stdout)

  t.AppendHeader(table.Row{"id", "Date", "Gun", "Shots", "Notes"})

  // Loop over records and add them to the table
  for record.Next() {
    recordCount++
    record.Scan(&id, &Date, &Gun, &Shots, &Notes)
    totalSlice = append(totalSlice, Shots)
    t.AppendRows([]table.Row{{id, Date, Gun, Shots, Notes}})
  }

  // adds up the slice to tell you the total number of shots 
  for _, num := range totalSlice {
    total += num
  }

  t.AppendFooter(table.Row{"", "", "", "Total:", total})
  t.SetStyle(table.StyleLight)

  // To have a line separate all rows, uncomment next line.
  //t.Style().Options.SeparateRows = true

  if recordCount > 0 {
    t.Render()
  } else {
    fmt.Println("\nIt appears there are no entries matching that query. Please try again.\n")
  }
}

