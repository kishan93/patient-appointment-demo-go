package app_test

import (
	"patient-appointment-demo-go/internal/app"
	"reflect"
	"testing"
)

func TestConfigWithPort(t *testing.T) {
	c := app.ConfigWithPort(8000)

    if reflect.TypeOf(c) != reflect.TypeOf(app.AppConfig{}) {
        t.Errorf("app.ConfigWithPort did not return correct type")
    }

    if c.Port != 8000 {
        t.Errorf("found incorrect value for Port on app config")
    }

}

func TestNewApp(t *testing.T) {
    a := app.New(app.AppConfig{})


    if reflect.TypeOf(a) != reflect.TypeOf(app.App{}) {
        t.Errorf("app.New did not return correct type")
    }
}
