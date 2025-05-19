'use client';
import SidebarCloseButton from '@/components/Sidebar/SidebarCloseButton';
import SidebarFloatingButton from '@/components/Sidebar/SidebarFloatingButton';
import SidebarItem from '@/components/Sidebar/SidebarItem';
import { useSelectedPatientContext } from '@/context/SelectedPatientContext';
import { usePathname } from 'next/navigation';
import React, { ReactNode, useState, useEffect, useRef } from 'react';

interface SidebarItem {
    id: string;
    label: string;
    href: string;
    badge?: string;
}

interface RootLayoutProps {
    children: ReactNode;
}

const sidebarItems: SidebarItem[] = [
    { id: 'dashboard', label: 'Dashboard', href: '/dashboard' },
    { id: 'newPatient', label: 'Novo Paciente', href: '/dashboard/newPatient', badge: 'v1' },
    { id: 'listPatients', label: 'Lista de Pacientes', href: '/dashboard/listPatients', badge: 'v1' },
];

export default function RootLayout({ children }: RootLayoutProps) {
    const currentPath = usePathname();
    const [sidebarOpen, setSidebarOpen] = useState(false);
    const [isSmallScreen, setIsSmallScreen] = useState(false);
    const { selectedPatient } = useSelectedPatientContext();
    const sidebarRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const mediaQuery = window.matchMedia('(max-width: 640px)');

        const handleMediaChange = (e: MediaQueryListEvent) => {
            setIsSmallScreen(e.matches);
            if (e.matches) setSidebarOpen(false);
            else setSidebarOpen(true);
        };

        setIsSmallScreen(mediaQuery.matches);
        setSidebarOpen(!mediaQuery.matches);

        mediaQuery.addEventListener('change', handleMediaChange);
        return () => mediaQuery.removeEventListener('change', handleMediaChange);
    }, []);

    // Handle click outside to close sidebar
    useEffect(() => {
        if (!(sidebarOpen && isSmallScreen)) return;

        const handleClickOutside = (event: MouseEvent) => {
            if (
                sidebarRef.current &&
                !sidebarRef.current.contains(event.target as Node)
            ) {
                setSidebarOpen(false);
            }
        };

        document.addEventListener('mousedown', handleClickOutside);
        return () => {
            document.removeEventListener('mousedown', handleClickOutside);
        };
    }, [sidebarOpen, isSmallScreen]);

    // If we have a selected patient, we want to have more options on the sidebar
    useEffect(() => {
        if (selectedPatient) {
            if (!sidebarItems.some(item => item.id === 'patientDetails')) {
                sidebarItems.push({
                    id: 'patientDetails',
                    label: 'Detalhes do Paciente',
                    href: `/dashboard/patient`,
                });
            }
            // newDiagnostic
            if (!sidebarItems.some(item => item.id === 'newDiagnostic')) {
                sidebarItems.push({
                    id: 'newDiagnostic',
                    label: 'Novo Exame',
                    href: `/dashboard/patient/newDiagnostic`,
                });
            }
        }
    }, [selectedPatient]);

    const toggleSidebar = () => setSidebarOpen((open) => !open);

    return (
        <div className="flex">
            <aside
                ref={sidebarRef}
                id="default-sidebar"
                className={`fixed top-0 left-0 z-40 w-64 h-screen transition-transform bg-gray-50 dark:bg-gray-800
          ${sidebarOpen ? 'translate-x-0' : '-translate-x-full'}
          sm:translate-x-0
        `}
                aria-label="Sidebar"
                aria-hidden={!sidebarOpen && isSmallScreen}
            >
                <div className="h-full px-3 py-4 overflow-y-auto flex flex-col justify-between">
                    <ul className="space-y-2 font-medium">
                        {sidebarItems.map(({ id, label, href, badge }) => {
                            const isActive = currentPath === href;
                            return (
                                <SidebarItem
                                    key={id}
                                    id={id}
                                    href={href}
                                    label={label}
                                    isActive={isActive}
                                    badge={badge}
                                />
                            );
                        })}
                    </ul>

                    {/* Toggle button inside sidebar for small screens */}
                    {isSmallScreen && (
                        <SidebarCloseButton
                            sidebarOpen={sidebarOpen}
                            toggleSidebar={toggleSidebar}
                        />
                    )}
                </div>
            </aside>

            {isSmallScreen && !sidebarOpen && (
                <SidebarFloatingButton
                    toggleSidebar={toggleSidebar}
                    isSmallScreen={isSmallScreen}
                    sidebarOpen={sidebarOpen}
                />
            )}

            {/* Main content */}
            <main className={`flex-1 p-1 pt-14 sm:pt-4 transition-margin duration-300 flex flex-col items-center min-h-screen ${sidebarOpen && !isSmallScreen ? 'sm:ml-64' : ''}`}>
                <div className="grid grid-cols-12 gap-4 w-full">
                    {children}
                </div>
            </main>
        </div>
    );
}
