// polkadot RPC, that the Polkadot.js extension talks to
export const WS_PROVIDER_URLS_PROD = process.env.WS_RPC_PROVIDER_URL;

// bridge app, that sits in the kusama network and responsible for the NFT transferring
export const MOVR_BRIDGE_ADDRESS_KUSAMA = process.env.KUSAMA_BRIDGE_ADDRESS;

// MOVR contract, that is responsible for game logic, for the deposit (preorder) and other stuff
export const MOONBASE_CONTRACT = process.env.MOONBASE_CONTRACT;

// Chain ID. In order to avoid sending tokens to real network, we force Metamask to switch to
// exactly this network. Must be the numeric value. Specifically for the Moonbeam, please refer to the docs:
// https://docs.moonbeam.network/builders/get-started/moonbase/#chain-id
const moonbeamNetworkNumericID = process.env.MOONBEAM_NETWORK_ID;
export const TEST_MOONBEAM_NETWORK_ID = `0x${Number(moonbeamNetworkNumericID).toString(16)}`;