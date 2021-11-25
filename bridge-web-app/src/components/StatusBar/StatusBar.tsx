import React from 'react';
import styles from "./status.bar.module.scss";
import { RotateCw, CheckCircle, AlertTriangle } from "react-feather";

export enum StatusType {
    PROGRESS = 1,
    DONE,
    ERROR,
    HIDDEN
};

export interface BarProps {
    message: string;
    type: StatusType;
}

function _newProps(msg: string, t: StatusType): BarProps {
    return {
        message: msg,
        type: t,
    } as BarProps;
}

export function asError(msg: string): BarProps {
    return _newProps(msg, StatusType.ERROR);
}
export function asDone(msg: string): BarProps {
    return _newProps(msg, StatusType.DONE);
}
export function asProgress(msg: string): BarProps {
    return _newProps(msg, StatusType.PROGRESS);
}

/**
 * Common bar with loading indicator and message, saying the current status.
 * The styling is depending on its type (error, progress, done.
 */
function StatusBar(props: BarProps) {

    const classForType = ():string => {
        switch(props.type) {
            case StatusType.DONE:
                return styles.done;
            case StatusType.ERROR:
                return styles.error;
            case StatusType.PROGRESS:
                return styles.progress;
            case StatusType.HIDDEN:
                return "";
        }
    };

    const iconForType = ():React.ReactNode => {
        switch(props.type) {
            case StatusType.DONE:
                return  <CheckCircle className={styles.svg}/>;
            case StatusType.ERROR:
                return  <AlertTriangle className={styles.svg}/>;
            case StatusType.PROGRESS:
                return  <RotateCw className={styles.svg}/>;
            case StatusType.HIDDEN:
                return <span />;
        }
    }

    return (<div>
        {props.type !== StatusType.HIDDEN &&
            <div className={styles.progressMessage + " " + classForType()}>

                <div className={styles.loadingSpinner}>
                    { iconForType() }
                </div>

                <div className={styles.message}>
                    { props.message }
                </div>
            </div>
        }
    </div>);
}

export default StatusBar;