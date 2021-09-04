/**
 * 基础数据 API 集合类
 * 集成Abstract
 * @date 2020-1-14
 */
import Abstract from './abstract';
import { GetDemo, ILogin, PostDemo } from './typings';

class Auth extends Abstract {
  /**
     * get示例
     */
  getDemo(params: GetDemo) {
    return this.getReq({ url: 'Basic.GetDemo', params });
  }

  /**
     * post示例
     */
  login(data: ILogin) {
    return this.postReq({ url: '/login', data });
  }
}

// 单列模式返回对象
let instance;
export default (() => {
  if (instance) return instance;
  instance = new Auth();
  return instance;
})();
