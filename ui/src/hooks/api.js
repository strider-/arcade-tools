const version = "/api/v1"
const baseUrl = process.env.REACT_APP_ARCADE_API_URL.replace(/^\//, '')

const apiGet = async (endpoint) => {
    const resp = await fetch(`${baseUrl}/${version}/${endpoint.replace(/^\//, '')}`);
    const result = await resp.json();

    return { code: resp.status, result };
};

const apiPost = async (endpoint, formData) => {
    const resp = await fetch(`${baseUrl}/${version}/${endpoint.replace(/^\//, '')}`, {
        method: 'POST',
        body: formData
    });

    return resp.ok
}

export const getRelays = async () => {
    await apiGet('/relays')
}

export const getRelay = async (relayId) => {
    await apiGet(`/relay/${relayId}`)
}

export const setRelay = async (relayId, enabled) => {
    await apiPost(`/relay/${relayId}`, { state: enabled ? 1 : 0 })
}

export const nowPlaying = async () => {
    await apiGet('/now-playing')
}