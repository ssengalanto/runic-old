package log_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/runic/pkg/constants"
	"github.com/ssengalanto/runic/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	assertResult := func(t *testing.T, result *log.Log, err error) {
		assert.NotNil(t, result)
		require.NoError(t, err)
	}

	assertError := func(t *testing.T, result *log.Log, err error) {
		assert.Nil(t, result)
		require.Error(t, err)
	}

	testCases := []struct {
		name   string
		env    string
		assert func(t *testing.T, result *log.Log, err error)
	}{
		{
			name:   "development environment",
			env:    constants.DevEnv,
			assert: assertResult,
		},
		{
			name:   "staging environment",
			env:    constants.StgEnv,
			assert: assertResult,
		},
		{
			name:   "production environment",
			env:    constants.ProdEnv,
			assert: assertResult,
		},
		{
			name:   "invalid environment",
			env:    "invalid",
			assert: assertError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slog, err := log.New(tc.env)
			tc.assert(t, slog, err)
		})
	}
}

func TestLogger(t *testing.T) {
	tests := []struct {
		level zapcore.Level
		logFn func(log.Log, string, ...log.KeyValue)
	}{
		{zapcore.InfoLevel, func(l log.Log, msg string, fields ...log.KeyValue) { l.Info(msg, fields...) }},
		{zapcore.ErrorLevel, func(l log.Log, msg string, fields ...log.KeyValue) { l.Error(msg, fields...) }},
		{zapcore.DebugLevel, func(l log.Log, msg string, fields ...log.KeyValue) { l.Debug(msg, fields...) }},
		{zapcore.WarnLevel, func(l log.Log, msg string, fields ...log.KeyValue) { l.Warn(msg, fields...) }},
		// {zapcore.FatalLevel, func(l log.Log, msg string, fields ...log.KeyValue) { l.Fatal(msg, fields...) }},
		// {zapcore.PanicLevel, func(l log.Log, msg string, fields ...log.KeyValue) { l.Panic(msg, fields...) }},
	}

	for _, tc := range tests {
		t.Run(tc.level.String(), func(t *testing.T) {
			msg := gofakeit.Word()
			logger, observedLogs := log.NewTestInstance(tc.level)
			tc.logFn(logger, msg)

			allLogs := observedLogs.All()
			assert.Equal(t, msg, allLogs[0].Message)
			assert.Equal(t, 1, observedLogs.Len())
		})
	}
}

func TestField(t *testing.T) {
	field := log.Field("key", "value")
	assert.Equal(t, "key", field.Key)
	assert.Equal(t, "value", field.Value)
}
