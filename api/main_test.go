package api

import (
	"os"
	"testing"
	"time"

	db "github.com/Nevets9/simplebank/db/sqlc"
	"github.com/Nevets9/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server{
	config := util.Config{
		// TokenAssyemetricKey: util.RandomString(32),
		TokenSecret: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}