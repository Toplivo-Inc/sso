import { Client } from './model'
import { DataTable } from './table'
import { columns } from './columns'

export const clients: Client[] = [
    {
        id: "00000000-0000-0000-0000-000000000000",
        name: "Example",
    },
    {
        id: "00000001-0000-0000-0000-000000000000",
        name: "Lazy",
        description: "Fucking app",
    }
]

async function getData(): Promise<Client[]> {
    return clients
}


export async function Clients() {
    const data = await getData()

    return (
        <div className="p-10 bg-zinc-900 w-[100%] h-[100%]">
            <DataTable columns={columns} data={data} />
        </div>
    )
}
