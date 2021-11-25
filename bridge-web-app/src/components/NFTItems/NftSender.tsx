import React, {FC, useState} from 'react';
import styles from "../DepositMaker/deposit.module.scss";
import {useSelector} from "react-redux";
import {getSelectedMetamaskAddress, getSelectedNFT, getSelectedPolkadotAddress} from "../../store/removSelectors";
import {ApiPromise, WsProvider} from "@polkadot/api";
import {MOONBASE_CONTRACT, MOVR_BRIDGE_ADDRESS_KUSAMA, WS_PROVIDER_URLS_PROD} from "../../constants";
import {web3FromAddress} from "@polkadot/extension-dapp";
import StatusBar, {asDone, asError, asProgress, BarProps, StatusType} from "../StatusBar/StatusBar";

const NftSender: FC = () => {

    const selectedPolkadotAddress = useSelector(getSelectedPolkadotAddress);
    const selectedMetamaskAddress = useSelector(getSelectedMetamaskAddress);
    const nftID = useSelector(getSelectedNFT);

    const [finished, setFinished] = useState<boolean>(false);
    const [currentState, setCurrentState] = useState<BarProps>({message: "", type: StatusType.HIDDEN});

    async function sendNFT() {

        // https://polkadot.js.org/docs/api/start/api.query
        setCurrentState(asProgress("Connect to chain"));

        // connect to Polkadot
        const wsProvider = new WsProvider(WS_PROVIDER_URLS_PROD);
        let polkadotApi = await ApiPromise.create({ provider: wsProvider });
        await polkadotApi.isReady;

        const [chain, nodeName, nodeVersion] = await Promise.all([
            polkadotApi.rpc.system.chain(),
            polkadotApi.rpc.system.name(),
            polkadotApi.rpc.system.version()
        ]);
        setCurrentState(asProgress(`You are connected to chain ${chain} using ${nodeName} v${nodeVersion}`));

        // address from the Polkadot.js wallet, sender
        const acc = await web3FromAddress(selectedPolkadotAddress);

        let extrinsics = [
            polkadotApi.tx.system.remark(`RMRK::SEND::2.0.0::${nftID}::${MOVR_BRIDGE_ADDRESS_KUSAMA}`),
            polkadotApi.tx.system.remark(`RMRK::BRIDGE::2.0.0::MOVR::${selectedMetamaskAddress}`),
        ];
        const transaction = polkadotApi.tx.utility.batchAll(extrinsics);

        await transaction.signAndSend(selectedPolkadotAddress, { signer: acc.signer }, (result) => {
            if (result.status.isInBlock) {
                // Block addition is pending
                setCurrentState(asProgress("Transaction is submitted. Please wait..."));
            } else if (result.status.isFinalized) {
                // Block addition is finalised
                setFinished(true);
                setCurrentState(asDone("NFT transfer was finished successfully"));
            }
        }).catch((err: any) => {
            setCurrentState(asError("NFT transfer failed. Reason is: " + err.message));
        });
    }

    return (
        <div>

            <StatusBar message={currentState.message} type={currentState.type} />

            {finished ?
                <div>
                    <p>
                        <span className={styles.success}>Done!</span> You may explore contract transactions <a href={"https://moonbase.moonscan.io/address/" + MOONBASE_CONTRACT}>by link</a>
                    </p>
                </div>
                :
                <input
                    disabled={currentState.type === StatusType.PROGRESS }
                    className={styles.orderButton}
                    type={"button"}
                    value={currentState.type === StatusType.PROGRESS ? "Progress..." : "Send NFT to Moonriver"}
                    onClick={sendNFT}/>
            }
        </div>);
};

export default NftSender;