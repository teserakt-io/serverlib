// Copyright 2020 Teserakt AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

// DBType defines the different supported database types
type DBType string

func (d DBType) String() string {
	return string(d)
}

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

// IsInsecure return true whenever the secure connection type is insecure
func (m DBSecureConnectionType) IsInsecure() bool {
	return m == DBSecureConnectionInsecure
}

// IsSelfSigned return true whenever the secure connection type is self signed
func (m DBSecureConnectionType) IsSelfSigned() bool {
	return m == DBSecureConnectionSelfSigned
}

func (m DBSecureConnectionType) String() string {
	return string(m)
}
