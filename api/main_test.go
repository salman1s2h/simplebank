package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/salman1s2h/simplebank/db/sqlc"
	"github.com/salman1s2h/simplebank/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Env{
		TOKEN_SYMMETRIC_KEY:   util.RandomString(32),
		ACCESS_TOKEN_DURATION: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
