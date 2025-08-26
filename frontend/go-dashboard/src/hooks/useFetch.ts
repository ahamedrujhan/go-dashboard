import {useEffect, useState} from "react";
import api from "../api/axios.ts";

interface UseFetchOptions {
    immediate?: boolean; // whether to fetch immediately
}

export function useFetch<T = any>(url:string, options:UseFetchOptions = {}) {
    const {immediate = true} = options;

    const [data, setData] = useState< T | null>(null)
    const [loading, setLoading] = useState<boolean>(immediate)
    const [error, setError] = useState<any>(null)

    const fetchData = async () => {
        try {
            setLoading(true)
            const response = await api.get<T>(url)
            setData(response.data)
        } catch (err : any) {
            setError(err)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        if (immediate) {
            fetchData()
        }
    }, [url]);

    return {data, loading, error, refetch: fetchData}
}

