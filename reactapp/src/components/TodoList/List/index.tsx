import React, { FC, ReactElement } from 'react';
import { ITodo } from '../typings';
import TdItem from './item';

interface IProps {
    todoList: ITodo[];
}

const TdList: FC<IProps> = ({ todoList }): ReactElement => (
  <div className="to-list">
    {todoList.map((item: ITodo) => <TdItem item={item} key={item.id} />)}
  </div>
);

export default TdList;
