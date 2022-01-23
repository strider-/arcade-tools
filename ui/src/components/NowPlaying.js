import { Card, CardMedia, Skeleton } from '@mui/material';
import useNowPlaying from '../hooks/useNowPlaying';

const NowPlaying = () => {
    const { isOnline, isLoading, isInMenus, game } = useNowPlaying()

    if (isLoading) {
        return <Skeleton variant="rectangular" width={303} height={450} sx={{ my: 2 }} />
    }

    var img

    if (!isOnline) {
        img = "/offline.png"
    } else if (isInMenus) {
        img = "/menus.png"
    } else {
        img = `/${game.romName}.png`
    }

    return (
        <Card sx={{ minWidth: 275, minHeight: 450, my: 2 }} variant="outlined">
            <CardMedia
                component="img"
                alt={game?.romName || 'Current Status'}
                image={img} />
        </Card>
    )
}

export default NowPlaying;
