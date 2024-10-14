package submitserver

import (
	"fmt"
	"net/http"

	"github.com/kubeflow/spark-operator/api/v1beta2"
)

type SubmitRequest struct {
	// TODO: Define the fields for the submit request
	body v1beta2.SparkApplication
}

func (s *SparkSubmitServer) Submit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit")
}

func (s *SparkSubmitServer) AppState(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit")
}

func (s *SparkSubmitServer) AppStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit")
}

func (s *SparkSubmitServer) DescribeApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit")
}

func (s *SparkSubmitServer) AppLogs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit")
}

func (s *SparkSubmitServer) DeleteApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Submit")
}

func (s *SparkSubmitServer) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}
