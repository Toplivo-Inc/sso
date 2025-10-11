import { ConsoleSidebar } from '@/components/console-sidebar'
import { SidebarProvider, } from '@/components/ui/sidebar'
import { Header } from '@/widgets/Header'

interface ConsoleLayoutProps {
    children?: React.ReactNode;
}

export function ConsoleLayout({ children }: ConsoleLayoutProps) {
    return (
        <div className="grid grid-rows-[auto_1fr] grid-cols-[auto_1fr] h-screen">
            <Header className='row-start-1 col-span-2 z-30' />
            <SidebarProvider>
                <ConsoleSidebar className='row-start-2 h-full border-2' />
                <main className='row-start-2 col-start-2 overflow-auto'>
                    {children}
                </main>
            </SidebarProvider>
        </div>
    )
}
