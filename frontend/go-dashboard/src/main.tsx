import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers';
// import './index.css'
import App from './App.tsx'
import {createTheme, ThemeProvider} from "@mui/material";

const theme = createTheme();

createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <ThemeProvider theme={theme}>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
        <App />
          </LocalizationProvider>
      </ThemeProvider>
  </StrictMode>,
)
