package config

import "testing"

func TestDBSecureConnectionType(t *testing.T) {
	t.Run("PostgresSSLMode returns expected values", func(t *testing.T) {
		if DBSecureConnectionEnabled.PostgresSSLMode() != PostgresSSLModeFull {
			t.Errorf(
				"Expected PostgresSSLMode to return %v, got %v",
				PostgresSSLModeFull,
				DBSecureConnectionEnabled.PostgresSSLMode(),
			)
		}

		if DBSecureConnectionSelfSigned.PostgresSSLMode() != PostgresSSLModeRequire {
			t.Errorf(
				"Expected PostgresSSLMode to return %v, got %v",
				PostgresSSLModeRequire,
				DBSecureConnectionSelfSigned.PostgresSSLMode(),
			)
		}

		if DBSecureConnectionInsecure.PostgresSSLMode() != PostgresSSLModeDisable {
			t.Errorf(
				"Expected PostgresSSLMode to return %v, got %v",
				PostgresSSLModeDisable,
				DBSecureConnectionInsecure.PostgresSSLMode(),
			)
		}

		if DBSecureConnectionEmpty.PostgresSSLMode() != PostgresSSLModeFull {
			t.Errorf(
				"Expected PostgresSSLMode to return %v, got %v",
				PostgresSSLModeFull,
				DBSecureConnectionEmpty.PostgresSSLMode(),
			)
		}
	})
}
