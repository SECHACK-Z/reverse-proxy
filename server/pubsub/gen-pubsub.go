// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pubsub

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

func (ps *__AccessPubSub) Unsub(subscribeID string) bool {
	if _, ok := ps.subs[subscribeID]; ok {
		delete(ps.subs, subscribeID)
		return true
	}
	return false
}

func (ps *__AccessPubSub) Pub(event Access) {
	for _, f := range ps.subs {
		go f(event)
	}
}

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

func (ps *__SystemPubSub) Unsub(subscribeID string) bool {
	if _, ok := ps.subs[subscribeID]; ok {
		delete(ps.subs, subscribeID)
		return true
	}
	return false
}

func (ps *__SystemPubSub) Pub(event System) {
	for _, f := range ps.subs {
		go f(event)
	}
}

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

func (ps *__UpdateConfigPubSub) Unsub(subscribeID string) bool {
	if _, ok := ps.subs[subscribeID]; ok {
		delete(ps.subs, subscribeID)
		return true
	}
	return false
}

func (ps *__UpdateConfigPubSub) Pub(event UpdateConfig) {
	for _, f := range ps.subs {
		go f(event)
	}
}

type __HealthCheckPubSub struct {
	subs map[string]func(HealthCheck)
	c    chan HealthCheck
}

var HealthCheckEvent = &__HealthCheckPubSub{
	subs: make(map[string]func(HealthCheck)),
	c:    make(chan HealthCheck, 10),
}

func (ps *__HealthCheckPubSub) Sub(f func(et HealthCheck)) string {
	subID := randomStr(5)
	for _, ok := ps.subs[subID]; ok; _, ok = ps.subs[subID] {
		subID = randomStr(5)
	}
	ps.subs[subID] = f
	return subID
}

func (ps *__HealthCheckPubSub) Unsub(subscribeID string) bool {
	if _, ok := ps.subs[subscribeID]; ok {
		delete(ps.subs, subscribeID)
		return true
	}
	return false
}

func (ps *__HealthCheckPubSub) Pub(event HealthCheck) {
	for _, f := range ps.subs {
		go f(event)
	}
}

type __GetWebhookPubSub struct {
	subs map[string]func(GetWebhook)
	c    chan GetWebhook
}

var GetWebhookEvent = &__GetWebhookPubSub{
	subs: make(map[string]func(GetWebhook)),
	c:    make(chan GetWebhook, 10),
}

func (ps *__GetWebhookPubSub) Sub(f func(et GetWebhook)) string {
	subID := randomStr(5)
	for _, ok := ps.subs[subID]; ok; _, ok = ps.subs[subID] {
		subID = randomStr(5)
	}
	ps.subs[subID] = f
	return subID
}

func (ps *__GetWebhookPubSub) Unsub(subscribeID string) bool {
	if _, ok := ps.subs[subscribeID]; ok {
		delete(ps.subs, subscribeID)
		return true
	}
	return false
}

func (ps *__GetWebhookPubSub) Pub(event GetWebhook) {
	for _, f := range ps.subs {
		go f(event)
	}
}

type __DeployPubSub struct {
	subs map[string]func(Deploy)
	c    chan Deploy
}

var DeployEvent = &__DeployPubSub{
	subs: make(map[string]func(Deploy)),
	c:    make(chan Deploy, 10),
}

func (ps *__DeployPubSub) Sub(f func(et Deploy)) string {
	subID := randomStr(5)
	for _, ok := ps.subs[subID]; ok; _, ok = ps.subs[subID] {
		subID = randomStr(5)
	}
	ps.subs[subID] = f
	return subID
}

func (ps *__DeployPubSub) Unsub(subscribeID string) bool {
	if _, ok := ps.subs[subscribeID]; ok {
		delete(ps.subs, subscribeID)
		return true
	}
	return false
}

func (ps *__DeployPubSub) Pub(event Deploy) {
	for _, f := range ps.subs {
		go f(event)
	}
}
