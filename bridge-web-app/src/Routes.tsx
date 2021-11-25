import React, { FC } from 'react';
import { Route, Routes } from 'react-router-dom';
import StepByStepToMOVR from "./components/StepByStepWidget/StepByStepToMOVR";
import StepByStepBackToRMRK from "./components/StepByStepWidget/StepByStepBackToRMRK";

const AppRouter: FC = () => (
 <Routes>
     <Route path="/to-movr" element={<StepByStepToMOVR />} />
     <Route path="/to-rmrk" element={<StepByStepBackToRMRK />}/>
 </Routes>
);

export default AppRouter;