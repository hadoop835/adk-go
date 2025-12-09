// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adka2a

import (
	"context"

	"github.com/a2aproject/a2a-go/a2asrv"

	"google.golang.org/genai"

	"google.golang.org/adk/session"
)

type ExecutorContext interface {
	context.Context

	SessionID() string
	UserID() string
	AgentName() string
	ReadonlyState() session.ReadonlyState
	UserContent() *genai.Content
	RequestContext() *a2asrv.RequestContext
}

type executorContext struct {
	context.Context
	meta        invocationMeta
	session     session.ReadonlyState
	userContent *genai.Content
}

var _ ExecutorContext = (*executorContext)(nil)

func newExecutorContext(ctx context.Context, meta invocationMeta, session session.ReadonlyState, userContent *genai.Content) ExecutorContext {
	return &executorContext{
		Context:     ctx,
		meta:        meta,
		session:     session,
		userContent: userContent,
	}
}

func (ec *executorContext) SessionID() string {
	return ec.meta.sessionID
}

func (ec *executorContext) UserID() string {
	return ec.meta.userID
}

func (ec *executorContext) AgentName() string {
	return ec.meta.agentName
}

func (ec *executorContext) ReadonlyState() session.ReadonlyState {
	return ec.session
}

func (ec *executorContext) RequestContext() *a2asrv.RequestContext {
	return ec.meta.reqCtx
}

func (ec *executorContext) UserContent() *genai.Content {
	return ec.userContent
}
