import React, {FC, useEffect, useState} from 'react';
import styles from './movr.nft.selector.module.scss';
import  {
    EthersProvider,
    setUpMetamask,
    switchMetamaskToMoonbeamNetwork
} from "../../wallets/connectors";
import {useDispatch, useSelector} from "react-redux";
import {ethers} from "ethers";
import {setSelectedMetamaskAddress} from "../../store/removReducerSlice";
import {getSelectedMetamaskAddress} from "../../store/removSelectors";
import {MOONBASE_CONTRACT, TEST_MOONBEAM_NETWORK_ID} from "../../constants";
import contractMoonbaseJSON from "../../contracts/moonbase.contract.json";
import NftItems from "../NFTItems/NftItems";
import StatusBar, {asDone, asError, asProgress, BarProps, StatusType} from "../StatusBar/StatusBar";

const MovrNftSelector: FC = () => {
    const dispatch = useDispatch();
    const selectedMetamaskAddress = useSelector(getSelectedMetamaskAddress);
    const [currentState, setCurrentState] = useState<BarProps>({message: "", type: StatusType.HIDDEN});
    const [imageUuids, setImageUuids] = useState<string[]>([]);

    const handleAccountsChanged = (accounts: any) => {
        if (accounts.length === 0) {
            setCurrentState(asError("Please, login with Metamask"));
        } else if (accounts[0] !== selectedMetamaskAddress || selectedMetamaskAddress === "") {
            let validMetamaskUserAddress = ethers.utils.getAddress(accounts[0]);
            dispatch(setSelectedMetamaskAddress(validMetamaskUserAddress));
        }
    };

    async function connectToMetamask(): Promise<[string, EthersProvider]> {
        return setUpMetamask(handleAccountsChanged)
            .then((selectedAccount: string) => {
                let promiseProvider = switchMetamaskToMoonbeamNetwork(setCurrentState);
                return Promise.all([selectedAccount, promiseProvider]);
            })
            .catch((err) => {
                setCurrentState(asError("error connecting to Metamask: " + err));
                return Promise.reject(err);
            })
    }

    useEffect(() =>{
        setCurrentState(asProgress("Connect to the Metamask extension"));
        connectToMetamask()
            .then(([selectedAccount, provider]) => {

                // create instance of Contract
                setCurrentState(asProgress("Connect to the MOVR contract"));
                const contract = new ethers.Contract(MOONBASE_CONTRACT, contractMoonbaseJSON, provider);

                const chainId = provider.send('eth_chainId',[]);

                return Promise.all([selectedAccount, contract, chainId]);
            })
            .then(([selectedAccount, contract, chainId]) => {

                // the last check, that we send funds to correct network:
                if (chainId !== TEST_MOONBEAM_NETWORK_ID) {
                    return Promise.reject("Please, select the testnet.moonbeam.network in your Metamask (in the top right corner) and try again");
                }

                // 3) call function and get UIds, this call is read-only, shouldn't take a long tie
                let validMetamaskAddress = ethers.utils.getAddress(selectedAccount);

                setCurrentState(asProgress("Perform call to contract..."));
                let promiseMovrUids = contract.getUserProperty(validMetamaskAddress);

                return Promise.all([
                    promiseMovrUids,
                    contract
                ]);
            })
            .then(([movrUIDs, contract]) => {

                // here we call contract for every item,
                // and convert MOVR UID to RMRK UID
                setCurrentState(asProgress("Load RMRK ID's for your images..."));
                return Promise.all(
                    movrUIDs.map((movrUID: number) => contract.getUID(movrUID))
                );
            })
            .then((rmrkUIDs) => {
                let strArray = rmrkUIDs as string[];
                setImageUuids(strArray);
                if (strArray.length === 0) {
                    setCurrentState(asDone("NFTs are loaded, but empty array returned"))
                } else {
                    setCurrentState(asDone("NFTs are loaded"))
                }
            })
            .catch((err:any) => {
                console.log(err);
                setCurrentState(asError("error loading NFTs: " + err.message));
            });

    }, []);

    return(<div className={styles.NftSelector}>

        <StatusBar message={currentState.message} type={currentState.type} />

        <NftItems imageUUIDs={imageUuids} moveToStep={2} />

    </div>);
}

export default MovrNftSelector;