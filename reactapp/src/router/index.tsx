import {HashRouter} from 'react-router-dom';
import {createHashHistory, History} from 'history';

export const history: History = createHashHistory();

export default class MyRouter extends HashRouter {
  history: History = history;
}
