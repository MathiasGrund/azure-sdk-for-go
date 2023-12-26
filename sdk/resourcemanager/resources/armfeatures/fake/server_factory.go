//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armfeatures.ClientFactory type.
type ServerFactory struct {
	Server                                 Server
	FeatureServer                          FeatureServer
	SubscriptionFeatureRegistrationsServer SubscriptionFeatureRegistrationsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armfeatures.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armfeatures.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                      *ServerFactory
	trMu                                     sync.Mutex
	trServer                                 *ServerTransport
	trFeatureServer                          *FeatureServerTransport
	trSubscriptionFeatureRegistrationsServer *SubscriptionFeatureRegistrationsServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "Client":
		initServer(s, &s.trServer, func() *ServerTransport { return NewServerTransport(&s.srv.Server) })
		resp, err = s.trServer.Do(req)
	case "FeatureClient":
		initServer(s, &s.trFeatureServer, func() *FeatureServerTransport { return NewFeatureServerTransport(&s.srv.FeatureServer) })
		resp, err = s.trFeatureServer.Do(req)
	case "SubscriptionFeatureRegistrationsClient":
		initServer(s, &s.trSubscriptionFeatureRegistrationsServer, func() *SubscriptionFeatureRegistrationsServerTransport {
			return NewSubscriptionFeatureRegistrationsServerTransport(&s.srv.SubscriptionFeatureRegistrationsServer)
		})
		resp, err = s.trSubscriptionFeatureRegistrationsServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
