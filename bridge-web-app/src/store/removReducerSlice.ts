import {createSlice, current, PayloadAction} from '@reduxjs/toolkit';

export interface RemovState {
    selectedNFT: string; // ID of NFT
    currentStepNumber: number; // from 1 to 4
    selectedPolkadotWalletAddress: string,
    selectedMetamaskWalletAddress: string,
}

const initialState: RemovState = {
    selectedNFT: "",
    currentStepNumber: 1,
    selectedPolkadotWalletAddress: "",
    selectedMetamaskWalletAddress: "",
}

export const removSlice = createSlice({
    name: 'remov',
    initialState,
    reducers: {

        resetStore: (state) => {
            return {
                ...initialState
            };
        },

        setSelectedNFT: (state, action: PayloadAction<string>) => {
            return {
                ...state,
                selectedNFT: action.payload,
            }
        },

        setCurrentStepNumber: (state, action: PayloadAction<number>) => {
            return {
                ...state,
                currentStepNumber: action.payload,
            }
        },

        setSelectedPolkadotAddress: (state, action:PayloadAction<string>) => {
            return {
                ...state,
                selectedPolkadotWalletAddress: action.payload,
            }
        },

        setSelectedMetamaskAddress: (state, action: PayloadAction<string>) => {
            return {
                ...state,
                selectedMetamaskWalletAddress: action.payload,
            }
        },

    }
});

export const {
    resetStore,
    setSelectedNFT,
    setCurrentStepNumber,
    setSelectedPolkadotAddress,
    setSelectedMetamaskAddress,
} = removSlice.actions;

export default removSlice.reducer;