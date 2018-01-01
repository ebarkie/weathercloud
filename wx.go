// Copyright (c) 2017-2018 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package weathercloud

import (
	"github.com/ebarkie/request"
	"github.com/ebarkie/weatherlink/units"
)

// Wx represents weather observations.
type Wx struct {
	request.Data
}

func tens(f float64) int {
	return int(f * 10.0)
}

// Bar is atmospheric pressure in inches.
func (w *Wx) Bar(in float64) {
	w.SetInt("bar", tens(units.Pressure(in*units.Inches).Hectopascals()))
}

// DailyRain is rain so far today (local time) in inches.
func (w *Wx) DailyRain(in float64) {
	w.SetInt("rain", tens(units.Length(in*units.Inches).Millimeters()))
}

// DailyET is the evapotranspiration so far today (local time) in
// inches.
func (w *Wx) DailyET(in float64) {
	w.SetInt("et", tens(units.Length(in*units.Inches).Millimeters()))
}

// DewPoint is the outdoor dew point in degrees Fahrenheit.
func (w *Wx) DewPoint(f float64) {
	w.SetInt("dew", tens(units.Fahrenheit(f).Celsius()))
}

// HeatIndex is the outdoor heat index in degrees Fahrenheit.
func (w *Wx) HeatIndex(f float64) {
	w.SetInt("heat", tens(units.Fahrenheit(f).Celsius()))
}

// InHumidity is the indoor humidity percentage (0-100).
func (w *Wx) InHumidity(p int) {
	w.SetInt("humin", p)
}

// InTemp is the indoor temperature in degrees Fahrenheit.
func (w *Wx) InTemp(f float64) {
	w.SetInt("tempin", tens(units.Fahrenheit(f).Celsius()))
}

// LeafWetness is the leaf wetness index.
func (w *Wx) LeafWetness(p int) {
	w.SetIntf(request.Indexed{Format: "leafwet#", Begin: 1, Width: 2}, p)
}

// OutHumidity is the outdoor humidity percentage (0-100).
func (w *Wx) OutHumidity(p int) {
	w.SetIntf(request.Indexed{Format: "hum#", Begin: 1, Zero: 1, Width: 2}, p)
}

// OutTemp is outdoor temperature in degrees Fahrenheit.
func (w *Wx) OutTemp(f float64) {
	w.SetIntf(request.Indexed{Format: "temp#", Begin: 1, Zero: 1, Width: 2}, tens(units.Fahrenheit(f).Celsius()))
}

// RainRate is rain over the past hour or the accumulated
// rainfall for the past 60 minutes in inches.
func (w *Wx) RainRate(in float64) {
	w.SetInt("rainrate", tens(units.Length(in*units.Inches).Millimeters()))
}

// SolarRad is solar radiation in watts per square meter.
func (w *Wx) SolarRad(wm2 int) {
	w.SetInt("solarrad", wm2*10)
}

// SoilMoist is the soil moisture in centibars of tension.
func (w *Wx) SoilMoist(cb int) {
	w.SetIntf(request.Indexed{Format: "soilmoist#", Begin: 1, Width: 2}, cb)
}

// UVIndex is the UltraViolet light index.
func (w *Wx) UVIndex(i float64) {
	w.SetInt("uvi", tens(i))
}

// WindChill is the wind chill in degrees Fahrenheit.
func (w *Wx) WindChill(f float64) {
	w.SetInt("chill", tens(units.Fahrenheit(f).Celsius()))
}

// WindDir is instantaneous wind direction in degrees (0-359).
func (w *Wx) WindDir(deg int) {
	w.SetIntf(request.Indexed{Format: "wdir#", Begin: 1, Zero: 1, Width: 2}, deg)
}

// WindDirAvg is the 10 minute average wind direction in
// degrees (0-359).
func (w *Wx) WindDirAvg(deg int) {
	w.SetIntf(request.Indexed{Format: "wdiravg#", Begin: 1, Zero: 1, Width: 2}, deg)
}

// WindGustSpeed is the 10 minute wind gust speed in miles per hour.
func (w *Wx) WindGustSpeed(mph float64) {
	w.SetIntf(request.Indexed{Format: "wspdhi#", Begin: 1, Zero: 1, Width: 2}, tens(units.Speed(mph*units.MPH).MPS()))
}

// WindSpeed is the instantaneous wind speed in miles per hour.
func (w *Wx) WindSpeed(mph float64) {
	w.SetIntf(request.Indexed{Format: "wspd#", Begin: 1, Zero: 1, Width: 2}, tens(units.Speed(mph*units.MPH).MPS()))
}

// WindSpeedAvg is the 10 minute average wind speed in miles
// per hour.
func (w *Wx) WindSpeedAvg(mph float64) {
	w.SetIntf(request.Indexed{Format: "wspdavg#", Begin: 1, Zero: 1, Width: 2}, tens(units.Speed(mph*units.MPH).MPS()))
}
