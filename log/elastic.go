package log

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/olivere/elastic"
)

type elasticLogger struct {
	logger        log.Logger
	esClient      *elastic.Client
	indexBaseName string
}

var _ log.Logger = &elasticLogger{}

// WithElasticSearch wrap around a gokit logger to make it forward logs to elasticSearch,
// using the provided elastic.Client.
// Index will be created from given indexBaseName, with the date appended
// like so: <indexBaseName>-YYYY-MM-DD, for easier index management.
func WithElasticSearch(logger log.Logger, esClient *elastic.Client, indexBaseName string) (log.Logger, error) {
	return &elasticLogger{
		logger:        logger,
		esClient:      esClient,
		indexBaseName: indexBaseName,
	}, nil
}

// Log starts a goroutine responsible of publishing the logged keyvals to elasticsearch.
// It then calls the wrapped logger if provided.
func (l *elasticLogger) Log(keyvals ...interface{}) error {
	// this goroutine send the logged values to elasticsearch.
	// done in the background to not slow down the application execution.
	// on elasticsearch failure, the error will be only logged on the wrapped logger.
	go func() {
		// we create and use this logger just to ease the json convertion from keyvals into buf
		buf := bytes.NewBuffer(nil)
		jsonLogger := log.NewJSONLogger(buf)
		jsonLogger = log.With(jsonLogger, "@timestamp", log.DefaultTimestamp)

		if err := jsonLogger.Log(keyvals...); err != nil {
			l.logger.Log("msg", "failed to log keyvals to buffer", "error", err, "data", keyvals)
		}

		index := fmt.Sprintf("%s-%s", l.indexBaseName, time.Now().Format("2006.01.02"))
		_, err := l.esClient.Index().Index(index).Type("log").
			BodyString(buf.String()).
			Refresh("true").
			Do(context.Background())

		// if elasticSearch logging fail, we log the error on the standard logger only.
		if err != nil {
			l.logger.Log("msg", "failed to log to elasticsearch", "error", err, "data", keyvals)
			return
		}
	}()

	// Default gokit logger caller is fetched from 3 levels deep in callstack
	// we need 2 more levels to keep proper caller displaying.
	logger := log.With(l.logger, "caller", log.Caller(5))
	if err := logger.Log(keyvals...); err != nil {
		return err
	}

	return nil
}
