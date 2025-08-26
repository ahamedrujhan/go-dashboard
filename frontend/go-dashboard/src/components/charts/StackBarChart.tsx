import * as React from 'react';
import { PieChart } from '@mui/x-charts/PieChart';
import {useFetch} from "../../hooks/useFetch.ts";
import {BarChart} from "@mui/x-charts/BarChart";


export default function StackBarChart(data) {

    var regions = data?.data?.regions || []
    var sold = data?.data?.sold || []
    var revenue = data?.data?.revenue || []

    var soldData = regions.map((region, idx) => ({ label: region, value: sold[idx] }))
    var revenueData = regions.map((region, idx) => ({ label: region, value: revenue[idx] }))

    return (
        // <PieChart
        //     series={[
        //         {
        //             innerRadius: 0,
        //             outerRadius: 80,
        //             data: soldData,
        //             highlightScope: { fade: 'global', highlight: 'item' },
        //             faded: { innerRadius: 30, additionalRadius: -30, color: 'gray' },
        //         },
        //         {
        //             innerRadius: 100,
        //             outerRadius: 120,
        //             data: revenueData,
        //             // highlightScope: { fade: 'global', highlight: 'item' },
        //             // faded: { innerRadius: 30, additionalRadius: -30, color: 'gray' },
        //         },
        //     ]}
        //     height={300}
        //     hideLegend
        // />
        <BarChart
            xAxis={[{ data: regions }]}
            series={[{ data: sold, label: "Sold Quantity",  id: 'sold', stack: 'total' },
                { data: revenue, label: "Revenue",  id: 'revenue', stack: 'total'  }]}
            height={300}

        />

    );
}
