import { useState } from "react"
import { useQuery } from "react-query"
import { getRelay, setRelay } from "./api"

const useRelay = (relayId) => {
    const [enabled, setEnabled] = useState()
    const { data, isLoading } = useQuery(["v1", "relay", { id: relayId }], async () => await getRelay(relayId),
        {
            onSettled: (data) => setEnabled(data?.result?.state === 1 ? true : false)
        })
    const { isOnline, result: { name } = {} } = data || {};

    const toggle = async () => {
        if (isOnline) {
            await setRelay(relayId, !enabled)
            setEnabled(!enabled)
        }
    }

    return { isOnline, enabled, name, isLoading, toggle }
}

export default useRelay;