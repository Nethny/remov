import detectEthereumProvider from "@metamask/detect-provider";
import {ethers} from "ethers";
import {asError, asProgress, BarProps} from "../components/StatusBar/StatusBar";
import {TEST_MOONBEAM_NETWORK_ID} from "../constants";

declare let window: any;

export type EthersProvider = ethers.providers.StaticJsonRpcProvider | ethers.providers.Web3Provider;

/**
 * Returns current provider.
 * If the Metamask is installed, then we use standard window.ethereum object, injected by
 * Metamask extension. If the Metamask is not installed, then we can use RPC provider with
 * limited functionality. It probably won't work so far, may be in the future we will re-use this
 * provider, but for now we assume that Metamask is always installed
 */
const web3 = (): EthersProvider => {
    if (typeof window.ethereum !== "undefined") {
        // We are in the browser and MetaMask is running
        return new ethers.providers.Web3Provider(window.ethereum);
    } else {
        // We are on the server *OR* the user is not running metamask
        return new ethers.providers.StaticJsonRpcProvider("https://rpc.testnet.moonbeam.network", {
            chainId: 1287,
            name: "moonbase-alpha",
        });
    }
};
export default web3;

/**
 * Connect to the Metamask and request user's permission to operate. Here we return selectedAccount in the Promise (in
 * addition to dispatch() method), because metamask set up and contract connect will be handled in the save useEffect()
 * and the state is not being propagated, so we can't always use dispatch()
 *
 * @param dispatch link to dispatch()
 * @param onAccountChange
 */
export async function setUpMetamask(onAccountChange: (accounts: any) => void): Promise<string>{
    const provider = await detectEthereumProvider();
    if (provider && window.ethereum) {

        // Popup Metamask for request permission
        // Docs: https://docs.metamask.io/guide/rpc-api.html#permissions
        return window.ethereum
            .request({ method: 'eth_requestAccounts' })
            .then((a: any) => {
                let selectedAddress = ethers.utils.getAddress(a[0]);
                console.log("Request is gathered for Metamask account " + selectedAddress);
                return Promise.resolve(selectedAddress);
            })
            .then((selectedAccount: string) => {

                // (attempt to fix bug) we need to explicitly "use" some address so that in
                // the future "eth_accounts" would return at least one account
                window.ethereum
                    .request({ method: 'eth_accounts' })
                    .then(onAccountChange)
                    .catch((err:any) => {
                        console.error(err);
                    });

                window.ethereum.on('accountsChanged', onAccountChange);
                return Promise.resolve(selectedAccount);
            });
    }

    return Promise.reject("Metamask is not installed");
};

export async function switchMetamaskToMoonbeamNetwork(updateState: (props: BarProps) => void): Promise<EthersProvider> {

    const provider = web3();
    const chainId = await provider.send('eth_chainId',[]);

    if (chainId == TEST_MOONBEAM_NETWORK_ID) {
        return Promise.resolve(provider);
    }

    try {
        // https://github.com/rekmarks/EIPs/blob/3326-create/EIPS/eip-3326.md
        await window.ethereum.request({
            method: 'wallet_switchEthereumChain',
            params: [{ chainId: TEST_MOONBEAM_NETWORK_ID }],
        }).then(() => {
            window.location.reload();
        });
    } catch (switchError: any) {

        if (switchError.code === 4001) {
            updateState(asError("User rejected the new network adding. Please, add new network to your Metamask and try again"));
            return Promise.reject();
        }

        if (switchError.code === 4902) {
            // This error code indicates that the chain has not been added to MetaMask.
            // So we request user to add this network to his/her extension.
            // When the new Moonbeam network is added, then let's switch to it, user should approve this action
            try {
                updateState(asProgress("Add new network to the Metamask, follow the instruction by your Metamask extension"));

                await window.ethereum.request({
                    method: 'wallet_addEthereumChain',
                    params: [{
                        chainId: TEST_MOONBEAM_NETWORK_ID,
                        rpcUrls: ['https://rpc.testnet.moonbeam.network'],
                        chainName: 'moonbase-alpha'
                    }],
                }).then(() => {
                    window.location.reload();
                });
            } catch (addError) {
                // handle "add" error
                updateState(asError("Error " + addError));
                return Promise.reject(addError);
            }
        }
    }
    return Promise.resolve(provider);
}