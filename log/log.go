package log

import "github.com/RichardKnop/logging"

var (
    // 日志格式
    logger = logging.New(nil, nil , new(logging.ColouredFormatter))

    INFO    = logger[logging.INFO]

    WARNING = logger[logging.WARNING]

    ERROR   = logger[logging.ERROR]

    FATAL   = logger[logging.FATAL]
)

func Set(l logging.LoggerInterface) {
    INFO    = l
    WARNING = l
    ERROR   = l
    FATAL   = l
}