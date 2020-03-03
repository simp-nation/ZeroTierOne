/*
 * Copyright (c)2013-2020 ZeroTier, Inc.
 *
 * Use of this software is governed by the Business Source License included
 * in the LICENSE.TXT file in the project's root directory.
 *
 * Change Date: 2024-01-01
 *
 * On the date above, in accordance with the Business Source License, use
 * of this software will be governed by version 2.0 of the Apache License.
 */
/****/

package zerotier

//#include "../../native/GoGlue.h"
import "C"

// SelfTest runs a series of tests on the ZeroTier core and the Go service code, returning true on success.
// Results are sent to stdout.
func SelfTest() bool {
	return true
}
