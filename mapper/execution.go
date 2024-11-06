package mapper

import (
	pb "github.com/nuzur/extension-sdk/idl/gen"
	nemgen "github.com/nuzur/extension-sdk/proto_deps/nem/idl/gen"
)

func MapExecutionToGetResponse(
	in *nemgen.ExtensionExecution,
	step *pb.ExecutionResponseTypeStepData,
	final *pb.ExecutionResponseTypeFinalData,
) *pb.GetExecutionResponse {
	switch in.Status {
	case nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_CANCELLED:
		return &pb.GetExecutionResponse{
			Status:    pb.ExecutionStatus_EXECUTION_STATUS_CANCELLED,
			StatusMsg: in.StatusMsg,
			Type:      pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_INVALID,
			// calculate timeelspased
		}
	case nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_FAILED:
		return &pb.GetExecutionResponse{
			Status:    pb.ExecutionStatus_EXECUTION_STATUS_FAILED,
			StatusMsg: in.StatusMsg,
			Type:      pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_INVALID,
			// calculate timeelspased
		}
	case nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_INPROGRESS:
		return &pb.GetExecutionResponse{
			Status:                pb.ExecutionStatus_EXECUTION_STATUS_INPROGRESS,
			StatusMsg:             in.StatusMsg,
			Type:                  pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_INVALID,
			CurrentStepIdentifier: step.GetStepIdentifier(),
			Data: &pb.ExecutionResponseTypeData{
				Step: step,
			},
			// calculate timeelspased
		}
	case nemgen.ExtensionExecutionStatus_EXTENSION_EXECUTION_STATUS_SUCCEEDED:
		return &pb.GetExecutionResponse{
			Status:    pb.ExecutionStatus_EXECUTION_STATUS_SUCCEEDED,
			StatusMsg: in.StatusMsg,
			Type:      pb.ExecutionResponseType_EXECUTION_RESPONSE_TYPE_FINAL,
			Data: &pb.ExecutionResponseTypeData{
				Final: final,
			},
			// calculate timeelspased
		}
	}

	return &pb.GetExecutionResponse{
		Status: pb.ExecutionStatus_EXECUTION_STATUS_INVALID,
	}
}
