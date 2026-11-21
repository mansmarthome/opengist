package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thomiceli/opengist/internal/i18n"
	"github.com/thomiceli/opengist/internal/web/test"
)

func TestLocaleFromAcceptLanguage(t *testing.T) {
	s := test.Setup(t)
	defer test.Teardown(t)

	localeFor := func(acceptLanguage string) string {
		s.RequestWithHeaders(t, "GET", "/-/login", nil, 200, map[string]string{"Accept-Language": acceptLanguage})
		locale, ok := s.ContextData("locale")
		require.True(t, ok)
		return locale.(*i18n.Locale).Code
	}

	require.Equal(t, "ru-RU", localeFor("fr-FR,fr;q=0.9,en;q=0.8"))
	require.Equal(t, "ru-RU", localeFor(""))
}
