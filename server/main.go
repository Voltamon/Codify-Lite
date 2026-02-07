package main

import (
	"codify-lite/backend"
	"codify-lite/backend/logger"
	"log"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	// Use local database directory
	// path relative to the server executable or working directory
	// In dev (wails dev), we run from server/ so ../database/pocketbase is correct
	dataDir := "../database/pocketbase"

	// Ensure data directory exists
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	pb := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: dataDir,
	})

	// Initialize custom app logic
	appLogger := logger.InitC()

	// Bootstrap PocketBase to initialize DB
	if err := pb.Bootstrap(); err != nil {
		log.Fatal("Failed to bootstrap PocketBase:", err)
	}

	if pb.DB() == nil {
		log.Fatal("pb.DB() is nil")
	}
	dbConn := pb.DB().DB()
	if dbConn == nil {
		log.Fatal("dbConn is nil")
	}

	// Create user_preferences table
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS user_preferences (
		key   TEXT NOT NULL PRIMARY KEY,
		value TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal("Failed to create user_preferences table:", err)
	}

	myApp := backend.NewApp(appLogger, dbConn)

	// Register custom routes/hooks BEFORE firing serving
	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Example: add a custom internal route
		e.Router.GET("/api/hello", func(c echo.Context) error {
			return c.String(200, "Hello from PocketBase!")
		})

		return nil
	})

	// Run PocketBase in a separate goroutine so it doesn't block Wails
	go func() {
		// Explicitly set the args to avoid picking up Wails flags
		pb.RootCmd.SetArgs([]string{"serve", "--http=127.0.0.1:8090"})

		if err := pb.Start(); err != nil {
			log.Println("PocketBase Start failed:", err)
		}
	}()

	// Wails Application
	err = wails.Run(&options.App{
		Title:            "Codify Lite",
		Width:            1024,
		Height:           768,
		AssetServer:      &assetserver.Options{},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        myApp.Startup,
		Bind: []interface{}{
			myApp,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
