package security

import (
	//log "github.com/Sirupsen/logrus"

	api_operation "github.com/james-nesbitt/radi-api/operation"
)

/**
 * Authorization Rules
 */

// A single executable rule
type AuthorizeOperationRule interface {
	AuthorizeOperation(api_operation.Operation) RuleResult
}

// An ordered list of Rules
type AuthorizeOperationRules interface {
	Set(string, AuthorizeOperationRule)
	Get(string) (AuthorizeOperationRule, bool)
	Order() []string

	Merge(AuthorizeOperationRules)

	AuthorizeOperation(api_operation.Operation) RuleResult
}

// The response from a rule execution
type RuleResult interface {
	Rule() string    // Id of rule that triggered the result
	Message() string // Message about result that can be shown in UI
	Allow() bool     // Does the result explicitly allow the operation
	Deny() bool      // Does the result explicitly deny the opration
}

// Constructor for SimpleRuleResult
func New_SimpleRuleResult(id string, message string, status int) *SimpleRuleResult {
	return &SimpleRuleResult{
		id:      id,
		message: message,
		status:  status,
	}
}

// Simple integer based RuleResult
type SimpleRuleResult struct {
	id      string
	message string
	status  int
}

// Convert this into a RuleResult interface
func (result *SimpleRuleResult) RuleResult() RuleResult {
	return RuleResult(result)
}

// Does the result explicitly pass
func (result *SimpleRuleResult) Rule() string {
	return result.id
}

// Does the result explicitly pass
func (result *SimpleRuleResult) Message() string {
	return result.message
}

// Does the result explicitly pass
func (result *SimpleRuleResult) Allow() bool {
	return result.status > 0
}

// Does the result explicitly fail
func (result *SimpleRuleResult) Deny() bool {
	return result.status < 0
}

// An ordered set of rules
type SimpleAuthorizeOperationRules struct {
	rules map[string]AuthorizeOperationRule
	order []string
}

// Safe lazy initializer
func (rules *SimpleAuthorizeOperationRules) safe() {
	if rules.order == nil {
		rules.rules = map[string]AuthorizeOperationRule{}
		rules.order = []string{}
	}
}

// Merge one rules into another
func (rules *SimpleAuthorizeOperationRules) Merge(merge AuthorizeOperationRules) {
	rules.safe()
	for _, id := range merge.Order() {
		rule, _ := merge.Get(id)
		rules.Set(id, rule)
	}
}

// Add one rule
func (rules *SimpleAuthorizeOperationRules) Set(id string, setRule AuthorizeOperationRule) {
	rules.safe()
	if _, found := rules.rules[id]; !found {
		rules.order = append(rules.order, id)
	}
	rules.rules[id] = setRule
}

// Get a Rule by ID
func (rules *SimpleAuthorizeOperationRules) Get(id string) (AuthorizeOperationRule, bool) {
	rules.safe()
	rules.safe()
	rule, found := rules.rules[id]
	return rule, found
}

// Get an ordered list of Rule ids
func (rules *SimpleAuthorizeOperationRules) Order() []string {
	rules.safe()
	rules.safe()
	return rules.order
}

// Authorize operations by returning the result from the first rule that explicitly allows or denies the operation
func (rules *SimpleAuthorizeOperationRules) AuthorizeOperation(op api_operation.Operation) RuleResult {
	rules.safe()

	//log.WithFields(log.Fields{"op": op.Id(), "count": len(rules.rules), "rules.Order()": rules.Order()}).Info("SimpleAuthorizeOperationRules: Auth op")
	for _, id := range rules.Order() {
		rule, _ := rules.Get(id)

		//log.WithFields(log.Fields{"id": id, "rule": rule}).Info("SimpleAuthorizeOperationRules: Checking Rule")
		result := rule.AuthorizeOperation(op)

		// return any result that explicitly allows or denies the operation
		if result.Allow() || result.Deny() {
			return result
		}
	}
	return New_SimpleRuleResult("security.norulematch", "no rule matched", 0).RuleResult()
}
