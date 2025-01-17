package server

import (
	"net/http"

	"github.com/dfroberg/m3ufilter/config"
	"github.com/dfroberg/m3ufilter/logger"
	"github.com/dfroberg/m3ufilter/m3u"
	"github.com/mileusna/crontab"
)

var log = logger.Get()

func Serve(appConfig *config.Config) {
	conf := &httpState{
		playlists: &m3u.Streams{},
		lock:      false,
		appConfig: appConfig,
		crontab:   crontab.New(),
	}

	log.Info("Scheduling cronjob to periodically update playlist.")
	scheduleJob(conf, appConfig.Core.UpdateSchedule)

	log.Info("Parsing for the first time...")
	conf.crontab.RunAll()

	log.Info("Starting server")
	http.Handle("/playlist.m3u", httpHandler{conf, getPlaylist})
	http.Handle("/epg.xml", httpHandler{conf, getEpg})
	http.Handle("/update", httpHandler{conf, postUpdate})

	server := &http.Server{Addr: appConfig.Core.ServerListen}
	log.Fatal(server.ListenAndServe())
}

func scheduleJob(conf *httpState, schedule string) {
	conf.crontab.MustAddJob(schedule, func() {
		updatePlaylist(conf)
	})
}
