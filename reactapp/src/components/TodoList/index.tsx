import React, {FC, ReactElement, useCallback, useEffect, useReducer, useState} from 'react'
import TdInput from "./Input";
import TdList from "./List";
import {ACTION_TYPE, IState, ITodo} from "./typings";
import {todoReducer} from "./reducer";

const initialState: IState = {
    todoList: []
}

const TodoList: FC = (): ReactElement => {
    // const [todoList, setTodoList] = useState<ITodo[]>([])
    const [state, dispatch] = useReducer(todoReducer, initialState)
    const addTodo = useCallback((todo: ITodo) => {
        // setTodoList(todoList => [...todoList, todo])
        dispatch({
            type: ACTION_TYPE.ADD_TODO,
            payload: todo,
        })
    }, [])

    useEffect(() => {
        console.log(state.todoList);
    }, [state.todoList])

    return (
        <div>
            <TdInput
                addTodo={addTodo}
                todoList={state.todoList}
            />
            <TdList todoList={state.todoList}/>
        </div>
    )
}

export default TodoList;