import {CircularProgress, Box} from "@mui/material";
import Typography from "@mui/material/Typography";

export default function Loader() {
    return (
        <Box sx={{
            display:"flex",
            justifyContent:"center",
            alignItems:"center",
            p:10
        }}>
            <CircularProgress />
            <Typography variant={"body1"} sx={{ml:2}}>Loading...</Typography>
        </Box>
    )
}