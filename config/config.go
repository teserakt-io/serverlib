package config

// DBType defines the different supported database types
type DBType string

const (
	// DBTypeEmpty defines an empty database type
	DBTypeEmpty DBType = ""
	// DBTypePostgres defines the PostgreSQL database type
	DBTypePostgres DBType = "postgres"
	// DBTypeSQLite defines the SQLite database type
	DBTypeSQLite DBType = "sqlite3"
)

// DBSecureConnectionType defines the different database connection security options
type DBSecureConnectionType string

const (
	// DBSecureConnectionEmpty defines an empty DBSecureConnectionType
	DBSecureConnectionEmpty DBSecureConnectionType = ""
	// DBSecureConnectionEnabled is used to enable SSL on the database connection
	DBSecureConnectionEnabled DBSecureConnectionType = "enabled"
	// DBSecureConnectionSelfSigned is used to allow SSL self signed certificates on the database connection
	DBSecureConnectionSelfSigned DBSecureConnectionType = "selfsigned"
	// DBSecureConnectionInsecure is used to disable SSL on database connection
	DBSecureConnectionInsecure DBSecureConnectionType = "insecure"

	// PostgresSSLModeFull is used to enable full certificate checks on postgres
	PostgresSSLModeFull = "sslmode=verify-full"
	// PostgresSSLModeRequire is used to allow self signed certificates on postgres
	PostgresSSLModeRequire = "sslmode=require"
	// PostgresSSLModeDisable is used to disable encryption on postgres
	PostgresSSLModeDisable = "sslmode=disable"
)

// PostgresSSLMode returns the corresponding SSLMode from SecureConnectionType
// defaulting to the most secure one.
func (m DBSecureConnectionType) PostgresSSLMode() string {
	switch m {
	case DBSecureConnectionSelfSigned:
		return PostgresSSLModeRequire
	case DBSecureConnectionInsecure:
		return PostgresSSLModeDisable
	default: // DBSecureConnectionEnabled and anything else
		return PostgresSSLModeFull
	}
}

func (m DBSecureConnectionType) String() string {
	return string(m)
}
