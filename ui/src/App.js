import { createTheme, ThemeProvider } from '@mui/material/styles';
import { CssBaseline } from '@mui/material';
import Box from '@mui/material/Box';

const theme = createTheme({ palette: { mode: "dark" } });

function App() {
    return (
        <ThemeProvider theme={theme}>
            <CssBaseline />
            <Box sx={{ flexGrow: 1 }} m={2}>

            </Box>
        </ThemeProvider>
    );
}

export default App;
