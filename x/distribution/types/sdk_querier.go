package types

import (
	context "context"

	types1 "github.com/cosmos/cosmos-sdk/x/distribution/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	grpc "google.golang.org/grpc"
)

func RegisterSDKQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/cosmos.distribution.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_ValidatorOutstandingRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorOutstandingRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorOutstandingRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorOutstandingRewards(ctx, req.(*QueryValidatorOutstandingRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_ValidatorCommission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorCommissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorCommission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorCommission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorCommission(ctx, req.(*QueryValidatorCommissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_ValidatorSlashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryValidatorSlashesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorSlashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/ValidatorSlashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorSlashes(ctx, req.(*QueryValidatorSlashesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegationRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.QueryDelegationRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationRewards(ctx, req.(*types1.QueryDelegationRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegationTotalRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegationTotalRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegationTotalRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegationTotalRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegationTotalRewards(ctx, req.(*QueryDelegationTotalRewardsRequest))
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
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorValidators(ctx, req.(*QueryDelegatorValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_DelegatorWithdrawAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDelegatorWithdrawAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DelegatorWithdrawAddress(ctx, req.(*QueryDelegatorWithdrawAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_CommunityPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommunityPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CommunityPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/CommunityPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CommunityPool(ctx, req.(*QueryCommunityPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Query_TokenizeShareRecordReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTokenizeShareRecordRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenizeShareRecordReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.distribution.v1beta1.Query/TokenizeShareRecordReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenizeShareRecordReward(ctx, req.(*QueryTokenizeShareRecordRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SDK_Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.distribution.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _SDK_Query_Params_Handler,
		},
		{
			MethodName: "ValidatorOutstandingRewards",
			Handler:    _SDK_Query_ValidatorOutstandingRewards_Handler,
		},
		{
			MethodName: "ValidatorCommission",
			Handler:    _SDK_Query_ValidatorCommission_Handler,
		},
		{
			MethodName: "ValidatorSlashes",
			Handler:    _SDK_Query_ValidatorSlashes_Handler,
		},
		{
			MethodName: "DelegationRewards",
			Handler:    _SDK_Query_DelegationRewards_Handler,
		},
		{
			MethodName: "DelegationTotalRewards",
			Handler:    _SDK_Query_DelegationTotalRewards_Handler,
		},
		{
			MethodName: "DelegatorValidators",
			Handler:    _SDK_Query_DelegatorValidators_Handler,
		},
		{
			MethodName: "DelegatorWithdrawAddress",
			Handler:    _SDK_Query_DelegatorWithdrawAddress_Handler,
		},
		{
			MethodName: "CommunityPool",
			Handler:    _SDK_Query_CommunityPool_Handler,
		},
		{
			MethodName: "TokenizeShareRecordReward",
			Handler:    _SDK_Query_TokenizeShareRecordReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/distribution/v1beta1/query.proto",
}
