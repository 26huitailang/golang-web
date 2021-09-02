import React, {FC, ReactElement} from "react";
import {ITodo} from "../typings";
import TdItem from "./item"

interface IProps {
    todoList: ITodo[];
}

const TdList: FC<IProps> = ({todoList}): ReactElement => {
    return (
        <div className="to-list">
            {todoList.map((item: ITodo) => <TdItem item={item}/>)}
        </div>
    )
}

export default TdList;