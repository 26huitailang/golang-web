import React, { FC, ReactElement, useState } from "react";

const Login: FC = (): ReactElement => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const handleSubmit = (event: any) => {
    event.preventDefault()
    console.log(username, password)
    //TODO 账户密码校验
    // todo 请求服务器，获取token，变更app的认证状态
  }
  return (
    <div>
      <form onSubmit={event => handleSubmit(event)}>
        <input type="text" className="username" value={username}
          onChange={event => setUsername(event.target.value)}></input>
        <input type="password" className="password" value={password}
          onChange={event => setPassword(event.target.value)}></input>
        <button type="submit">登录</button>
      </form>
    </div>
  )
}

export default Login;