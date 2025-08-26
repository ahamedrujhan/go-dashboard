import { useState, useEffect } from "react";
import { LineChart } from '@mui/x-charts/LineChart';
import { Box, TextField } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import dayjs, { Dayjs } from 'dayjs';

export default function MonthlySales(data) {
    const months = data?.data?.months || []
    const soldQuantity = data?.data?.soldQuantity || []
    const totalRevenue = data?.data?.totalRevenue || []
    const margin = { right: 24 };

    // Convert month strings to Date objects for DatePicker
    const monthDates = months.map(m => {
        var date = dayjs().year(m.split("-")[0]).month(m.split("-")[1]-1).date(1)
        // console.log(date)

        return (date)

    });
    const [fromDate, setFromDate] = useState<Dayjs | null>(dayjs(monthDates[0]));
    const [toDate, setToDate] = useState<Dayjs | null>(dayjs(monthDates[monthDates.length - 1])); // latest





    // Filter data based on selected dates
    const filteredIndexes = monthDates
        .map((date, idx) => ({ date, idx }))
        .filter(d =>
            (!fromDate || !d.date.isAfter(fromDate)) && // keep if after or equal fromDate
            (!toDate   || !d.date.isBefore(toDate))      // keep if before or equal toDate
        )
        .map(d => d.idx);


    const filteredMonths = filteredIndexes.map(idx => months[idx]);
    const filteredRevenue = filteredIndexes.map(idx => totalRevenue[idx]);
    const filteredQuantity = filteredIndexes.map(idx => soldQuantity[idx]);




    return (
<>
        <Box sx={{ display: "flex", gap: 2, mb: 2 }}>
            <DatePicker
                label="From"
                value={fromDate}
                onChange={(newValue) => setFromDate(newValue)}
                renderInput={(params) => <TextField {...params} />}
            />

            <DatePicker
                label="To"
                value={toDate}
                onChange={(newValue) => setToDate(newValue)}
                renderInput={(params) => <TextField {...params} />}
            />

        </Box>

        <LineChart
            height={300}
            series={[
                // { data: soldQuantity, label: 'sold' },
                { data: filteredRevenue, label: 'revenue' },
            ]}
            xAxis={[{ scaleType: 'point', data: filteredMonths }]}
            // yAxis={[{ width: 50 }]}
            margin={margin}

        />
</>
    )
}