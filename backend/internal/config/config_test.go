package config_test

import (
	"os"
	"testing"

	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	host := "localhost"
	port := "3000"
	tz := "UTC"
	env := "development"
	image := "echo-service"
	mgouri := "echo-service"
	mgoconame := "echo-service"
	mgodbname := "echo-service"
	jwtsecret := "fake-secret-jwt"

	os.Setenv("HOST", host)
	os.Setenv("PORT", port)
	os.Setenv("ENV", env)
	os.Setenv("TZ", tz)
	os.Setenv("IMAGE", image)

	os.Setenv("MONGO_URI", mgouri)
	os.Setenv("MONGO_CONNECTION_NAME", mgoconame)
	os.Setenv("MONGO_DATABASE_NAME", mgodbname)

	os.Setenv("JWT_SECRET", jwtsecret)

	cfg := config.New()
	assert.NotNil(t, cfg)
	assert.Equal(t, host, cfg.Host)
	assert.Equal(t, port, cfg.Port)
	assert.Equal(t, env, cfg.ENV)
	assert.Equal(t, tz, cfg.TZ)
	assert.Equal(t, image, cfg.Image)

	assert.Equal(t, mgouri, cfg.Mongos[0].URI)
	assert.Equal(t, mgoconame, cfg.Mongos[0].ConnectionName)
	assert.Equal(t, mgodbname, cfg.Mongos[0].DatabaseName)

	assert.Equal(t, jwtsecret, cfg.JWTSecret)
}
