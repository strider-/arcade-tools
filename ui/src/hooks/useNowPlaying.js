import { useState } from "react";
import { useQuery } from "react-query";
import { killRetroArch, nowPlaying } from "./api";

const interval = 5000

const useNowPlaying = () => {
    const [refetchInterval, setRefetchInterval] = useState(interval)

    const { data, isLoading } = useQuery(["v1", "now-playing"], async () => await nowPlaying(), {
        refetchInterval: refetchInterval,
        onError: () => setRefetchInterval(0)
    })

    const { isOnline, status, result: game } = data || {};

    const kill = async () => {
        await killRetroArch()
    }

    return { isOnline, isLoading, isInMenus: status === 204, game, kill }
}

export default useNowPlaying;
