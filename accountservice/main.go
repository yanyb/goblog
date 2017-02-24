package main

import (
        "github.com/callistaenterprise/goblog/accountservice/service"
        "github.com/callistaenterprise/goblog/accountservice/dbclient"
        "flag"
        "github.com/spf13/viper"
        "github.com/callistaenterprise/goblog/accountservice/config"
        "github.com/callistaenterprise/goblog/accountservice/messaging"
        "os/signal"
        "os"
        "syscall"
        log "github.com/Sirupsen/logrus"
        "bytes"
        "fmt"
)

var appName = "accountservice"

type PlainFormatter struct {

}
func (f PlainFormatter) Format(e *log.Entry) ([]byte, error) {
        var b *bytes.Buffer
        if e.Buffer != nil {
                b = e.Buffer
        } else {
                b = &bytes.Buffer{}
        }
        fmt.Fprintf(b, "%s", e.Message)
        b.WriteByte('\n')
        return b.Bytes(), nil
}

func initLogger() {
        if viper.GetString("profile") != "dev" {
                log.SetFormatter(&PlainFormatter{})
        }
}

func main() {
        log.Info("Starting %v\n", appName)
        parseFlags()
        initLogger()
        config.LoadConfiguration(viper.GetString("configServerUrl"), appName, viper.GetString("profile"))
        initializeBoltClient()
        initializeMessaging()
        handleSigterm(func() {

        })
        service.StartWebServer(viper.GetString("server_port"))
}

func parseFlags() {
        profile := flag.String("profile", "test", "Environment profile, something similar to spring profiles")
        configServerUrl := flag.String("configServerUrl", "http://configserver:8888", "Address to config server")

        flag.Parse()
        viper.Set("profile", *profile)
        viper.Set("configServerUrl", *configServerUrl)
}

func initializeMessaging() {
        if !viper.IsSet("broker_url") {
                panic("No 'broker_url' set in configuration, cannot start")
        }
        service.MessagingClient = &messaging.MessagingClient{}
        service.MessagingClient.ConnectToBroker(viper.GetString("broker_url"))
}

func initializeBoltClient() {
        service.DBClient = &dbclient.BoltClient{}
        service.DBClient.OpenBoltDb()
        service.DBClient.Seed()
}

// Handles Ctrl+C or most other means of "controlled" shutdown gracefully. Invokes the supplied func before exiting.
func handleSigterm(handleExit func()) {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        signal.Notify(c, syscall.SIGTERM)
        go func() {
                <-c
                handleExit()
                os.Exit(1)
        }()
}
