import {
    configureStore,
} from '@reduxjs/toolkit';
import removSliceReducer from "./removReducerSlice";

export const store = configureStore({
    reducer: {
        removReducer: removSliceReducer,
    },
});

export type RootState = ReturnType<typeof store.getState>;
