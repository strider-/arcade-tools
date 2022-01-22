import { Card, CardMedia, Grid, Stack } from '@mui/material';
import LightModeIcon from '@mui/icons-material/LightMode';
import MonitorIcon from '@mui/icons-material/Monitor';
import { useState } from 'react';
import RelayButton from './components/RelayButton';

const Home = () => {
    const [isLit, setIsLit] = useState(true)
    const [hasDisplay, setHasDisplay] = useState(true)

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
                <Stack direction="row" spacing={2} sx={{ height: 80 }}>
                    <RelayButton
                        caption="Lighting"
                        icon={<LightModeIcon />}
                        onToggle={setIsLit}
                        enabled={isLit}
                    />
                    <RelayButton
                        caption="Monitor"
                        icon={<MonitorIcon />}
                        onToggle={setHasDisplay}
                        enabled={hasDisplay}
                    />
                </Stack>
            </Grid>
        </Grid>
    )
}

export default Home;
