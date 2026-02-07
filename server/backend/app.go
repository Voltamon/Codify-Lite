package backend

import (
	"context"
	"database/sql"
	"log/slog"

	"codify-lite/backend/db"
	"codify-lite/backend/validator"
)

// App struct
type App struct {
	ctx    context.Context
	logger *slog.Logger
	q      *db.Queries
}

// NewApp creates a new App application struct
func NewApp(logger *slog.Logger, conn db.DBTX) *App {
	return &App{
		logger: logger,
		q:      db.New(conn),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// GreetUser returns the saved name or empty string if not found
func (a *App) GreetUser() (string, error) {
	ctx := a.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	a.logger.Info("Greeting user")
	val, err := a.q.GetUserPreference(ctx, "name")
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		a.logger.Error("failed to get name", "error", err)
		return "", err
	}
	return val, nil
}

// SaveName validates and saves the user name
func (a *App) SaveName(name string) error {
	a.logger.Info("Saving name", "name", name)

	if err := validator.ValidateName(name); err != nil {
		a.logger.Warn("validation failed", "error", err)
		return err
	}

	err := a.q.SetUserPreference(a.ctx, db.SetUserPreferenceParams{
		Key:   "name",
		Value: name,
	})
	if err != nil {
		a.logger.Error("failed to save name", "error", err)
		return err
	}
	return nil
}
