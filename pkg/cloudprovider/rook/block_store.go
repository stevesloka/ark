/*
Copyright 2017 Heptio Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rook

import (
	"github.com/pkg/errors"

	"github.com/heptio/ark/pkg/cloudprovider"
)

const rookapiKey = "rookaApiUrl"

type blockStore struct {
}

func NewBlockStore() cloudprovider.BlockStore {
	return &blockStore{}
}

func (b *blockStore) Init(config map[string]string) error {
	region := config[rookapiKey]
	if region == "" {
		return errors.Errorf("missing %s in rook configuration", rookapiKey)
	}

	return nil
}

func (b *blockStore) CreateVolumeFromSnapshot(snapshotID, volumeType, volumeAZ string, iops *int64) (volumeID string, err error) {
	return "", nil
}

func (b *blockStore) GetVolumeInfo(volumeID, volumeAZ string) (string, *int64, error) {
	return "", nil, nil
}

func (b *blockStore) IsVolumeReady(poolName, volumeID string) (ready bool, err error) {
	return false, nil
}

func (b *blockStore) ListSnapshots(tagFilters map[string]string) ([]string, error) {
	return []string{}, nil
}

func (b *blockStore) CreateSnapshot(volumeID, volumeAZ string, tags map[string]string) (snapshotID string, err error) {
	return "", nil
}

func (b *blockStore) DeleteSnapshot(snapshotID string) error {
	return nil
}
