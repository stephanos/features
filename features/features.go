package features

import (
	update_activities "github.com/temporalio/features/features/update/activities"
	update_async_accepted "github.com/temporalio/features/features/update/async_accepted"
	update_basic "github.com/temporalio/features/features/update/basic"
	update_client_interceptor "github.com/temporalio/features/features/update/client_interceptor"
	update_deduplication "github.com/temporalio/features/features/update/deduplication"
	update_non_durable_reject "github.com/temporalio/features/features/update/non_durable_reject"
	update_self "github.com/temporalio/features/features/update/self"
	update_task_failure "github.com/temporalio/features/features/update/task_failure"
	update_validation_replay "github.com/temporalio/features/features/update/validation_replay"
	update_worker_restart "github.com/temporalio/features/features/update/worker_restart"
	harness "github.com/temporalio/features/harness/go/harness"
)

func init() {
	// Please keep list in alphabetical order
	harness.MustRegisterFeatures(
		//activity_basic_no_workflow_timeout.Feature,
		//activity_cancel_try_cancel.Feature,
		//activity_retry_on_error.Feature,
		//bugs_go_activity_start_race.Feature,
		//bugs_go_child_workflow_cancel_panic.Feature,
		//build_id_versioning_activity_and_child_on_correct_version.Feature,
		//build_id_versioning_continues_as_new_on_correct_version.Feature,
		//build_id_versioning_only_appropriate_worker_gets_task.Feature,
		//build_id_versioning_unversioned_worker_gets_unversioned_task.Feature,
		//build_id_versioning_unversioned_worker_no_task.Feature,
		//build_id_versioning_versions_added_while_worker_polling.Feature,
		//child_workflow_result.Feature,
		//child_workflow_signal.Feature,
		//client_http_proxy.Feature,
		//client_http_proxy_auth.Feature,
		//continue_as_new_continue_as_same.Feature,
		//data_converter_binary_protobuf.Feature,
		//data_converter_binary.Feature,
		//data_converter_codec.Feature,
		//data_converter_empty.Feature,
		//data_converter_failure.Feature,
		//data_converter_json_protobuf.Feature,
		//data_converter_json.Feature,
		//eager_activity_non_remote_activities_worker.Feature,
		//eager_workflow_successful_start.Feature,
		//query_successful_query.Feature,
		//query_timeout_due_to_no_active_workers.Feature,
		//query_unexpected_arguments.Feature,
		//query_unexpected_query_type_name.Feature,
		//query_unexpected_return_type.Feature,
		//reset_reset_and_delete.Feature,
		//schedule_backfill.Feature,
		//schedule_basic.Feature,
		//schedule_cron.Feature,
		//schedule_pause.Feature,
		//schedule_trigger.Feature,
		//signal_external.Feature,
		//telemetry_metrics.Feature,
		update_activities.Feature,
		update_async_accepted.Feature,
		update_basic.Feature,
		update_deduplication.Feature,
		update_client_interceptor.Feature,
		update_non_durable_reject.Feature,
		update_self.Feature,
		update_task_failure.Feature,
		update_validation_replay.Feature,
		update_worker_restart.Feature,
	)
}
