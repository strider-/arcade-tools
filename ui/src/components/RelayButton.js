import { Button, Skeleton } from '@mui/material';
import useRelay from '../hooks/useRelay';

const RelayButton = ({ relayId, icon, caption }) => {
    const { isOnline, enabled, isLoading, toggle } = useRelay(relayId)

    if (isLoading) {
        return <Skeleton variant="rectangular" />
    }

    return (
        <Button
            disabled={!isOnline && !isLoading}
            size="large"
            variant="contained"
            color={enabled ? 'success' : 'error'}
            startIcon={icon}
            onClick={toggle}>
            {caption}
        </Button>
    )
}

export default RelayButton;
