import React, {FC, useEffect, useState} from 'react';
import {useDispatch, useSelector} from "react-redux";
import {setCurrentStepNumber, setSelectedNFT} from '../../store/removReducerSlice';
import {getSelectedNFT} from "../../store/removSelectors";
import styles from "./nft.module.scss";
import LoadableImage from "../LoadableImage/LoadableImage";
import StatusBar, {StatusType} from "../StatusBar/StatusBar";
import {RMRK_KANARIA_API_ACCOUNT_BIRDS, RMRK_KANARIA_API_NFT_DETAILS} from "../../constants";

interface AccountBird {
    id: string // NFT id
    collectionId: string
    image: string // Image URL
    metadata_name: string
}
interface ImageDetails {
    id: string;
    image: string; // url
    name: string;
    details: string;
    // properties
}

interface NFTDetails {
    id: string;
    forsale: number; // price
    owner: string;
    rootowner: string;
    collectionId: string;
    metadata: string;
    sn: string;
    updated_at: string
    image: string; // URL
}

// expected one of these params
interface ItemsProps {
    walletAddress?: string;
    imageUUIDs?: string[];
    moveToStep: number;
}

/*
Widget with a list of available NTF images
 */
const NftItems: FC<ItemsProps> = (props: ItemsProps) => {
    const [images, setImages] = useState<ImageDetails[]>([])
    const [loading, setLoading] = useState<boolean>(true)

    const dispatch = useDispatch();

    const selected = useSelector(getSelectedNFT);

    useEffect(() => {

        // reset state
        setImages([]);

        // load all the items associated with given wallet
        if (props.walletAddress !== undefined && props.walletAddress !== "") {
            setLoading(true);
            getAccountBirds(props.walletAddress)
                .then((arrImages: ImageDetails[]) => {
                    setImages(arrImages);
                    setLoading(false);
                });
        }

        // load all the items associated with given UUID array
        else if (props.imageUUIDs !== undefined) {
            setLoading(true);
            let promises = props.imageUUIDs.map((uuid: string) => loadNFTDetails(uuid));
            Promise.all(promises)
                .then((NFTDetails: NFTDetails[]) => NFTDetails.map(nft => {
                    return {
                        id: nft.id,
                        image: nft.image, // url
                        name: nft.id,
                        details: nft.collectionId
                    } as ImageDetails
                }))
                .then((arrImages: ImageDetails[]) => {
                    setImages(arrImages);
                    setLoading(false);
                });
        }


    },[props.walletAddress, props.imageUUIDs]);

    function change(nextSelected: ImageDetails) {
        dispatch(setSelectedNFT(nextSelected.id));

        // and, the current step is done, let's move on to the next step
        dispatch(setCurrentStepNumber(props.moveToStep));
    }

    // Returns TRUE if at least one property is set and has valid value
    function isPropsSet():boolean {
        return (props.walletAddress !== undefined && props.walletAddress !== "")
            || (props.imageUUIDs !== undefined && props.imageUUIDs.length > 0);
    }

    return (
        <div className={styles.imageContainer}>
                <div className={styles.scrollable}>

                    {loading &&
                        <StatusBar message={"Loading image details"} type={StatusType.PROGRESS} />
                    }


                    {isPropsSet() &&
                        <div>

                            {!loading && images.length == 0 && <p>Sorry, no NFT found :(</p> }

                            {!loading && images.length > 0 &&
                                <ul className={styles.nftItem}>{images.map((img: ImageDetails) => {
                                    return (
                                        <li key={img.id}>
                                            <input type="radio"
                                                   name="images"
                                                   id={"nft_" + img.id}
                                                   checked={selected == img.id}
                                                   onChange={() => change(img)}
                                            />
                                            <label htmlFor={"nft_" + img.id}>
                                                <LoadableImage
                                                    className={styles.image}
                                                    name={img.name}
                                                    src={img.image}
                                                    id={img.id}
                                                />
                                            </label>
                                        </li>)
                                })}</ul>
                            }
                    </div>
                    }
                </div>
        </div>
    );
}

function getAccountBirds(address: string): Promise<ImageDetails[]> {
    return fetch(`${RMRK_KANARIA_API_ACCOUNT_BIRDS}/${address}`,{mode: 'cors'})
        // the JSON body is taken from the response
        .then(res => res.json())
        .then(res => {
            let items:AccountBird[] = res as AccountBird[];

            return Promise.resolve(items.map((ab) => {
                return {
                    id: ab.id,
                    image: ab.image, // url
                    name: ab.metadata_name,
                    details: ab.collectionId
                } as ImageDetails
            }));
        })
}

// load NFT details by its UUID
function loadNFTDetails(UUID: string): Promise<NFTDetails> {
    return fetch(`${RMRK_KANARIA_API_NFT_DETAILS}/${UUID}`,{mode: 'cors'})
        .then(res => res.json())
        .then((res) => res as NFTDetails)
}

export default NftItems;