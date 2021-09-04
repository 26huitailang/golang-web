import axios, { AxiosRequestConfig } from 'axios';
import { Method } from './typings';

interface PendingType {
    url?: string;
    method?: Method;
    params: any;
    data: any;
    // eslint-disable-next-line @typescript-eslint/ban-types
    cancel: Function;
}

const pending: Array<PendingType> = [];
const { CancelToken } = axios;

const instance = axios.create({
  timeout: 10000,
  responseType: 'json',
});

// 请求拦截器
instance.interceptors.request.use(
  (request) => request,
  (error) => Promise.reject(error),
);

// 响应拦截器
instance.interceptors.response.use(
  (response) => {
    const { code } = response.data;
    switch (code) {
      case [123, 123]:
        break;
      default:
        break;
    }
    return response;
  },
  (error) => {
    const { response } = error;
    switch (response.status) {
      case 401:
        break;
      case 403:
        break;
      case 500:
        break;
      case 503:
        break;
      default:
        break;
    }
    return Promise.reject(response);
  },
);

export default instance;
