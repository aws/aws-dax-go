package client

import (
	"github.com/aws/smithy-go/logging"
)

const (
	ClassificationDebug                        = logging.Debug
	ClassificationWarn                         = logging.Warn
	ClassificationError logging.Classification = "ERROR"
)
