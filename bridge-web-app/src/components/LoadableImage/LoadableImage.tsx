import React, {FC, useEffect, useState} from 'react';
import styles from './image.module.scss';
import {RefreshCw} from "react-feather";

enum Status {
    LOADING = 1,
    LOADED,
    FAILED,
};


interface ItemsProps {
    src: string;
    id: string;
    name: string;
    className: string;
}

const LoadableImage: FC<ItemsProps> = (props: ItemsProps) => {
    const [status, setStatus] = useState<Status>(Status.LOADING);

    function getImageStyle() {
        if (status == Status.LOADING || status == Status.FAILED) {
            return styles.loadableImage + " " + styles.hidden;
        }
        return styles.loadableImage + " " + styles.visible;
    }

    return(
        <div className={props.className + " " + styles.container}>

            <img
                className={ getImageStyle() }
                src={props.src}
                data-id={props.id}
                alt={props.name}
                onLoad={() => setStatus(Status.LOADED)}
                onError={() => setStatus(Status.FAILED)}
            />

            {status === Status.FAILED &&
                <p className={styles.error}>Error</p>
            }
            {status === Status.LOADING &&
                <div className={styles.loadingSpinner}>
                    <RefreshCw className={styles.icon} />
                </div>
            }
        </div>
    );
}

export default LoadableImage;