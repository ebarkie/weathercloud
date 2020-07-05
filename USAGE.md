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
func (d Device) Encode(obs ...query.Values) string
```
Encode returns the request URL for the specified observations. This is generally
used for testing and debugging.

#### func (*Device) Upload

```go
func (d *Device) Upload(obs ...query.Values) (err error)
```
Upload uploads the specified observations.

#### type Wx

```go
type Wx struct {
	query.Data
}
```

Wx represents weather observations.

#### func (*Wx) Bar

```go
func (w *Wx) Bar(in float64)
```
Bar is atmospheric pressure in inches.

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

#### func (*Wx) InHumidity

```go
func (w *Wx) InHumidity(p int)
```
InHumidity is the indoor humidity percentage (0-100).

#### func (*Wx) InTemp

```go
func (w *Wx) InTemp(f float64)
```
InTemp is the indoor temperature in degrees Fahrenheit.

#### func (*Wx) LeafWetness

```go
func (w *Wx) LeafWetness(p int)
```
LeafWetness is the leaf wetness index.

#### func (*Wx) OutHumidity

```go
func (w *Wx) OutHumidity(p int)
```
OutHumidity is the outdoor humidity percentage (0-100).

#### func (*Wx) OutTemp

```go
func (w *Wx) OutTemp(f float64)
```
OutTemp is outdoor temperature in degrees Fahrenheit.

#### func (*Wx) RainRate

```go
func (w *Wx) RainRate(in float64)
```
RainRate is rain over the past hour or the accumulated rainfall for the past 60
minutes in inches.

#### func (*Wx) SoilMoist

```go
func (w *Wx) SoilMoist(cb int)
```
SoilMoist is the soil moisture in centibars of tension.

#### func (*Wx) SolarRad

```go
func (w *Wx) SolarRad(wm2 int)
```
SolarRad is solar radiation in watts per square meter.

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

#### func (*Wx) WindDir

```go
func (w *Wx) WindDir(deg int)
```
WindDir is instantaneous wind direction in degrees (0-359).

#### func (*Wx) WindDirAvg

```go
func (w *Wx) WindDirAvg(deg int)
```
WindDirAvg is the 10 minute average wind direction in degrees (0-359).

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

#### func (*Wx) WindSpeedAvg

```go
func (w *Wx) WindSpeedAvg(mph float64)
```
WindSpeedAvg is the 10 minute average wind speed in miles per hour.
