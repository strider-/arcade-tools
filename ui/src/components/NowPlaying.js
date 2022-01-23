import { Card, CardMedia, Skeleton } from '@mui/material';
import useNowPlaying from '../hooks/useNowPlaying';

const getImage = (isOnline, isInMenus, game) => {
    if (!isOnline) {
        return ["Offline", "/offline.png"]
    } else if (isInMenus) {
        return ["In Menus", "/menus.png"]
    } else {
        return [game.romName, `/${game.romName}.png`]
    }
}

const NowPlaying = () => {
    const { isOnline, isLoading, isInMenus, game } = useNowPlaying()

    if (isLoading) {
        return <Skeleton variant="rectangular" width={303} height={450} sx={{ my: 2 }} />
    }

    const [imgAlt, imgPath] = getImage(isOnline, isInMenus, game)

    return (
        <Card sx={{ minWidth: 275, minHeight: 450, my: 2 }} variant="outlined">
            <CardMedia
                component="img"
                alt={imgAlt}
                image={imgPath} />
        </Card>
    )
}

export default NowPlaying;
