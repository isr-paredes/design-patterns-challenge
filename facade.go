package main
import (
    "fmt"
)
//Ligth
type LightInt interface {
	TurnOn()
	TurnOff()
}
type LightS struct{}
func(l *LightS) TurnOn(){
    fmt.Println("Light ON")
}
func(l *LightS) TurnOff(){
    fmt.Println("Light OFF")
}
//Thermostat
type ThermostatInt interface {
	SetTemperature(temp int)
	GetTemperature() int
}
type ThermostatS struct{
    temp int
}
func (t *ThermostatS) GetTemperature() int{
    fmt.Printf("Temperature is %d degrees \n",t.temp)
    return t.temp
}
func (t *ThermostatS) SetTemperature(temp int){
    t.temp=temp
    fmt.Printf("Temperature set to %d degrees \n", temp)
}
//SecurityCamera
type SecurityCameraInt interface {
    StartRecording()
	StopRecording()
}
type SecurityCameraS struct{}
func (s *SecurityCameraS) StartRecording(){
    fmt.Println("Is recording")
}
func (s *SecurityCameraS) StopRecording(){
    fmt.Println("Is NOT recording")
}


// Facade
type SmartHomeFacade struct {
	light        LightInt
	thermostat   ThermostatInt
	securityCamera SecurityCameraInt


}
func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		light:        &LightS{},
		thermostat:   &ThermostatS{},
		securityCamera: &SecurityCameraS{},
	}
}
// Facade functions 
func (f *SmartHomeFacade) TurnOnLights() {
	f.light.TurnOn()
}

func (f *SmartHomeFacade) TurnOffLights() {
	f.light.TurnOff()
}

func (f *SmartHomeFacade) SetTemperature(temp int) {
	f.thermostat.SetTemperature(temp)
}

func (f *SmartHomeFacade) GetTemperature() int {
	return f.thermostat.GetTemperature()
}

func (f *SmartHomeFacade) StartSecurityCamera() {
	f.securityCamera.StartRecording()
}

func (f *SmartHomeFacade) StopSecurityCamera() {
	f.securityCamera.StopRecording()
}

func main() {
    facade := NewSmartHomeFacade()
    facade.TurnOnLights()
	facade.SetTemperature(22)
	facade.StartSecurityCamera()

	facade.TurnOffLights()
	facade.StopSecurityCamera()
    facade.GetTemperature()
}