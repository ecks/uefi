// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io> and contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package efivario

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/0x5a17ed/uefi/efi/efiguid"
)

var testGuid = efiguid.MustFromString("3cd99f3f-4b2b-43eb-ac29-f0890a4772b7")

func tempDir() string {
	dir := os.Getenv("TMPDIR")
	if dir == "" {
		// Use /var/tmp here since that's not a tmpfs
		// at least on my machine.
		dir = "/var/tmp"
	}
	return dir
}

type FsContextTestSuite struct {
	suite.Suite

	context *FsContext
	tmpDir  string
}

func (s *FsContextTestSuite) SetupTest() {
	dir, err := ioutil.TempDir(tempDir(), "uefi-test")
	require.NoError(s.T(), err)

	s.tmpDir = dir
	s.context = &FsContext{afero.NewBasePathFs(afero.NewOsFs(), dir)}
}

func (s *FsContextTestSuite) TearDownTest() {
	require.NoError(s.T(), os.RemoveAll(s.tmpDir))
}

// TestSetNewVariable tests setting a new variable.
func (s *FsContextTestSuite) TestSetNewVariable() {
	// given that ...
	_, err := s.context.fs.Stat("TestVar-3CD99F3F-4B2B-43EB-AC29-F0890A4772B7")
	require.True(s.T(), errors.Is(err, afero.ErrFileNotFound))

	// when ...
	err = s.context.Set("TestVar", testGuid, RuntimeAccess, []byte{0x00})
	require.NoError(s.T(), err)

	// then ...
	_, err = s.context.fs.Stat("TestVar-3CD99F3F-4B2B-43EB-AC29-F0890A4772B7")
	require.NoError(s.T(), err)
}

// TestSetExistingVariable tests changing an existing variable.
func (s *FsContextTestSuite) TestSetExistingVariable() {
	// given that ...
	f, err := s.context.fs.Create("TestVar-3CD99F3F-4B2B-43EB-AC29-F0890A4772B7")
	require.NoError(s.T(), err)
	defer func() { require.NoError(s.T(), f.Close()) }()

	// when ...
	err = s.context.Set("TestVar", testGuid, RuntimeAccess, []byte{0x00})
	require.NoError(s.T(), err)

	// then ...
	_, err = s.context.fs.Stat("TestVar-3CD99F3F-4B2B-43EB-AC29-F0890A4772B7")
	require.NoError(s.T(), err)
}

// TestDelete tests deleting an existing variable.
func (s *FsContextTestSuite) TestDelete() {
	// given that ...
	f, err := s.context.fs.Create("TestVar-3CD99F3F-4B2B-43EB-AC29-F0890A4772B7")
	require.NoError(s.T(), err)
	defer func() { require.NoError(s.T(), f.Close()) }()

	// when ...
	err = s.context.Delete("TestVar", testGuid)
	require.NoError(s.T(), err)

	// then ...
	_, err = s.context.fs.Stat("TestVar-3CD99F3F-4B2B-43EB-AC29-F0890A4772B7")
	require.True(s.T(), errors.Is(err, afero.ErrFileNotFound))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFsContextTestSuite(t *testing.T) {
	suite.Run(t, new(FsContextTestSuite))
}
