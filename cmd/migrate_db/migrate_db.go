package migrate_db

import (
	"database/sql"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/forbole/bdjuno/v4/database"
	parsecmdtypes "github.com/forbole/juno/v4/cmd/parse/types"
	"github.com/forbole/juno/v4/logging"
	"github.com/forbole/juno/v4/types/config"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/spf13/cobra"
)

var migrations = &migrate.EmbedFileSystemMigrationSource{
	FileSystem: database.Migrations,
	Root:       "schema",
}

func NewMigrateDBCmd(parseCfg *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "migrate-db",
		Short:             "Migrate the database schema",
		PersistentPreRunE: runPersistentPreRuns(parsecmdtypes.ReadConfigPreRunE(parseCfg)),
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "up",
			Short: "migrate db up",
			RunE: func(cmd *cobra.Command, args []string) error {
				context, err := parsecmdtypes.GetParserContext(config.Cfg, parseCfg)
				if err != nil {
					return err
				}
				db := database.Cast(context.Database)
				return migrateUp(db.SQL.DB, context.Logger)
			},
		},
	)

	cmd.AddCommand(
		&cobra.Command{
			Use:   "down",
			Short: "migrate db down",
			RunE: func(cmd *cobra.Command, args []string) error {
				context, err := parsecmdtypes.GetParserContext(config.Cfg, parseCfg)
				if err != nil {
					return err
				}
				db := database.Cast(context.Database)
				return migrateDown(db.SQL.DB, context.Logger)
			},
		},
	)

	return cmd
}

func migrateUp(rawDB *sql.DB, log logging.Logger) error {
	applied, err := migrate.Exec(rawDB, "postgres", migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	log.Info("migrations applied", map[string]interface{}{
		"applied": applied,
	})
	return nil
}

func migrateDown(rawDB *sql.DB, log logging.Logger) error {
	applied, err := migrate.Exec(rawDB, "postgres", migrations, migrate.Down)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	log.Info("migrations applied", map[string]interface{}{
		"applied": applied,
	})
	return nil
}

func runPersistentPreRuns(preRun func(_ *cobra.Command, _ []string) error) func(_ *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if root := cmd.Root(); root != nil {
			if root.PersistentPreRunE != nil {
				err := root.PersistentPreRunE(root, args)
				if err != nil {
					return err
				}
			}
		}

		return preRun(cmd, args)
	}
}
