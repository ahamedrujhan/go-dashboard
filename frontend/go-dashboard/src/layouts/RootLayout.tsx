// Root layout

import Box from "@mui/material/Box";
import {Outlet} from "react-router-dom";
export default function RootLayout() {

return (
    <Box sx={{
        display:"flex",
        minHeight: "100vh",
        p:1
    }}>
        {/*any modifications to root layout you can add*/}

        <main>
            <Outlet/>
        </main>
    </Box>
)

}