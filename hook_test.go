package logrus_stackdriver

import (
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/google-api-go-wrapper/config"
)

const hooksTestLemgth = 1000

func newHookForTesting() (*StackdriverHook, error) {
	return NewWithConfig("aaaa", "bbbb", config.Config{
		Email:      "foo@example.com",
		PrivateKey: "YOUR_KEY",
	})

}

func TestStackdriverHook_Fire(t *testing.T) {
	hooks, err := newHookForTesting()
	fmt.Printf("%+v", hooks)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
	}
	entryList := make([]logrus.Entry, 0, hooksTestLemgth)
	for i := 0; i < hooksTestLemgth; i++ {
		entryList = append(entryList, *logrus.NewEntry(nil))
	}
	for _, e := range entryList {
		go func(h *StackdriverHook, ee logrus.Entry) { h.Fire(&ee) }(hooks, e)
	}
}

func TestStackdriverHook_SetLevels(t *testing.T) {
	hooks, err := newHookForTesting()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < hooksTestLemgth; i++ {
		go func(h *StackdriverHook) { hooks.SetLevels(defaultLevels) }(hooks)
	}
}

func TestStackdriverHook_AddFilter(t *testing.T) {
	hooks, err := newHookForTesting()
	if err != nil {
		fmt.Println(err)
	}
	filterFunc := func(x interface{}) interface{} { return x }
	for i := 0; i < hooksTestLemgth; i++ {
		go func(h *StackdriverHook) { h.AddFilter("filter", filterFunc) }(hooks)
	}
}

func TestStackdriverHook_AddIgnore(t *testing.T) {
	hooks, err := newHookForTesting()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < hooksTestLemgth; i++ {
		go func(h *StackdriverHook) { h.AddIgnore("ignore") }(hooks)
	}
}

func TestStackdriverHook_SetLabels(t *testing.T) {
	hooks, err := newHookForTesting()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < hooksTestLemgth; i++ {
		go func(h *StackdriverHook) { hooks.SetLabels(map[string]string{"label1": "label1"}) }(hooks)
	}
}

func TestStackdriverHook_Async(t *testing.T) {
	hooks, err := newHookForTesting()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < hooksTestLemgth; i++ {
		go func(h *StackdriverHook) { h.Async() }(hooks)
	}
}
