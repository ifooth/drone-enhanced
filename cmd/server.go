package main

import (
	"net/http"

	pluginConfig "github.com/drone/drone-go/plugin/config"
	"github.com/ifooth/drone-ci-enhanced/plugin/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func ServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "serve extension api",
		Long:  `serve extension api for drone ci`,
	}

	conf := &serverConfig{}
	conf.registerFlag(cmd)

	cmd.Run = func(cmd *cobra.Command, args []string) {
		runServerCmd(conf)
	}
	return cmd
}

func runServerCmd(conf *serverConfig) {
	spec := new(envSpec)
	if err := envconfig.Process("", spec); err != nil {
		logrus.Fatal(err)
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if spec.Secret == "" {
		logrus.Fatalln("missing secret key")
	}

	handler := pluginConfig.Handler(config.NewConfigPlugin(), spec.Secret, logrus.StandardLogger())

	logrus.Infof("server listening on address %s", conf.httpAddress)

	http.Handle("/api/v1/plugin/config", handler)
	logrus.Fatal(http.ListenAndServe(conf.httpAddress, nil))
}

type envSpec struct {
	Debug  bool   `envconfig:"PLUGIN_DEBUG"`
	Secret string `envconfig:"PLUGIN_SECRET"`
}

type serverConfig struct {
	httpAddress string
}

func (s *serverConfig) registerFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&s.httpAddress, "http-address", "127.0.0.1:8080", "Listen host:port for HTTP endpoints.")
}
