package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/ensomnatt/webfetch/internal/config"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/sirupsen/logrus"
)

type Server struct {
	r *http.ServeMux
	srv http.Server
	cfg *config.Config
}

type Data struct {
	Platform string
	Kernel string
	Uptime string
	CPUModel string
	CPUUsage string
	DiskUsage string
	RAMUsage string
}

func NewServer(addr string) *Server {
	r := http.NewServeMux()
	return &Server{
		r: r,
		srv: http.Server{
			Addr: addr,
			Handler: r,
		},
		cfg: config.NewConfig(),
	}
}

func (s *Server) Start() error {
	s.r.HandleFunc("/", s.PageHandler)
	s.r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(s.cfg.AppDir))))
	return s.srv.ListenAndServe()
}

func (s *Server) PageHandler(w http.ResponseWriter, r *http.Request) {
	hostInfoStat, err := host.Info()
	CPUUsage, err := cpu.Percent(time.Second, false)
	CPUInfoStat, err := cpu.Info()
	diskUsage, err := disk.Usage("/")
	RAMUsage, err := mem.VirtualMemory()
	if err != nil {
		logrus.Errorf("error with getting infoStat: %v", err)
	}

	uptimeHours := int(hostInfoStat.Uptime) / 3600
	uptimeMinutes := int(hostInfoStat.Uptime % 3600) / 60
	uptimeSeconds := int(hostInfoStat.Uptime) % 60

	data := Data{
		Platform: hostInfoStat.Platform,
		Kernel: hostInfoStat.KernelVersion,
		Uptime: fmt.Sprintf("%02d:%02d:%02d", uptimeHours, uptimeMinutes, uptimeSeconds),
		CPUModel: CPUInfoStat[0].ModelName,
		CPUUsage: fmt.Sprintf("%v%%", int(CPUUsage[0])),
		DiskUsage: fmt.Sprintf("%.2fGB/%.2fGB", float64(diskUsage.Used)/1e9, float64(diskUsage.Total)/1e9),
		RAMUsage: fmt.Sprintf("%.2fGB/%.2fGB", float64(RAMUsage.Used)/1e9, float64(RAMUsage.Total)/1e9),
	}

	template.Must(template.ParseFiles(filepath.Join(s.cfg.AppDir, "index.html"))).Execute(w, data)
}
