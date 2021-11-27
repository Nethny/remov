export const WS_PROVIDER_URLS_PROD = process.env.WS_RPC_PROVIDER_URL || "wss://kusama-rpc.polkadot.io";
export const MOVR_BRIDGE_ADDRESS_KUSAMA = process.env.KUSAMA_BRIDGE_ADDRESS || "";
export const MOONBASE_CONTRACT = process.env.MOONBASE_CONTRACT || "";
const moonbeamNetworkNumericID:number = parseInt(process.env.MOONBEAM_NETWORK_ID || "0", 10);
export const TEST_MOONBEAM_NETWORK_ID = `0x${Number(moonbeamNetworkNumericID).toString(16)}`;
export const TRANSACTION_VALUE_IN_ETHER = process.env.VALUE_OF_CONTRACT_TRANSACTION_ETHER || "0";
export const RMRK_KANARIA_API_ACCOUNT_BIRDS = process.env.RMRK_KANARIA_API_ACCOUNT_BIRDS;
export const RMRK_KANARIA_API_NFT_DETAILS = process.env.RMRK_KANARIA_API_NFT_DETAILS;