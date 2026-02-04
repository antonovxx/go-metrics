package metrics

import (
	models "antonovxx/go-metrics/internal/model"
	"antonovxx/go-metrics/internal/repository"
	"net/http"
	"strconv"
	"strings"
)

func Update(storage *repository.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/update/")
		parts := strings.Split(path, "/")

		if len(parts) != 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		metricType := parts[0]
		metricName := parts[1]
		metricValue := parts[2]

		if metricName == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch metricType {
		case models.Gauge:
			value, err := strconv.ParseFloat(metricValue, 64)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			storage.UpdateGauge(metricName, value)

		case models.Counter:
			value, err := strconv.ParseInt(metricValue, 10, 64)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			storage.UpdateCounter(metricName, value)

		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
