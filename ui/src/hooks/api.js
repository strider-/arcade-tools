const version = "api/v1"
const baseUrl = process.env.REACT_APP_ARCADE_API_URL.replace(/^\//, '')
const requestTimeout = 3000

const fetchWithTimeout = async (url, opts) => {
    try {
        const controller = new AbortController()

        setTimeout(() => controller.abort(), requestTimeout)
        const resp = await fetch(url, { signal: controller.signal, ...opts })

        return { isOnline: true, status: resp.status, result: await resp.json() }
    } catch (error) {
        return { isOnline: false }
    }
}

const apiGet = async (endpoint) => {
    return await fetchWithTimeout(`${baseUrl}/${version}/${endpoint.replace(/^\//, '')}`);
};

const apiPost = async (endpoint, payload) => {
    return await fetchWithTimeout(`${baseUrl}/${version}/${endpoint.replace(/^\//, '')}`, {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    });
}

export const getRelays = async () =>
    await apiGet('/relays')

export const getRelay = async (relayId) =>
    await apiGet(`/relays/${relayId}`)

export const setRelay = async (relayId, enabled) =>
    await apiPost(`/relays/${relayId}`, { state: enabled ? 1 : 0 })

export const nowPlaying = async () =>
    await apiGet('/now-playing')
