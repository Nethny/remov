import React, {FC, useEffect} from 'react';
import styles from "./stebbystep.module.scss";
import {useDispatch, useSelector} from "react-redux";
import {getCurrentStepNumber} from "../../store/removSelectors";
import MovrNftSelector from "../MovrNFTSelector/MovrNftSelector";
import PolkadotWallet from "../PolkadotWallet/PolkadotWallet";
import {resetStore} from "../../store/removReducerSlice";
import NftRefund from "../NFTItems/NftRefund";

const StepByStepBackToRMRK: FC = () => {

    const dispatch = useDispatch();
    const currentStep = useSelector(getCurrentStepNumber);

    function isActive(step: number):string {
        if (currentStep >= step) {
            return styles.step + " " + styles.active;
        }
        return styles.step;
    }

    useEffect(()=>{
        dispatch(resetStore());
    },[]);

    return (<div className={styles.stepByStep}>
        <ol>
            <li className={isActive(1)}>
                <div className={styles.content}>
                    <h1>Select Images</h1>
                    <p>Please, connect your Metamask extension and choose NFT's you'd like to refund back to the RMRK network</p>

                    <MovrNftSelector />

                </div>
            </li>

            <li className={isActive(2)}>
                <div className={styles.content}>
                    <h1>Select your target Remark account</h1>
                    <p>Please, connect your Polkadot.js extension and select wallet where you want to refund your NFT</p>

                    <PolkadotWallet moveTo={3} />

                </div>
            </li>

            <li className={isActive(3)}>
                <div className={styles.content}>
                    <h1>Refund your NFTs</h1>
                    <p>Now its time to move your NTFs back to Remark network from MOVR</p>

                    <NftRefund />

                </div>
            </li>

        </ol>
    </div>);
}

export default StepByStepBackToRMRK;