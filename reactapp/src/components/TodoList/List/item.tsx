import React, { FC, ReactElement } from 'react';
import { ITodo } from '../typings';

interface IProps {
    item: ITodo
}

const TdItem: FC<IProps> = ({ item }): ReactElement => (
  <div className="td-item">
    {item.id}
    -
    {item.content}
    -
    {item.completed}
  </div>
);

export default TdItem;
