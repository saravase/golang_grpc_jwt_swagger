package auth

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	Auth        *Auth
	AccessRoles map[string][]string
}

func NewAuthInterceptor(auth *Auth, accessRoles map[string][]string) *AuthInterceptor {

	return &AuthInterceptor{
		Auth:        auth,
		AccessRoles: accessRoles,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		log.Println("--> Unary Interceptor ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler) error {
		log.Println("--> Stream Interceptor ", info.FullMethod)

		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {

	log.Println("==== Authorize Check === ", method)
	// Check method availablity
	accessibleRoles, found := interceptor.AccessRoles[method]
	if !found {
		log.Println("==== Found Failed === ", accessibleRoles)
		return nil
	}

	log.Println("==== Roles === ", accessibleRoles)

	// Check Meta data availablity
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return status.Error(codes.Unauthenticated, "Meta data couldn't provided")
	}

	log.Println("==== Meta Data === ", md)

	// Check Access Token availablity
	authValues := md["authorization"]
	if len(authValues) == 0 {
		return status.Error(codes.Unauthenticated, "Access token is not provided")
	}

	log.Println("==== Auth Values === ", authValues)
	// Verify access token
	accessToken := authValues[0]
	claims, err := interceptor.Auth.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Access token is invalid. Reason: %v", err)
	}

	log.Println("==== Access Token === ", accessToken)

	// Check method access permission
	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	return status.Error(codes.Unauthenticated, "Method access permission denied")

}
