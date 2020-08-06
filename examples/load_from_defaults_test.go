package tests

import (
	"github.com/Sayed-Soroush-Hashemi/GoConfMan/pkg/goconfman"
	"math"
	"testing"
)

func TestLoadFromDefaults(t *testing.T) {
	g := GlobalConfig{}
	goconfman.LoadFromDefaults(&g)

	if g.IntegerValue != 42 {
		t.Errorf("g.IntegerValue = %d != 42", g.IntegerValue)
	}
	if math.Abs(float64(g.FloatValue-3.14)) > 1e-8 {
		t.Errorf("g.FloatValue = %f != 3.14", g.FloatValue)
	}
	if g.StringValue != "in global config" {
		t.Errorf("g.StringValue = %s != \"in global config\"", g.StringValue)
	}

	if g.LocalConfig.IntegerValue != 420 {
		t.Errorf("g.LocalConfig.IntegerValue = %d != 420", g.LocalConfig.IntegerValue)
	}
	if math.Abs(float64(g.LocalConfig.FloatValue-31.4)) > 1e-8 {
		t.Errorf("g.LocalConfig.FloatValue = %f != 31.4", g.LocalConfig.FloatValue)
	}
	if g.LocalConfig.StringValue != "in local config" {
		t.Errorf("g.LocalConfig.StringValue = %s != \"in local config\"", g.LocalConfig.StringValue)
	}

	if g.NonGoConfManConfig.LocalConfig.IntegerValue != 420 {
		t.Errorf("g.NonGoConfManConfig.LocalConfig.IntegerValue = %d != 420", g.NonGoConfManConfig.LocalConfig.IntegerValue)
	}
	if math.Abs(float64(g.NonGoConfManConfig.LocalConfig.FloatValue-31.4)) > 1e-8 {
		t.Errorf("g.NonGoConfManConfig.LocalConfig.FloatValue = %f != 31.4", g.NonGoConfManConfig.LocalConfig.FloatValue)
	}
	if g.NonGoConfManConfig.LocalConfig.StringValue != "in local config" {
		t.Errorf("g.NonGoConfManConfig.LocalConfig.StringValue = %s != \"in local config\"", g.NonGoConfManConfig.LocalConfig.StringValue)
	}
}