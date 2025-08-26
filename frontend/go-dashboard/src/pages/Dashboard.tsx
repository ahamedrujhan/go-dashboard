import Box from "@mui/material/Box";
import Tile from "../components/Tile.tsx";
import TableHead from "../components/TableHead.tsx";
import {Grid} from "@mui/material";
import {lazy, Suspense} from "react";
import MonthlySalesSkeliton from "../components/skelitons/MonthlySalesSkeliton.tsx";
import {useFetch} from "../hooks/useFetch.ts";
import StackBarChart from "../components/charts/StackBarChart.tsx";
import Typography from "@mui/material/Typography";
import Loader from "../components/skelitons/Loader.tsx";
import RevenueByCountry from "../components/DataGrid/RevenueByCountry.tsx";

const MonthlySales = lazy(() => import("../components/charts/MonthlySales.tsx")as Promise<{ default: React.ComponentType<any> }>);
const TopProducts = lazy(() => import("../components/charts/BarCharts.tsx")as Promise<{ default: React.ComponentType<any> }>);
const Top30Regions = lazy(() => import("../components/charts/StackBarChart.tsx")as Promise<{ default: React.ComponentType<any> }>);

export default function Dashboard() {

    // fetch the data
    const {data, loading, error, refetch} = useFetch("/monthly-revenue")
    const {
        data: monthlyRevenueData,
        loading: monthlyRevenueLoading,
        error: monthlyRevenueError,
        refetch: refetchMonthlyRevenue } = useFetch("/top-20-products")
    const {
        data: regionsData,
        loading: regionsLoading,
        error: regionsError,
        refetch: refetchRegions } = useFetch("/top-30-regions")



    return (
        <Box>
            <Grid container spacing={2}
                  // sx={{border:"solid black"}}
            >
                <Grid item size={12}>
                    <Typography align={"center"} variant={"h3"} sx={{fontWeight: 'bold'}} >Sales Analytics DashBoard </Typography>
                </Grid>
                {/*Monthly Sales*/}
                <Grid item size={12}>

                    <Tile>
                        <Grid container spacing={2}>
                           <Grid item size={12}>
                                   <TableHead Heading={"Monthly Sales"}/>
                           </Grid>
                            <Grid item size={12}>
                                {loading ?  <Loader/> :
                                    <Suspense fallback={<MonthlySalesSkeliton />}>
                                        <MonthlySales data={data?.data} />
                                    </Suspense>}
                            </Grid>
                        </Grid>
                    </Tile>
                </Grid>

                {/*Top 20 Products*/}
                <Grid item size={12}>

                    <Tile>
                        <Grid item size={12}>
                            <TableHead Heading={"Top 20 Products"}/>
                        </Grid>
                        <Grid item size={12}>
                            {monthlyRevenueLoading ?  <Loader/> :
                                <Suspense fallback={<MonthlySalesSkeliton />}>
                                    <TopProducts data={monthlyRevenueData?.data}/>
                                </Suspense>}

                        </Grid>
                    </Tile>
                </Grid>
                {/*Top 30 Regions by Revenue*/}
                <Grid item size={12}>
                    <Tile>
                        <Grid item size={12}>
                            <TableHead Heading={"Top 30 Regions by Revenue"}/>
                        </Grid>
                        <Grid item size={12}>
                            {regionsLoading ?  <Loader/> :
                                <Suspense fallback={<MonthlySalesSkeliton />}>
                                    <Top30Regions data={regionsData?.data} />
                                </Suspense>}
                        </Grid>
                    </Tile>
                </Grid>
                <Grid item size={12}>
                    <Tile>
                        <Grid item size={12}>
                            <TableHead Heading={"Revenue By Country and Product"}/>
                        </Grid>
                        <Grid item size={12}>

                           <RevenueByCountry/>
                        </Grid>
                    </Tile>
                </Grid>

            </Grid>


        </Box>
    )
}