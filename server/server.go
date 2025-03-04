package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ensomnatt/webfetch/config"
	"github.com/ensomnatt/webfetch/sysinfo"
)

type Server struct {
	r   *http.ServeMux
	srv http.Server
  cfg *config.Config
}

type Data struct {
	FastFetch []sysinfo.KeyValue
}

func NewServer(addr string) *Server {
	r := http.NewServeMux()
	return &Server{
		r: r,
		srv: http.Server{
			Addr:    addr,
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
	tmpl, err := template.ParseFiles(filepath.Join(s.cfg.AppDir, "index.html"))
	if err != nil {
		_ = fmt.Errorf("error with parsing files: %v", err)
		http.Error(w, "error with parsing files", http.StatusInternalServerError)
		return
	}

	data := Data{
		FastFetch: sysinfo.GetSystemInfo(),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		_ = fmt.Errorf("error with executing template: %v", err)
		http.Error(w, "error with executing template", http.StatusInternalServerError)
		return
	}
}
