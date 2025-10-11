import { ColumnDef } from "@tanstack/react-table";
import { Client } from "./model";

export const columns: ColumnDef<Client>[] = [
	{
		accessorKey: "id",
		header: "Client ID",
	},
	{
		accessorKey: "name",
		header: "Client Name",
	},
	{
		accessorKey: "description",
		header: "Description",
	},
]
