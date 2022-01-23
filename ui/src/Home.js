import { Grid, Stack } from '@mui/material';
import LightModeIcon from '@mui/icons-material/LightMode';
import MonitorIcon from '@mui/icons-material/Monitor';
import RelayButton from './components/RelayButton';
import NowPlaying from './components/NowPlaying';

const Home = () => {
    return (
        <Grid
            container
            spacing={0}
            direction="column"
            alignItems="center"
            justifyContent="center"
            style={{ minHeight: '90vh' }}>
            <Grid item xs={3}>
                <NowPlaying />
                <Stack direction="row" spacing={2} sx={{ height: 80 }}>
                    <RelayButton relayId={1} icon={<LightModeIcon />} />
                    <RelayButton relayId={2} icon={<MonitorIcon />} />
                </Stack>
            </Grid>
        </Grid>
    )
}

export default Home;
