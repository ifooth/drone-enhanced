package main

import (
	"net/http"

	pluginConfig "github.com/drone/drone-go/plugin/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ifooth/drone-enhanced/plugin/config"
	"github.com/ifooth/drone-enhanced/providers"
)

func ServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Serve drone's extension api",
		Long:  `Serve extension api for drone ci`,
	}

	conf := &serverConfig{}
	conf.registerFlag(cmd)

	cmd.Run = func(cmd *cobra.Command, args []string) {
		runServerCmd(conf)
	}
	return cmd
}

func runServerCmd(conf *serverConfig) {
	logFormatter := new(logrus.TextFormatter)
	logFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logFormatter.FullTimestamp = true
	logrus.SetFormatter(logFormatter)

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

	if spec.ProviderToken == "" {
		logrus.Fatalln("missing PROVIDER_TOKEN")
	}

	var provider providers.Provider
	var err error

	switch spec.Provider {
	case "GITEA":
		if spec.ProviderServer == "" {
			logrus.Fatalln("missing PROVIDER_SERVER")
		}

		cred := &providers.GiteaCredential{Server: spec.ProviderServer, Token: spec.ProviderToken, Debug: spec.Debug}
		provider, err = providers.NewGiteaClient(cred)
		if err != nil {
			logrus.Fatal(err)
		}
	default:
		logrus.Fatalln("missing PROVIDER or not support")
	}

	configPluginHandler := pluginConfig.Handler(config.NewConfigPlugin(provider), spec.Secret, logrus.StandardLogger())

	logrus.Infof("server listening on address %s", conf.httpAddress)

	http.Handle("/api/v1/plugin/config", configPluginHandler)
	logrus.Fatal(http.ListenAndServe(conf.httpAddress, nil))
}

type envSpec struct {
	Debug          bool   `envconfig:"PLUGIN_DEBUG"`
	Secret         string `envconfig:"PLUGIN_SECRET"`
	Provider       string `envconfig:"PROVIDER"`
	ProviderServer string `envconfig:"PROVIDER_SERVER"`
	ProviderToken  string `envconfig:"PROVIDER_TOKEN"`
}

type serverConfig struct {
	httpAddress string
}

func (s *serverConfig) registerFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&s.httpAddress, "http-address", "127.0.0.1:8080", "Listen host:port for HTTP endpoints.")
}
