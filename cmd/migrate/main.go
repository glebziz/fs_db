package main

import (
	"bufio"
	"context"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

//go:embed migrations/*.sql
var migrations embed.FS

func main() {
	ctx := context.Background()

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <sqlite_db> <badger_db>\n", os.Args[0])
		os.Exit(1)
	}

	sqlitePath := os.Args[1]
	badgerPath := os.Args[2]

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Running a migration may overwrite the latest version of the file contents if the badger database was not empty.")

loop:
	for {
		var choice string
		fmt.Print("Are you sure you want to continue? [y/N] ")
		choice, err := scanner.ReadString('\n')
		if err != nil {
			slog.Error("Read string",
				slog.Any("err", err),
			)
			os.Exit(1)
		}

		choice = strings.TrimSpace(choice)
		switch strings.ToLower(choice) {
		case "y", "yes":
			break loop
		case "n", "no", "":
			fmt.Println("Aborting...")
			os.Exit(0)
		}
	}

	db, err := setupSqlite(ctx, sqlitePath)
	if err != nil {
		fmt.Printf("Setup sqlite error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	r, err := setupRepos(badgerPath)
	if err != nil {
		fmt.Printf("Setup repos error: %v\n", err)
		os.Exit(1)
	}
	defer r.p.Close()

	err = migrateAll(ctx, db, r)
	if err != nil {
		fmt.Printf("Migrate all error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Migration complete.")
}
