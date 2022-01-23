import { Card, CardMedia, Grid, Stack } from '@mui/material';
import LightModeIcon from '@mui/icons-material/LightMode';
import MonitorIcon from '@mui/icons-material/Monitor';
import RelayButton from './components/RelayButton';

const Home = () => {
    return (
        <Grid
            container
            spacing={0}
            direction="column"
            alignItems="center"
            justifyContent="center"
            style={{ minHeight: '90vh' }}
        >
            <Grid item xs={3}>
                <Card sx={{ minWidth: 275, minHeight: 450, my: 2 }} variant="outlined">
                    <CardMedia
                        component="img"
                        alt="placeholder"
                        image="/npph.jpg" />
                </Card>
                <Stack direction="row" spacing={2}>
                    <RelayButton
                        relayId={1}
                        caption="Lighting"
                        icon={<LightModeIcon />}
                    />
                    <RelayButton
                        relayId={2}
                        caption="Monitor"
                        icon={<MonitorIcon />}
                    />
                </Stack>
            </Grid>
        </Grid>
    )
}

export default Home;
