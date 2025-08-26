import * as React from 'react';
import Box from '@mui/material/Box';
import { DataGrid } from '@mui/x-data-grid';
import { useEffect, useState } from 'react';
import api from "../../api/axios.ts";
import type {GridColDef, GridPaginationModel} from '@mui/x-data-grid';

const columns: GridColDef[] = [
    { field: 'Country', headerName: 'Country', width: 200, headerAlign: 'center', align: 'center' },
    { field: 'ProductName', headerName: 'ProductName', width: 250, headerAlign: 'center', align: 'center' },
    { field: 'TotalRevenue', headerName: 'TotalRevenue', type: 'number', width: 200, headerAlign: 'center', align: 'center' },
    { field: 'TransactionCount', headerName: 'TransactionCount', type: 'number', width: 200, headerAlign: 'center', align: 'center' },
];

export default function RevenueByCountry() {
    const [rows, setRows] = useState<any[]>([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<any>(null);

    const [paginationModel, setPaginationModel] = useState<GridPaginationModel>({
        page: 0,
        pageSize: 50,
    });
    const [rowCount, setRowCount] = useState(0);

    useEffect(() => {
        const fetchData = async () => {
            setLoading(true);
            try {
                const res = await api.get(`country-product-revenue?page=${paginationModel.page + 1}&perPage=${paginationModel.pageSize}`);
                const data = res.data.data || [];

                setRows(
                    data.map((row: any, index: number) => ({
                        ...row,
                        id: `${paginationModel.page * paginationModel.pageSize + index}`, // unique ID
                    }))
                );

                setRowCount(res.data.totalRecords || 0); // Total rows from API
                setError(null);
            } catch (err) {
                setError(err);
                setRows([]);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [paginationModel]);

    return (
        <Box sx={{ height: 600, width: '100%' }}>
            <DataGrid
                rows={rows}
                columns={columns}
                pagination
                paginationMode="server"
                rowCount={rowCount}
                pageSizeOptions={[10, 25, 50, 100]}
                paginationModel={paginationModel}
                onPaginationModelChange={setPaginationModel}
                loading={loading}
                sortModel={[
                    {
                        field: 'TotalRevenue',
                        sort: 'desc',
                    },
                ]}
                disableRowSelectionOnClick
            />
        </Box>
    );
}
