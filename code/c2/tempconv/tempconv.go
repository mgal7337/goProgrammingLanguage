//Package tempconv converts celcius to fahrenheit and opposite
package tempconv

import "fmt"

//Celcius is a representant of Celcius scale
type Celcius float64

//Fahrenheit is a representant of Fahrenheit scale
type Fahrenheit float64

//Kelvin is a representant of Kelvin scale
type Kelvin float64

const (
	//AbsoluteZeroC is the lowest possible value in Celcius
	AbsoluteZeroC Celcius = -273.15
	//FreezingC is the value when water freezes
	FreezingC Celcius = 0
	//BoilingC is the value when water boils
	BoilingC Celcius = 100
)

func (c Celcius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }

//CToF converts Celcius to Fahrenheit
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC converts Fahrenheit to Celcius
func FToC(f Fahrenheit) Celcius { return Celcius(f-32) * 5 / 9 }

//CToK converts Celcius To Kelvin
func CToK(c Celcius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

//KToC converts Kelvin To Celcius
func KToC(k Kelvin) Celcius { return Celcius(k + Kelvin(AbsoluteZeroC)) }
