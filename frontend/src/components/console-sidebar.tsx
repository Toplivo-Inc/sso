import { Sidebar, SidebarContent, SidebarGroup, SidebarGroupContent, SidebarGroupLabel, SidebarMenu, SidebarMenuButton, SidebarMenuItem } from "@/components/ui/sidebar"
import { cn } from "@/lib/utils"

const items = [
    {
        title: "Clients",
        url: "clients",
    },
    {
        title: "Client scopes",
        url: "scopes",
    },
    {
        title: "Users",
        url: "users",
    },
    {
        title: "Sessions",
        url: "sessions",
    },
]

interface ConsoleSidebarProps extends React.ComponentProps<"div"> { }

export function ConsoleSidebar({ className, ...props }: ConsoleSidebarProps) {
    return (
        <Sidebar variant="floating" collapsible="none" className={cn(className)} {...props}>
            <SidebarContent>
                <SidebarGroup>
                    <SidebarGroupLabel>Application</SidebarGroupLabel>
                    <SidebarGroupContent>
                        <SidebarMenu>
                            {items.map((item) => (
                                <SidebarMenuItem key={item.title}>
                                    <SidebarMenuButton asChild>
                                        <a href={item.url}>
                                            <span>{item.title}</span>
                                        </a>
                                    </SidebarMenuButton>
                                </SidebarMenuItem>
                            ))}
                        </SidebarMenu>
                    </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
        </Sidebar>
    )
}
