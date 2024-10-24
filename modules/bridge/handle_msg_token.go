package bridge

import (
	juno "github.com/forbole/juno/v4/types"
	bridge "github.com/hyle-team/bridgeless-core/x/bridge/types"
)

// handleMsgInsertToken allows to properly handle a MsgInsertToken
func (m *Module) handleMsgInsertToken(tx *juno.Tx, msg *bridge.MsgInsertToken) error {
	return nil
}

// handleMsgDeleteToken allows to properly handle a MsgDeleteToken
func (m *Module) handleMsgDeleteToken(tx *juno.Tx, msg *bridge.MsgDeleteToken) error {
	return nil
}
