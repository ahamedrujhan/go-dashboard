import { BarChart } from '@mui/x-charts/BarChart';
import {Checkbox, FormControlLabel, FormGroup} from "@mui/material";
import Box from "@mui/material/Box";
import {useState} from "react";
export default function BarCharts(data) {
    const [showSold, setShowSold] = useState(true)
    const [showStock, setShowStock] = useState(true)
    const products = data.data?.products || []
    const soldQuantity = data.data?.soldQuantity || []
    const stockQuantity = data.data?.stockQuantity || []

    const series = [
        showSold && { data: soldQuantity, label: "Sold Quantity", color:"#4254FB" },
        showStock && { data: stockQuantity, label: "Stock Quantity", color:"#FFB422" }
    ].filter(Boolean)

    return (
        <>
            <Box sx={{display:"flex", justifyContent:"center"}}>
                <FormControlLabel control={<Checkbox checked={showSold} onChange={() => setShowSold(!showSold)} />} label="Sold" />
                <FormControlLabel control={<Checkbox checked={showStock} onChange={() => setShowStock(!showStock)} />} label="Stock"  />
            </Box>

        <BarChart
            xAxis={[{ data: products }]}
            series={series}
            height={300}
        />
        </>
    )
}