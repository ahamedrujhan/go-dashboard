import {Paper} from "@mui/material";


export default function Tile({children}) {
    return (
        <Paper elevation={3} sx={{padding: 2, margin: 2, borderRadius: '10px'}}>
            {children}
        </Paper>
    )
}