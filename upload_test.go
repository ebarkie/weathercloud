// Copyright (c) 2017-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package weathercloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadEncode(t *testing.T) {
	a := assert.New(t)

	d := Device{WID: "0123", Key: "fedc"}

	wx := &Wx{}
	wx.Bar(29.86)
	a.Equal("http://api.weathercloud.net/v01/set?bar=10111&key=fedc&type=902&version=1.0&wid=0123", d.Encode(wx))

	wx.OutTemp(20)
	a.Equal("http://api.weathercloud.net/v01/set?bar=10111&key=fedc&temp=-66&type=902&version=1.0&wid=0123", d.Encode(wx))
}
