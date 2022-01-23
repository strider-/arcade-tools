import { useState } from "react"
import { useQuery } from "react-query"
import { getRelay, setRelay } from "./api"

const useRelay = (relayId) => {
    const [enabled, setEnabled] = useState()
    const { data, isLoading } = useQuery(["v1", "relay", { id: relayId }], async () => await getRelay(relayId),
        {
            onSettled: (data) => setEnabled(data.result.state)
        })
    const { isOnline } = data || {};

    const toggle = async () => {
        if (isOnline) {
            await setRelay(relayId, enabled ^ 1)
            setEnabled(enabled ^ 1)
        }
    }

    return { isOnline, enabled: enabled === 1, isLoading, toggle }
}

export default useRelay;