// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pubsub

// TODO: Close処理
type __AccessPubSub struct {
	subs map[string]func(Access)
	c    chan Access
}

var AccessEvent = &__AccessPubSub{
	subs: make(map[string]func(Access)),
	c:    make(chan Access, 10),
}

func (ps *__AccessPubSub) Sub(f func(et Access)) string {
	subID := randomStr(5)
	for _, ok := ps.subs[subID]; ok; _, ok = ps.subs[subID] {
		subID = randomStr(5)
	}
	ps.subs[subID] = f
	return subID
}

func (ps *__AccessPubSub) Pub(event Access) {
	for _, f := range ps.subs {
		go f(event)
	}
}

// TODO: Close処理
type __SystemPubSub struct {
	subs map[string]func(System)
	c    chan System
}

var SystemEvent = &__SystemPubSub{
	subs: make(map[string]func(System)),
	c:    make(chan System, 10),
}

func (ps *__SystemPubSub) Sub(f func(et System)) string {
	subID := randomStr(5)
	for _, ok := ps.subs[subID]; ok; _, ok = ps.subs[subID] {
		subID = randomStr(5)
	}
	ps.subs[subID] = f
	return subID
}

func (ps *__SystemPubSub) Pub(event System) {
	for _, f := range ps.subs {
		go f(event)
	}
}

// TODO: Close処理
type __UpdateConfigPubSub struct {
	subs map[string]func(UpdateConfig)
	c    chan UpdateConfig
}

var UpdateConfigEvent = &__UpdateConfigPubSub{
	subs: make(map[string]func(UpdateConfig)),
	c:    make(chan UpdateConfig, 10),
}

func (ps *__UpdateConfigPubSub) Sub(f func(et UpdateConfig)) string {
	subID := randomStr(5)
	for _, ok := ps.subs[subID]; ok; _, ok = ps.subs[subID] {
		subID = randomStr(5)
	}
	ps.subs[subID] = f
	return subID
}

func (ps *__UpdateConfigPubSub) Pub(event UpdateConfig) {
	for _, f := range ps.subs {
		go f(event)
	}
}