import React, { FC } from 'react';
import { NavLink } from 'react-router-dom';
import styles from './navbar.module.scss';

type NavbarOption = {
    id: number;
    children: React.ReactNode;
    href: string;
};

type NavbarProps = {
    options: NavbarOption[];
};

const NavBar: FC<NavbarProps> = ({ options }) => (
    <nav className={styles.navbar}>
        {options.map((option) => (
            <NavLink key={option.id}
                     to={option.href}
                     className={({ isActive }) => styles.navLink + (isActive ? " " + styles.active : "") + " " + (option.id == 1 ? styles.left : styles.right) }>
                {option.children}
            </NavLink>
        ))}
    </nav>
);

export default NavBar;