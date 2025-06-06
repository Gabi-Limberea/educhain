import React, { ReactNode, FC, useState, useEffect } from 'react';
import '../static/home.css';

interface LinkItem {
    href: string;
    label: string;
}

interface MenuItem {
    id: string;
    label: string;
}

interface LayoutProps {
    leftLinks: LinkItem[];
    children: ReactNode;
    menuItems: MenuItem[];
}

const Layout: FC<LayoutProps> = ({ leftLinks, children, menuItems }) => {
const [activeId, setActiveId] = useState<string>(menuItems[0]?.id);

useEffect(() => {
    const observer = new IntersectionObserver(
    (entries) => {
        entries.forEach(entry => {
        if (entry.isIntersecting) {
            setActiveId(entry.target.id);
        }
        });
    },
    { root: null, rootMargin: '0px', threshold: 0.6 }
    );

    menuItems.forEach(item => {
    const el = document.getElementById(item.id);
    if (el) observer.observe(el);
    });

    return () => { observer.disconnect(); };
}, [menuItems]);

const scrollToSection = (id: string) => {
    const el = document.getElementById(id);
    if (el) el.scrollIntoView({ behavior: 'smooth', block: 'start' });
};

return (
    <div className="layout">
    <aside className="left">
        <h2>Utilities</h2>
        <ul>
        {leftLinks.map(link => (
            <li key={link.href}><a href={link.href}>{link.label}</a></li>
        ))}
        </ul>
    </aside>

    <main className="main">
        {children}
    </main>

    <aside className="right">
        <h2>Contents</h2>
        <ul>
        {menuItems.map(item => (
            <li key={item.id} className={activeId === item.id ? 'active' : ''}>
            <button onClick={() => scrollToSection(item.id)}>
                {item.label}
            </button>
            </li>
        ))}
        </ul>
    </aside>
    </div>
)};
  
  
export default Layout;
