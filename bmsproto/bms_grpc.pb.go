// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: bmsproto/bms.proto

package bms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BmsDatabaseCrudClient is the client API for BmsDatabaseCrud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BmsDatabaseCrudClient interface {
	CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	AddMovie(ctx context.Context, in *NewMovie, opts ...grpc.CallOption) (*Movie, error)
	AddBooking(ctx context.Context, in *NewBooking, opts ...grpc.CallOption) (*Booking, error)
	AddPayment(ctx context.Context, in *NewPayment, opts ...grpc.CallOption) (*Payment, error)
	GetMovies(ctx context.Context, in *EmptyMovie, opts ...grpc.CallOption) (*Movies, error)
	GetMovieByPreference(ctx context.Context, in *MoviePreference, opts ...grpc.CallOption) (*Movies, error)
	GetListOfShowsByTheatre(ctx context.Context, in *Theatre, opts ...grpc.CallOption) (*Shows, error)
	GetListOfBookingsByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Bookings, error)
	UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type bmsDatabaseCrudClient struct {
	cc grpc.ClientConnInterface
}

func NewBmsDatabaseCrudClient(cc grpc.ClientConnInterface) BmsDatabaseCrudClient {
	return &bmsDatabaseCrudClient{cc}
}

func (c *bmsDatabaseCrudClient) CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) AddMovie(ctx context.Context, in *NewMovie, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/AddMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) AddBooking(ctx context.Context, in *NewBooking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/AddBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) AddPayment(ctx context.Context, in *NewPayment, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/AddPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) GetMovies(ctx context.Context, in *EmptyMovie, opts ...grpc.CallOption) (*Movies, error) {
	out := new(Movies)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/GetMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) GetMovieByPreference(ctx context.Context, in *MoviePreference, opts ...grpc.CallOption) (*Movies, error) {
	out := new(Movies)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/GetMovieByPreference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) GetListOfShowsByTheatre(ctx context.Context, in *Theatre, opts ...grpc.CallOption) (*Shows, error) {
	out := new(Shows)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/GetListOfShowsByTheatre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) GetListOfBookingsByUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Bookings, error) {
	out := new(Bookings)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/GetListOfBookingsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bmsDatabaseCrudClient) UpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/bms.BmsDatabaseCrud/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BmsDatabaseCrudServer is the server API for BmsDatabaseCrud service.
// All implementations must embed UnimplementedBmsDatabaseCrudServer
// for forward compatibility
type BmsDatabaseCrudServer interface {
	CreateUser(context.Context, *NewUser) (*User, error)
	AddMovie(context.Context, *NewMovie) (*Movie, error)
	AddBooking(context.Context, *NewBooking) (*Booking, error)
	AddPayment(context.Context, *NewPayment) (*Payment, error)
	GetMovies(context.Context, *EmptyMovie) (*Movies, error)
	GetMovieByPreference(context.Context, *MoviePreference) (*Movies, error)
	GetListOfShowsByTheatre(context.Context, *Theatre) (*Shows, error)
	GetListOfBookingsByUser(context.Context, *User) (*Bookings, error)
	UpdateUser(context.Context, *User) (*User, error)
	mustEmbedUnimplementedBmsDatabaseCrudServer()
}

// UnimplementedBmsDatabaseCrudServer must be embedded to have forward compatible implementations.
type UnimplementedBmsDatabaseCrudServer struct {
}

func (UnimplementedBmsDatabaseCrudServer) CreateUser(context.Context, *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) AddMovie(context.Context, *NewMovie) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) AddBooking(context.Context, *NewBooking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBooking not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) AddPayment(context.Context, *NewPayment) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPayment not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) GetMovies(context.Context, *EmptyMovie) (*Movies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) GetMovieByPreference(context.Context, *MoviePreference) (*Movies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByPreference not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) GetListOfShowsByTheatre(context.Context, *Theatre) (*Shows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfShowsByTheatre not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) GetListOfBookingsByUser(context.Context, *User) (*Bookings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfBookingsByUser not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) UpdateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedBmsDatabaseCrudServer) mustEmbedUnimplementedBmsDatabaseCrudServer() {}

// UnsafeBmsDatabaseCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BmsDatabaseCrudServer will
// result in compilation errors.
type UnsafeBmsDatabaseCrudServer interface {
	mustEmbedUnimplementedBmsDatabaseCrudServer()
}

func RegisterBmsDatabaseCrudServer(s grpc.ServiceRegistrar, srv BmsDatabaseCrudServer) {
	s.RegisterService(&BmsDatabaseCrud_ServiceDesc, srv)
}

func _BmsDatabaseCrud_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).CreateUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/AddMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).AddMovie(ctx, req.(*NewMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_AddBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBooking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).AddBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/AddBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).AddBooking(ctx, req.(*NewBooking))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_AddPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).AddPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/AddPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).AddPayment(ctx, req.(*NewPayment))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).GetMovies(ctx, req.(*EmptyMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_GetMovieByPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoviePreference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).GetMovieByPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/GetMovieByPreference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).GetMovieByPreference(ctx, req.(*MoviePreference))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_GetListOfShowsByTheatre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Theatre)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).GetListOfShowsByTheatre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/GetListOfShowsByTheatre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).GetListOfShowsByTheatre(ctx, req.(*Theatre))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_GetListOfBookingsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).GetListOfBookingsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/GetListOfBookingsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).GetListOfBookingsByUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _BmsDatabaseCrud_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BmsDatabaseCrudServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bms.BmsDatabaseCrud/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BmsDatabaseCrudServer).UpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// BmsDatabaseCrud_ServiceDesc is the grpc.ServiceDesc for BmsDatabaseCrud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BmsDatabaseCrud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bms.BmsDatabaseCrud",
	HandlerType: (*BmsDatabaseCrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _BmsDatabaseCrud_CreateUser_Handler,
		},
		{
			MethodName: "AddMovie",
			Handler:    _BmsDatabaseCrud_AddMovie_Handler,
		},
		{
			MethodName: "AddBooking",
			Handler:    _BmsDatabaseCrud_AddBooking_Handler,
		},
		{
			MethodName: "AddPayment",
			Handler:    _BmsDatabaseCrud_AddPayment_Handler,
		},
		{
			MethodName: "GetMovies",
			Handler:    _BmsDatabaseCrud_GetMovies_Handler,
		},
		{
			MethodName: "GetMovieByPreference",
			Handler:    _BmsDatabaseCrud_GetMovieByPreference_Handler,
		},
		{
			MethodName: "GetListOfShowsByTheatre",
			Handler:    _BmsDatabaseCrud_GetListOfShowsByTheatre_Handler,
		},
		{
			MethodName: "GetListOfBookingsByUser",
			Handler:    _BmsDatabaseCrud_GetListOfBookingsByUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _BmsDatabaseCrud_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bmsproto/bms.proto",
}
