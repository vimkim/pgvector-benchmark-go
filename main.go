package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Command-line flags
	var (
		db    = flag.String("database", "", "Database name")
		dim   = flag.Int("dim", 256, "Vector dimension")
		rows  = flag.Int("rows", 75000, "Number of rows")
		table = flag.String("table", "", "Table name")
		user  = flag.String("u", "", "Database user")
	)
	flag.Parse()

	if *db == "" || *table == "" || *user == "" {
		fmt.Println("Missing required arguments. Usage: --database <db> --table <table> -u <user>")
		os.Exit(1)
	}

	// Generate SQL content for pgvector test.
	// This query:
	// 1. Retrieves a reference vector from the table.
	// 2. Uses EXPLAIN ANALYZE to benchmark a cosine distance calculation.
	sqlFilename := fmt.Sprintf("test_pgvector_%d_%d.sql", *dim, *rows)
	sqlContent := fmt.Sprintf(`
EXPLAIN (ANALYZE, BUFFERS, VERBOSE)
WITH ref AS (
    SELECT vec AS ref_vec FROM %s LIMIT 1
)
    SELECT vec <=> ref_vec AS cosine_distance 
    FROM %s, ref;
`, *table, *table)

	// Write SQL to file
	err := os.WriteFile(sqlFilename, []byte(sqlContent), 0644)
	if err != nil {
		fmt.Printf("Error writing SQL file: %v\n", err)
		os.Exit(1)
	}

	// Build psql command.
	// Note: Ensure that your PostgreSQL environment is configured for password-less authentication or that you have set the appropriate environment variables.
	cmd := exec.Command("psql", "-U", *user, "-d", *db, "-f", sqlFilename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running: psql -U %s -d %s -f %s\n", *user, *db, sqlFilename)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running psql: %v\n", err)
		os.Exit(1)
	}
}
