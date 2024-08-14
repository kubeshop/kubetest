package lite

import (
	"encoding/json"
	"fmt"
)

type ActionResult struct {
	Ref   string `json:"r"`
	Value string `json:"v"`
}

type ActionDeclare struct {
	Condition string   `json:"c"`
	Ref       string   `json:"r"`
	Parents   []string `json:"p,omitempty"`
}

type ActionExecute struct {
	Ref      string `json:"r"`
	Negative bool   `json:"n,omitempty"`
	Toolkit  bool   `json:"t,omitempty"`
	Pure     bool   `json:"p,omitempty"`
}

type ActionPause struct {
	Ref string `json:"r"`
}

type ActionTimeout struct {
	Ref     string `json:"r"`
	Timeout string `json:"t"`
}

type ActionRetry struct {
	Ref   string `json:"r"`
	Count int32  `json:"c,omitempty"`
	Until string `json:"u,omitempty"`
}

type ActionSetup struct {
	CopyInit     bool `json:"i,omitempty"`
	CopyToolkit  bool `json:"t,omitempty"`
	CopyBinaries bool `json:"b,omitempty"`
}

type ActionType string

const (
	// Declarations
	ActionTypeDeclare ActionType = "declare"
	ActionTypePause   ActionType = "pause"
	ActionTypeResult  ActionType = "result"
	ActionTypeTimeout ActionType = "timeout"
	ActionTypeRetry   ActionType = "retry"

	// Operations
	ActionTypeContainerTransition ActionType = "container"
	ActionTypeCurrentStatus       ActionType = "status"
	ActionTypeStart               ActionType = "start"
	ActionTypeEnd                 ActionType = "end"
	ActionTypeSetup               ActionType = "setup"
	ActionTypeExecute             ActionType = "execute"
)

type LiteContainerConfig struct {
	Command    *[]string `json:"command,omitempty"`
	Args       *[]string `json:"args,omitempty"`
	WorkingDir *string   `json:"workingDir,omitempty"`
}

type LiteActionContainer struct {
	Config LiteContainerConfig `json:"c"`
}

// LiteAction is lightweight version of Action,
// that is intended to use directly in the Init Process.
// It's not including original ContainerConfig,
// as it requires additional 40MB of structs in the binary.
type LiteAction struct {
	CurrentStatus *string              `json:"s,omitempty"`
	Start         *string              `json:"S,omitempty"`
	End           *string              `json:"E,omitempty"`
	Setup         *ActionSetup         `json:"_,omitempty"`
	Declare       *ActionDeclare       `json:"d,omitempty"`
	Result        *ActionResult        `json:"r,omitempty"`
	Container     *LiteActionContainer `json:"c,omitempty"`
	Execute       *ActionExecute       `json:"e,omitempty"`
	Timeout       *ActionTimeout       `json:"t,omitempty"`
	Pause         *ActionPause         `json:"p,omitempty"`
	Retry         *ActionRetry         `json:"R,omitempty"`
}

func (a *LiteAction) Type() ActionType {
	if a.Declare != nil {
		return ActionTypeDeclare
	} else if a.Pause != nil {
		return ActionTypePause
	} else if a.Result != nil {
		return ActionTypeResult
	} else if a.Timeout != nil {
		return ActionTypeTimeout
	} else if a.Retry != nil {
		return ActionTypeRetry
	} else if a.Container != nil {
		return ActionTypeContainerTransition
	} else if a.CurrentStatus != nil {
		return ActionTypeCurrentStatus
	} else if a.Start != nil {
		return ActionTypeStart
	} else if a.End != nil {
		return ActionTypeEnd
	} else if a.Setup != nil {
		return ActionTypeSetup
	} else if a.Execute != nil {
		return ActionTypeExecute
	}
	v, e := json.Marshal(a)
	panic(fmt.Sprintf("unknown action type: %s, %v", string(v), e))
}
