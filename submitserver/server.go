package submitserver

import "github.com/kubeflow/spark-operator/api/v1beta2"

type SparkApplicationRest interface {
	SubmitJob(application v1beta2.SparkApplication) v1beta2.SparkApplicationStatus
	DescribeJob(appName string) v1beta2.SparkApplication
	ApplicationState(appName string) v1beta2.SparkApplicationStatus
	ApplicationStatus(appName string) v1beta2.ApplicationState
	ApplicationLogs(appName string) string
	DeleteJob(application v1beta2.SparkApplication) bool
	// TODO: Add any other necessary methods for the server
}
