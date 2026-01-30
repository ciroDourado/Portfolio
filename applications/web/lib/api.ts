const getBaseUrl = () => {
    if (typeof window === 'undefined') {
        return process.env.API_INTERNAL_URL;
    }
    return process.env.NEXT_PUBLIC_API_URL;
};

export async function fetchAPI<T>(endpoint: string): Promise<T> {
    const baseUrl = getBaseUrl();
    const url = `${baseUrl}${endpoint}`;

    const res = await fetch(url, {
        next: {revalidate: 60},
    });

    if (!res.ok) {
        throw new Error(`Erro na API: ${res.statusText}`);
    }

    return res.json();
}