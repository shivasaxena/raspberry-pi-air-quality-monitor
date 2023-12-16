package PMS7003

import "fmt"

type PMS7003SensorValue struct {
	// StartByte1, StartByte2                                                                    byte
	Framelen                                                                                  uint16 // :2
	PM10CF10Standard, PM25CF10Standard, PM100CF10Standard                                     uint16 // 2:4, 4:6, 6:8
	PM10Atmospheric, PM25Atmospheric, PM100Atmospheric                                        uint16
	ParticlesGT03, ParticlesGT05, ParticlesGT10, ParticlesGT25, ParticlesGT50, ParticlesGT100 uint16
	Unused                                                                                    uint16
	Checksum                                                                                  uint16
}

func (t PMS7003SensorValue) String() string {
	return fmt.Sprintf(
		"[PM10 CF10 Standard:%d, PM25 CF10 Standard:%d, PM100 CF10 Standard:%d, "+
			"PM10 Atmospheric:%d, PM25 Atmospheric:%d, PM100 Atmospheric:%d, "+
			"Particles GT 03:%d,Particles GT 05:%d, Particles GT 10:%d, Particles GT 25:%d, Particles GT 50:%d, Particles GT 100:%d]",
		t.PM10CF10Standard, t.PM25CF10Standard, t.PM100CF10Standard,
		t.PM10Atmospheric, t.PM25Atmospheric, t.PM100Atmospheric,
		t.ParticlesGT03, t.ParticlesGT05, t.ParticlesGT10, t.ParticlesGT25, t.ParticlesGT50, t.ParticlesGT100)
}
