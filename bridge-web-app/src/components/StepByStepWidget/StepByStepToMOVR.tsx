import React, {FC, useEffect} from 'react';
import {useDispatch, useSelector} from "react-redux";
import styles from "./stebbystep.module.scss";
import PolkadotWallet from "../PolkadotWallet/PolkadotWallet";
import {getCurrentStepNumber, getSelectedPolkadotAddress} from "../../store/removSelectors";
import DepositMaker from "../DepositMaker/DepositMaker";
import NftSender from "../NFTItems/NftSender";
import NftItems from "../NFTItems/NftItems";
import {resetStore} from "../../store/removReducerSlice";
import {ExternalLink} from "react-feather";


const StepByStepToMOVR: FC = () => {

    const currentStep = useSelector(getCurrentStepNumber);
    const selectedWalletAddress = useSelector(getSelectedPolkadotAddress);
    const dispatch = useDispatch();

    useEffect(()=>{
        dispatch(resetStore());
    },[]);

    function isActive(step: number):string {
        if (currentStep >= step) {
            return styles.step + " " + styles.active;
        }
        return styles.step;
    }

    return(<div className={styles.stepByStep}>
        <ol>
            <li className={isActive(1)}>
                <div className={styles.content}>
                    <h1>Select your Remark account and load NFT</h1>
                    <p>Please, connect your <a href={"https://polkadot.js.org/extension/"} target={"_blank"}>Polkadot.js extension <ExternalLink className={styles.icon} /></a> and select active wallet account</p>

                    <PolkadotWallet moveTo={2} />

                </div>
            </li>

            <li className={isActive(2)}>
                <div className={styles.content}>
                    <h1>Select NFTs</h1>
                    <p>Please, select NFTs found for selected account, that you want to transfer to the Moonriver network</p>

                    <NftItems walletAddress={selectedWalletAddress} moveToStep={3} />

                </div>
            </li>

            <li className={isActive(3)}>
                <div className={styles.content}>
                    <h1>Deposit</h1>
                    <p>Please, make a deposit to transfer your NFT to the Moonbase network.</p>

                    <DepositMaker />

                </div>
            </li>

            <li className={isActive(4)}>
                <div className={styles.content}>
                    <h1>Transfer NFT token to Moonriver network</h1>
                    <p>The last step is to transfer your NFT to the Moonbase network</p>

                    <NftSender />

                </div>
            </li>

        </ol>
    </div>);
};

export default StepByStepToMOVR;