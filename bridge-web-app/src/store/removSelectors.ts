import { RootState } from './index';


export const getSelectedNFT = (state: RootState): string  =>
    state.removReducer.selectedNFT;
export const getCurrentStepNumber = (state: RootState): number  =>
    state.removReducer.currentStepNumber;
export const getSelectedPolkadotAddress = (state: RootState): string =>
    state.removReducer.selectedPolkadotWalletAddress;
export const getSelectedMetamaskAddress = (state: RootState): string =>
    state.removReducer.selectedMetamaskWalletAddress;