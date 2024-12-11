package routes

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/damongolding/immich-kiosk/internal/config"
	"github.com/damongolding/immich-kiosk/internal/utils"
	"github.com/labstack/echo/v4"
)

// Redirect returns an echo.HandlerFunc that handles URL redirections based on configured redirect paths.
// It takes a baseConfig parameter containing the application configuration including redirect mappings.
//
// If the requested redirect name exists in the RedirectsMap, it redirects to the mapped URL.
// Otherwise, it redirects to the root path "/".
//
// The function returns a temporary (307) redirect in both cases.
func Redirect(baseConfig *config.Config) echo.HandlerFunc {

	return func(c echo.Context) error {

		redirectCount, err := c.Cookie(redirectCountHeader)
		if err != nil {
			redirectCount = &http.Cookie{Value: "0"}
		}

		count := 0
		if redirectCount != nil {
			var err error
			count, err = strconv.Atoi(redirectCount.Value)
			if err != nil {
				count = 0
			}
		}

		// Check if maximum redirects exceeded
		if count >= maxRedirects {
			cookie := &http.Cookie{
				Name:   redirectCountHeader,
				Value:  "",
				MaxAge: -1,
			}
			c.SetCookie(cookie)
			return echo.NewHTTPError(http.StatusTooManyRequests, "Too many redirects")
		}

		redirectName := c.Param("redirect")
		if redirectName == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Redirect name is required")
		}

		if redirectItem, exists := baseConfig.Kiosk.RedirectsMap[redirectName]; exists {

			if strings.EqualFold(redirectItem.Type, "internal") {

				parsedUrl, err := url.Parse(redirectItem.URL)
				if err != nil {
					log.Error("parse internal redirect URL",
						"url", redirectItem.URL,
						"redirect", redirectName,
						"error", err)

					return echo.NewHTTPError(http.StatusInternalServerError, "Invalid redirect URL")
				}

				for key, values := range parsedUrl.Query() {
					for _, value := range values {
						c.QueryParams().Add(key, value)
					}
				}

				// Update the request URL with the new query parameters
				newURL := c.Request().URL
				queryParams := c.QueryParams()
				newURL.RawQuery = queryParams.Encode()
				c.Request().URL = newURL

				return Home(baseConfig)(c)
			}

			c.SetCookie(&http.Cookie{
				Name:  redirectCountHeader,
				Value: strconv.Itoa(count + 1),
			})

			mergedRedirect := mergeRequestQueries(c.QueryParams(), redirectItem)
			if _, err := url.Parse(mergedRedirect.URL); err != nil {
				log.Error("Invalid merged redirect URL", "error", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process redirect")
			}

			return c.Redirect(http.StatusTemporaryRedirect, redirectItem.URL)
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

// mergeRequestQueries combines query parameters from an incoming request with those
// already present in a redirect URL. It takes the request query parameters and a
// redirect configuration item as input, and returns an updated redirect configuration
// with merged query parameters in its URL.
//
// If parsing the redirect URL fails, the original redirect item is returned unchanged.
// Otherwise, it:
// 1. Extracts queries from both the request and redirect URL
// 2. Merges them using utils.MergeQueries
// 3. Updates the redirect URL with the combined query string
func mergeRequestQueries(requestQueries url.Values, redirectItem config.Redirect) config.Redirect {
	redirectURL, err := url.Parse(redirectItem.URL)
	if err != nil {
		log.Error("parse redirect URL", "url", redirectItem.URL, "err", err)
		return redirectItem
	}

	redirectQueries := redirectURL.Query()

	merged := utils.MergeQueries(requestQueries, redirectQueries)

	redirectURL.RawQuery = merged.Encode()
	redirectItem.URL = redirectURL.String()

	return redirectItem
}
