# weathercloud

```go
import "github.com/ebarkie/weathercloud"
```

Package weathercloud implements the Weathercloud device upload protocol.

## Usage

```go
var URL = "http://api.weathercloud.net/v01"
```
URL is the base API URL.

#### type Device

```go
type Device struct {
	WID             string    // Device ID (assigned during creation)
	Key             string    // Device key (assigned during creation)
	SoftwareType    uint16    // Software type
	SoftwareVersion string    // Software version
	Time            time.Time // Upload time (default is now)
}
```

Device represents a weather device.

#### func (Device) Encode

```go
func (d Device) Encode(obs ...observations) string
```
Encode returns the request URL for the specified observations. This is generally
used for testing and debugging.

#### func (*Device) Upload

```go
func (d *Device) Upload(obs ...observations) (err error)
```
Upload uploads the specified observations.

#### type Wx

```go
type Wx struct {
	request.Data
}
```

Wx represents weather observations.

#### func (*Wx) Barometer

```go
func (w *Wx) Barometer(in float64)
```
Barometer is atmospheric pressure in inches.

#### func (*Wx) DailyET

```go
func (w *Wx) DailyET(in float64)
```
DailyET is the evapotranspiration so far today (local time) in inches.

#### func (*Wx) DailyRain

```go
func (w *Wx) DailyRain(in float64)
```
DailyRain is rain so far today (local time) in inches.

#### func (*Wx) DewPoint

```go
func (w *Wx) DewPoint(f float64)
```
DewPoint is the outdoor dew point in degrees Fahrenheit.

#### func (*Wx) HeatIndex

```go
func (w *Wx) HeatIndex(f float64)
```
HeatIndex is the outdoor heat index in degrees Fahrenheit.

#### func (*Wx) IndoorHumidity

```go
func (w *Wx) IndoorHumidity(p int)
```
IndoorHumidity is the indoor humidity percentage (0-100).

#### func (*Wx) IndoorTemperature

```go
func (w *Wx) IndoorTemperature(f float64)
```
IndoorTemperature is the indoor temperature in degrees Fahrenheit.

#### func (*Wx) LeafWetness

```go
func (w *Wx) LeafWetness(p int)
```
LeafWetness is the leaf wetness index.

#### func (*Wx) OutdoorHumidity

```go
func (w *Wx) OutdoorHumidity(p int)
```
OutdoorHumidity is the outdoor humidity percentage (0-100).

#### func (*Wx) OutdoorTemperature

```go
func (w *Wx) OutdoorTemperature(f float64)
```
OutdoorTemperature is outdoor temperature in degrees Fahrenheit.

#### func (*Wx) RainRate

```go
func (w *Wx) RainRate(in float64)
```
RainRate is rain over the past hour or the accumulated rainfall for the past 60
minutes in inches.

#### func (*Wx) SoilMoisture

```go
func (w *Wx) SoilMoisture(cb int)
```
SoilMoisture is the soil moisture in centibars of tension.

#### func (*Wx) SolarRadiation

```go
func (w *Wx) SolarRadiation(wm2 int)
```
SolarRadiation is solar radiation in watts per square meter.

#### func (*Wx) UVIndex

```go
func (w *Wx) UVIndex(i float64)
```
UVIndex is the UltraViolet light index.

#### func (*Wx) WindChill

```go
func (w *Wx) WindChill(f float64)
```
WindChill is the wind chill in degrees Fahrenheit.

#### func (*Wx) WindDirection

```go
func (w *Wx) WindDirection(deg int)
```
WindDirection is instantaneous wind direction in degrees (0-359).

#### func (*Wx) WindDirectionAverage

```go
func (w *Wx) WindDirectionAverage(deg int)
```
WindDirectionAverage is the 10 minute average wind direction in degrees (0-359).

#### func (*Wx) WindGustSpeed

```go
func (w *Wx) WindGustSpeed(mph float64)
```
WindGustSpeed is the 10 minute wind gust speed in miles per hour.

#### func (*Wx) WindSpeed

```go
func (w *Wx) WindSpeed(mph float64)
```
WindSpeed is the instantaneous wind speed in miles per hour.

#### func (*Wx) WindSpeedAverage

```go
func (w *Wx) WindSpeedAverage(mph float64)
```
WindSpeedAverage is the 10 minute average wind speed in miles per hour.
