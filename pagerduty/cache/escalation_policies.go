package cache

import (
	"sync"

	"github.com/petetanton/go-pagerduty-sdk/pagerduty/model"
)

type EscalationPolicyCache struct {
	m     sync.RWMutex
	users map[string]*model.EscalationPolicy
}

func NewEscalationPolicyCache() *EscalationPolicyCache {
	return &EscalationPolicyCache{
		m:     sync.RWMutex{},
		users: map[string]*model.EscalationPolicy{},
	}
}

func (uc *EscalationPolicyCache) ReadEscalationPolicy(id string) *model.EscalationPolicy {
	uc.m.RLock()
	defer uc.m.RUnlock()
	return uc.users[id]
}

func (uc *EscalationPolicyCache) WriteEscalationPolicy(policy *model.EscalationPolicy) {
	uc.m.Lock()
	defer uc.m.Unlock()
	uc.users[policy.Id] = policy
}

func (uc *EscalationPolicyCache) WriteEscalationPolicies(policies []*model.EscalationPolicy) {
	uc.m.Lock()
	defer uc.m.Unlock()
	for _, user := range policies {
		uc.users[user.Id] = user
	}
}

func (uc *EscalationPolicyCache) RemoveEscalationPolicy(id string) {
	uc.m.Lock()
	defer uc.m.Unlock()
	uc.users[id] = nil
}
