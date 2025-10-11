import { ConsoleLayout } from '@/components/console-layout';
import { Outlet } from 'react-router';

export const consoleRoutes = [
    {
        path: "/console",
        element: (
            <ConsoleLayout>
                <Outlet />
            </ConsoleLayout>
        ),
        children: [
            {
                path: "/console/clients",
                lazy: async () => {
                    const { Clients } = await import('@/pages/clients');
                    return { Component: Clients };
                }
            }
        ],
    },
]
