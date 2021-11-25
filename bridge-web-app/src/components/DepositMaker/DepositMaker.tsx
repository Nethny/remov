import React, {FC, useEffect, useState} from 'react';
import {ethers} from "ethers";
import styles from './deposit.module.scss';
import {TransactionReceipt, TransactionResponse} from "@ethersproject/abstract-provider";
import {useDispatch, useSelector} from 'react-redux';
import {getSelectedNFT, getSelectedMetamaskAddress} from "../../store/removSelectors";
import {setCurrentStepNumber, setSelectedMetamaskAddress} from '../../store/removReducerSlice';
import {DollarSign} from "react-feather";
import { setUpMetamask, switchMetamaskToMoonbeamNetwork} from "../../wallets/connectors";
import {MOONBASE_CONTRACT, TEST_MOONBEAM_NETWORK_ID} from "../../constants";
import contractMoonbaseJSON from "../../contracts/moonbase.contract.json";
import StatusBar, {asDone, asError, asProgress, BarProps, StatusType} from "../StatusBar/StatusBar";

const DepositMaker: FC = () => {

    const dispatch = useDispatch();
    const nftID = useSelector(getSelectedNFT);
    const selectedMetamaskAddress = useSelector(getSelectedMetamaskAddress);

    const [finished, setFinished] = useState<boolean>(false);
    const [providerDetected, setProviderDetected] = useState<boolean>(false);
    const [hash, setHash] = useState<string>("");

    // local status panel that shows user in human-readable form what is going on at current step
    const [currentState, setCurrentState] = useState<BarProps>({message: "", type: StatusType.HIDDEN});

    const handleAccountsChanged = (accounts: any) => {
        if (accounts.length === 0) {
            setProviderDetected(false);
            setCurrentState(asProgress("Please, login with Metamask"));

        } else if (accounts[0] !== selectedMetamaskAddress) {
            let validMetamaskUserAddress = ethers.utils.getAddress(accounts[0]);
            dispatch(setSelectedMetamaskAddress(validMetamaskUserAddress))
            setProviderDetected(true);
        }
    };

    /**
     * Perform money transaction to the MOVR contract
     */
    async function makeDeposit() {
        setCurrentState(asProgress("Switch to MoonBase network. " +
            "Your Metamask should ask you to connect to Moonbase, please allow"));

        // switch Metamask to Moonbase chain
        await switchMetamaskToMoonbeamNetwork(setCurrentState)
            .then((provider) => {

                setCurrentState(asProgress("Connect to our MOVR Bridge Contract with a signer"));
                const signer = provider.getSigner();
                const contract = new ethers.Contract(MOONBASE_CONTRACT, contractMoonbaseJSON, provider);
                let contractWithSigner = contract.connect(signer);

                setCurrentState(asProgress("Calling the contract. Please confirm it with Metamask " +
                    "(should open in a separate window)"));
                let feeData = provider.getFeeData();

                const chainId = provider.send('eth_chainId',[]);

                return Promise.all([
                    contractWithSigner,
                    feeData,
                    chainId
                ]);
            })
            .then(([contractWithSigner, feeData, chainId]) => {
                
                // call the contract's remote function
                let gasPrice = feeData.gasPrice;
                let overrides = {
                    gasLimit: 10000000,
                    gasPrice: gasPrice,
                    value: ethers.utils.parseUnits('0.021', 'ether'),
                };

                let validMetamaskAddress = ethers.utils.getAddress(selectedMetamaskAddress);

                // the last check, that we send funds to correct network:
                if (chainId !== TEST_MOONBEAM_NETWORK_ID) {
                    return Promise.reject("Please, select the testnet.moonbeam.network in your Metamask (in the top right corner) and try again");
                }

                // call our contract
                return contractWithSigner.createRMRKOrder(validMetamaskAddress, nftID, overrides);

            })
            .then((res:TransactionResponse) => {
                setCurrentState(asProgress("Operation was sent to contract mining, please wait for the result. " +
                    "It may take some time."));
                // wait for the contract function is done, may be very long
                return res.wait()
            })
            .then((trRe: TransactionReceipt) => {
                dispatch(setCurrentStepNumber(4));
                setCurrentState(asDone("Deposit was made successfully"));
                setFinished(true);
                setHash(trRe.transactionHash);
            })
            .catch((err:any) => {
                console.log(err);
                if (err.code === 4001) {
                    setCurrentState(asError("Transaction rejected by your wallet"));
                } else {
                    setCurrentState(asError("Failed to submit transaction. Error: " + err));
                }
            });
    };

    /**
     * On the initializing we connect to the Metamask where user selects some active account
     */
    useEffect(() => {
        setUpMetamask(handleAccountsChanged)
            .then((selectedAddress: string) => {
                dispatch(setSelectedMetamaskAddress(selectedAddress));
                return setProviderDetected(true);
            })
            .catch(() => setProviderDetected(false))
    },[]);

    return(<div>

            <StatusBar message={currentState.message} type={currentState.type} />

            {providerDetected ?
                <div>
                    {finished ?
                        <div>
                            <span className={styles.success}>Transaction was finished.</span> <br />
                            You may find details <a href={"https://moonbase.subscan.io/tx/" + hash} target={"_blank"}>by link</a>
                        </div>
                        :

                        <button
                            disabled={currentState.type === StatusType.PROGRESS }
                            className={styles.orderButton}
                            type={"button"}
                            onClick={makeDeposit}>
                            { currentState.type !== StatusType.PROGRESS ?
                                <span>
                                        <DollarSign className={styles.icon} /> <span className={styles.text}>Pay deposit</span>
                                    </span>
                                :
                                <span>Progress...</span> }
                        </button>
                    }

                </div>
                :
                <p style={{"color":"red"}}>Provider is not detected</p> }

    </div>);
}

export default DepositMaker;