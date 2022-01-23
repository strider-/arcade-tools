import { Button, Skeleton } from '@mui/material';
import useRelay from '../hooks/useRelay';

const RelayButton = ({ relayId, icon }) => {
    const { isOnline, enabled, name, isLoading, toggle } = useRelay(relayId)

    if (isLoading) {
        return <Skeleton variant="rectangular" width={144} height={80} />
    }

    return (
        <Button
            sx={{ minWidth: 144 }}
            disabled={!isOnline && !isLoading}
            size="large"
            variant="contained"
            color={enabled ? 'success' : 'error'}
            startIcon={icon}
            onClick={toggle}>
            {name}
        </Button>
    )
}

export default RelayButton;
