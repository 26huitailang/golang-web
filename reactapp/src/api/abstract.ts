import instance from './instance';
import { AxiosRequest, CustomResponse } from './typings';

class Abstract {
    // 外部传入的baseUrl
    protected baseURL: string | undefined = process.env.REACT_APP_BASEURL;
    // 自定义header头
    // eslint-disable-next-line @typescript-eslint/ban-types
    protected headers: object= {
      ContentType: 'application/json;charset=UTF-8',
    }

    // eslint-disable-next-line max-len
    private apiAxios({ baseURL = this.baseURL, headers = this.headers, method, url, data, params, responseType }: AxiosRequest): Promise<CustomResponse> {
      // url解析
      return new Promise((resolve, reject) => {
        instance({
          baseURL,
          headers,
          method,
          url,
          params,
          data,
          responseType,
        }).then((res) => {
          // 200:服务端业务处理正常结束
          if (res.status === 200) {
            if (res.data.code === 0) {
              resolve({ code: res.data.code, message: 'success', data: res.data?.data, origin: res.data });
            } else {
              // TODO: message
              console.log(`${url}请求失败`);
              resolve({ code: res.data.code, message: res.data?.errorMessage || (`${url}请求失败`), data: res.data?.data, origin: res.data });
            }
          } else {
            resolve({ code: res.data.code, message: res.data?.errorMessage || (`${url}请求失败`), data: null });
          }
        }).catch((err) => {
          const message = err?.data?.errorMessage || err?.message || (`${url}请求失败`);
          // TODO: message
          console.log(`${url}请求失败`);
          // eslint-disable-next-line
          reject({ status: false, message, data: null });
        });
      });
    }

    /**
     * GET类型的网络请求
     */
    protected getReq({ baseURL, headers, url, data, params, responseType }: AxiosRequest) {
      return this.apiAxios({ baseURL, headers, method: 'GET', url, data, params, responseType });
    }

    /**
     * POST类型的网络请求
     */
    protected postReq({ baseURL, headers, url, data, params, responseType }: AxiosRequest) {
      return this.apiAxios({ baseURL, headers, method: 'POST', url, data, params, responseType });
    }

    /**
     * PUT类型的网络请求
     */
    protected putReq({ baseURL, headers, url, data, params, responseType }: AxiosRequest) {
      return this.apiAxios({ baseURL, headers, method: 'PUT', url, data, params, responseType });
    }

    /**
     * DELETE类型的网络请求
     */
    protected deleteReq({ baseURL, headers, url, data, params, responseType }: AxiosRequest) {
      return this.apiAxios({ baseURL, headers, method: 'DELETE', url, data, params, responseType });
    }
}

export default Abstract;
