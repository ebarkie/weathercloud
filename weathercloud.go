// Copyright (c) 2017-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package weathercloud implements the Weathercloud device upload
// protocol.
package weathercloud

import "time"

// softwareType is assigned by Weathercloud.  Please consult them
// before hijacking it.
const softwareType = uint16(902)

// URL is the base API URL.
var URL = "http://api.weathercloud.net/v01"

// Device represents a weather device.
type Device struct {
	WID             string    // Device ID (assigned during creation)
	Key             string    // Device key (assigned during creation)
	SoftwareType    uint16    // Software type
	SoftwareVersion string    // Software version
	Time            time.Time // Upload time (default is now)
}
