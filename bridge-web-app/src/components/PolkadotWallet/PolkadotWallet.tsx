import React, {FC, useEffect, useState} from 'react';
import {
    web3Accounts,
    web3Enable,
} from '@polkadot/extension-dapp';
import {Keyring} from "@polkadot/api";
import Select from 'react-select';
import {ActionMeta, OnChangeValue} from "react-select/dist/declarations/src/types";
import styles from "./polkadot.wallet.module.scss";
import {useDispatch} from "react-redux";
import {setCurrentStepNumber, setSelectedPolkadotAddress} from "../../store/removReducerSlice";

const enum Status {
    NotLoaded = 0,
    ExtensionIsNotInstalled = 1,
    FoundUserAccount,
    GeneralError
}
function statusDescription(status: Status):string {
    switch (status) {
        case Status.NotLoaded:
            return "Loading has not been started yet";
        case Status.ExtensionIsNotInstalled:
            return "The Polkadot.js extension is not installed or user is not authorized";
        case Status.FoundUserAccount:
            return "User accounts found";
        case Status.GeneralError:
            return "General error";
    }
}

interface Option {
    value: string;
    label: string;
}

function convertToKusamaAddress(address: string) {
    if (address === undefined) {
        return ""
    }
    const keyring = new Keyring({ type: 'sr25519' });
    const publicKey = keyring.decodeAddress(address);
    const kusamaAddress = keyring.encodeAddress(publicKey, 2); // 2 for Kusama SS58
    return kusamaAddress;
}

interface WalletProps {
    moveTo: number;
}

const PolkadotWallet: FC<WalletProps> = (props: WalletProps) => {

    const dispatch = useDispatch();

    const [status, setStatus] = useState<Status>(Status.NotLoaded);
    const [loaded, setLoaded] = useState<boolean>(false);
    const [walletAccounts, setWalletAccounts] = useState<Option[]>([]);

    function handleWalletSelect(newValue: OnChangeValue<Option, boolean>, actionMeta: ActionMeta<Option>)  {
        let nVal = newValue?.valueOf() as Option;
        saveSelectedOption(nVal);
    };

    function saveSelectedOption(address: Option) {
        dispatch(setSelectedPolkadotAddress(convertToKusamaAddress(address.value)));
        dispatch(setCurrentStepNumber(props.moveTo));
    }

    useEffect(() => {
        async function connectToLocalWallet(): Promise<Status> {

            const allInjected = await web3Enable('bridge');
            if (allInjected.length === 0) {
                // no extension installed, or the user did not accept the authorization
                // in this case we should inform the use and give a link to the extension
                return Promise.reject(Status.ExtensionIsNotInstalled);
            }

            setLoaded(true);

            // we are now informed that the user has at least one extension and that we
            // will be able to show and use accounts
            const allAccounts = await web3Accounts();
            let selectableAccounts: Option[] = new Array();
            for (let a of allAccounts) {
                selectableAccounts.push({
                    label: a.meta.name + ": " + a.address.substring(0, 10),
                    value: a.address
                } as Option);
            }
            setWalletAccounts(selectableAccounts);

            return Promise.resolve(Status.FoundUserAccount);
        }

        connectToLocalWallet()
            .then((st) => setStatus(st))
            .catch((status) => {
                setLoaded(true);
                setStatus(status)
            })
    }, []);

    return(
        <div>
            {!loaded ?
                <p>loading...</p>
                :
                <div>{status === Status.FoundUserAccount ?
                    <div className={styles.walletSelector}>

                        <p>Select your wallet</p>
                        <Select
                            className={styles.selectBox}
                            placeholder={"Please select wallet account"}
                            onChange={handleWalletSelect}
                            options={walletAccounts}
                        />

                    </div>
                    :
                    <p style={{"color": "red"}}>Not connected, status: {statusDescription(status)}</p> }
                </div>
            }
        </div>
    )
}

export default PolkadotWallet;