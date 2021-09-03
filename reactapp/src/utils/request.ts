import axios from 'axios';
import { IResponse } from '../api/typings';

export type Method = 'GET' | 'POST' | 'PUT' | 'DELETE'
export type ResponseType = 'blob' | 'json' | 'text' | 'stream'

export interface AxiosRequest {
    baseURL?: string;
    url: string;
    data?: any;
    params?: any;
    method?: Method;
    headers?: any;
    timeout?: number;
    responseType?: ResponseType;
}

export interface AxiosResponse {
    data: IResponse;
    headers: any;
    request?: any;
    status: number;
    statusText: string;
    config: AxiosRequest;
}

export interface CustomResponse {
    readonly status: boolean;
    readonly message: string;
    data: any;
    origin?: any;
}

export interface GetDemo {
    id: number;
    str: string;
}

export interface PostDemo {
    id: number;
    list: Array<{
        id: number;
        version: number;
    }>;
}

axios.create({
  baseURL: 'http://localhost:8001/api',
  headers: {
    'Content-Type': 'application/json',
  },
});
