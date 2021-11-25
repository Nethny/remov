import React, {FC, useState} from 'react';
import StatusBar, {asDone, asError, asProgress, BarProps, StatusType} from "../StatusBar/StatusBar";
import styles from "../DepositMaker/deposit.module.scss";
import {MOONBASE_CONTRACT} from "../../constants";
import {CornerDownLeft, Check} from "react-feather";
import {switchMetamaskToMoonbeamNetwork} from "../../wallets/connectors";
import {ethers} from "ethers";
import contractMoonbaseJSON from "../../contracts/moonbase.contract.json";
import {TransactionReceipt, TransactionResponse} from "@ethersproject/abstract-provider";
import {useSelector} from "react-redux";
import {getSelectedNFT, getSelectedPolkadotAddress} from "../../store/removSelectors";

const NftRefund: FC = () => {
    const [currentState, setCurrentState] = useState<BarProps>({message: "", type: StatusType.HIDDEN});
    const [finished, setFinished] = useState<boolean>(false);
    const [hash, setHash] = useState<string>("");

    const nftID = useSelector(getSelectedNFT);
    const selectedPolkadotAddress = useSelector(getSelectedPolkadotAddress);

    async function refundNFT() {

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

                return Promise.all([
                    contractWithSigner,
                    feeData
                ]);
            })
            .then(([contractWithSigner, feeData]) => {

                // call the contract's remote function
                let gasPrice = feeData.gasPrice;
                let overrides = {
                    gasLimit: 10000000,
                    gasPrice: gasPrice,
                    value: ethers.utils.parseUnits('0.021', 'ether'),
                };

                // call our contract
                return contractWithSigner.createMOVROrder(selectedPolkadotAddress, nftID, overrides);

            })
            .then((res:TransactionResponse) => {
                setCurrentState(asProgress("Operation was sent to contract mining, please wait for the result. " +
                    "It may take some time, up to 5 minutes."))

                // wait for the contract function is done, may be very long, because it is processed by miners
                return res.wait()
            })
            .then((trRe: TransactionReceipt) => {
                setCurrentState(asDone("NFT were successfully refunded"));
                setFinished(true);
                setHash(trRe.transactionHash);
            })
            .catch((err:any) => {
                console.log(err);
                setCurrentState(asError("error " + err));
            });
    }

    return(
        <div>
            <StatusBar message={currentState.message} type={currentState.type} />

            {finished ?
                <div>
                    <span className={styles.success}>Transaction was finished.</span> <br />
                    You may find details <a href={"https://moonbase.subscan.io/tx/" + hash} target={"_blank"}>by link</a>,
                    but please be aware that you transaction will be visible within some short amount of time.
                </div>
                :
                <button
                    disabled={currentState.type === StatusType.PROGRESS}
                    className={styles.orderButton}
                    type={"button"}
                    onClick={refundNFT}>
                        { currentState.type !== StatusType.PROGRESS ?
                            <span>
                                <CornerDownLeft className={styles.icon} />
                                <span className={styles.text}>Refund NFT to back Remark</span>
                            </span>
                        :
                            <span>Progress...</span> }
                </button>
            }
        </div>
    );
};
export default NftRefund;