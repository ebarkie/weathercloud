// Copyright (c) 2017-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package weathercloud

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type observations interface {
	Clear()
	Values() map[string]string
}

// createRequest builds the HTTP request.
func (d Device) createRequest(obs ...observations) *http.Request {
	req, _ := http.NewRequest("GET", URL+"/set", nil)

	// Create mandatory query parameters.
	q := req.URL.Query()
	q.Add("wid", d.WID)
	q.Add("key", d.Key)
	if d.SoftwareType == 0 {
		q.Add("type", strconv.Itoa(int(softwareType)))
	} else {
		q.Add("type", strconv.Itoa(int(d.SoftwareType)))
	}
	if d.SoftwareVersion == "" {
		q.Add("version", "1.0")
	} else {
		q.Add("version", d.SoftwareVersion)
	}
	if !d.Time.IsZero() {
		q.Add("date", d.Time.In(time.UTC).Format("20060102"))
		q.Add("time", d.Time.In(time.UTC).Format("1504"))
	}

	// Add observations to query parameters.
	for _, o := range obs {
		for k, v := range o.Values() {
			q.Add(k, v)
		}
	}

	req.URL.RawQuery = q.Encode()

	return req
}

// Encode returns the request URL for the specified observations.  This
// is generally used for testing and debugging.
func (d Device) Encode(obs ...observations) string {
	return d.createRequest(obs...).URL.String()
}

// Upload uploads the specified observations.
func (d *Device) Upload(obs ...observations) (err error) {
	// Clear payload(s) after upload attempt.
	defer func() {
		for _, o := range obs {
			o.Clear()
		}
	}()

	// Initiate HTTP request.
	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(d.createRequest(obs...))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check HTTP return status code.
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("HTTP request returned non-OK status code %d", resp.StatusCode)
		return
	}

	return
}
