package encoder

import (
	"github.com/aaguilartablada/grabana/encoder/golang"
	"github.com/aaguilartablada/grafana-sdk"
	"go.uber.org/zap"
)

func ToGolang(logger *zap.Logger, dashboard sdk.Board) (string, error) {
	golangEncoder := golang.NewEncoder(logger)

	return golangEncoder.EncodeDashboard(dashboard)
}
