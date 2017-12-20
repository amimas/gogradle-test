package main

import (
	"io"

	"github.com/Shopify/sarama"
)

type logPair struct {
	stream   io.ReadCloser
	info     logger.Info
	producer sarama.AsyncProducer
}
