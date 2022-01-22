import { Button } from '@mui/material';

const RelayButton = ({ enabled, icon, caption, onToggle }) => {
    return (
        <Button
            size="large"
            variant="contained"
            color={enabled ? 'success' : 'error'}
            startIcon={icon}
            onClick={() => onToggle(!enabled)}>
            {caption}
        </Button>
    )
}

export default RelayButton;
