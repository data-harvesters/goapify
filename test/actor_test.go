package test

import (
	"testing"

	"github.com/data-harvesters/goapify"
)

func TestActor(t *testing.T) {
	a := goapify.NewActor()
	_ = a

	a.ProxyConfiguration.Proxy(goapify.UseRandomProxy())
}
