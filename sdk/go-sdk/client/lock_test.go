package client

/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"context"
	"github.com/stretchr/testify/assert"
	"mosn.io/layotto/components/lock"
	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
	"testing"
)

func TestTryLock(t *testing.T) {
	ctx := context.Background()
	t.Run("try lock", func(t *testing.T) {
		request := runtimev1pb.TryLockRequest{
			ResourceId: "lock_test",
			LockOwner:  "layotto",
		}
		lock, err := testClient.TryLock(ctx, &request)
		assert.Nil(t, err)
		assert.NotNil(t, lock.Success, true)
	})
}

func TestUnLock(t *testing.T) {
	ctx := context.Background()
	t.Run("try lock", func(t *testing.T) {
		request := runtimev1pb.TryLockRequest{
			ResourceId: "lock_test",
			LockOwner:  "layotto",
		}
		lock, err := testClient.TryLock(ctx, &request)
		assert.Nil(t, err)
		assert.NotNil(t, lock.Success, true)
	})

	t.Run("unlock", func(t *testing.T) {
		request := runtimev1pb.UnlockRequest{
			ResourceId: "lock_test",
			LockOwner:  "layotto1",
		}
		unlock, err := testClient.Unlock(ctx, &request)
		assert.Nil(t, err)
		assert.NotNil(t, unlock.Status, runtimev1pb.UnlockResponse_Status(lock.LOCK_BELONG_TO_OTHERS))
	})

	t.Run("unlock", func(t *testing.T) {
		request := runtimev1pb.UnlockRequest{
			ResourceId: "lock_test",
			LockOwner:  "layotto",
		}
		unlock, err := testClient.Unlock(ctx, &request)
		assert.Nil(t, err)
		assert.NotNil(t, unlock.Status, runtimev1pb.UnlockResponse_Status(lock.SUCCESS))
	})

	t.Run("unlock", func(t *testing.T) {
		request := runtimev1pb.UnlockRequest{
			ResourceId: "lock_test",
			LockOwner:  "layotto",
		}
		unlock, err := testClient.Unlock(ctx, &request)
		assert.Nil(t, err)
		assert.NotNil(t, unlock.Status, runtimev1pb.UnlockResponse_Status(lock.LOCK_UNEXIST))
	})
}
