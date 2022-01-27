import { Card, CardHeader, CardMedia, IconButton, Skeleton } from '@mui/material';
import useNowPlaying from '../hooks/useNowPlaying';
import CancelIcon from '@mui/icons-material/Cancel';
import { useEffect, useState } from 'react';

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
    const { isOnline, isLoading, isInMenus, game, kill } = useNowPlaying()
    const [killDisabled, setKillDisabled] = useState(false)

    useEffect(() => {
        setKillDisabled(() => game === null)
    }, [game])

    if (isLoading) {
        return <Skeleton variant="rectangular" width={303} height={450} sx={{ my: 2 }} />
    }

    const [imgAlt, imgPath] = getImage(isOnline, isInMenus, game)

    const stopGame = () => {
        setKillDisabled(true)
        kill()
    }

    return (
        <Card sx={{ minWidth: 275, minHeight: 450, my: 2 }} variant="outlined">
            {game && <CardHeader
                title="Now Playing"
                subheader={game.romName}
                action={
                    <IconButton color="error"
                        onClick={stopGame}
                        disabled={killDisabled}>
                        <CancelIcon />
                    </IconButton>
                }
            />}
            <CardMedia
                component="img"
                alt={imgAlt}
                image={imgPath} />
        </Card>
    )
}

export default NowPlaying;
