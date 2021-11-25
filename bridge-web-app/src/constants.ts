export const WS_PROVIDER_URLS_PROD = 'wss://kusama-rpc.polkadot.io'; // polkadot
export const MOVR_BRIDGE_ADDRESS_KUSAMA = "HeReMovoG8uXkhsdJTWCjy7yhmaPFSeZtekqt54m6mEXQya"; // here we send NFT tokens, RMRK contract
export const MOONBASE_CONTRACT = '0xab51e296f3cafaBeb20b26EfcFEC02Be52F26C85'; // here we send funds (deposit)
const moonbeamNetworkNumericID = 1287; // https://docs.moonbeam.network/builders/get-started/moonbase/#chain-id
export const TEST_MOONBEAM_NETWORK_ID = `0x${Number(moonbeamNetworkNumericID).toString(16)}`;