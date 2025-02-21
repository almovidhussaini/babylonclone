package types

import "github.com/amovidhussaini/ybtcclone/types"

func NewEventSlashedFinalityProvider(evidence *Evidence) *EventSlashedFinalityProvider {
	return &EventSlashedFinalityProvider{
		Evidence: evidence,
	}
}

func NewEventJailedFinalityProvider(fpPk *types.BIP340PubKey) *EventJailedFinalityProvider {
	return &EventJailedFinalityProvider{PublicKey: fpPk.MarshalHex()}
}
