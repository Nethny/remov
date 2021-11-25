import React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import AppRouter from './Routes';
import { store } from './store';
import { Provider } from 'react-redux';
import { NavBar } from './components/AppNavbar';
import styles from "./App.module.scss";
import { ArrowRight } from "react-feather";

function App() {
  return (
      <Provider store={store}>
          <Router>

              <div className={styles.appBackground}>
                  <div className={styles.divider}>
                      <div className={styles.dividerPusher}></div>
                      <svg data-name="Layer 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1200 120"
                           preserveAspectRatio="none">
                          <path d="M1200 120L0 16.48 0 0 1200 0 1200 120z" className={styles.shapeFill}></path>
                      </svg>
                  </div>

                <div className={styles.app}>

                    <div className={styles.logo}>
                        <div><a href={"https://remov.app/"}>ReMov</a></div>
                        <p>your bridge between RMRK and Movr</p>
                    </div>

                  <header className={styles.navbar}>
                    <NavBar
                      options={[
                          { id: 1, children: <span>RMRK <ArrowRight /> Movr</span>, href: '/to-movr' },
                          { id: 2, children: <span>Movr <ArrowRight /> RMRK</span>, href: '/to-rmrk' },
                       ]} />
                    </header>

                    <AppRouter />
                </div>

              </div>

              <div className={styles.footer}>
                  <div className={styles.text}>
                      Powered by Remov team in 2021
                  </div>
              </div>

          </Router>
      </Provider>
  );
}

export default App;
