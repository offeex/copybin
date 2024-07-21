package handler

import (
	"context"
	"copybin/internal/config"
	"copybin/internal/impl/pasta"
	"copybin/internal/impl/user"
	"copybin/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctx := context.Background()

	// Db setup
	pgContainer, err := postgres.Run(
		ctx,
		"postgres",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("poggers"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	if err != nil {
		t.Fatalf("failed to create container: %v", err)
		return
	}

	if err := pgContainer.Start(ctx); err != nil {
		t.Fatalf("failed to start container: %v", err)
		return
	}

	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %v", err)
		}
	}()

	connString := pgContainer.MustConnectionString(ctx)

	if err != nil {
		t.Fatalf("failed to get endpoint: %v", err)
		return
	}

	db := storage.New(connString)
	h := &Handler{
		cfg: &config.Config{
			JWTSecret: "de1258d7f7cd665c9aa9e7c60322777ab32bae88e8c746ebfd8bd72e0a81de91",
		},
		pastaStore: pasta.NewStore(db),
		userStore:  user.NewStore(db),
	}
	router := gin.New()
	router.Use(h.AuthMiddleware)
	router.GET("/me", h.Me)

	if _, err := h.userStore.CreateIfNotExists("digas"); err != nil {
		t.Fatalf("failed to create user: %v", err)
		return
	}

	req, _ := http.NewRequest(
		"GET",
		"/me",
		nil,
	)
	req.Header.Set(
		"Authorization",
		"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjb3B5YmluIiwic3ViIjoiMSIsImV4cCI6MTcyMTY1OTYyNywiaWF0IjoxNzIxNTczMjI3fQ.wd41jmhtJzx0sW1KWzp5CkllDbOW5ShL9grLtika6jw",
	)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf(
			"handler returned wrong status code: got %v",
			w.Code,
		)
	}

	t.Logf("Response: %s", w.Body.String())
}
