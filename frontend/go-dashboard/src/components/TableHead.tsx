import Typography from "@mui/material/Typography";

export default function TableHead({Heading}) {
    return (
        <Typography variant="h6" sx={{fontWeight: 'bold'}} align={'center'}>
            {Heading}
        </Typography>
    )
}