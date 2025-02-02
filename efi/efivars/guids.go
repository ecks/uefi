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

package efivars

import (
	"github.com/ecks/uefi/efi/efiguid"
)

var (
	GlobalVariable = efiguid.New(0x8be4df61, 0x93ca, 0x11d2, [8]byte{0xaa, 0x0d, 0x00, 0xe0, 0x98, 0x03, 0x2b, 0x8c})
)
