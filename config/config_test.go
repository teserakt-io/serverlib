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

	t.Run("IsSelfSigned return true only on self signed type", func(t *testing.T) {
		testCases := map[DBSecureConnectionType]bool{
			DBSecureConnectionSelfSigned: true,
			DBSecureConnectionEmpty:      false,
			DBSecureConnectionInsecure:   false,
			DBSecureConnectionEnabled:    false,
		}

		for secureCnx, expectedResult := range testCases {
			if secureCnx.IsSelfSigned() != expectedResult {
				t.Errorf(
					"Expected IsSelfSigned to return %v for type %v, got %v",
					expectedResult,
					secureCnx,
					secureCnx.IsSelfSigned(),
				)
			}
		}
	})

	t.Run("IsInsecure return true only on insecure type", func(t *testing.T) {
		testCases := map[DBSecureConnectionType]bool{
			DBSecureConnectionSelfSigned: false,
			DBSecureConnectionEmpty:      false,
			DBSecureConnectionInsecure:   true,
			DBSecureConnectionEnabled:    false,
		}

		for secureCnx, expectedResult := range testCases {
			if secureCnx.IsInsecure() != expectedResult {
				t.Errorf(
					"Expected IsInsecure to return %v for type %v, got %v",
					expectedResult,
					secureCnx,
					secureCnx.IsInsecure(),
				)
			}
		}
	})
}
