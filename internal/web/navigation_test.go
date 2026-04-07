package web

import (
	"encoding/json"
	"html/template"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guohuiyuan/music-lib/model"
)

func newTestTemplate(t *testing.T) *template.Template {
	t.Helper()

	return template.Must(template.New("").Funcs(template.FuncMap{
		"artistTokens":       splitArtistTokens,
		"albumID":            songAlbumID,
		"playlistDetailURL":  playlistDetailURL,
		"playlistExtraValue": playlistExtraValue,
		"tojson": func(v interface{}) string {
			if v == nil {
				return ""
			}
			b, err := json.Marshal(v)
			if err != nil {
				return ""
			}
			return string(b)
		},
	}).ParseFS(templateFS,
		"templates/pages/*.html",
		"templates/partials/*.html",
	))
}

func TestRenderIndexPlaylistCardsUseAjaxNavigation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.SetHTMLTemplate(newTestTemplate(t))
	router.GET(RoutePrefix, func(c *gin.Context) {
		renderIndex(c, nil, []model.Playlist{
			{
				ID:         "123",
				Name:       "Top Hits",
				TrackCount: 18,
				Creator:    "Tester",
				Source:     "qq",
				Cover:      "https://example.com/cover.jpg",
			},
		}, "", []string{"qq"}, "", collectionContentPlaylist, "", "", "", false, "", nil)
	})

	req := httptest.NewRequest("GET", RoutePrefix, nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	body := rec.Body.String()
	if rec.Code != 200 {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if strings.Contains(body, `onclick="location.href=`) {
		t.Fatalf("rendered html still uses location.href navigation: %s", body)
	}
	if !strings.Contains(body, `onclick="navigateTo('`) {
		t.Fatalf("rendered html missing navigateTo playlist navigation: %s", body)
	}
}

func TestAppJSIncludesAjaxNavigationEntryPoints(t *testing.T) {
	content, err := templateFS.ReadFile("templates/static/js/app.js")
	if err != nil {
		t.Fatalf("ReadFile(app.js): %v", err)
	}

	js := string(content)
	if !strings.Contains(js, "async function navigateTo(url, options = {})") {
		t.Fatal("app.js missing navigateTo function")
	}
	if !strings.Contains(js, "function bindPageNavigationEvents()") {
		t.Fatal("app.js missing bindPageNavigationEvents function")
	}
	if !strings.Contains(js, "initializePageContent(document);") {
		t.Fatal("app.js missing initializePageContent bootstrap")
	}
}
