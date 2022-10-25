package service

import "server.pluto.com/utils"

type SeverRecsource struct {
}

func (severRecsource *SeverRecsource) GetServerResources() (server *utils.Server) {
	var s utils.Server
	s.Os = utils.InitOs()
	s.Cpu, _ = utils.InitCpu()
	s.Disk, _ = utils.InitDisk()
	s.Ram, _ = utils.InitRAM()
	return &s
}
