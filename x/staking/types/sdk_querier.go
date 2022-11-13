package types

import (
	context "context"

	grpc1 "github.com/gogo/protobuf/grpc"
	grpc "google.golang.org/grpc"
)

func RegisterSDKQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_SDK_Query_serviceDesc, srv)
}

func _SDK_Query_Validators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Validators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validators(ctx, req.(*QueryValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_Validator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Validator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validator(ctx, req.(*QueryValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_ValidatorDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/ValidatorDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorDelegations(ctx, req.(*QueryValidatorDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_ValidatorUnbondingDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorUnbondingDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorUnbondingDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/ValidatorUnbondingDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorUnbondingDelegations(ctx, req.(*QueryValidatorUnbondingDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_Delegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Delegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Delegation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Delegation(ctx, req.(*QueryDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_UnbondingDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnbondingDelegationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnbondingDelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/UnbondingDelegation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnbondingDelegation(ctx, req.(*QueryUnbondingDelegationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegatorDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/DelegatorDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorDelegations(ctx, req.(*QueryDelegatorDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegatorUnbondingDelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorUnbondingDelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorUnbondingDelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/DelegatorUnbondingDelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorUnbondingDelegations(ctx, req.(*QueryDelegatorUnbondingDelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_Redelegations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRedelegationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Redelegations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Redelegations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Redelegations(ctx, req.(*QueryRedelegationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegatorValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/DelegatorValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(ctx, req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegatorValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/DelegatorValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidator(ctx, req.(*QueryDelegatorValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_HistoricalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHistoricalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HistoricalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/HistoricalInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HistoricalInfo(ctx, req.(*QueryHistoricalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Pool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_TokenizeShareRecordById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenizeShareRecordByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenizeShareRecordById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/TokenizeShareRecordById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenizeShareRecordById(ctx, req.(*QueryTokenizeShareRecordByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_TokenizeShareRecordByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenizeShareRecordByDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenizeShareRecordByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/TokenizeShareRecordByDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenizeShareRecordByDenom(ctx, req.(*QueryTokenizeShareRecordByDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_TokenizeShareRecordsOwned_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenizeShareRecordsOwnedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenizeShareRecordsOwned(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/TokenizeShareRecordsOwned",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenizeShareRecordsOwned(ctx, req.(*QueryTokenizeShareRecordsOwnedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_AllTokenizeShareRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTokenizeShareRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllTokenizeShareRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/AllTokenizeShareRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllTokenizeShareRecords(ctx, req.(*QueryAllTokenizeShareRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_LastTokenizeShareRecordId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastTokenizeShareRecordIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastTokenizeShareRecordId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/LastTokenizeShareRecordId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastTokenizeShareRecordId(ctx, req.(*QueryLastTokenizeShareRecordIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_TotalTokenizeSharedAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalTokenizeSharedAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalTokenizeSharedAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.staking.v1beta1.Query/TotalTokenizeSharedAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalTokenizeSharedAssets(ctx, req.(*QueryTotalTokenizeSharedAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SDK_Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.staking.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validators",
			Handler:    _SDK_Query_Validators_Handler,
		},
		{
			MethodName: "Validator",
			Handler:    _SDK_Query_Validator_Handler,
		},
		{
			MethodName: "ValidatorDelegations",
			Handler:    _SDK_Query_ValidatorDelegations_Handler,
		},
		{
			MethodName: "ValidatorUnbondingDelegations",
			Handler:    _SDK_Query_ValidatorUnbondingDelegations_Handler,
		},
		{
			MethodName: "Delegation",
			Handler:    _SDK_Query_Delegation_Handler,
		},
		{
			MethodName: "UnbondingDelegation",
			Handler:    _SDK_Query_UnbondingDelegation_Handler,
		},
		{
			MethodName: "DelegatorDelegations",
			Handler:    _SDK_Query_DelegatorDelegations_Handler,
		},
		{
			MethodName: "DelegatorUnbondingDelegations",
			Handler:    _SDK_Query_DelegatorUnbondingDelegations_Handler,
		},
		{
			MethodName: "Redelegations",
			Handler:    _SDK_Query_Redelegations_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _SDK_Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorValidator",
			Handler:    _SDK_Query_DelegatorValidator_Handler,
		},
		{
			MethodName: "HistoricalInfo",
			Handler:    _SDK_Query_HistoricalInfo_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _SDK_Query_Pool_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _SDK_Query_Params_Handler,
		},
		{
			MethodName: "TokenizeShareRecordById",
			Handler:    _SDK_Query_TokenizeShareRecordById_Handler,
		},
		{
			MethodName: "TokenizeShareRecordByDenom",
			Handler:    _SDK_Query_TokenizeShareRecordByDenom_Handler,
		},
		{
			MethodName: "TokenizeShareRecordsOwned",
			Handler:    _SDK_Query_TokenizeShareRecordsOwned_Handler,
		},
		{
			MethodName: "AllTokenizeShareRecords",
			Handler:    _SDK_Query_AllTokenizeShareRecords_Handler,
		},
		{
			MethodName: "LastTokenizeShareRecordId",
			Handler:    _SDK_Query_LastTokenizeShareRecordId_Handler,
		},
		{
			MethodName: "TotalTokenizeSharedAssets",
			Handler:    _SDK_Query_TotalTokenizeSharedAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staking/v1beta1/query.proto",
}
