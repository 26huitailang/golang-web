import request from '../utils/request';
import { IResponse } from './typings';

export const loginReq = (body: any): IResponse => {
  request.post('/login').then(data: any => {

  })
};
